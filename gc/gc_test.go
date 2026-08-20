/*
 * License:
 * Copyright (c) 2019, VSHN AG, <info@vshn.ch>
 * Licensed under "BSD 3-Clause". See LICENSE file.
 */

package gc

import (
	"testing"

	"github.com/saremox/go-icinga2-client/icinga2"
	"github.com/saremox/signalilo/config"
	"github.com/stretchr/testify/assert"
)

// TestCollectServiceMissingKeepFor ensures that a service without a
// 'keep_for' var doesn't cause collectService to panic; regression test for
// a bug where an unchecked type assertion on svc.Vars["keep_for"] could
// crash the whole process from within the GC background goroutine.
func TestCollectServiceMissingKeepFor(t *testing.T) {
	c := config.NewMockConfiguration(1)
	svc := icinga2.Service{
		Name: "test.host!test-service",
		Vars: icinga2.Vars{},
	}

	assert.NotPanics(t, func() {
		err := collectService(svc, c, nil)
		assert.NoError(t, err)
	})
}

// TestCollectServiceInvalidKeepFor ensures that a 'keep_for' var of the
// wrong type doesn't cause collectService to panic.
func TestCollectServiceInvalidKeepFor(t *testing.T) {
	c := config.NewMockConfiguration(1)
	svc := icinga2.Service{
		Name: "test.host!test-service",
		Vars: icinga2.Vars{
			"keep_for": "not-a-number",
		},
	}

	assert.NotPanics(t, func() {
		err := collectService(svc, c, nil)
		assert.NoError(t, err)
	})
}
