package storage

import (
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) GetPublicUrl(relativeFilePath string) string {
	resp := s.client.Storage.GetPublicUrl(s.bucketId, relativeFilePath, storage_go.UrlOptions{})

	return resp.SignedURL
}
