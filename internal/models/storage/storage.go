package storage

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	storage_go "github.com/supabase-community/storage-go"
)


type Storage struct {
	logger jsonlog.Logger
	client *storage_go.Client
}

func New(logger jsonlog.Logger, client *storage_go.Client) *Storage {
	return &Storage{
		logger: logger,
		client: client,
	}
}