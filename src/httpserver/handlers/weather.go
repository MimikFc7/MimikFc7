package handlers

import (
	"encoding/json"
	"fmt"
	"httpserver/httpsrv"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Weather struct {
	Id      int    `json:"id"`
	Feeling string `json:"feeling"`
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func (self *Weather) RandId() {
	self.Id = rand.Int()
}

func (self *Weather) RandFeeling() {
	self.Feeling = String(10)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	wE := Weather{}
	wE.RandFeeling()
	wE.RandId()
	b, err := json.Marshal(wE)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, string(b))
}

func GetWeatherEP() httpsrv.EPHandler {
	h1 := httpsrv.EPHandler{
		URL:        "/weather",
		HandleFunc: handleRequest,
	}
	return h1
}
