// This file was auto-generated by Fern from our API Definition.

package documentindexes

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

// Used to retrieve a list of Document Indexes.
func (c *Client) List(ctx context.Context, request *vellumclientgo.DocumentIndexesListRequest) (*vellumclientgo.PaginatedDocumentIndexReadList, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/document-indexes"

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
	if request.Search != nil {
		queryParams.Add("search", fmt.Sprintf("%v", *request.Search))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", request.Status))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *vellumclientgo.PaginatedDocumentIndexReadList
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

// Creates a new document index.
func (c *Client) Create(ctx context.Context, request *vellumclientgo.DocumentIndexCreateRequest) (*vellumclientgo.DocumentIndexRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/document-indexes"

	var response *vellumclientgo.DocumentIndexRead
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

// Used to retrieve a Document Index given its ID or name.
//
// Either the Document Index's ID or its unique name
func (c *Client) Retrieve(ctx context.Context, id string) (*vellumclientgo.DocumentIndexRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/document-indexes/%v", id)

	var response *vellumclientgo.DocumentIndexRead
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

// Used to fully update a Document Index given its ID or name.
//
// Either the Document Index's ID or its unique name
func (c *Client) Update(ctx context.Context, id string, request *vellumclientgo.DocumentIndexUpdateRequest) (*vellumclientgo.DocumentIndexRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/document-indexes/%v", id)

	var response *vellumclientgo.DocumentIndexRead
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPut,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Used to delete a Document Index given its ID or name.
//
// Either the Document Index's ID or its unique name
func (c *Client) Destroy(ctx context.Context, id string) error {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/document-indexes/%v", id)

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodDelete,
			Headers: c.header,
		},
	); err != nil {
		return err
	}
	return nil
}

// Used to partial update a Document Index given its ID or name.
//
// Either the Document Index's ID or its unique name
func (c *Client) PartialUpdate(ctx context.Context, id string, request *vellumclientgo.PatchedDocumentIndexUpdateRequest) (*vellumclientgo.DocumentIndexRead, error) {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/document-indexes/%v", id)

	var response *vellumclientgo.DocumentIndexRead
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

// Adds a previously uploaded Document to the specified Document Index.
//
// Either the Vellum-generated ID or the originally supplied external_id that uniquely identifies the Document you'd like to add.
// Either the Vellum-generated ID or the originally specified name that uniquely identifies the Document Index to which you'd like to add the Document.
func (c *Client) AddDocument(ctx context.Context, documentId string, id string) error {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/document-indexes/%v/documents/%v", documentId, id)

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: c.header,
		},
	); err != nil {
		return err
	}
	return nil
}

// Removes a Document from a Document Index without deleting the Document itself.
//
// Either the Vellum-generated ID or the originally supplied external_id that uniquely identifies the Document you'd like to remove.
// Either the Vellum-generated ID or the originally specified name that uniquely identifies the Document Index from which you'd like to remove a Document.
func (c *Client) RemoveDocument(ctx context.Context, documentId string, id string) error {
	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/document-indexes/%v/documents/%v", documentId, id)

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodDelete,
			Headers: c.header,
		},
	); err != nil {
		return err
	}
	return nil
}
