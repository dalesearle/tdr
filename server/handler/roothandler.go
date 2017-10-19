package handler

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/dalesearle/tdr/server/api"
)

const Path string = "/home/dsearle/go/src/github.com/dalesearle/tdr/angular/dist/"
type rootHandler struct {
}

func NewRootHandler() *rootHandler {
	return new(rootHandler)
}

func (handler *rootHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if strings.HasSuffix(req.URL.String(), "/") {
		postIndex(&resp)
	} else {
		http.StripPrefix("/", http.FileServer(http.Dir(Path))).ServeHTTP(resp, req)
	}
}

func postIndex(resp *http.ResponseWriter) {
	//tmpl, err := template.ParseFiles(api.Static_Path + "login.html")
	tmpl, err := template.ParseFiles(Path+ "index.html")
	if !api.Errored(*resp, err) {
		tmpl.Execute(*resp, "")
	}
}