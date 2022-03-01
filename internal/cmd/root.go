package cmd

import (
	"github.com/amirhnajafiz/go-kit/internal/decode"
	"github.com/amirhnajafiz/go-kit/internal/endpoints"
	"github.com/amirhnajafiz/go-kit/internal/http/handler"
	"log"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func Execute() {
	svc := handler.Handler{}

	uppercaseHandler := httpTransport.NewServer(
		endpoints.MakeUppercaseEndpoint(svc),
		decode.DecodeUppercaseRequest,
		decode.EncodeResponse,
	)

	lowercaseHandler := httpTransport.NewServer(
		endpoints.MakeLowercaseEndpoint(svc),
		decode.DecodeLowercaseRequest,
		decode.EncodeResponse,
	)

	concatHandler := httpTransport.NewServer(
		endpoints.MakeConcatenateEndpoint(svc),
		decode.DecodeConcatenateRequest,
		decode.EncodeResponse,
	)

	splitHandler := httpTransport.NewServer(
		endpoints.MakeSplitEndpoint(svc),
		decode.DecodeSplitRequest,
		decode.EncodeResponse,
	)

	countHandler := httpTransport.NewServer(
		endpoints.MakeCountEndpoint(svc),
		decode.DecodeCountRequest,
		decode.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/lowercase", lowercaseHandler)
	http.Handle("/concatenate", concatHandler)
	http.Handle("/split", splitHandler)
	http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
