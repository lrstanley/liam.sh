// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package httpware

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/markbates/goth/gothic"
)

func SecurityText(c *gin.Context) {
	c.String(http.StatusOK, strings.TrimLeft(`
Contact: mailto:me@liamstanley.io
Contact: https://liam.sh/chat
Contact: https://github.com/lrstanley
Encryption: https://github.com/lrstanley.gpg
Preferred-Languages: en
`, "\n"))
}

func AuthRequired(c *gin.Context) {
	if GetUser(c) == nil {
		Error(c, http.StatusUnauthorized, errors.New("unauthorized"))
		c.Abort()
		return
	}
	c.Next()
}

func InjectUser(db *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key, err := gothic.GetFromSession(models.UserSessionKey, c.Request)
		if err != nil {
			c.Next()
			return
		}

		userInt, err := strconv.Atoi(key)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			c.Abort()
			return
		}

		u, err := db.User.Get(c.Request.Context(), userInt)
		if err != nil {
			if ent.IsNotFound(err) {
				c.Next()
				return
			}
			Error(c, http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), models.UserContextKey, u)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetUser(c *gin.Context) *ent.User {
	u, _ := c.Request.Context().Value(models.UserContextKey).(*ent.User)
	return u
}
