// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ai

import (
	"net/http"
	"time"

	"github.com/apex/log"
)

// Example if we wanted to dump the full request or response: https://github.com/tmc/langchaingo/blob/main/httputil/debug_transport.go

type httpTransport struct {
	headers map[string]string
	chain   http.RoundTripper
	logger  log.Interface
}

func (h *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range h.headers {
		req.Header.Add(key, value)
	}

	// Log request(s).
	rlogger := h.logger.WithFields(log.Fields{
		"method": req.Method,
		"url":    req.URL.String(),
	})

	start := log.Now()
	rlogger.Debug("sending request")

	var resp *http.Response
	var err error

	if h.chain != nil {
		resp, err = h.chain.RoundTrip(req)
	} else {
		resp, err = http.DefaultTransport.RoundTrip(req)
	}

	// Log response(s).
	if err != nil {
		rlogger.Error("request failed")
	} else if resp.StatusCode >= 400 {
		rlogger.WithFields(log.Fields{
			"status": resp.StatusCode,
			"took":   time.Since(start),
		}).Error("request failed")
	} else {
		rlogger.WithFields(log.Fields{
			"status": resp.StatusCode,
			"took":   time.Since(start),
		}).Debug("request completed")
	}

	return resp, err
}
