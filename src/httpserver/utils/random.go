package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const randLen = 10

var charslist = []string{}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)

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

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandInt() int {
	return rand.Int()
}

func RandString() string {

	if len(charslist) < 1 {
		charslist = readLines("utils/word_rus.txt")
	}

	if len(charslist) > 1 {
		return charslist[rand.Intn(len(charslist)-1)+0]
	}

	return StringWithCharset(randLen, charset)
}
