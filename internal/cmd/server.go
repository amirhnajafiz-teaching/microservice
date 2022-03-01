package cmd

import (
	"github.com/amirhnajafiz/go-kit/internal/endpoints"
	"github.com/amirhnajafiz/go-kit/internal/http/handler"
	"github.com/amirhnajafiz/go-kit/internal/http/middleware"
	"github.com/gin-gonic/gin"
	httpTransport "github.com/go-kit/kit/transport/http"
)

func New() *gin.Engine {
	router := gin.Default()
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

	router.POST("/uppercase", uppercaseHandler)
	router.POST("/lowercase", lowercaseHandler)
	router.POST("/concatenate", concatHandler)
	router.POST("/split", splitHandler)
	router.POST("/count", countHandler)

	return router
}
