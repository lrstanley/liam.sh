// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/v52/github"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubgist"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

const gistInterval = 120 * time.Minute

func GistRunner(ctx context.Context) error {
	logger := log.FromContext(ctx).WithField("runner", "github_gists")
	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	var err error
	if SyncOnStart {
		err = database.RunWithTx(ctx, logger, db, getGists)
		if err != nil {
			logger.WithError(err).Error("failed to get gists")
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(gistInterval):
			err = database.RunWithTx(ctx, logger, db, getGists)
			if err != nil {
				logger.WithError(err).Error("failed to get gists")
			}
		}
	}
}

func getGists(ctx context.Context, logger log.Interface, db *ent.Tx) error {
	gists, err := fetchGists(ctx, logger, db)
	if err != nil {
		return fmt.Errorf("failed to fetch gists: %w", err)
	}

	var req *http.Request
	buf := &bytes.Buffer{}

	for _, gist := range gists {
		for _, file := range gist.GetFiles() {
			content, _ := db.GithubGist.Query().Where(
				githubgist.GistID(gist.GetID()),
				githubgist.UpdatedAt(gist.GetUpdatedAt().Time),
				githubgist.Name(file.GetFilename()),
			).Limit(1).Select(githubgist.FieldContent).String(ctx)

			if content == "" {
				req, err = http.NewRequest(http.MethodGet, file.GetRawURL(), http.NoBody)
				if err != nil {
					return fmt.Errorf("failed to create raw gist request: %w", err)
				}

				buf.Reset()
				_, err = RestClient.Do(ctx, req, buf)
				if err != nil {
					return fmt.Errorf("failed to fetch gist: %w", err)
				}

				content = buf.String()
			}

			file.Content = &content

			_, err = storeGist(ctx, db, gist, file)
			if err != nil {
				return fmt.Errorf("failed to store gist: %v %v; %w", gist.GetID(), file.GetFilename(), err)
			}
		}
	}

	err = removeGists(ctx, db, gists)
	if err != nil {
		return fmt.Errorf("failed to remove gists: %w", err)
	}

	logger.WithField("gists", len(gists)).Info("fetched newest gists")
	return nil
}

// fetchGists fetches all gists for the authenticated user from Github. It
// will also iterate through all pages, returning all gists in their entirety.
func fetchGists(ctx context.Context, logger log.Interface, db *ent.Tx) (allGists []*github.Gist, err error) {
	opts := &github.GistListOptions{
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

		var gists []*github.Gist

		logger.WithField("page", opts.ListOptions.Page).Info("querying gists")
		gists, resp, err = RestClient.Gists.List(ctx, "", opts)
		if err != nil {
			return nil, err
		}

		for _, gist := range gists {
			allGists = append(allGists, gist)
		}

		if resp.NextPage < 1 {
			break
		}
		opts.ListOptions.Page = resp.NextPage
	}

	return allGists, nil
}

var reLatestRawURL = regexp.MustCompile(`/raw/[^/]+/`)

// storeGist fetches the gist from github and stores it in the database.
func storeGist(ctx context.Context, db *ent.Tx, gist *github.Gist, file github.GistFile) (id int, err error) {
	q := db.GithubGist.Create().
		SetGistID(gist.GetID()).
		SetHTMLURL(gist.GetHTMLURL()).
		SetPublic(gist.GetPublic()).
		SetCreatedAt(gist.GetCreatedAt().Time).
		SetUpdatedAt(gist.GetUpdatedAt().Time).
		SetDescription(gist.GetDescription()).
		SetOwner(gist.GetOwner()).
		SetName(file.GetFilename()).
		SetType(file.GetType()).
		SetLanguage(file.GetLanguage()).
		SetSize(int64(file.GetSize())).
		SetRawURL(reLatestRawURL.ReplaceAllString(file.GetRawURL(), "/raw/"))

	if file.GetContent() != "" {
		q = q.SetContent(file.GetContent())
	}

	return q.OnConflictColumns(githubgist.FieldGistID, githubgist.FieldName).UpdateNewValues().ID(ctx)
}

// removeGists removes all gists that are in the database, but not
// on github.
func removeGists(ctx context.Context, db *ent.Tx, gists []*github.Gist) (err error) {
	var storedGists []string
	err = db.GithubGist.Query().Select(githubgist.FieldGistID).Scan(ctx, &storedGists)
	if err != nil {
		return err
	}

	for _, id := range storedGists {
		contains := false
		for _, gist := range gists {
			if gist.GetID() == id {
				contains = true
				break
			}
		}
		if contains {
			continue
		}

		_, err = db.GithubGist.Delete().
			Where(githubgist.GistID(id)).
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
