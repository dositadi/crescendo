package authservice

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceST = "Render signup page f(n) under authservice"
)

func (a *AuthService) RenderSessionTimeOutPage() error {
	fs := []string{
		"internal/web/static/auth/session_expired.html",
	}

	temp := template.Must(template.New("session_expired.html").ParseFS(a.embedded.Get(), fs...))

	data := struct {
		LoginUrl string
	}{
		LoginUrl: utils.LOGIN.String(),
	}

	if err := temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Error executing template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceST,
		})
		return e
	}
	return nil
}
