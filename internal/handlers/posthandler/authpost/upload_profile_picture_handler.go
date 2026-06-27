package authpost

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/authservice"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceUp = "authpost.UploadProfilePicture()"
)

func (a *Auth) UploadProfilePicture(w http.ResponseWriter, r *http.Request) {
	user := a.getUser(r)
	fmt.Println(user.Email)
	r.Body = http.MaxBytesReader(w, r.Body, (4<<20)+(1<<10))
	// 1 << 20 = 1mb | 1mb << 2 = 1mb * 2^n = 1mb * 2^2 = 1mb * 4 = 4mb.(4 << 20 = 4 * 2^20 = 4mb) 1 << 10 = 1 * 2^n = 1 * 2^10 = 1024

	page := authservice.New(w, a.embedded, a.logger, r, a.usermodel)

	if err := r.ParseMultipartForm(4 << 20); err != nil {
		e := helper.WrapError("File size exceeds limit (4mb)", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})

		if err = page.RenderInfo(authservice.InfoAvatar, authservice.Info{
			Title:   "File too large",
			Message: "The picture you choose is above 4mb. Choose a picture of lesser size and try again.",
		}, true); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceUp,
			})
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return


		
	}

	file, _, err := r.FormFile(utils.AVATAR_KEY)
	if err != nil {
		e := helper.WrapError("Error getting file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		if err = page.RenderInfo(authservice.InfoAvatar, authservice.Info{
			Title:   "Something wrong happened.",
			Message: "Kindly try again.",
		}, true); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceUp,
			})
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	defer file.Close()

	fileName := genFilePath(user.Id)

	exist, err := a.storage.Exists(user.Id, strings.Split(fileName, "/")[1])
	if err != nil {
		e := helper.WrapError("Error checking if file exists", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		if err = page.RenderInfo(authservice.InfoAvatar, authservice.Info{
			Title:   "Server Error",
			Message: "Something wrong happened. Kindly check your network connection and try again.",
		}, true); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceUp,
			})
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// If avatar exists already in the store update instead
	if exist {
		if err = a.storage.UpdateProfilePicture(fileName, file); err != nil {
			e := helper.WrapError("Error updating file", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUp,
			})
			if err = page.RenderInfo(authservice.InfoAvatar, authservice.Info{
				Title:   "Server Error",
				Message: "Something wrong happened. Kindly check your network connection and try again.",
			}, true); err != nil {
				a.logger.PrintError(err.Error(), map[string]string{
					"Source": sourceUp,
				})
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	} else {
		if err = a.storage.UploadProfilePicture(fileName, file); err != nil {
			e := helper.WrapError("Error uploading file", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUp,
			})
			if err = page.RenderInfo(authservice.InfoAvatar, authservice.Info{
				Title:   "Server Error",
				Message: "Something wrong happened. Kindly check your network connection and try again.",
			}, true); err != nil {
				a.logger.PrintError(err.Error(), map[string]string{
					"Source": sourceUp,
				})
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	avatarUrl := fmt.Sprintf("%s?t=%v", a.storage.GetPublicUrl(fileName), time.Now().Unix())

	userUpdate := data.UpdateUser{
		AvatarUrl: &avatarUrl,
	}

	if err := a.usermodel.Update(user.Id, userUpdate); err != nil {
		e := helper.WrapError("Error uploading file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		if err = page.RenderInfo(authservice.InfoAvatar, authservice.Info{
			Title:   "Server Error",
			Message: "Something wrong happened. Kindly check your network connection and try again.",
		}, true); err != nil {
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceUp,
			})
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	a.logger.PrintInfo("User uploaded avatar", map[string]string{
		"Source":   sourceUp,
		"User":     user.Username,
		"Filename": fileName,
	})
	if err = page.RenderInfo(authservice.InfoAvatar, authservice.Info{
		Title:   "Success",
		Message: "Your changes have been saved.",
	}, false); err != nil {
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceUp,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func genFilePath(userId string) string {
	return fmt.Sprintf("%s/profile-avatar.png", userId)
}
