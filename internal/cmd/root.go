package cmd

import (
	"github.com/amirhnajafiz/go-kit/internal/http/middleware"
	"log"
	"net/http"

	"github.com/amirhnajafiz/go-kit/internal/endpoints"
	"github.com/amirhnajafiz/go-kit/internal/http/handler"
	httpTransport "github.com/go-kit/kit/transport/http"
)

func Execute() {
	svc := handler.Handler{}

	uppercaseHandler := httpTransport.NewServer(
		endpoints.MakeUppercaseEndpoint(svc),
		middleware.DecodeUppercaseRequest,
		middleware.EncodeResponse,
	)

	lowercaseHandler := httpTransport.NewServer(
		endpoints.MakeLowercaseEndpoint(svc),
		middleware.DecodeLowercaseRequest,
		middleware.EncodeResponse,
	)

	concatHandler := httpTransport.NewServer(
		endpoints.MakeConcatenateEndpoint(svc),
		middleware.DecodeConcatenateRequest,
		middleware.EncodeResponse,
	)

	splitHandler := httpTransport.NewServer(
		endpoints.MakeSplitEndpoint(svc),
		middleware.DecodeSplitRequest,
		middleware.EncodeResponse,
	)

	countHandler := httpTransport.NewServer(
		endpoints.MakeCountEndpoint(svc),
		middleware.DecodeCountRequest,
		middleware.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/lowercase", lowercaseHandler)
	http.Handle("/concatenate", concatHandler)
	http.Handle("/split", splitHandler)
	http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
