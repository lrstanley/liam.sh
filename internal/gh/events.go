// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"encoding/json"
	"log/slog"
	"slices"
	"time"

	"github.com/google/go-github/v82/github"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

var allowedEvents = []string{
	"CreateEvent",
	"DeleteEvent",
	"ForkEvent",
	"IssueCommentEvent",
	"IssuesEvent",
	"MemberEvent",
	"PublicEvent",
	"PullRequestEvent",
	"PullRequestReviewEvent",
	"PullRequestReviewCommentEvent",
	"PullRequestReviewThreadEvent",
	"PushEvent",
	"ReleaseEvent",
	"WatchEvent",
}

// EventsRunner fetches all events for the authenticated user from Github, storing
// them in the database.
func EventsRunner(ctx context.Context, logger *slog.Logger) error {
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}

	opts := &github.ListOptions{
		PerPage: 100,
		Page:    1,
	}

	var allEvents []*github.Event
	var resp *github.Response
	var user *github.User
	var err error

	user, _, err = RestClient.Users.Get(ctx, "")
	if err != nil {
		logger.ErrorContext(ctx, "failed to get user", "error", err)
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		time.Sleep(5 * time.Second)

		var events []*github.Event

		logger.InfoContext(ctx, "querying events", "page", opts.Page)
		events, resp, err = RestClient.Activity.ListEventsPerformedByUser(ctx, user.GetLogin(), false, opts)
		if err != nil {
			return err
		}

		for _, event := range events {
			isAllowed := slices.Contains(allowedEvents, event.GetType())

			if !isAllowed {
				continue
			}

			allEvents = append(allEvents, event)
		}

		if resp.NextPage < 1 {
			break
		}
		opts.Page = resp.NextPage
	}

	var exists bool
	var count int

	for _, event := range allEvents {
		exists, err = db.GithubEvent.Query().Where(githubevent.EventIDEqualFold(event.GetID())).Exist(ctx)
		if err != nil {
			return err
		}

		if exists {
			continue
		}

		payload := map[string]any{}
		err = json.Unmarshal([]byte(event.GetRawPayload()), &payload)
		if err != nil {
			return err
		}

		count++

		err = db.GithubEvent.Create().
			SetEventID(event.GetID()).
			SetEventType(event.GetType()).
			SetCreatedAt(event.GetCreatedAt().Time).
			SetPublic(event.GetPublic()).
			SetActorID(event.GetActor().GetID()).
			SetActor(event.GetActor()).
			SetRepoID(event.GetRepo().GetID()).
			SetRepo(event.GetRepo()).
			SetPayload(payload).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	logger.InfoContext(ctx, "fetched newest events", "events", count)
	return nil
}
