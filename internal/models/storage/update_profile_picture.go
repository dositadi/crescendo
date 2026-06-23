package storage

import (
	"io"

	"github.com/dositadi/groupie-tracker/internal/helper"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) UpdateProfilePicture(relativeFilePath string, file io.Reader) error {
	contentType := "image/png"
	_, err := s.client.Storage.UpdateFile(s.bucketId, relativeFilePath, file, storage_go.FileOptions{
		ContentType: &contentType,
	})
	if err != nil {
		e := helper.WrapError("Profile picture upload error", err)
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": "storage.UpdateProfilePicture()",
		})
		return e
	}
	return nil
}
