// This file was auto-generated by Fern from our API Definition.

package deployments

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	vellumclientgo "github.com/vellum-ai/vellum-client-go"
	core "github.com/vellum-ai/vellum-client-go/core"
	io "io"
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

// Used to list all Prompt Deployments.
func (c *Client) List(ctx context.Context, request *vellumclientgo.DeploymentsListRequest) (*vellumclientgo.PaginatedSlimDeploymentReadList, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/deployments"

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

	var response *vellumclientgo.PaginatedSlimDeploymentReadList
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

// Used to retrieve a Prompt Deployment given its ID or name.
//
// Either the Deployment's ID or its unique name
func (c *Client) Retrieve(ctx context.Context, id string) (*vellumclientgo.DeploymentRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/deployments/%v", id)

	var response *vellumclientgo.DeploymentRead
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

// Retrieve a Deployment Release Tag by tag name, associated with a specified Deployment.
//
// A UUID string identifying this deployment.
// The name of the Release Tag associated with this Deployment that you'd like to retrieve.
func (c *Client) RetrieveDeploymentReleaseTag(ctx context.Context, id string, name string) (*vellumclientgo.DeploymentReleaseTagRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/deployments/%v/release-tags/%v", id, name)

	var response *vellumclientgo.DeploymentReleaseTagRead
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

// Updates an existing Release Tag associated with the specified Deployment.
//
// A UUID string identifying this deployment.
// The name of the Release Tag associated with this Deployment that you'd like to update.
func (c *Client) UpdateDeploymentReleaseTag(ctx context.Context, id string, name string, request *vellumclientgo.PatchedDeploymentReleaseTagUpdateRequest) (*vellumclientgo.DeploymentReleaseTagRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/deployments/%v/release-tags/%v", id, name)

	var response *vellumclientgo.DeploymentReleaseTagRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPatch,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Given a set of input variable values, compile the exact payload that Vellum would send to the configured model provider
// for execution if the execute-prompt endpoint had been invoked. Note that this endpoint does not actually execute the
// prompt or make an API call to the model provider.
//
// This endpoint is useful if you don't want to proxy LLM provider requests through Vellum and prefer to send them directly
// to the provider yourself. Note that no guarantees are made on the format of this API's response schema, other than
// that it will be a valid payload for the configured model provider. It's not recommended that you try to parse or
// derive meaning from the response body and instead, should simply pass it directly to the model provider as is.
//
// We encourage you to seek advise from Vellum Support before integrating with this API for production use.
func (c *Client) RetrieveProviderPayload(ctx context.Context, request *vellumclientgo.DeploymentProviderPayloadRequest) (*vellumclientgo.DeploymentProviderPayloadResponse, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/deployments/provider-payload"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(vellumclientgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 403:
			value := new(vellumclientgo.ForbiddenError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 404:
			value := new(vellumclientgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 500:
			value := new(vellumclientgo.InternalServerError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *vellumclientgo.DeploymentProviderPayloadResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
