package request

// For each method, we define request structs

type UppercaseRequest struct {
	S string `json:"s"`
}

type LowercaseRequest struct {
	S string `json:"s"`
}

type ConcatenateRequest struct {
	S string `json:"s"`
	C string `json:"c"`
}

type SplitRequest struct {
	S string `json:"s"`
	K string `json:"k"`
}

type CountRequest struct {
	S string `json:"s"`
}
