package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", indexHandler)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/home/dsearle/go/src/github.com/dalesearle/tdr/web/html/"))))
	log.Fatal(http.ListenAndServeTLS("localhost:8081", "/home/dsearle/go/src/github.com/dalesearle/tdr/web/cert.pem", "/home/dsearle/go/src/github.com/dalesearle/tdr/web/key.pem", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	serve("/home/dsearle/go/src/github.com/dalesearle/tdr/web/html/index.html", w)
}

func cssHandler(w http.ResponseWriter, req *http.Request) {
	serve("/home/dsearle/go/src/github.com/dalesearle/tdr/web/html/resources/siege.css", w)
}

func serve(pattern string, w http.ResponseWriter) {

	fmt.Println(pattern)
	content, err := ioutil.ReadFile(pattern)
	if err != nil {
		fmt.Println("Unable to locate pattern ", pattern)
	} else {
		w.Write(content)
	}

}
