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

	if err := StartHTTPServer(os.Getenv("VCAP_APP_HOST"), os.Getenv("VCAP_APP_PORT")); err != nil {
		log.Println("Error starting server", err)
	}

}
