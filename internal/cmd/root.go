package cmd

import (
	"github.com/amirhnajafiz/go-kit/internal/decode"
	"github.com/amirhnajafiz/go-kit/internal/endpoints"
	"github.com/amirhnajafiz/go-kit/internal/http/handler"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func Execute() {
	svc := handler.Handler{}

	uppercaseHandler := httptransport.NewServer(
		endpoints.MakeUppercaseEndpoint(svc),
		decode.DecodeUppercaseRequest,
		decode.EncodeResponse,
	)

	lowercaseHandler := httptransport.NewServer(
		endpoints.MakeLowercaseEndpoint(svc),
		decode.DecodeLowercaseRequest,
		decode.EncodeResponse,
	)

	concatHandler := httptransport.NewServer(
		endpoints.MakeConcatenateEndpoint(svc),
		decode.DecodeConcatenateRequest,
		decode.EncodeResponse,
	)

	splitHandler := httptransport.NewServer(
		endpoints.MakeSplitEndpoint(svc),
		decode.DecodeSplitRequest,
		decode.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
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
