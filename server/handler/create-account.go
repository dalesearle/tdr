package handler

import (
	"net/http"

	"fmt"

	"database/sql"

	"github.com/dalesearle/tdr/server/api"
	_ "github.com/go-sql-driver/mysql"
)

type acctHandler struct{}

func NewAccountHandler() *acctHandler {
	return new(acctHandler)
}

func (handler *acctHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

func (handler *acctHandler) handlePost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	if params, decoded := api.NewAccountBuilder().Decode(req.Body); decoded {
		createAccount(resp, &params)
	} else {
		api.PostStatus(resp, http.StatusInternalServerError)
	}
}

func createAccount(resp http.ResponseWriter, params *api.Account) {
	if emailIsUnique(resp, params.Email()) && passwordEncrypts(resp, params) && accountInserts(resp, params) {
		api.PostStatus(resp, http.StatusCreated)
	}
}

func emailIsUnique(resp http.ResponseWriter, email string) bool {
	_, err := api.NewAccountBuilder().FetchAccount(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}
	}
	api.PostStatus(resp, http.StatusUnauthorized)
	return false
}

func passwordEncrypts(resp http.ResponseWriter, acct *api.Account) bool {
	err := acct.EncryptPassword()
	if err != nil {
		api.PostStatus(resp, http.StatusInternalServerError)
	}
	return err == nil
}

func accountInserts(resp http.ResponseWriter, acct *api.Account) bool {
	err := acct.Insert()
	if err != nil {
		api.PostStatus(resp, http.StatusInternalServerError)
		return false
	}
	return true
}
