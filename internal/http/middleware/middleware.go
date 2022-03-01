package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	request2 "github.com/amirhnajafiz/go-kit/internal/http/request"
)

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request2.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeLowercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request2.LowercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeConcatenateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request2.ConcatenateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeSplitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request2.SplitRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request2.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
