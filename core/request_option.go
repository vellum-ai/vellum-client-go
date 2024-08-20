// This file was auto-generated by Fern from our API Definition.

package core

import (
	fmt "fmt"
	http "net/http"
)

// RequestOption adapts the behavior of the client or an individual request.
type RequestOption interface {
	applyRequestOptions(*RequestOptions)
}

// RequestOptions defines all of the possible request options.
//
// This type is primarily used by the generated code and is not meant
// to be used directly; use the option package instead.
type RequestOptions struct {
	BaseURL     string
	HTTPClient  HTTPClient
	HTTPHeader  http.Header
	MaxAttempts uint
	ApiKey      string
}

// NewRequestOptions returns a new *RequestOptions value.
//
// This function is primarily used by the generated code and is not meant
// to be used directly; use RequestOption instead.
func NewRequestOptions(opts ...RequestOption) *RequestOptions {
	options := &RequestOptions{
		HTTPHeader: make(http.Header),
	}
	for _, opt := range opts {
		opt.applyRequestOptions(options)
	}
	return options
}

// ToHeader maps the configured request options into a http.Header used
// for the request(s).
func (r *RequestOptions) ToHeader() http.Header {
	header := r.cloneHeader()
	if r.ApiKey != "" {
		header.Set("X_API_KEY", fmt.Sprintf("%v", r.ApiKey))
	}
	return header
}

func (r *RequestOptions) cloneHeader() http.Header {
	headers := r.HTTPHeader.Clone()
	headers.Set("X-Fern-Language", "Go")
	headers.Set("X-Fern-SDK-Name", "github.com/vellum-ai/vellum-client-go")
	headers.Set("X-Fern-SDK-Version", "v0.8.2")
	return headers
}

// BaseURLOption implements the RequestOption interface.
type BaseURLOption struct {
	BaseURL string
}

func (b *BaseURLOption) applyRequestOptions(opts *RequestOptions) {
	opts.BaseURL = b.BaseURL
}

// HTTPClientOption implements the RequestOption interface.
type HTTPClientOption struct {
	HTTPClient HTTPClient
}

func (h *HTTPClientOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPClient = h.HTTPClient
}

// HTTPHeaderOption implements the RequestOption interface.
type HTTPHeaderOption struct {
	HTTPHeader http.Header
}

func (h *HTTPHeaderOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPHeader = h.HTTPHeader
}

// MaxAttemptsOption implements the RequestOption interface.
type MaxAttemptsOption struct {
	MaxAttempts uint
}

func (m *MaxAttemptsOption) applyRequestOptions(opts *RequestOptions) {
	opts.MaxAttempts = m.MaxAttempts
}

// ApiKeyOption implements the RequestOption interface.
type ApiKeyOption struct {
	ApiKey string
}

func (a *ApiKeyOption) applyRequestOptions(opts *RequestOptions) {
	opts.ApiKey = a.ApiKey
}
