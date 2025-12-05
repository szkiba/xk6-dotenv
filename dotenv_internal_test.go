package dotenv

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const k6env = "K6_ENV"

func TestMain(m *testing.M) {
	if err := os.Chdir("testdata"); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func unset(t *testing.T) {
	t.Helper()

	require.NoError(t, os.Unsetenv("global"))
	require.NoError(t, os.Unsetenv("common"))
	require.NoError(t, os.Unsetenv("local"))
	require.NoError(t, os.Unsetenv("localonly"))
}

func Test_default(t *testing.T) {
	t.Setenv(k6env, "")
	unset(t)
	load()

	require.Equal(t, ".env", os.Getenv("global"))
	require.Equal(t, ".env.development", os.Getenv("common"))
	require.Equal(t, ".env.development.local", os.Getenv("local"))
	require.Equal(t, ".env.local", os.Getenv("localonly"))
}

func Test_disabled(t *testing.T) {
	t.Setenv(k6env, "false")
	unset(t)
	load()

	require.Empty(t, os.Getenv("global"))
	require.Empty(t, os.Getenv("common"))
	require.Empty(t, os.Getenv("local"))
	require.Empty(t, os.Getenv("localonly"))
}

func Test_test(t *testing.T) {
	t.Setenv(k6env, "test")
	unset(t)
	load()

	require.Equal(t, ".env", os.Getenv("global"))
	require.Equal(t, ".env.test", os.Getenv("common"))
	require.Equal(t, ".env.test.local", os.Getenv("local"))
	require.Empty(t, os.Getenv("localonly"))
}

func Test_production(t *testing.T) {
	t.Setenv(k6env, "production")
	unset(t)
	load()

	require.Equal(t, ".env", os.Getenv("global"))
	require.Equal(t, ".env.production", os.Getenv("common"))
	require.Equal(t, ".env.production.local", os.Getenv("local"))
	require.Equal(t, ".env.local", os.Getenv("localonly"))
}

func Test_development(t *testing.T) {
	t.Setenv(k6env, "development")
	unset(t)
	load()

	require.Equal(t, ".env", os.Getenv("global"))
	require.Equal(t, ".env.development", os.Getenv("common"))
	require.Equal(t, ".env.development.local", os.Getenv("local"))
	require.Equal(t, ".env.local", os.Getenv("localonly"))
}
