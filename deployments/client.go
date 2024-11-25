// This file was auto-generated by Fern from our API Definition.

package deployments

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	vellumclientgo "github.com/vellum-ai/vellum-client-go"
	core "github.com/vellum-ai/vellum-client-go/core"
	option "github.com/vellum-ai/vellum-client-go/option"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Used to list all Prompt Deployments.
func (c *Client) List(
	ctx context.Context,
	request *vellumclientgo.DeploymentsListRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.PaginatedSlimDeploymentReadList, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/deployments"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.PaginatedSlimDeploymentReadList
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Used to retrieve a Prompt Deployment given its ID or name.
func (c *Client) Retrieve(
	ctx context.Context,
	// Either the Deployment's ID or its unique name
	id string,
	opts ...option.RequestOption,
) (*vellumclientgo.DeploymentRead, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/deployments/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.DeploymentRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve a specific Deployment History Item by either its UUID or the name of a Release Tag that points to it.
func (c *Client) DeploymentHistoryItemRetrieve(
	ctx context.Context,
	// Either the UUID of Deployment History Item you'd like to retrieve, or the name of a Release Tag that's pointing to the Deployment History Item you'd like to retrieve.
	historyIdOrReleaseTag string,
	// A UUID string identifying this deployment.
	id string,
	opts ...option.RequestOption,
) (*vellumclientgo.DeploymentHistoryItem, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v1/deployments/%v/history/%v",
		id,
		historyIdOrReleaseTag,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.DeploymentHistoryItem
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// List Release Tags associated with the specified Prompt Deployment
func (c *Client) ListDeploymentReleaseTags(
	ctx context.Context,
	// Either the Prompt Deployment's ID or its unique name
	id string,
	request *vellumclientgo.ListDeploymentReleaseTagsRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.PaginatedDeploymentReleaseTagReadList, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/deployments/%v/release-tags", id)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.PaginatedDeploymentReleaseTagReadList
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve a Deployment Release Tag by tag name, associated with a specified Deployment.
func (c *Client) RetrieveDeploymentReleaseTag(
	ctx context.Context,
	// A UUID string identifying this deployment.
	id string,
	// The name of the Release Tag associated with this Deployment that you'd like to retrieve.
	name string,
	opts ...option.RequestOption,
) (*vellumclientgo.DeploymentReleaseTagRead, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v1/deployments/%v/release-tags/%v",
		id,
		name,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.DeploymentReleaseTagRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Updates an existing Release Tag associated with the specified Prompt Deployment.
func (c *Client) UpdateDeploymentReleaseTag(
	ctx context.Context,
	// A UUID string identifying this deployment.
	id string,
	// The name of the Release Tag associated with this Deployment that you'd like to update.
	name string,
	request *vellumclientgo.PatchedDeploymentReleaseTagUpdateRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.DeploymentReleaseTagRead, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v1/deployments/%v/release-tags/%v",
		id,
		name,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.DeploymentReleaseTagRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
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
func (c *Client) RetrieveProviderPayload(
	ctx context.Context,
	request *vellumclientgo.DeploymentProviderPayloadRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.DeploymentProviderPayloadResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/deployments/provider-payload"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

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
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
