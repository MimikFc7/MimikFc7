package userinfo

import (
	"encoding/json"
	"fmt"
	"httpserver/httpsrv"
	"httpserver/models/userinfo"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	wE := userinfo.UserInfo{}
	wE.SetValues()
	b, err := json.Marshal(wE)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, string(b))
}

func GetEP() httpsrv.EPHandler {

	h1 := httpsrv.EPHandler{
		URL:        "/userinfo",
		HandleFunc: handleRequest,
	}
	return h1
}
