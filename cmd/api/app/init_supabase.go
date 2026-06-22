package app

import (
	"os"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/supabase-community/supabase-go"
)

func (a *App) initSupabase() *supabase.Client {
	client, err := supabase.NewClient(a.config.supabaseUrl, a.config.supabaseSecretKet, &supabase.ClientOptions{})
	if err != nil {
		e := helper.WrapError("Unable to connect to supabase", err)
		a.logger.PrintFatal(e.Error(), map[string]string{
			"Source": "app.initSupabase()",
		})
		os.Exit(1)
	}

	a.logger.PrintInfo("Supabase connected successfully", map[string]string{
		"Source": "app.initSupabase()",
	})
	return client
}
