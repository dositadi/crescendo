package storage

import (
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

func (s *Storage) DeleteProfilePicture(relativeFilePath ...string) error {
	_, err := s.client.RemoveFile(utils.USER_PROFILE_BUCKET_ID, relativeFilePath)
	if err != nil {
		e := helper.WrapError("Profile picture upload error", err)
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": "storage.DeleteProfilePicture()",
		})
		return e
	}
	return nil
}
