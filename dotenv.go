// Package dotenv contains k6 dotenv extension.
package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	load()
}

//nolint:errcheck,gosec
func load() {
	env := os.Getenv(envVar) //nolint:forbidigo

	if env == disabled {
		return
	}

	if env == "" {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")

	if env != "test" {
		godotenv.Load(".env.local")
	}

	godotenv.Load(".env." + env)

	godotenv.Load()
}

const (
	envVar   = "K6_ENV"
	disabled = "false"
)
