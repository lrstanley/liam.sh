// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"strconv"

	"ariga.io/entcache"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/user"
	"github.com/markbates/goth"
)

// Validate authService implements chix.AuthService.
var _ chix.AuthService[ent.User, int] = (*authService)(nil)

func NewAuthService(db *ent.Client, admin int) *authService {
	return &authService{
		db:    db,
		admin: admin,
	}
}

type authService struct {
	db    *ent.Client
	admin int
}

func (s *authService) Get(ctx context.Context, id int) (*ent.User, error) {
	return s.db.User.Get(ctx, id)
}

func (s *authService) Set(ctx context.Context, guser *goth.User) (id int, err error) {
	uid, err := strconv.Atoi(guser.UserID)
	if err != nil {
		return 0, err
	}

	if uid != s.admin {
		return 0, chix.ErrAccessDenied
	}

	q := s.db.User.Create().
		SetUserID(uid).
		SetName(guser.Name).
		SetLogin(guser.NickName).
		SetEmail(guser.Email).
		SetAvatarURL(guser.AvatarURL).
		SetLocation(guser.Location).
		SetBio(guser.Description)

	if data, ok := guser.RawData["html_url"]; ok {
		if url, ok := data.(string); ok {
			q = q.SetHTMLURL(url)
		}
	}

	return q.OnConflictColumns(user.FieldUserID).Ignore().UpdateNewValues().ID(entcache.Evict(ctx))
}

func (s *authService) Roles(ctx context.Context, id int) ([]string, error) {
	if ok, _ := s.db.User.Query().Where(user.IDEQ(id)).Exist(ctx); ok {
		return []string{"admin"}, nil
	}
	return nil, nil
}
