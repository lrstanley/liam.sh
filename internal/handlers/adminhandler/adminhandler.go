// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package adminhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/httpware"
	"github.com/markbates/goth/gothic"
)

func New(db *ent.Client, r *gin.RouterGroup) {
	h := &handler{
		db: db,
	}

	r.Use(httpware.AuthRequired)

	r.GET("/ping", h.ping)
	r.GET("/users", h.getUsers)
}

type handler struct {
	db *ent.Client
}

func (h *handler) ping(c *gin.Context) {
	uid, err := gothic.GetFromSession("user", c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"user":    uid,
	})
}

func (h *handler) getUsers(c *gin.Context) {
	users, err := h.db.User.Query().All(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
