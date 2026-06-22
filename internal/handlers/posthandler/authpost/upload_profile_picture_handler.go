package authpost

import (
	"fmt"
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceUp = "authpost.UploadProfilePicture()"
)

func (a *Auth) UploadProfilePicture(w http.ResponseWriter, r *http.Request) {
	user := a.getUser(r)
	fmt.Println(user.Email)
	r.Body = http.MaxBytesReader(w, r.Body, (1<<20<<2)+(1<<10))
	// 1 << 20 = 1mb | 1mb << 2 = 1mb * 2^n = 1mb * 2^2 = 1mb * 4 = 4mb. 1 << 10 = 1 * 2^n = 1 * 2^10 = 1024

	if err := r.ParseMultipartForm(1 << 20 << 2); err != nil {
		e := helper.WrapError("File size exceeds limit (4mb)", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		// Send an error response here
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile(utils.AVATAR_KEY)
	if err != nil {
		e := helper.WrapError("Error getting file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
		// Send an error response here
		return
	}

	defer file.Close()

	fileName := genFilePath(user.Id)

	if err = a.storage.UploadProfilePicture(fileName, file); err != nil {
		e := helper.WrapError("Error uploading file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
		// Send an error response here
		return
	}

	avatarUrl := a.storage.GetPublicUrl(fileName)

	userUpdate := data.UpdateUser{
		AvatarUrl: &avatarUrl,
	}

	if err := a.usermodel.Update(user.Id, userUpdate); err != nil {
		e := helper.WrapError("Error uploading file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
		// Send an error response here
		return
	}

	a.logger.PrintInfo("User uploaded avatar", map[string]string{
		"Source":   sourceUp,
		"User":     user.Username,
		"Filename": fileName,
	})
}

func genFilePath(userId string) string {
	return fmt.Sprintf("%s/profile-avatar.png", userId)
}
