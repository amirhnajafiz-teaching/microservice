package cmd

import (
	"log"
	"net/http"
)

func Execute() {
	router := New()

	log.Fatal(http.ListenAndServe(":8080", router))
}
