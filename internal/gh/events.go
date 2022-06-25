// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"encoding/json"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/v44/github"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

const eventsInterval = 30 * time.Minute

func EventsRunner(ctx context.Context) error {
	logger := log.FromContext(ctx).WithField("runner", "github_events")
	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}

	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	err := database.RunWithTx(ctx, logger, db, getEvents)
	if err != nil {
		logger.WithError(err).Error("failed to get events")
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(eventsInterval):
			err = database.RunWithTx(ctx, logger, db, getEvents)
			if err != nil {
				logger.WithError(err).Error("failed to get events")
			}
		}
	}
}

var (
	allowedEvents = []string{
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
)

func getEvents(ctx context.Context, logger log.Interface, db *ent.Tx) error {
	opts := &github.ListOptions{
		PerPage: 100,
		Page:    1,
	}

	var allEvents []*github.Event
	var resp *github.Response
	var user *github.User
	var err error

	user, _, err = Client.Users.Get(ctx, "")
	if err != nil {
		logger.WithError(err).Error("failed to get user")
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

		logger.WithField("page", opts.Page).Info("querying events")
		events, resp, err = Client.Activity.ListEventsPerformedByUser(ctx, user.GetLogin(), false, opts)
		if err != nil {
			return err
		}

		for _, event := range events {
			isAllowed := false
			for _, allowed := range allowedEvents {
				if event.GetType() == allowed {
					isAllowed = true
					break
				}
			}

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

		payload := map[string]interface{}{}
		err = json.Unmarshal([]byte(event.GetRawPayload()), &payload)
		if err != nil {
			return err
		}

		count++

		err = db.GithubEvent.Create().
			SetEventID(event.GetID()).
			SetEventType(event.GetType()).
			SetCreatedAt(event.GetCreatedAt()).
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

	logger.WithField("events", count).Info("fetched newest events")
	return nil
}
