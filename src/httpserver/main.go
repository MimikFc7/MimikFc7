// gosample2 project main.go
package main

import (
	"httpserver/handlers"
	"httpserver/httpsrv"
)

const (
	LISTEN_PORT = 8080
)

var jobs chan int

func main() {

	h1 := handlers.GetWeatherEP()
	var hadlers = []httpsrv.EPHandler{h1}
	httpServer := httpsrv.HTTPServer{
		Name:     "MainWeb",
		Port:     LISTEN_PORT,
		Handlers: hadlers,
	}
	httpServer.StartServer()
}
