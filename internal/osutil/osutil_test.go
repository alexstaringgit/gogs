// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package osutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFile(t *testing.T) {
	tests := []struct {
		path   string
		expVal bool
	}{
		{
			path:   "osutil.go",
			expVal: true,
		}, {
			path:   "../osutil",
			expVal: false,
		}, {
			path:   "not_found",
			expVal: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, IsFile(test.path))
		})
	}
}

func TestIsDir(t *testing.T) {
	tests := []struct {
		path   string
		expVal bool
	}{
		{
			path:   "osutil.go",
			expVal: false,
		}, {
			path:   "../osutil",
			expVal: true,
		}, {
			path:   "not_found",
			expVal: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, IsDir(test.path))
		})
	}
}

func TestIsExist(t *testing.T) {
	tests := []struct {
		path   string
		expVal bool
	}{
		{
			path:   "osutil.go",
			expVal: true,
		}, {
			path:   "../osutil",
			expVal: true,
		}, {
			path:   "not_found",
			expVal: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, IsExist(test.path))
		})
	}
}

func TestCurrentUsername(t *testing.T) {
	// Make sure it does not blow up
	CurrentUsername()
}

func TestCurrentUsernamePrefersEnvironmentVariable(t *testing.T) {
	// Some users/scripts expect that they can change the current username via environment variables
	if userBak, ok := os.LookupEnv("USER"); ok {
		defer os.Setenv("USER", userBak)
	} else {
		defer os.Unsetenv("USER")
	}

	if err := os.Setenv("USER", "__TESTING::USERNAME"); err != nil {
		t.Skip("Could not set the USER environment variable:", err)
	}
	assert.Equal(t, "__TESTING::USERNAME", CurrentUsername())
}
