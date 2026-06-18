package legalpages

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type LegalPage struct {
	logger   jsonlog.Logger
	embedded groupietracker.Embedded
}

func New(logger jsonlog.Logger, embedded groupietracker.Embedded) *LegalPage {
	return &LegalPage{
		logger:   logger,
		embedded: embedded,
	}
}
