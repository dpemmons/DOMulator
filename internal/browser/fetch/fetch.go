package fetch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/dpemmons/DOMulator/testing"
)

// FetchAPI provides the fetch() implementation for JavaScript
type FetchAPI struct {
	client       *http.Client
	networkMocks *testing.NetworkMocks
}

// New creates a new FetchAPI instance
func New(networkMocks *testing.NetworkMocks) *FetchAPI {
	return &FetchAPI{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		networkMocks: networkMocks,
	}
}

// CreateFetchFunction creates the fetch() function for JavaScript binding
func (f *FetchAPI) CreateFetchFunction(vm *goja.Runtime) func(goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			panic(vm.NewTypeError("fetch requires a URL"))
		}

		url := call.Arguments[0].String()
		var options *RequestOptions

		// Parse options if provided
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			options = f.parseRequestOptions(vm, call.Arguments[1])
		} else {
			options = &RequestOptions{
				Method:  "GET",
				Headers: make(map[string]string),
			}
		}

		// Create and return a Promise
		promise, resolve, reject := vm.NewPromise()

		go func() {
			response, err := f.fetch(url, options)
			if err != nil {
				reject(vm.NewGoError(err))
				return
			}

			jsResponse := f.createResponseObject(vm, response)
			resolve(jsResponse)
		}()

		return vm.ToValue(promise)
	}
}

// RequestOptions represents the options for a fetch request
type RequestOptions struct {
	Method  string
	Headers map[string]string
	Body    string
}

// parseRequestOptions parses JavaScript request options into Go struct
func (f *FetchAPI) parseRequestOptions(vm *goja.Runtime, value goja.Value) *RequestOptions {
	options := &RequestOptions{
		Method:  "GET",
		Headers: make(map[string]string),
	}

	if value == nil || goja.IsNull(value) || goja.IsUndefined(value) {
		return options
	}

	obj := value.ToObject(vm)
	if obj == nil {
		return options
	}

	// Parse method
	if methodVal := obj.Get("method"); !goja.IsUndefined(methodVal) {
		options.Method = strings.ToUpper(methodVal.String())
	}

	// Parse headers
	if headersVal := obj.Get("headers"); !goja.IsUndefined(headersVal) {
		headersObj := headersVal.ToObject(vm)
		if headersObj != nil {
			for _, key := range headersObj.Keys() {
				value := headersObj.Get(key)
				if !goja.IsUndefined(value) {
					options.Headers[key] = value.String()
				}
			}
		}
	}

	// Parse body
	if bodyVal := obj.Get("body"); !goja.IsUndefined(bodyVal) {
		options.Body = bodyVal.String()
	}

	return options
}

// fetch performs the actual HTTP request
func (f *FetchAPI) fetch(url string, options *RequestOptions) (*Response, error) {
	// Check for mock first
	if f.networkMocks != nil {
		if mockResponse := f.networkMocks.Match(options.Method, url); mockResponse != nil {
			return &Response{
				Status:     mockResponse.Status,
				StatusText: http.StatusText(mockResponse.Status),
				Headers:    mockResponse.Headers,
				Body:       mockResponse.Body,
				URL:        url,
				Ok:         mockResponse.Status >= 200 && mockResponse.Status < 300,
			}, nil
		}
	}

	// Create HTTP request
	var body io.Reader
	if options.Body != "" {
		body = strings.NewReader(options.Body)
	}

	req, err := http.NewRequestWithContext(context.Background(), options.Method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// Perform request
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Convert headers
	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &Response{
		Status:     resp.StatusCode,
		StatusText: resp.Status,
		Headers:    headers,
		Body:       string(bodyBytes),
		URL:        url,
		Ok:         resp.StatusCode >= 200 && resp.StatusCode < 300,
	}, nil
}

// createResponseObject creates a JavaScript Response object
func (f *FetchAPI) createResponseObject(vm *goja.Runtime, response *Response) *goja.Object {
	obj := vm.NewObject()

	// Basic properties
	obj.Set("status", response.Status)
	obj.Set("statusText", response.StatusText)
	obj.Set("ok", response.Ok)
	obj.Set("url", response.URL)

	// Headers object
	headersObj := vm.NewObject()
	for key, value := range response.Headers {
		headersObj.Set(key, value)
	}
	obj.Set("headers", headersObj)

	// Body reading methods
	obj.Set("text", func(call goja.FunctionCall) goja.Value {
		promise, resolve, _ := vm.NewPromise()
		resolve(vm.ToValue(response.Body))
		return vm.ToValue(promise)
	})

	obj.Set("json", func(call goja.FunctionCall) goja.Value {
		promise, resolve, reject := vm.NewPromise()

		go func() {
			var result interface{}
			err := response.JSON(&result)
			if err != nil {
				reject(vm.NewGoError(err))
				return
			}
			resolve(vm.ToValue(result))
		}()

		return vm.ToValue(promise)
	})

	obj.Set("arrayBuffer", func(call goja.FunctionCall) goja.Value {
		promise, resolve, _ := vm.NewPromise()
		buffer := vm.NewArrayBuffer([]byte(response.Body))
		resolve(buffer)
		return vm.ToValue(promise)
	})

	obj.Set("blob", func(call goja.FunctionCall) goja.Value {
		promise, resolve, _ := vm.NewPromise()
		// For now, just return the body as text
		// In a full implementation, this would be a proper Blob object
		resolve(vm.ToValue(response.Body))
		return vm.ToValue(promise)
	})

	obj.Set("formData", func(call goja.FunctionCall) goja.Value {
		promise, _, reject := vm.NewPromise()
		// For now, reject as FormData parsing is complex
		// This will be implemented when we add FormData API
		reject(vm.NewTypeError("FormData response parsing not yet implemented"))
		return vm.ToValue(promise)
	})

	// Clone method
	obj.Set("clone", func(call goja.FunctionCall) goja.Value {
		return f.createResponseObject(vm, response)
	})

	return obj
}
