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

	if currentPass != "" {
		userPass, err := a.usermodel.GetWithEmail(user.Email)
		if err != nil {
			e := helper.WrapError("User fetch error", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
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

		if err = a.compareHashedPassword(userPass.HashedPassword, []byte(currentPass)); err != nil {
			e := helper.WrapError("The current password is not correct", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})

			if err = page.RenderInfo(authservice.InfoForm, authservice.Info{
				Title:   "Password mismatch.",
				Message: "The current password you provided is not correct.",
			}, true); err != nil {
				a.logger.PrintError(err.Error(), map[string]string{
					"Source": sourcePr,
				})
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		if newPass == "" {
			e := helper.WrapError("The new password field cannot be empty", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			if err = page.RenderInfo(authservice.InfoForm, authservice.Info{
				Title:   "Empty field.",
				Message: "The new password field cannot be empty.",
			}, true); err != nil {
				a.logger.PrintError(err.Error(), map[string]string{
					"Source": sourcePr,
				})
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		if confirmPass == "" {
			e := helper.WrapError("The confirm password field cannot be empty", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			if err = page.RenderInfo(authservice.InfoForm, authservice.Info{
				Title:   "Empty field.",
				Message: "The confirm password field cannot be empty.",
			}, true); err != nil {
				a.logger.PrintError(err.Error(), map[string]string{
					"Source": sourcePr,
				})
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		if newPass != confirmPass {
			e := helper.WrapError("Password mismatch. New password and confirm password do not match", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourcePr,
			})
			if err = page.RenderInfo(authservice.InfoForm, authservice.Info{
				Title:   "Password mismatch.",
				Message: "New password and confirm password do not match.",
			}, true); err != nil {
				a.logger.PrintError(err.Error(), map[string]string{
					"Source": sourcePr,
				})
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
	} else {
		e := "Current password field cannot be empty"
		a.logger.PrintError(e, map[string]string{
			"Source": sourcePr,
		})
		if err := page.RenderInfo(authservice.InfoForm, authservice.Info{
			Title:   "Empty field.",
			Message: "Current password field cannot be empty.",
		}, true); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourcePr,
			})
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	var updateUser data.UpdateUser

	if username != "" {
		updateUser.Username = &username
	}

	hashedPassword, err := a.hashPassword([]byte(newPass))
	if err != nil {
		e := helper.WrapError("Password hash error.", err)
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
	updateUser.HashedPassword = hashedPassword

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
		Message: "Your changes have been saved.",
	}, false); err != nil {
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourcePr,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
