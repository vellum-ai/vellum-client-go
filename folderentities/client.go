// This file was auto-generated by Fern from our API Definition.

package folderentities

import (
	context "context"
	vellumclientgo "github.com/vellum-ai/vellum-client-go"
	core "github.com/vellum-ai/vellum-client-go/core"
	option "github.com/vellum-ai/vellum-client-go/option"
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

// List all folder entities within a specified folder.
func (c *Client) List(
	ctx context.Context,
	request *vellumclientgo.FolderEntitiesListRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.PaginatedFolderEntityList, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/folder-entities"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.PaginatedFolderEntityList
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

// Add an entity to a specific folder or root directory.
//
// Adding an entity to a folder will remove it from any other folders it might have been a member of.
func (c *Client) AddEntityToFolder(
	ctx context.Context,
	// The ID of the folder to which the entity should be added. This can be a UUID of a folder, or the name of a root
	// directory. Supported root directories include:
	// - PROMPT_SANDBOX
	// - WORKFLOW_SANDBOX
	// - DOCUMENT_INDEX
	// - TEST_SUITE
	folderId string,
	request *vellumclientgo.AddEntityToFolderRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/folders/%v/add-entity", folderId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

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
		},
	); err != nil {
		return err
	}
	return nil
}
