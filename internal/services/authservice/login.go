package authservice

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

const (
	sourceR = "Render login page f(n) under authservice pkg"
)

func (a *AuthService) RenderLoginPage() error {
	fs := []string{
		"internal/web/static/auth/login.html",
	}

	temp, err := template.New("login").ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceR,
		})
		return e
	}

	if err = temp.Execute(a.responseWriter, nil); err != nil {
		e := helper.WrapError("Error executing template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceR,
		})
		return e
	}
	return nil
}
