package storage

import (
	"fmt"
	"io"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) UpdateProfilePicture(relativeFilePath string, file io.Reader) error {
	resp, err := s.client.UpdateFile(utils.USER_PROFILE_BUCKET_ID, relativeFilePath, file, storage_go.FileOptions{})
	if err != nil {
		e := helper.WrapError("Profile picture upload error", err)
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": "storage.UpdateProfilePicture()",
		})
		return e
	}
	fmt.Println(resp.Data)
	return nil
}
