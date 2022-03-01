package endpoints

import (
	"context"

	"github.com/amirhnajafiz/go-kit/internal/http/handler"
	request2 "github.com/amirhnajafiz/go-kit/internal/http/request"
	response2 "github.com/amirhnajafiz/go-kit/internal/http/response"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)

func MakeUppercaseEndpoint(svc handler.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(request2.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return response2.UppercaseResponse{V: v, Err: err.Error()}, nil
		}

		return response2.UppercaseResponse{V: v, Err: ""}, nil
	}
}

func MakeLowercaseEndpoint(svc handler.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(request2.LowercaseRequest)
		v, err := svc.Lowercase(req.S)
		if err != nil {
			return response2.LowercaseResponse{V: v, Err: err.Error()}, nil
		}

		return response2.LowercaseResponse{V: v, Err: ""}, nil
	}
}

func MakeCountEndpoint(svc handler.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(request2.CountRequest)
		v := svc.Count(req.S)

		return response2.CountResponse{V: v}, nil
	}
}
