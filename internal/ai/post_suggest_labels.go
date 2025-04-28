// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ai

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

const maxPostContentBytes = 1000 * 512 // 512KB

type SuggestedLabel struct {
	Name       string `json:"name"`
	ExistingID *int   `json:"existing_id,omitempty"`
}

func (s *service) PostSuggestLabels(ctx context.Context, content string) ([]SuggestedLabel, error) {
	if len(content) > maxPostContentBytes {
		return nil, fmt.Errorf("content is too large: %d bytes", len(content))
	}
	if len(content) == 0 {
		return nil, nil
	}

	// TODO: check to see if OpenRouter charges when we give up on requests.
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	client, err := s.newClient(ctx, WithJSONSchema(&openai.ResponseFormatJSONSchemaProperty{
		Type: "object",
		Properties: map[string]*openai.ResponseFormatJSONSchemaProperty{
			"labels": {
				Type:        "array",
				Description: "List of labels recommended for the post. Labels should be in the format of kebab case, e.g. 'label-name'.",
				Items: &openai.ResponseFormatJSONSchemaProperty{
					Type:        "string",
					Description: "Label name. Should be in the format of kebab case, e.g. 'label-name'.",
				},
			},
		},
		AdditionalProperties: false,
		Required:             []string{"labels"},
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	exampleLabels, err := s.db.Label.Query().Select(label.FieldName).Strings(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query labels: %w", err)
	}

	type rawSuggestedLabels struct {
		Labels []string `json:"labels"`
	}

	results, _, err := GetJSONSchemaResponse[rawSuggestedLabels](
		ctx, client, []llms.MessageContent{
			llms.TextParts(
				llms.ChatMessageTypeSystem,
				"You are an expert at structured data extraction. You will be given unstructured text in markdown format from a blog post and you should return what labels/tags are relevant to the post in the provided structure. Each label must be distinct (not describe or imply the same thing). Here are some labels which are used by other posts, which can also be used for this post:",
				strings.Join(exampleLabels, ", "),
			),
			llms.TextParts(llms.ChatMessageTypeHuman, content),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	suggestedLabels := make([]SuggestedLabel, 0, len(results.Labels))

	var filteredLabels []string

	for _, l := range results.Labels {
		if err = label.NameValidator(l); err != nil {
			log.FromContext(ctx).WithError(err).WithField("label", l).Warn("ignoring suggested invalid label name")
		}
		filteredLabels = append(filteredLabels, l)
	}

	if len(filteredLabels) == 0 {
		return nil, nil
	}

	chunked := database.Chunked(ctx, 100, s.db.Label.Query().Where(label.NameIn(filteredLabels...)))

	for l, err := range chunked {
		if err != nil {
			return nil, fmt.Errorf("failed to query labels: %w", err)
		}

		suggestedLabels = append(suggestedLabels, SuggestedLabel{
			Name:       l.Name,
			ExistingID: &l.ID,
		})
	}

	for _, l := range filteredLabels {
		if !slices.ContainsFunc(suggestedLabels, func(s SuggestedLabel) bool { return s.Name == l }) {
			suggestedLabels = append(suggestedLabels, SuggestedLabel{
				Name: l,
			})
		}
	}

	// Sort by name.
	slices.SortFunc(suggestedLabels, func(a, b SuggestedLabel) int {
		return strings.Compare(a.Name, b.Name)
	})

	return suggestedLabels, nil
}
