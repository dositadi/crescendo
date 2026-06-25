package authpost

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/authservice"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourcePr = "authpost.UpdateProfileHandler()"
)

func (a *Auth) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(utils.USERNAME_KEY)
	currentPass := r.FormValue(utils.PASSWORD_KEY)
	newPass := r.FormValue(utils.NEW_PASSWORD_KEY)
	confirmPass := r.FormValue(utils.CONFIRM_PASS_KEY)
	user := a.getUser(r)
	page := authservice.New(w, a.embedded, a.logger, r, a.usermodel)

	newHashPassword, info := a.validatePassword(currentPass, newPass, confirmPass, user.Email)
	if info != nil {
		if err := page.RenderInfo(authservice.InfoForm, *info, true); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourcePr,
			})
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var updateUser data.UpdateUser

	if username != "" && username != user.Username {
		updateUser.Username = &username
	}

	if newHashPassword != nil {
		updateUser.HashedPassword = newHashPassword
	}

	if err := a.usermodel.Update(user.Id, updateUser); err != nil {
		e := helper.WrapError("Error uploading file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})

		if err = page.RenderInfo(authservice.InfoForm, authservice.Info{
			Title:   "Something wrong happened.",
			Message: "Kindly try again.",
		}, true); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourcePr,
			})
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	a.logger.PrintInfo("User Profile updated successfully", map[string]string{
		"Source": sourcePr,
		"User":   user.Username,
	})

	if err := page.RenderInfo(authservice.InfoForm, authservice.Info{
		Title:   "Success.",
		Message: a.constructFeedbackMessage(updateUser),
	}, false); err != nil {
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourcePr,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *Auth) constructFeedbackMessage(data data.UpdateUser) string {
	switch {
	case data.HashedPassword != nil && data.Username != nil:
		return "Password and username changes have been saved."
	case data.HashedPassword != nil:
		return "Password changes have been saved."
	case data.Username != nil:
		return "Username changes have been saved."
	default:
		return "No new changes effected"
	}
}

func (a *Auth) validatePassword(currentPass, newPass, confirmPass, userEmail string) ([]byte, *authservice.Info) {
	if currentPass == "" {
		return nil, nil
	} else {
		userPass, err := a.usermodel.GetWithEmail(userEmail)
		if err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourcePr,
			})
			return nil, &authservice.Info{
				Title:   "Something wrong happened.",
				Message: "Kindly try again.",
			}
		}

		if err = a.compareHashedPassword(userPass.HashedPassword, []byte(currentPass)); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourcePr,
			})
			return nil, &authservice.Info{
				Title:   "Password mismatch.",
				Message: "The current password you provided is not correct.",
			}
		}

		if newPass == "" {
			return nil, &authservice.Info{
				Title:   "Empty field.",
				Message: "The new password field cannot be empty.",
			}
		}

		if len(newPass) < 8 {
			return nil, &authservice.Info{
				Title:   "Invalid Password.",
				Message: "The new password must be at least 8 characters long.",
			}
		}

		if confirmPass == "" {
			return nil, &authservice.Info{
				Title:   "Empty field.",
				Message: "The confirm password field cannot be empty.",
			}
		}

		if newPass != confirmPass {
			return nil, &authservice.Info{
				Title:   "Password mismatch.",
				Message: "New password and confirm password do not match.",
			}
		}

		newHashedPass, err := a.hashPassword([]byte(newPass))
		if err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourcePr,
			})
			return nil, &authservice.Info{
				Title:   "Something wrong happened.",
				Message: "Kindly try again.",
			}
		}

		if err = a.compareHashedPassword(newHashedPass, []byte(currentPass)); nil == err {
			return nil, &authservice.Info{
				Title:   "Not valid.",
				Message: "Your new password is the same as your old password.",
			}
		}
		return newHashedPass, nil
	}
}
