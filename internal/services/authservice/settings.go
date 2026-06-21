package authservice

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

const (
	sourceS = "authservice.RenderSettingsPage()"
)

func (a *AuthService) RenderSettingsPage() error {
	fs := []string{
		"internal/web/static/pages/settings_page.html",
	}

	data := struct{}{}

	temp, err := template.New("settings_page.html").ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceS,
		})
		return e
	}

	if err := temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceS,
		})
		return e
	}
	return nil
}
