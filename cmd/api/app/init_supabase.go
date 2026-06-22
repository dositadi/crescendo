package app

import (
	"github.com/dositadi/groupie-tracker/internal/utils"
	storage_go "github.com/supabase-community/storage-go"
)

func (a *App) initSupabase(supabaseUrl, secretKey string) *storage_go.Client {
	client := storage_go.NewClient(supabaseUrl, secretKey, nil)

	client.CreateBucket(utils.USER_PROFILE_BUCKET_ID, storage_go.BucketOptions{})

	return client
}
