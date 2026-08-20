/*
 * License:
 * Copyright (c) 2019, VSHN AG, <info@vshn.ch>
 * Licensed under "BSD 3-Clause". See LICENSE file.
 */

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInitializeReturnsErrorOnUnreachableIcinga is a regression test for a
// bug where Initialize only logged an Icinga-client creation failure and
// kept going, leaving callers with a nil Icinga client and no way to fail
// fast. Callers that require a working Icinga connection (i.e. the serve
// command) rely on this error being returned so they can abort startup
// instead of panicking on a nil client later.
func TestInitializeReturnsErrorOnUnreachableIcinga(t *testing.T) {
	mockCfg := &MockConfiguration{
		config: SignaliloConfig{
			IcingaConfig: icingaConfig{
				URL: []string{"127.0.0.1:1"},
			},
		},
	}
	mockCfg.logger = MockLogger(1)

	err := Initialize(mockCfg)
	assert.Error(t, err)
	assert.Nil(t, mockCfg.GetIcingaClient())
}
