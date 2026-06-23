package authservice

import (
	"html/template"
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

type Error string

const (
	NAME_ERROR     Error = "Name Error"
	EMAIL_ERROR    Error = "Email error"
	PASSWORD_ERROR Error = "Password error"
	TERMS_ERROR    Error = "Terms Error"
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
		Message string
	}{
		Message: message,
	}

	a.responseWriter.WriteHeader(http.StatusBadRequest)

	switch e {
	case NAME_ERROR:
		err = temp.ExecuteTemplate(a.responseWriter, "username-error", data)
	case PASSWORD_ERROR:
		err = temp.ExecuteTemplate(a.responseWriter, "password-error", data)
	case EMAIL_ERROR:
		err = temp.ExecuteTemplate(a.responseWriter, "email-error", data)
	case TERMS_ERROR:
		err = temp.ExecuteTemplate(a.responseWriter, "term-error", data)
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

type SettingInfo string

const (
	InfoAvatar = SettingInfo("info-avatar")
	InfoForm   = SettingInfo("info")

	sourceInfo = "authservice.RenderInfo()"
)

type Info struct {
	Title   string
	Message string
}

func (a *AuthService) RenderInfo(id SettingInfo, info Info, isError bool) error {
	fs := []string{
		"internal/web/static/partials/pages/settings_page.html",
	}

	data := struct {
		Id   string
		Info Info
	}{
		Id:   string(id),
		Info: info,
	}

	temp, err := template.New("settings_page.html").ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceInfo,
		})
		return e
	}

	if isError {
		err = temp.ExecuteTemplate(a.responseWriter, "error", data)
	} else {
		err = temp.ExecuteTemplate(a.responseWriter, "success", data)
	}
	if err != nil {
		e := helper.WrapError("Error executing template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceInfo,
		})
		return e
	}
	return nil
}
