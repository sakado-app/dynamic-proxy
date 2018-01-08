package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const Port = 15774
const TokenLength = 64

var token string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var ip string

func main() {
	token = genToken(TokenLength)

	http.HandleFunc("/", handle)

	fmt.Printf("Token : %s\n", token)
	fmt.Printf("Listening on %d\n", Port)
	http.ListenAndServe(":"+strconv.Itoa(Port), nil)
}

func genToken(length int) string {
	rand.Seed(time.Now().UnixNano())
	
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func handle(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[1:]

	if path == "set" {
		if token != request.URL.Query().Get("token") {
			fmt.Fprintln(writer, "Wrong token")
			return
		}

		remote := request.RemoteAddr
		index := strings.Index(remote, ":")

		ip = remote[:index]

		fmt.Printf("I.P. set to '%s'\n", ip)
		fmt.Fprintln(writer, "OK")
	} else if path == "get" {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(writer, ip)
	} else {
		fmt.Fprintln(writer, "Unknown request")
	}
}
