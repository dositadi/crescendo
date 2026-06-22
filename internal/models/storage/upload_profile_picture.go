package storage

import (
	"fmt"
	"io"

	"github.com/dositadi/groupie-tracker/internal/helper"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) UploadProfilePicture(fileRelativePath string, file io.Reader) error {
	contentType := "image/png"

	_, err := s.client.Storage.UploadFile(s.bucketId, fileRelativePath, file, storage_go.FileOptions{
		ContentType: &contentType,
	})

	if err != nil {
		e := helper.WrapError("Profile picture upload error", err)
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": "storage.UploadProfilePicture()",
		})
		return e
	}
	s.logger.PrintInfo(fmt.Sprintf("%s uploaded successfully", fileRelativePath), map[string]string{
		"Source": "storage.UploadProfilePicture()",
	})
	return nil
}
