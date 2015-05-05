package main

import (
	"log"
	"os"

	. "github.com/oleiade/lane"
)

var responseQueue *Queue

func main() {

	responseQueue = NewQueue()

	accessKey := os.Getenv("IVONA_ACCESS_KEY")
	secretKey := os.Getenv("IVONA_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		log.Println("main: environment variables not set")
		return
	}

	InitIvona(accessKey, secretKey)

	go RunResponseWorker()

	if err := StartHTTPServer("127.0.0.1", "9574"); err != nil {
		log.Println("Error starting server", err)
	}

}
