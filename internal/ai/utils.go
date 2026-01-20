// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ai

import (
	"context"
	"encoding/json"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func WithJSONSchema(schema *openai.ResponseFormatJSONSchemaProperty) openai.Option {
	return openai.WithResponseFormat(&openai.ResponseFormat{
		Type: "json_schema",
		JSONSchema: &openai.ResponseFormatJSONSchema{
			Name:   "object",
			Strict: true,
			Schema: schema,
		},
	})
}

func GetJSONSchemaResponse[T any](
	ctx context.Context,
	client llms.Model,
	messages []llms.MessageContent,
	options ...llms.CallOption,
) (value *T, resp *llms.ContentResponse, err error) {
	options = append(options, llms.WithJSONMode())

	// TODO: builtin retry support here?

	resp, err = client.GenerateContent(ctx, messages, options...)
	if err != nil {
		return nil, nil, err
	}

	var result T

	err = json.Unmarshal([]byte(resp.Choices[0].Content), &result)
	if err != nil {
		return nil, nil, err
	}
	return &result, resp, nil
}
