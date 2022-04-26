// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package httpware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, statusCode int, err error) bool {
	if statusCode == http.StatusNotFound && err == nil {
		err = errors.New("the requested resource was not found")
	}

	if err == nil {
		return false
	}

	// TODO: check for db errors here.
	// if statusCode == http.StatusInternalServerError && models.IsClientError(err) {
	// 	if models.IsNotFound(err) {
	// 		statusCode = http.StatusNotFound
	// 	} else {
	// 		// if it's internal server error, override since we know it's a client error.
	// 		statusCode = http.StatusBadRequest
	// 	}
	// }

	statusText := http.StatusText(statusCode)

	if strings.HasPrefix(c.Request.URL.Path, "/api/") {
		c.JSON(statusCode, gin.H{
			"error":      err.Error(),
			"type":       statusText,
			"code":       statusCode,
			"request_id": requestid.Get(c),
			"timestamp":  time.Now().UTC().Format(time.RFC3339),
		})
	} else {
		c.Writer.WriteHeader(statusCode)
		http.Error(c.Writer, fmt.Sprintf(
			"%s: %s (id: %s)",
			statusText,
			err.Error(),
			requestid.Get(c),
		), statusCode)
	}

	if statusCode >= 500 {
		log.FromContext(c.Request.Context()).WithFields(log.Fields{
			"status":      statusCode,
			"status_text": statusText,
		}).Errorf("http error: %v", err)
	}

	return true
}

func ErrorCatcher(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()
	if err == nil {
		return
	}

	_ = Error(c, c.Writer.Status(), err)
}
