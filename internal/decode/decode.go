package decode

import (
	"context"
	"encoding/json"
	"net/http"

	request2 "github.com/amirhnajafiz/go-kit/internal/http/request"
)

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request2.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request2.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
