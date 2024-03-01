// This file was auto-generated by Fern from our API Definition.

package workflowdeployments

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

func (c *Client) List(ctx context.Context, request *vellumclientgo.WorkflowDeploymentsListRequest) (*vellumclientgo.PaginatedSlimWorkflowDeploymentList, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/workflow-deployments"

	queryParams := make(url.Values)
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	if request.Ordering != nil {
		queryParams.Add("ordering", fmt.Sprintf("%v", *request.Ordering))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", request.Status))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *vellumclientgo.PaginatedSlimWorkflowDeploymentList
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

// Used to retrieve a workflow deployment given its ID or name.
//
// Either the Workflow Deployment's ID or its unique name
func (c *Client) Retrieve(ctx context.Context, id string) (*vellumclientgo.WorkflowDeploymentRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/workflow-deployments/%v", id)

	var response *vellumclientgo.WorkflowDeploymentRead
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
