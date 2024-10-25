// This file was auto-generated by Fern from our API Definition.

package client

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	vellumclientgo "github.com/vellum-ai/vellum-client-go"
	adhoc "github.com/vellum-ai/vellum-client-go/adhoc"
	core "github.com/vellum-ai/vellum-client-go/core"
	deployments "github.com/vellum-ai/vellum-client-go/deployments"
	documentindexes "github.com/vellum-ai/vellum-client-go/documentindexes"
	documents "github.com/vellum-ai/vellum-client-go/documents"
	folderentities "github.com/vellum-ai/vellum-client-go/folderentities"
	metricdefinitions "github.com/vellum-ai/vellum-client-go/metricdefinitions"
	option "github.com/vellum-ai/vellum-client-go/option"
	sandboxes "github.com/vellum-ai/vellum-client-go/sandboxes"
	testsuiteruns "github.com/vellum-ai/vellum-client-go/testsuiteruns"
	testsuites "github.com/vellum-ai/vellum-client-go/testsuites"
	workflowdeployments "github.com/vellum-ai/vellum-client-go/workflowdeployments"
	workflowsandboxes "github.com/vellum-ai/vellum-client-go/workflowsandboxes"
	workspacesecrets "github.com/vellum-ai/vellum-client-go/workspacesecrets"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	AdHoc               *adhoc.Client
	Deployments         *deployments.Client
	DocumentIndexes     *documentindexes.Client
	Documents           *documents.Client
	FolderEntities      *folderentities.Client
	MetricDefinitions   *metricdefinitions.Client
	Sandboxes           *sandboxes.Client
	TestSuiteRuns       *testsuiteruns.Client
	TestSuites          *testsuites.Client
	WorkflowDeployments *workflowdeployments.Client
	WorkflowSandboxes   *workflowsandboxes.Client
	WorkspaceSecrets    *workspacesecrets.Client
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
		header:              options.ToHeader(),
		AdHoc:               adhoc.NewClient(opts...),
		Deployments:         deployments.NewClient(opts...),
		DocumentIndexes:     documentindexes.NewClient(opts...),
		Documents:           documents.NewClient(opts...),
		FolderEntities:      folderentities.NewClient(opts...),
		MetricDefinitions:   metricdefinitions.NewClient(opts...),
		Sandboxes:           sandboxes.NewClient(opts...),
		TestSuiteRuns:       testsuiteruns.NewClient(opts...),
		TestSuites:          testsuites.NewClient(opts...),
		WorkflowDeployments: workflowdeployments.NewClient(opts...),
		WorkflowSandboxes:   workflowsandboxes.NewClient(opts...),
		WorkspaceSecrets:    workspacesecrets.NewClient(opts...),
	}
}

// An internal-only endpoint that's subject to breaking changes without notice. Not intended for public use.
func (c *Client) ExecuteCode(
	ctx context.Context,
	request *vellumclientgo.CodeExecutorRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.CodeExecutorResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/execute-code"

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
		}
		return apiError
	}

	var response *vellumclientgo.CodeExecutorResponse
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

// Executes a deployed Prompt and returns the result.
func (c *Client) ExecutePrompt(
	ctx context.Context,
	request *vellumclientgo.ExecutePromptRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.ExecutePromptResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/execute-prompt"

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

	var response *vellumclientgo.ExecutePromptResponse
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

// Executes a deployed Prompt and streams back the results.
func (c *Client) ExecutePromptStream(
	ctx context.Context,
	request *vellumclientgo.ExecutePromptStreamRequest,
	opts ...option.RequestOption,
) (*core.Stream[vellumclientgo.ExecutePromptEvent], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/execute-prompt-stream"

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

	streamer := core.NewStreamer[vellumclientgo.ExecutePromptEvent](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Headers:         headers,
			Client:          options.HTTPClient,
			Request:         request,
			ErrorDecoder:    errorDecoder,
		},
	)
}

// Executes a deployed Workflow and returns its outputs.
func (c *Client) ExecuteWorkflow(
	ctx context.Context,
	request *vellumclientgo.ExecuteWorkflowRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.ExecuteWorkflowResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/execute-workflow"

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

	var response *vellumclientgo.ExecuteWorkflowResponse
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

// Executes a deployed Workflow and streams back its results.
func (c *Client) ExecuteWorkflowStream(
	ctx context.Context,
	request *vellumclientgo.ExecuteWorkflowStreamRequest,
	opts ...option.RequestOption,
) (*core.Stream[vellumclientgo.WorkflowStreamEvent], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/execute-workflow-stream"

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

	streamer := core.NewStreamer[vellumclientgo.WorkflowStreamEvent](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Headers:         headers,
			Client:          options.HTTPClient,
			Request:         request,
			ErrorDecoder:    errorDecoder,
		},
	)
}

// Generate a completion using a previously defined deployment.
//
// Important: This endpoint is DEPRECATED and has been superseded by
// [execute-prompt](/api-reference/api-reference/execute-prompt).
func (c *Client) Generate(
	ctx context.Context,
	request *vellumclientgo.GenerateBodyRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.GenerateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/generate"

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

	var response *vellumclientgo.GenerateResponse
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

// Generate a stream of completions using a previously defined deployment.
//
// Important: This endpoint is DEPRECATED and has been superseded by
// [execute-prompt-stream](/api-reference/api-reference/execute-prompt-stream).
func (c *Client) GenerateStream(
	ctx context.Context,
	request *vellumclientgo.GenerateStreamBodyRequest,
	opts ...option.RequestOption,
) (*core.Stream[vellumclientgo.GenerateStreamResponse], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/generate-stream"

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

	streamer := core.NewStreamer[vellumclientgo.GenerateStreamResponse](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Headers:         headers,
			Client:          options.HTTPClient,
			Request:         request,
			ErrorDecoder:    errorDecoder,
		},
	)
}

// Perform a search against a document index.
func (c *Client) Search(
	ctx context.Context,
	request *vellumclientgo.SearchRequestBodyRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.SearchResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/search"

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

	var response *vellumclientgo.SearchResponse
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

// Used to submit feedback regarding the quality of previously generated completions.
func (c *Client) SubmitCompletionActuals(
	ctx context.Context,
	request *vellumclientgo.SubmitCompletionActualsRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/submit-completion-actuals"

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
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

// Used to submit feedback regarding the quality of previous workflow execution and its outputs.
//
// **Note:** Uses a base url of `https://predict.vellum.ai`.
func (c *Client) SubmitWorkflowExecutionActuals(
	ctx context.Context,
	request *vellumclientgo.SubmitWorkflowExecutionActualsRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://predict.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/submit-workflow-execution-actuals"

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
