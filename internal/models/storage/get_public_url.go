package storage

import (
	"github.com/dositadi/groupie-tracker/internal/utils"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) GetPublicUrl(relativeFilePath string) string {
	resp := s.client.GetPublicUrl(utils.USER_PROFILE_BUCKET_ID, relativeFilePath, storage_go.UrlOptions{})

	return resp.SignedURL
}
