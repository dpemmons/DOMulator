package fetch

import (
	"encoding/json"
	"fmt"
)

// Response represents an HTTP response in the fetch API
type Response struct {
	Status     int
	StatusText string
	Headers    map[string]string
	Body       string
	URL        string
	Ok         bool
}

// JSON parses the response body as JSON
func (r *Response) JSON(v interface{}) error {
	if r.Body == "" {
		return fmt.Errorf("response body is empty")
	}

	err := json.Unmarshal([]byte(r.Body), v)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
}

// Text returns the response body as a string
func (r *Response) Text() string {
	return r.Body
}

// ArrayBuffer returns the response body as bytes
func (r *Response) ArrayBuffer() []byte {
	return []byte(r.Body)
}
