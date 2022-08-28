package cmd

import (
	"net/http"

	"github.com/amirhnajafiz/micro-go/internal/endpoints"
	"github.com/amirhnajafiz/micro-go/internal/http/handler"
	"github.com/amirhnajafiz/micro-go/internal/http/middleware"
	httpTransport "github.com/go-kit/kit/transport/http"
)

func New() *http.ServeMux {
	router := http.NewServeMux()
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

	router.Handle("/uppercase", uppercaseHandler)
	router.Handle("/lowercase", lowercaseHandler)
	router.Handle("/concatenate", concatHandler)
	router.Handle("/split", splitHandler)
	router.Handle("/count", countHandler)

	return router
}
