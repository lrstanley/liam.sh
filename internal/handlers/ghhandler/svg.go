// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghhandler

import (
	_ "embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/gh"
)

//go:embed project.svg
var rawSVG string

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
?accent=<string:github>   :: Decorator accents (see below for options).
  - computer
  - computer2
  - construction
  - docker
  - gear
  - git
  - github
  - go
  - gopher
  - rocket
  - secure
  - terminal
  - terraform
  - tool
`

type svgParams struct {
	Title       string                `form:"title"       validate:"printascii,max=20"`
	Description string                `form:"description" validate:"printascii,max=100"`
	Layout      string                `form:"layout"      validate:"required,oneof=all left right"`
	Accent      string                `form:"accent"      validate:"required,printascii"`
	Bg          string                `form:"bg"          validate:"required,printascii"`
	BgColor     string                `form:"bgcolor"     validate:"required,iscolor"`
	Height      int                   `form:"h"           validate:"gte=200,lte=1000"`
	Width       int                   `form:"w"           validate:"gte=600,lte=2000"`
	FontScale   float64               `form:"font"        validate:"gte=0.5,lte=3"`
	Repo        *ent.GithubRepository `form:"-"           json:"-"`
}

func (h *handler) getProjectSVG(w http.ResponseWriter, r *http.Request) {
	params := &svgParams{
		Accent:    "github",
		Layout:    "all",
		Bg:        "geometric",
		BgColor:   "#000000",
		Height:    200,
		Width:     961,
		FontScale: 1,
	}

	if !chix.Bind(w, r, params) {
		return
	}

	repoFullName := chi.URLParam(r, "owner") + "/" + chi.URLParam(r, "repo")

	if (params.Title == "" || params.Description == "") && repoFullName == "/" {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, svgDocs)
		return
	}

	if repoFullName != "/" {
		repo, err := h.db.GithubRepository.Query().Where(
			githubrepository.FullNameEqualFold(repoFullName),
		).First(r.Context())

		if chix.Error(w, r, http.StatusInternalServerError, err) {
			return
		}

		params.Repo = repo
	}

	params.Description = strings.TrimSpace(gh.ReStripEmoji.ReplaceAllString(params.Description, " "))

	w.Header().Set("Content-Type", "image/svg+xml")
	err := h.projectSVG.Execute(w, params)
	if err != nil {
		chix.Error(w, r, http.StatusInternalServerError, err)
	}
}
