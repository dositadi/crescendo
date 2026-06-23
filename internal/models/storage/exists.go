package storage

import (
	"github.com/dositadi/groupie-tracker/internal/helper"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) Exists(folder string, target string) (bool, error) {
	objects, err := s.client.Storage.ListFiles(s.bucketId, folder, storage_go.FileSearchOptions{Limit: 1, Offset: 0})
	if err != nil {
		e := helper.WrapError("Objects fetch error", err)
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": "storage.GetFiles()",
		})
		return false, e
	}
	for _, object := range objects {
		if object.Name == target {
			s.logger.PrintInfo("File exists", map[string]string{
				"Source": "storage.Exists()",
				"Target": target,
			})
			return true, nil
		}
	}
	s.logger.PrintInfo("File does not exist", map[string]string{
		"Source": "storage.Exists()",
		"Target": target,
	})
	return false, nil
}
