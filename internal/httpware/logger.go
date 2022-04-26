// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package httpware

import (
	"time"

	"github.com/apex/log"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// StructuredLogger wraps each request and writes a log entry with
// extra info. StructuredLogger also injects a logger into the request
// context that can be used by children middleware business logic.
func StructuredLogger(logger log.Interface, private bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		logEntry := logger.WithField("src", "httphandler")

		// RequestID middleware must be loaded before this is loaded into
		// the chain.
		if id := requestid.Get(c); id != "" {
			logEntry = logEntry.WithField("request_id", id)
		}

		start := time.Now()
		c.Next()

		finish := time.Now()

		if !private {
			logEntry = logEntry.WithField("remote_ip", c.ClientIP())
		}

		if u := GetUser(c); u != nil {
			logEntry = logEntry.WithFields(log.Fields{
				"user_id":    u.ID,
				"user_login": u.Login,
			})
		}

		logEntry.WithFields(log.Fields{
			"host":        c.Request.Host,
			"proto":       c.Request.Proto,
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"user_agent":  c.Request.Header.Get("User-Agent"),
			"status":      c.Writer.Status(),
			"duration_ms": float64(finish.Sub(start).Nanoseconds()) / 1000000.0,
			"bytes_in":    c.Request.Header.Get("Content-Length"),
			"bytes_out":   c.Writer.Size(),
		}).Info("handled request")
	}
}
