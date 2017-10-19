package api

import (
	"fmt"
	"net/http"

	"os"
)

func Errored(resp http.ResponseWriter, err error) bool {
	if err != nil {
		fmt.Println(err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func RedirectToAppError(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/app_error", http.StatusSeeOther)
}

func IsTesting() bool {
	for _, arg := range os.Args {
		if arg == "testing" {
			return true
		}
	}
	return false
}

func PostStatus(resp http.ResponseWriter, status int) {
	http.Error(resp, http.StatusText(status), status)
}
