// This file was auto-generated by Fern from our API Definition.

package workflows

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	vellumclientgo "github.com/vellum-ai/vellum-client-go"
	core "github.com/vellum-ai/vellum-client-go/core"
	option "github.com/vellum-ai/vellum-client-go/option"
	io "io"
	multipart "mime/multipart"
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

func (c *Client) Pull(
	ctx context.Context,
	// The ID of the Workflow to pull from
	id string,
	request *vellumclientgo.WorkflowsPullRequest,
	opts ...option.RequestOption,
) (io.Reader, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/workflows/%v/pull", id)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

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

	response := bytes.NewBuffer(nil)
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
			Response:        response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Push(
	ctx context.Context,
	artifact io.Reader,
	request *vellumclientgo.WorkflowPushRequest,
	opts ...option.RequestOption,
) (*vellumclientgo.WorkflowPushResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vellum.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/workflows/push"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *vellumclientgo.WorkflowPushResponse
	requestBuffer := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(requestBuffer)
	if artifact != nil {
		artifactFilename := "artifact_filename"
		if named, ok := artifact.(interface{ Name() string }); ok {
			artifactFilename = named.Name()
		}
		artifactPart, err := writer.CreateFormFile("artifact", artifactFilename)
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(artifactPart, artifact); err != nil {
			return nil, err
		}
	}
	if err := core.WriteMultipartJSON(writer, "exec_config", request.ExecConfig); err != nil {
		return nil, err
	}
	if request.WorkflowSandboxId != nil {
		if err := writer.WriteField("workflow_sandbox_id", fmt.Sprintf("%v", *request.WorkflowSandboxId)); err != nil {
			return nil, err
		}
	}
	if request.DeploymentConfig != nil {
		if err := core.WriteMultipartJSON(writer, "deployment_config", request.DeploymentConfig); err != nil {
			return nil, err
		}
	}
	if request.DryRun != nil {
		if err := writer.WriteField("dry_run", fmt.Sprintf("%v", *request.DryRun)); err != nil {
			return nil, err
		}
	}
	if request.Strict != nil {
		if err := writer.WriteField("strict", fmt.Sprintf("%v", *request.Strict)); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	headers.Set("Content-Type", writer.FormDataContentType())

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
			Request:         requestBuffer,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
