package handler

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Lowercase(string) (string, error)
	Concat(string, string) (string, error)
	Split(string, string) ([]string, error)
	Count(string) int
}
