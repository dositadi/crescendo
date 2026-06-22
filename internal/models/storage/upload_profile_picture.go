package storage

import (
	"fmt"
	"io"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) UploadProfilePicture(fileRelativePath string, file io.Reader) error {
	resp, err := s.client.UploadFile(utils.USER_PROFILE_BUCKET_ID, fileRelativePath, file, storage_go.FileOptions{})
	if err != nil {
		e := helper.WrapError("Profile picture upload error", err)
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": "storage.UploadProfilePicture()",
		})
		return e
	}
	fmt.Println(resp.Data)
	return nil
}
