package authservice

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

type Error string

const (
	NAME_ERROR     Error = "Name Error"
	EMAIL_ERROR    Error = "Email error"
	PASSWORD_ERROR Error = "Password error"
	sourceEr             = "Render auth error f(n) under authservice"
)

func (a *AuthService) RenderAuthError(e Error, message string) error {
	fs := []string{
		"internal/web/static/partials/auth/errors.html",
	}

	temp, err := template.New("errors.html").ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceEr,
		})
		return e
	}

	data := struct {
		Error string
	}{
		Error: message,
	}

	switch e {
	case NAME_ERROR:
		err = temp.ExecuteTemplate(a.responseWriter, "username-error", data)
	case PASSWORD_ERROR:
		err = temp.ExecuteTemplate(a.responseWriter, "password-error", data)
	case EMAIL_ERROR:
		err = temp.ExecuteTemplate(a.responseWriter, "email-error", data)
	}
	if err != nil {
		e := helper.WrapError("Error executing template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceEr,
		})
		return e
	}
	return nil
}
