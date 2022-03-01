package cmd

import (
	"net/http"
)

func New() http.Server {
	return http.Server{
		Addr: ":8080",
	}
}
