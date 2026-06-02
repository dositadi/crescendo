package authservice

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

const (
	sourceRS = "Render signup page f(n) under authservice"
)

func (a *AuthService) RenderSignupPage() error {
	fs := []string{
		"internal/web/static/auth/signup.html",
	}

	temp, err := template.New("signup.html").ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRS,
		})
		return e
	}

	if err = temp.Execute(a.responseWriter, nil); err != nil {
		e := helper.WrapError("Error executing template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRS,
		})
		return e
	}
	return nil
}
