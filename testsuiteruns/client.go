// This file was auto-generated by Fern from our API Definition.

package testsuiteruns

import (
	context "context"
	fmt "fmt"
	vellumclientgo "github.com/vellum-ai/vellum-client-go"
	core "github.com/vellum-ai/vellum-client-go/core"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Trigger a Test Suite and create a new Test Suite Run
func (c *Client) Create(ctx context.Context, request *vellumclientgo.TestSuiteRunCreateRequest) (*vellumclientgo.TestSuiteRunRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/test-suite-runs"

	var response *vellumclientgo.TestSuiteRunRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve a specific Test Suite Run by ID
//
// A UUID string identifying this test suite run.
func (c *Client) Retrieve(ctx context.Context, id string) (*vellumclientgo.TestSuiteRunRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/test-suite-runs/%v", id)

	var response *vellumclientgo.TestSuiteRunRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// A UUID string identifying this test suite run.
func (c *Client) ListExecutions(ctx context.Context, id string, request *vellumclientgo.TestSuiteRunsListExecutionsRequest) (*vellumclientgo.PaginatedTestSuiteRunExecutionList, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/test-suite-runs/%v/executions", id)

	queryParams := make(url.Values)
	for _, value := range request.Expand {
		queryParams.Add("expand", fmt.Sprintf("%v", *value))
	}
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *vellumclientgo.PaginatedTestSuiteRunExecutionList
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
