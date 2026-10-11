package main

import (
	"email-sender/pkg/server"
	"log"
	"net/http"
)

func main() {
	serverHanlder, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", serverHanlder); err != nil {
		log.Fatal(err)
	}

}
