package authservice

import (
	"fmt"
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceS = "authservice.RenderSettingsPage()"
)

func (a *AuthService) RenderSettingsPage() error {
	fs := []string{
		"internal/web/static/pages/settings_page.html",
	}

	user := a.getUser()

	data := struct {
		PrivacyPageUrl, TermsPageUrl, UploadUrl string

		AvatarKey string

		User data.User
	}{
		PrivacyPageUrl: fmt.Sprintf("%s?%s=%s", utils.PRIVACY.String(), utils.PAGE_KEY, a.request.URL.EscapedPath()),
		TermsPageUrl:   fmt.Sprintf("%s?%s=%s", utils.TERMS.String(), utils.PAGE_KEY, a.request.URL.EscapedPath()),
		UploadUrl:      utils.UPLOAD_PROFILE.String(),

		User: user,

		AvatarKey: utils.AVATAR_KEY,
	}

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
