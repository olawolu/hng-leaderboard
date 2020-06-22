package main

import (
	"net/http"
	"os"

	"github.com/olawolu/hng-leaderboard/uploader"
)

var CONN_PORT = os.Getenv("PORT")
var port string

func init() {
	port = CONN_PORT
	if port == "" {
		port = ":8080"
	}
}

func main() {
	fs := http.FileServer(http.Dir("./frontend/"))
	http.Handle("/", fs)
	http.HandleFunc("/upload", uploader.UploadFile)
	http.ListenAndServe(port, nil)
}
