package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func Execute() {
	router := New()

	fmt.Println("Listening of 8080 ....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
