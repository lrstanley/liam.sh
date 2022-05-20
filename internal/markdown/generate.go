// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package markdown

import (
	"bytes"
	"context"

	chromahtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"mvdan.cc/xurls/v2"
)

var engine = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
				// chromahtml.LinkableLineNumbers(true, "ln"),
				chromahtml.TabWidth(4),
			),
		),
		extension.NewLinkify(
			extension.WithLinkifyAllowedProtocols([][]byte{
				[]byte("http:"),
				[]byte("https:"),
			}),
			extension.WithLinkifyURLRegexp(
				xurls.Strict(),
			),
		),
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
		html.WithXHTML(),
		html.WithHardWraps(),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
)

// Generate generates markdown from the given input.
func Generate(ctx context.Context, input string) (string, error) {
	buf := &bytes.Buffer{}

	if err := engine.Convert([]byte(input), buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
