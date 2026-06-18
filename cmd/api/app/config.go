package app

import (
	"os"

	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type config struct {
	postgresDbDsn string
	jwtkey      string
	opencageKey string
}

func (c *config) Init(logger jsonlog.Logger) {
	temp := config{
		postgresDbDsn: os.Getenv("POSTGRES_DB_DSN"),
		jwtkey:        os.Getenv("JWTKEY"),
		opencageKey:   os.Getenv("OPEN_CAGE_KEY"),
	}

	if temp.jwtkey == "" {
		logger.PrintFatal("JWT Key is not set in the environment", map[string]string{
			"Context": "Config init f(n) under app pkg",
			"Hint":    "Set the JWT key in the environment",
		})
		os.Exit(1)
	}
	if temp.opencageKey == "" {
		logger.PrintFatal("Opencage Key is not set in the environment", map[string]string{
			"Context": "Config init f(n) under app pkg",
			"Hint":    "Set the Opencage key in the environment",
		})
		os.Exit(1)
	}
	if temp.postgresDbDsn == "" {
		logger.PrintFatal("Postgres db dsn is not set in the environment", map[string]string{
			"Context": "Config init f(n) under app pkg",
			"Hint":    "Set the Postgres db dsn in the environment",
		})
		os.Exit(1)
	}

	*c = temp
}
