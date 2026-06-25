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

	path := a.request.FormValue(utils.PATH_KEY)
	if path == "" {
		path = utils.HOME.String()
	}

	user, err := a.userModel.GetWithEmail(a.getUser().Email)
	if err != nil {
		e := helper.WrapError("User fetch error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceS,
		})
		return e
	}

	data := struct {
		PrivacyPageUrl, TermsPageUrl, UploadUrl, PrevPageUrl, ProfileUpdateUrl, LogoutUrl, DeleteAccountUrl string

		AvatarKey, UsernameKey, CurrentPassKey, NewPassKey, ConfirmPassKey, AvatarRespId, FormRespId string

		User data.User
	}{
		PrivacyPageUrl:   fmt.Sprintf("%s?%s=%s", utils.PRIVACY.String(), utils.PAGE_KEY, a.request.URL.EscapedPath()),
		TermsPageUrl:     fmt.Sprintf("%s?%s=%s", utils.TERMS.String(), utils.PAGE_KEY, a.request.URL.EscapedPath()),
		UploadUrl:        utils.UPLOAD_PROFILE.String(),
		PrevPageUrl:      path,
		ProfileUpdateUrl: utils.UPDATE_USER_INFO.String(),
		DeleteAccountUrl: utils.DELETE_ACCOUNT.String(),
		LogoutUrl:        utils.LOGOUT.String(),

		User: user,

		AvatarKey:      utils.AVATAR_KEY,
		UsernameKey:    utils.USERNAME_KEY,
		CurrentPassKey: utils.PASSWORD_KEY,
		NewPassKey:     utils.NEW_PASSWORD_KEY,
		ConfirmPassKey: utils.CONFIRM_PASS_KEY,

		AvatarRespId: string(InfoAvatar),
		FormRespId:   string(InfoForm),
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
