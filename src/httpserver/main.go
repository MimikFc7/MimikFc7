// gosample2 project main.go
package main

import (
	"httpserver/handlers/userinfo"
	"httpserver/handlers/weather"
	"httpserver/httpsrv"
)

const (
	LISTEN_PORT = 8080
)

var jobs chan int

func main() {
	h1 := weather.GetEP()
	h2 := userinfo.GetEP()
	var hadlers = []httpsrv.EPHandler{h1, h2}
	httpServer := httpsrv.HTTPServer{
		Name:     "MainWeb",
		Port:     LISTEN_PORT,
		Handlers: hadlers,
	}
	httpServer.StartServer()
}
