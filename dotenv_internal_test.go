// MIT License
//
// Copyright (c) 2021 Iván Szkiba
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package dotenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const k6env = "K6_ENV"

func TestMain(m *testing.M) {
	os.Chdir("test/testdata") // nolint:errcheck
	os.Exit(m.Run())
}

func unset(t *testing.T) {
	t.Helper()

	assert.Nil(t, os.Unsetenv("global"))
	assert.Nil(t, os.Unsetenv("common"))
	assert.Nil(t, os.Unsetenv("local"))
	assert.Nil(t, os.Unsetenv("localonly"))
}

func Test_default(t *testing.T) { // nolint:paralleltest
	assert.Nil(t, os.Setenv(k6env, ""))
	unset(t)
	load()

	assert.Equal(t, ".env", os.Getenv("global"))
	assert.Equal(t, ".env.development", os.Getenv("common"))
	assert.Equal(t, ".env.development.local", os.Getenv("local"))
	assert.Equal(t, ".env.local", os.Getenv("localonly"))
}

func Test_disabled(t *testing.T) { // nolint:paralleltest
	assert.Nil(t, os.Setenv(k6env, "false"))
	unset(t)
	load()

	assert.Empty(t, os.Getenv("global"))
	assert.Empty(t, os.Getenv("common"))
	assert.Empty(t, os.Getenv("local"))
	assert.Empty(t, os.Getenv("localonly"))
}

func Test_test(t *testing.T) { // nolint:paralleltest
	os.Setenv(k6env, "test")
	unset(t)
	load()

	assert.Equal(t, ".env", os.Getenv("global"))
	assert.Equal(t, ".env.test", os.Getenv("common"))
	assert.Equal(t, ".env.test.local", os.Getenv("local"))
	assert.Empty(t, os.Getenv("localonly"))
}

func Test_production(t *testing.T) { // nolint:paralleltest
	assert.Nil(t, os.Setenv(k6env, "production"))
	unset(t)
	load()

	assert.Equal(t, ".env", os.Getenv("global"))
	assert.Equal(t, ".env.production", os.Getenv("common"))
	assert.Equal(t, ".env.production.local", os.Getenv("local"))
	assert.Equal(t, ".env.local", os.Getenv("localonly"))
}

func Test_development(t *testing.T) { // nolint:paralleltest
	assert.Nil(t, os.Setenv(k6env, "development"))
	unset(t)
	load()

	assert.Equal(t, ".env", os.Getenv("global"))
	assert.Equal(t, ".env.development", os.Getenv("common"))
	assert.Equal(t, ".env.development.local", os.Getenv("local"))
	assert.Equal(t, ".env.local", os.Getenv("localonly"))
}
