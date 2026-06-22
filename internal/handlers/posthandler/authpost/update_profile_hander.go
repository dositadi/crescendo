package authpost

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
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

	if currentPass != "" {
		userPass, err := a.usermodel.GetWithEmail(user.Email)
		if err != nil {
			e := helper.WrapError("User fetch error", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			return
		}

		if err = a.compareHashedPassword(userPass.HashedPassword, []byte(currentPass)); err != nil {
			e := helper.WrapError("The current password is not correct", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			return
		}

		if newPass == "" {
			e := helper.WrapError("The new password field cannot be empty", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			return
		}

		if confirmPass == "" {
			e := helper.WrapError("The confirm password field cannot be empty", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			return
		}

		if newPass != confirmPass {
			e := helper.WrapError("Password mismatch. New password and confirm password do not match", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			return
		}
	} else {
		e := "Current password field cannot be empty"
		a.logger.PrintError(e, map[string]string{
			"Source": sourcePr,
		})
		return
	}

	var updateUser data.UpdateUser

	if username != "" {
		updateUser.Username = &username
	}

	updateUser.HashedPassword = []byte(newPass)

	if err := a.usermodel.Update(user.Id, updateUser); err != nil {
		e := helper.WrapError("Error uploading file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
		// Send an error response here
		return
	}
	a.logger.PrintInfo("User Profile updated successfully", map[string]string{
		"Source": sourcePr,
		"User":   user.Username,
	})
}
