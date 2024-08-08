// This file was auto-generated by Fern from our API Definition.

package core

import (
	fmt "fmt"
	http "net/http"
)

// ClientOption adapts the behavior of the generated client.
type ClientOption func(*ClientOptions)

// ClientOptions defines all of the possible client options.
// This type is primarily used by the generated code and is
// not meant to be used directly; use ClientOption instead.
type ClientOptions struct {
	BaseURL    string
	HTTPClient HTTPClient
	HTTPHeader http.Header
	ApiKey     string
}

// NewClientOptions returns a new *ClientOptions value.
// This function is primarily used by the generated code and is
// not meant to be used directly; use ClientOption instead.
func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		HTTPClient: http.DefaultClient,
		HTTPHeader: make(http.Header),
	}
}

// ToHeader maps the configured client options into a http.Header issued
// on every request.
func (c *ClientOptions) ToHeader() http.Header {
	header := c.cloneHeader()
	header.Set("X_API_KEY", fmt.Sprintf("%v", c.ApiKey))
	return header
}

func (c *ClientOptions) cloneHeader() http.Header {
	headers := c.HTTPHeader.Clone()
	headers.Set("X-Fern-Language", "Go")
	headers.Set("X-Fern-SDK-Name", "github.com/vellum-ai/vellum-client-go")
	headers.Set("X-Fern-SDK-Version", "v0.7.9")
	return headers
}
