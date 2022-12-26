// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

type GithubStats struct {
	CommitsYear      int `json:"commits_year"`
	PullRequests     int `json:"pull_requests"`
	OpenIssues       int `json:"open_issues"`
	ClosedIssues     int `json:"closed_issues"`
	Issues           int `json:"all_issues"`
	Followers        int `json:"followers"`
	Repos            int `json:"repositories"`
	ContributedRepos int `json:"contributed_repositories"`
	Stars            int `json:"stars"`
}
