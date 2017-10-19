package handler

import (
	"net/http"
)

func handlePreFlight(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	resp.Header().Set("Access-Control-Allow-Methods", "POST")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	resp.Header().Set("Access-Control-Max-Age", "86400")
}
