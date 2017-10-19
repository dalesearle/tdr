package handler

import (
	"net/http"

	"fmt"

	"database/sql"

	"github.com/dalesearle/tdr/server/api"
	"golang.org/x/crypto/bcrypt"
)

type loginHandler struct{}

func NewLoginHandler() *loginHandler {
	return new(loginHandler)
}

func (handler *loginHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "OPTIONS":
		handlePreFlight(resp, req)
	case "POST":
		handler.handlePost(resp, req)
	default:
		fmt.Println("Unsupported http method.")
		api.PostStatus(resp, http.StatusBadRequest)
	}
}

func (handler *loginHandler) handlePost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	if params, decoded := api.NewAccountBuilder().Decode(req.Body); decoded {
		doLogin(resp, &params)
	} else {
		api.PostStatus(resp, http.StatusInternalServerError)
	}
}

func doLogin(resp http.ResponseWriter, params *api.Account) {
	if storedAcct, fetched := fetchAccount(resp, params.Email()); fetched {
		if loginAuthenticates(resp, storedAcct.Password(), params.Password()) {
			api.PostStatus(resp, http.StatusOK)
		}
	}
}

func fetchAccount(resp http.ResponseWriter, email string) (api.Account, bool) {
	acct, err := api.NewAccountBuilder().FetchAccount(email)
	if err != nil {
		if err == sql.ErrNoRows {
			api.PostStatus(resp, http.StatusNotFound)
		} else {
			api.PostStatus(resp, http.StatusInternalServerError)
		}
	}
	return acct, err == nil
}

func loginAuthenticates(resp http.ResponseWriter, hashedPwd, givenPwd string) bool {
	fmt.Println("")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(givenPwd))
	if err != nil {
		api.PostStatus(resp, http.StatusUnauthorized)
		return false
	}
	return true
}
