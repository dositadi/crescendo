package storage

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/supabase-community/supabase-go"
)

type Storage struct {
	logger   jsonlog.Logger
	client   *supabase.Client
	bucketId string
}

func New(logger jsonlog.Logger, client *supabase.Client) *Storage {
	return &Storage{
		logger:   logger,
		client:   client,
		bucketId: utils.USER_PROFILE_BUCKET_ID,
	}
}
