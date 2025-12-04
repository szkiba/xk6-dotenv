// Package dotenv contains k6 dotenv extension.
package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	load()
}

func load() {
	env := os.Getenv(envVar)

	if env == disabled {
		return
	}

	if env == "" {
		env = "development"
	}

	_ = godotenv.Load(".env." + env + ".local")

	if env != "test" {
		_ = godotenv.Load(".env.local")
	}

	_ = godotenv.Load(".env." + env)

	_ = godotenv.Load()
}

const (
	envVar   = "K6_ENV"
	disabled = "false"
)
