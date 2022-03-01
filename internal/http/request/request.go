package request

// For each method, we define request structs

type UppercaseRequest struct {
	S string `json:"s"`
}

type CountRequest struct {
	S string `json:"s"`
}
