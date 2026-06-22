package authpost

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceUp = "authpost.UploadProfilePicture()"
)

func (a *Auth) UploadProfilePicture(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20<<2+1<<10)
	// 1 << 20 = 1mb | 1mb << 2 = 1mb * 2^n = 1mb * 2^2 = 1mb * 4 = 4mb. 1 << 10 = 1 * 2^n = 1 * 2^10 = 1024

	if err := r.ParseMultipartForm(1 << 20 << 2); err != nil {
		e := helper.WrapError("File size exceeds limit (4mb)", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	file, _, err := r.FormFile(utils.AVATAR_KEY)
	if err != nil {
		e := helper.WrapError("Error getting file", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUp,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	defer file.Close()
}
