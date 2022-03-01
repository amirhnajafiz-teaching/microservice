package response

// For each method, we define response structs

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type LowercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type ConcatenateResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type SplitResponse struct {
	V   []string `json:"v"`
	Err string   `json:"err,omitempty"` // errors don't define JSON marshaling
}

type CountResponse struct {
	V int `json:"v"`
}
