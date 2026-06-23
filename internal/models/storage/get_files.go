package storage

import (
	"fmt"

	"github.com/dositadi/groupie-tracker/internal/helper"
	storage_go "github.com/supabase-community/storage-go"
)

func (s *Storage) GetFiles(filePath string) {
	objects, err := s.client.Storage.ListFiles(s.bucketId, filePath, storage_go.FileSearchOptions{Limit: 1})
	if err != nil {
		e := helper.WrapError("Profile picture upload error", err)
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": "storage.GetFiles()",
		})
		return
	}
	for _, object := range objects {
		fmt.Println(object.Name)
	}
}
