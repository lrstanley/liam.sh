// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghhandler

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/cache"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/gh"
)

var (
	//go:embed project.svg
	rawSVG string

	iconCache = cache.New[string, string](100, 1*time.Hour)
)

const svgDocs = `# SVG Generator Routes:
GET /-/svg?options                 :: Dynamic SVG Generator based off title/description params.
GET /-/svg/<owner>/<repo>?options  :: Must map to a repository in the database.

[parameters]

?title=<string>           :: Title of the project (not required if using owner/repo).
?description=<string>     :: Description of the project (not required if using owner/repo).
?h=<int:200>              :: Height of the SVG (in px).
?w=<int:961>              :: Width of the SVG (in px).
?font=<float:1.0>         :: Font scale.
?bg=<string:geometric>    :: Background type.
  - geometric
  - topography
?bgcolor=<string:#000000> :: Background color (hex, rgb, rgba, hsl).
?layout=<string:all>      :: Layout type.
  - all
  - left
  - right
?icon=<string:mdi>:<string:github> :: Iconify icon to use. See: https://icones.js.org/
?icon.height=<int>                 :: Height of the icon (in px).
?icon.width=<int>                  :: Width of the icon (in px).
?icon.flip=<string>                :: Flip the icon.
  - horizontal
  - vertical
?icon.rotate=<int>                 :: Rotate the icon (1=90deg, 2=180deg, 3=270deg).
?icon.color=<string>               :: Color of the icon (hex, rgb, rgba, hsl).
`

type svgParams struct {
	Title       string  `form:"title"       validate:"printascii,max=20"`
	Description string  `form:"description" validate:"printascii,max=100"`
	Layout      string  `form:"layout"      validate:"required,oneof=all left right"`
	Bg          string  `form:"bg"          validate:"required,printascii"`
	BgColor     string  `form:"bgcolor"     validate:"required,iscolor"`
	Height      int     `form:"h"           validate:"gte=200,lte=1000"`
	Width       int     `form:"w"           validate:"gte=600,lte=2000"`
	FontScale   float64 `form:"font"        validate:"gte=0.5,lte=3"`

	Icon       string `form:"icon"        validate:"printascii"`
	IconHeight int    `form:"icon.height" validate:"gte=0,lte=300"`
	IconWidth  int    `form:"icon.width"  validate:"gte=0,lte=300"`
	IconFlip   string `form:"icon.flip"   validate:"omitempty,oneof=horizontal vertical"`
	IconRotate int    `form:"icon.rotate" validate:"gte=0,lte=3"`
	IconColor  string `form:"icon.color"  validate:"omitempty,iscolor"`

	Repo    *ent.GithubRepository `form:"-"           json:"-"`
	IconSVG string                `form:"-"           json:"-"`
}

func (h *handler) getIconifySVG(ctx context.Context, p *svgParams) string {
	if p.Icon == "" {
		return ""
	}

	_url := fmt.Sprintf("https://api.iconify.design/%s.svg", p.Icon)
	query := make(url.Values)

	if p.IconHeight > 0 {
		query.Set("height", fmt.Sprintf("%d", p.IconHeight))
	}

	if p.IconWidth > 0 {
		query.Set("width", fmt.Sprintf("%d", p.IconWidth))
	}

	if p.IconFlip != "" {
		query.Set("flip", p.IconFlip)
	}

	if p.IconRotate > 0 {
		query.Set("rotate", fmt.Sprintf("%d", p.IconRotate))
	}

	if p.IconColor != "" {
		query.Set("color", p.IconColor)
	}

	if len(query) > 0 {
		_url += "?" + query.Encode()
	}

	if icon := iconCache.Get(_url); icon != "" {
		return icon
	}

	logger := log.FromContext(ctx)

	resp, err := http.Get(_url)
	if err != nil {
		logger.WithError(err).WithField("url", _url).Error("failed to get iconify svg")
		iconCache.Set(_url, "")
		return ""
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.WithError(err).WithField("url", _url).Error("failed to read iconify svg")
		return ""
	}

	icon := string(b)
	iconCache.Set(_url, icon)
	return icon
}

func (h *handler) getProjectSVG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "max-age=120, s-maxage=120")
	w.Header().Set("Expires", time.Now().Add(2*time.Hour).Format(http.TimeFormat))

	params := &svgParams{
		Layout:    "all",
		Bg:        "geometric",
		BgColor:   "#000000",
		Height:    200,
		Width:     961,
		FontScale: 1,
		Icon:      "mdi:github",
	}

	if err := chix.Bind(r, params); err != nil {
		chix.Error(w, r, err)
		return
	}

	repoFullName := chi.URLParam(r, "owner") + "/" + chi.URLParam(r, "repo")

	if (params.Title == "" || params.Description == "") && repoFullName == "/" || params.Icon == "" {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, svgDocs)
		return
	}

	if repoFullName != "/" {
		repo, err := h.db.GithubRepository.Query().Where(
			githubrepository.FullNameEqualFold(repoFullName),
		).First(r.Context())

		if chix.Error(w, r, err) {
			return
		}

		params.Repo = repo
	}

	params.Description = strings.TrimSpace(gh.ReStripEmoji.ReplaceAllString(params.Description, " "))

	if icon := h.getIconifySVG(r.Context(), params); icon != "" {
		params.IconSVG = icon
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	err := h.projectSVG.Execute(w, params)
	if err != nil {
		chix.Error(w, r, err)
	}
}
