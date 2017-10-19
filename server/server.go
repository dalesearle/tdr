package main

import (
	"log"
	"net/http"

	"github.com/dalesearle/tdr/server/handler"
)

const Base_URL = "localhost:8081"
const Cert string = "/home/dsearle/go/src/github.com/dalesearle/tdr/server/localhost.cert"
const Key =	"/home/dsearle/go/src/github.com/dalesearle/tdr/server/localhost.key"

func main() {
	http.Handle("/login", handler.NewLoginHandler())
	http.Handle("/create-account", handler.NewAccountHandler())
	http.Handle("/", handler.NewRootHandler())
	log.Fatal(http.ListenAndServeTLS(Base_URL, Cert, Key,nil))
}
