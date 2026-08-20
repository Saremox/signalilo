/*
 * Authors:
 * Simon Gerber <simon.gerber@vshn.ch>
 *
 * License:
 * Copyright (c) 2019, VSHN AG, <info@vshn.ch>
 * Licensed under "BSD 3-Clause". See LICENSE file.
 */

package webhook

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/saremox/go-icinga2-client/icinga2"
	"github.com/saremox/signalilo/config"
	"github.com/stretchr/testify/assert"
)

func mockEchoHandler(w http.ResponseWriter, r *http.Request) {
	asJSON(w, http.StatusOK, "ok")
}

func TestAsJSON(t *testing.T) {
	handler := http.HandlerFunc(mockEchoHandler)

	// verify response properties
	assert.HTTPSuccess(t, handler, "GET", "http://example.com/webhook", nil)
	response := assert.HTTPBody(handler, "GET", "http://example.com/webhook", nil)
	assert.JSONEq(t, response, `{ "Status": 200, "Message": "ok" }`)
}

func TestBearerTokenHeader(t *testing.T) {
	conf := config.NewMockConfiguration(1)
	req, _ := http.NewRequest(http.MethodPost, "https://example.com/webhook", nil)
	req.Header.Add("Authorization", "Bearer "+conf.GetConfig().AlertManagerConfig.BearerToken)
	err := checkBearerToken(req, conf)
	assert.NoError(t, err)
}

func TestBearerTokenQueryParam(t *testing.T) {
	conf := config.NewMockConfiguration(1)
	req, _ := http.NewRequest(http.MethodPost, "https://example.com/webhook?token="+conf.GetConfig().AlertManagerConfig.BearerToken, nil)
	err := checkBearerToken(req, conf)
	assert.NoError(t, err)
}

func TestBearerTokenMissing(t *testing.T) {
	conf := config.NewMockConfiguration(1)
	req, _ := http.NewRequest(http.MethodPost, "https://example.com/webhook", nil)
	err := checkBearerToken(req, conf)
	assert.Error(t, err)
}

// TestWebhookBodyTooLarge is a regression test ensuring an oversized
// request body is rejected instead of being read in full, which could
// otherwise let a slow or huge request tie up a handler goroutine or
// exhaust memory.
func TestWebhookBodyTooLarge(t *testing.T) {
	conf := config.NewMockConfiguration(1)
	conf.SetIcingaClient(icinga2.NewMockClient())

	oversized := strings.NewReader(strings.Repeat("a", maxWebhookBodyBytes+1))
	req := httptest.NewRequest(http.MethodPost,
		"https://example.com/webhook?token="+conf.GetConfig().AlertManagerConfig.BearerToken, oversized)
	rec := httptest.NewRecorder()

	Webhook(rec, req, conf)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
