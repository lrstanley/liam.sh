// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package authhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/user"
	"github.com/lrstanley/liam.sh/internal/httpware"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func New(db *ent.Client, r *gin.RouterGroup, config *models.ConfigGithub) {
	h := &handler{
		db:     db,
		config: config,
	}

	r.GET("/self", httpware.AuthRequired, h.getSelf)
	r.GET("/providers", h.providers)
	r.GET("/providers/:provider", h.provider)
	r.GET("/providers/:provider/callback", h.callback)
	r.GET("/logout", h.logout)
}

type handler struct {
	db     *ent.Client
	config *models.ConfigGithub
}

func (h *handler) providers(c *gin.Context) {
	providers := goth.GetProviders()
	var data []string
	for _, p := range providers {
		data = append(data, p.Name())
	}

	c.JSON(http.StatusOK, gin.H{"providers": data})
}

func (h *handler) provider(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, gothic.GetContextWithProvider(c.Request, c.Param("provider")))
}

func (h *handler) callback(c *gin.Context) {
	guser, err := gothic.CompleteUserAuth(c.Writer, gothic.GetContextWithProvider(c.Request, c.Param("provider")))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if guser.UserID != strconv.Itoa(h.config.User) {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	uid, err := strconv.Atoi(guser.UserID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := h.db.User.Create().
		SetUserID(uid).
		SetName(guser.Name).
		SetLogin(guser.NickName).
		SetEmail(guser.Email).
		SetAvatarURL(guser.AvatarURL).
		SetLocation(guser.Location).
		SetBio(guser.Description).
		OnConflictColumns(user.FieldUserID).Ignore().
		UpdateNewValues().
		IDX(c.Request.Context())

	if err := gothic.StoreInSession(models.UserSessionKey, strconv.Itoa(id), c.Request, c.Writer); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (h *handler) logout(c *gin.Context) {
	_ = gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusFound, "/")
}

func (h *handler) getSelf(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": httpware.GetUser(c)})
}
