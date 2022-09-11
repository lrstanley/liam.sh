// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/v44/github"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/ent/label"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

const repositoryInterval = 30 * time.Minute

func RepositoryRunner(ctx context.Context) error {
	logger := log.FromContext(ctx).WithField("runner", "github_repositories")
	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	var err error
	if SyncOnStart {
		err = database.RunWithTx(ctx, logger, db, getRepositories)
		if err != nil {
			logger.WithError(err).Error("failed to get repositories")
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(repositoryInterval):
			err = database.RunWithTx(ctx, logger, db, getRepositories)
			if err != nil {
				logger.WithError(err).Error("failed to get repositories")
			}
		}
	}
}

var ReStripEmoji = regexp.MustCompile(`\s*:[^:]+:\s*`)

func getRepositories(ctx context.Context, logger log.Interface, db *ent.Tx) error {
	repos, err := fetchRepositories(ctx, logger, db)
	if err != nil {
		return fmt.Errorf("failed to fetch repositories: %w", err)
	}

	user, _, err := Client.Users.Get(ctx, "")
	if err != nil {
		logger.WithError(err).Error("failed to get user")
		return err
	}

	lc := database.NewLabelCreator()

	var repoId int
	var releaseId int
	var releases []*github.RepositoryRelease

	for _, repo := range repos {
		repoId, err = storeRepository(ctx, db, lc, repo)
		if err != nil {
			return fmt.Errorf("failed to store repository: %v; %w", repo.GetName(), err)
		}

		// Only fetch releases/assets for the authenticated user.
		if repo.GetOwner().GetLogin() != user.GetLogin() {
			continue
		}

		if repo.GetFork() || repo.GetDisabled() || !repo.GetHasDownloads() {
			logger.WithField("repo", repo.GetFullName()).Info("skipping repository release checks (fork/disabled/no-downloads)")
			continue
		}

		// Check if archived. If so, and we already have releases, skip.
		if repo.GetArchived() {
			var exists bool
			exists, _ = db.GithubRelease.Query().
				Where(githubrelease.HasRepositoryWith(
					githubrepository.RepoID(int64(repoId))),
				).Exist(ctx)
			if exists {
				logger.WithField("repo", repo.GetFullName()).Info("skipping repository release checks (archived)")
				continue
			}
		}

		releases, err = fetchReleases(ctx, logger, db, repo)
		if err != nil {
			return fmt.Errorf("failed to fetch releases: %w", err)
		}

		for _, release := range releases {
			releaseId, err = storeRelease(ctx, db, repoId, release)
			if err != nil {
				return fmt.Errorf("failed to store release: %w", err)
			}

			for _, asset := range release.Assets {
				_, err = storeAsset(ctx, db, releaseId, asset)
				if err != nil {
					return fmt.Errorf("failed to store asset: %w", err)
				}
			}
		}
	}

	err = removeRepositories(ctx, db, repos)
	if err != nil {
		return fmt.Errorf("failed to remove repositories: %w", err)
	}

	logger.WithField("repos", len(repos)).Info("fetched newest repositories")
	return nil
}

// fetchRepositories fetches all repositories for the authenticated user from Github. It
// will also iterate through all pages, returning all repositories in their entirety.
func fetchRepositories(ctx context.Context, logger log.Interface, db *ent.Tx) (allRepos []*github.Repository, err error) {
	opts := &github.RepositoryListOptions{
		Visibility:  "all",
		Affiliation: "owner,collaborator",
		ListOptions: github.ListOptions{PerPage: 100, Page: 1},
	}

	var resp *github.Response

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		time.Sleep(1 * time.Second)

		var repos []*github.Repository

		logger.WithField("page", opts.ListOptions.Page).Info("querying repositories")
		repos, resp, err = Client.Repositories.List(ctx, "", opts)
		if err != nil {
			return nil, err
		}

		for _, repo := range repos {
			allRepos = append(allRepos, repo)
		}

		if resp.NextPage < 1 {
			break
		}
		opts.ListOptions.Page = resp.NextPage
	}

	return allRepos, nil
}

// storeRepository fetches the repository from github and stores it in the database.
func storeRepository(ctx context.Context, db *ent.Tx, lc *database.LabelCreator, repo *github.Repository) (id int, err error) {
	// Languages aren't embedded into the repository struct, so we have to
	// make a separate request for each repo to get the list of languages.
	languages, _, err := Client.Repositories.ListLanguages(ctx, repo.GetOwner().GetLogin(), repo.GetName())
	if err != nil {
		return id, err
	}

	inputLabels := repo.Topics
	for lang := range languages {
		inputLabels = append(inputLabels, lang)
	}

	// Query (or create if missing) all labels related to the repository.
	err = lc.Populate(ctx, db, inputLabels)
	if err != nil {
		return id, err
	}

	// Upsert repository.
	id, err = db.GithubRepository.Create().
		SetRepoID(repo.GetID()).
		SetName(repo.GetName()).
		SetFullName(repo.GetFullName()).
		SetOwnerLogin(repo.GetOwner().GetLogin()).
		SetOwner(repo.GetOwner()).
		SetPublic(!repo.GetPrivate()).
		SetHTMLURL(repo.GetHTMLURL()).
		SetDescription(strings.TrimSpace(ReStripEmoji.ReplaceAllString(repo.GetDescription(), " "))).
		SetFork(repo.GetFork()).
		SetHomepage(repo.GetHomepage()).
		SetStarCount(repo.GetStargazersCount()).
		SetDefaultBranch(repo.GetDefaultBranch()).
		SetIsTemplate(repo.GetIsTemplate()).
		SetHasIssues(repo.GetHasIssues()).
		SetArchived(repo.GetArchived()).
		SetPushedAt(repo.GetPushedAt().Time).
		SetCreatedAt(repo.GetCreatedAt().Time).
		SetUpdatedAt(repo.GetUpdatedAt().Time).
		SetLicense(repo.GetLicense()).
		OnConflictColumns(githubrepository.FieldRepoID).
		UpdateNewValues().ID(ctx)
	if err != nil {
		return id, err
	}

	// Query current labels.
	oldLabels, err := db.Label.Query().
		Where(label.HasGithubRepositoriesWith(githubrepository.ID(id))).
		IDs(ctx)
	if err != nil {
		return id, err
	}

	// Add/remove label edges for the given repository.
	add, remove := lc.Diff(lc.Get(inputLabels), oldLabels)
	err = db.GithubRepository.Update().
		RemoveLabelIDs(remove...).
		AddLabelIDs(add...).
		Where(githubrepository.ID(id)).
		Exec(ctx)

	return id, err
}

// removeRepositories removes all repositories that are in the database, but not
// on github.
func removeRepositories(ctx context.Context, db *ent.Tx, repos []*github.Repository) (err error) {
	var storedRepos []int64
	err = db.GithubRepository.Query().
		Select(githubrepository.FieldRepoID).
		Scan(ctx, &storedRepos)
	if err != nil {
		return err
	}

	for _, id := range storedRepos {
		contains := false
		for _, repo := range repos {
			if repo.GetID() == id {
				contains = true
				break
			}
		}
		if contains {
			continue
		}

		_, err = db.GithubRepository.Delete().
			Where(githubrepository.RepoID(id)).
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
