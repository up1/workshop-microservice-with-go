package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP service with " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("Error: " + err.Error())
	}
}
