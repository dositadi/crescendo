package herokuapp

import (
	opencage "github.com/dositadi/groupie-tracker/internal/client/open_cage"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type HerokuApp struct {
	logger   jsonlog.Logger
	opencage opencage.OpenCage
}

func New(logger jsonlog.Logger, opencage opencage.OpenCage) *HerokuApp {
	return &HerokuApp{
		logger:   logger,
		opencage: opencage,
	}
}

func (a *HerokuApp) Init() {
	a.mapArtistsInfo()
}

func (a *HerokuApp) Get() map[int]ArtistInfo {
	return byId
}
