package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"httpserver/httpsrv"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var charslist = []string{}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

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

	if charslist != nil && len(charslist) > 0 {
		self.Feeling = charslist[rand.Intn(len(charslist)-1)+0]
	} else {
		self.Feeling = String(10)
	}

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

	charslist = readLines("handlers/word_rus.txt")
	fmt.Println(charslist[len(charslist)-1])
	fmt.Println(len(charslist))

	h1 := httpsrv.EPHandler{
		URL:        "/weather",
		HandleFunc: handleRequest,
	}
	return h1
}
