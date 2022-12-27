// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/models"
	ghql "github.com/shurcooL/githubv4"
)

const statsInterval = 4 * time.Hour

var Stats atomic.Pointer[models.GithubStats]

func StatsRunner(ctx context.Context) error {
	logger := log.FromContext(ctx).WithField("runner", "github_stats")

	getStats(ctx, logger)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(statsInterval):
			getStats(ctx, logger)
		}
	}
}

type baseStatsQuery struct {
	Viewer struct {
		ContributionsCollection struct {
			TotalCommitContributions     int
			RestrictedContributionsCount int
		}
		RepositoriesContributedTo struct {
			TotalCount int
		} `graphql:"repositoriesContributedTo(contributionTypes: [COMMIT, ISSUE, PULL_REQUEST, REPOSITORY])"`
		PullRequests struct {
			TotalCount int
		}
		OpenIssues struct {
			TotalCount int
		} `graphql:"openIssues: issues(states: OPEN)"`
		ClosedIssues struct {
			TotalCount int
		} `graphql:"closedIssues: issues(states: CLOSED)"`
		Repositories struct {
			TotalCount int
		} `graphql:"repositories(ownerAffiliations: OWNER)"`
	}
}

type repoStarsQuery struct {
	Viewer struct {
		Repositories struct {
			Nodes []struct {
				Name       string
				Stargazers struct {
					TotalCount int
				}
			}
			PageInfo struct {
				HasNextPage bool
				EndCursor   ghql.String
			}
		} `graphql:"repositories(first: 100, ownerAffiliations: OWNER, orderBy: {field: STARGAZERS, direction: DESC}, after: $after)"`
	}
}

func getStats(ctx context.Context, logger log.Interface) {
	base := baseStatsQuery{}
	err := GraphClient.Query(ctx, &base, nil)
	if err != nil {
		logger.WithError(err).Error("failed to get stats")
		return
	}
	logger.Info("retrieved base github stats")

	stats := &models.GithubStats{
		CommitsYear:      base.Viewer.ContributionsCollection.TotalCommitContributions + base.Viewer.ContributionsCollection.RestrictedContributionsCount,
		PullRequests:     base.Viewer.PullRequests.TotalCount,
		OpenIssues:       base.Viewer.OpenIssues.TotalCount,
		ClosedIssues:     base.Viewer.ClosedIssues.TotalCount,
		Issues:           base.Viewer.OpenIssues.TotalCount + base.Viewer.ClosedIssues.TotalCount,
		Repos:            base.Viewer.Repositories.TotalCount,
		ContributedRepos: base.Viewer.RepositoriesContributedTo.TotalCount,
	}

	repoStarVariables := map[string]any{
		"after": (*ghql.String)(nil),
	}

	for {
		repoStars := repoStarsQuery{}
		err := GraphClient.Query(ctx, &repoStars, repoStarVariables)
		if err != nil {
			logger.
				WithField("cursor", repoStarVariables["after"]).
				WithError(err).
				Error("failed to get repo stars")
			return
		}

		logger.
			WithField("cursor", repoStarVariables["after"]).
			Info("retrieved repo star stats")

		for _, repo := range repoStars.Viewer.Repositories.Nodes {
			stats.Stars += repo.Stargazers.TotalCount
		}

		if !repoStars.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}

		repoStarVariables["after"] = ghql.NewString(repoStars.Viewer.Repositories.PageInfo.EndCursor)
	}

	Stats.Store(stats)
}
