// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/vellum-ai/vellum-client-go/core"
	time "time"
)

type DeploymentsListRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The initial index from which to return the results.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Which field to use when ordering the results.
	Ordering *string `json:"-" url:"ordering,omitempty"`
	// status
	Status *DeploymentsListRequestStatus `json:"-" url:"status,omitempty"`
}

type ListDeploymentReleaseTagsRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The initial index from which to return the results.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Which field to use when ordering the results.
	Ordering *string                                 `json:"-" url:"ordering,omitempty"`
	Source   *ListDeploymentReleaseTagsRequestSource `json:"-" url:"source,omitempty"`
}

type DeploymentProviderPayloadRequest struct {
	// The ID of the deployment. Must provide either this or deployment_name.
	DeploymentId *string `json:"deployment_id,omitempty" url:"-"`
	// The name of the deployment. Must provide either this or deployment_id.
	DeploymentName *string `json:"deployment_name,omitempty" url:"-"`
	// The list of inputs defined in the Prompt's deployment with their corresponding values.
	Inputs []*PromptDeploymentInputRequest `json:"inputs,omitempty" url:"-"`
	// Optionally specify a release tag if you want to pin to a specific release of the Workflow Deployment
	ReleaseTag *string                                   `json:"release_tag,omitempty" url:"-"`
	ExpandMeta *CompilePromptDeploymentExpandMetaRequest `json:"expand_meta,omitempty" url:"-"`
}

type CompilePromptDeploymentExpandMetaRequest struct {
	// If enabled, the response will include the model identifier representing the ML Model invoked by the Prompt.
	ModelName *bool `json:"model_name,omitempty" url:"model_name,omitempty"`
	// If enabled, the response will include the release tag of the Prompt Deployment.
	DeploymentReleaseTag *bool `json:"deployment_release_tag,omitempty" url:"deployment_release_tag,omitempty"`
	// If enabled, the response will include the ID of the Prompt Version backing the deployment.
	PromptVersionId *bool `json:"prompt_version_id,omitempty" url:"prompt_version_id,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CompilePromptDeploymentExpandMetaRequest) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CompilePromptDeploymentExpandMetaRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CompilePromptDeploymentExpandMetaRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CompilePromptDeploymentExpandMetaRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CompilePromptDeploymentExpandMetaRequest) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// The subset of the metadata tracked by Vellum during Prompt Deployment compilation that the request opted into with `expand_meta`.
type CompilePromptMeta struct {
	ModelName            *string `json:"model_name,omitempty" url:"model_name,omitempty"`
	DeploymentReleaseTag *string `json:"deployment_release_tag,omitempty" url:"deployment_release_tag,omitempty"`
	PromptVersionId      *string `json:"prompt_version_id,omitempty" url:"prompt_version_id,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CompilePromptMeta) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CompilePromptMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler CompilePromptMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CompilePromptMeta(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CompilePromptMeta) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type DeploymentHistoryItem struct {
	Id           string    `json:"id" url:"id"`
	DeploymentId string    `json:"deployment_id" url:"deployment_id"`
	Timestamp    time.Time `json:"timestamp" url:"timestamp"`
	// A human-readable label for the deployment
	Label string `json:"label" url:"label"`
	// A name that uniquely identifies this deployment within its workspace
	Name           string            `json:"name" url:"name"`
	InputVariables []*VellumVariable `json:"input_variables" url:"input_variables"`
	// A human-readable description of the deployment
	Description *string `json:"description,omitempty" url:"description,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeploymentHistoryItem) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeploymentHistoryItem) UnmarshalJSON(data []byte) error {
	type embed DeploymentHistoryItem
	var unmarshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
	}{
		embed: embed(*d),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*d = DeploymentHistoryItem(unmarshaler.embed)
	d.Timestamp = unmarshaler.Timestamp.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeploymentHistoryItem) MarshalJSON() ([]byte, error) {
	type embed DeploymentHistoryItem
	var marshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
	}{
		embed:     embed(*d),
		Timestamp: core.NewDateTime(d.Timestamp),
	}
	return json.Marshal(marshaler)
}

func (d *DeploymentHistoryItem) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeploymentProviderPayloadResponse struct {
	Payload *DeploymentProviderPayloadResponsePayload `json:"payload" url:"payload"`
	Meta    *CompilePromptMeta                        `json:"meta,omitempty" url:"meta,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeploymentProviderPayloadResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeploymentProviderPayloadResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeploymentProviderPayloadResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeploymentProviderPayloadResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeploymentProviderPayloadResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeploymentProviderPayloadResponsePayload struct {
	StringUnknownMap map[string]interface{}
	String           string
}

func (d *DeploymentProviderPayloadResponsePayload) UnmarshalJSON(data []byte) error {
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		d.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		d.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, d)
}

func (d DeploymentProviderPayloadResponsePayload) MarshalJSON() ([]byte, error) {
	if d.StringUnknownMap != nil {
		return json.Marshal(d.StringUnknownMap)
	}
	if d.String != "" {
		return json.Marshal(d.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", d)
}

type DeploymentProviderPayloadResponsePayloadVisitor interface {
	VisitStringUnknownMap(map[string]interface{}) error
	VisitString(string) error
}

func (d *DeploymentProviderPayloadResponsePayload) Accept(visitor DeploymentProviderPayloadResponsePayloadVisitor) error {
	if d.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(d.StringUnknownMap)
	}
	if d.String != "" {
		return visitor.VisitString(d.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", d)
}

type DeploymentReleaseTagDeploymentHistoryItem struct {
	Id        string    `json:"id" url:"id"`
	Timestamp time.Time `json:"timestamp" url:"timestamp"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeploymentReleaseTagDeploymentHistoryItem) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeploymentReleaseTagDeploymentHistoryItem) UnmarshalJSON(data []byte) error {
	type embed DeploymentReleaseTagDeploymentHistoryItem
	var unmarshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
	}{
		embed: embed(*d),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*d = DeploymentReleaseTagDeploymentHistoryItem(unmarshaler.embed)
	d.Timestamp = unmarshaler.Timestamp.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeploymentReleaseTagDeploymentHistoryItem) MarshalJSON() ([]byte, error) {
	type embed DeploymentReleaseTagDeploymentHistoryItem
	var marshaler = struct {
		embed
		Timestamp *core.DateTime `json:"timestamp"`
	}{
		embed:     embed(*d),
		Timestamp: core.NewDateTime(d.Timestamp),
	}
	return json.Marshal(marshaler)
}

func (d *DeploymentReleaseTagDeploymentHistoryItem) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeploymentReleaseTagRead struct {
	// The name of the Release Tag
	Name string `json:"name" url:"name"`
	// The source of how the Release Tag was originally created
	//
	// * `SYSTEM` - System
	// * `USER` - User
	Source ReleaseTagSource `json:"source" url:"source"`
	// The Deployment History Item that this Release Tag is associated with
	HistoryItem *DeploymentReleaseTagDeploymentHistoryItem `json:"history_item" url:"history_item"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeploymentReleaseTagRead) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeploymentReleaseTagRead) UnmarshalJSON(data []byte) error {
	type unmarshaler DeploymentReleaseTagRead
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeploymentReleaseTagRead(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeploymentReleaseTagRead) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type PaginatedDeploymentReleaseTagReadList struct {
	Count    *int                        `json:"count,omitempty" url:"count,omitempty"`
	Next     *string                     `json:"next,omitempty" url:"next,omitempty"`
	Previous *string                     `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*DeploymentReleaseTagRead `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PaginatedDeploymentReleaseTagReadList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedDeploymentReleaseTagReadList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedDeploymentReleaseTagReadList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedDeploymentReleaseTagReadList(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedDeploymentReleaseTagReadList) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PaginatedSlimDeploymentReadList struct {
	Count    *int                  `json:"count,omitempty" url:"count,omitempty"`
	Next     *string               `json:"next,omitempty" url:"next,omitempty"`
	Previous *string               `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*SlimDeploymentRead `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PaginatedSlimDeploymentReadList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedSlimDeploymentReadList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedSlimDeploymentReadList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedSlimDeploymentReadList(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedSlimDeploymentReadList) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type SlimDeploymentRead struct {
	Id      string    `json:"id" url:"id"`
	Created time.Time `json:"created" url:"created"`
	// A human-readable label for the deployment
	Label string `json:"label" url:"label"`
	// A name that uniquely identifies this deployment within its workspace
	Name string `json:"name" url:"name"`
	// The current status of the deployment
	//
	// * `ACTIVE` - Active
	// * `ARCHIVED` - Archived
	Status *EntityStatus `json:"status,omitempty" url:"status,omitempty"`
	// The environment this deployment is used in
	//
	// * `DEVELOPMENT` - Development
	// * `STAGING` - Staging
	// * `PRODUCTION` - Production
	Environment    *EnvironmentEnum  `json:"environment,omitempty" url:"environment,omitempty"`
	LastDeployedOn time.Time         `json:"last_deployed_on" url:"last_deployed_on"`
	InputVariables []*VellumVariable `json:"input_variables" url:"input_variables"`
	// A human-readable description of the deployment
	Description *string `json:"description,omitempty" url:"description,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SlimDeploymentRead) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SlimDeploymentRead) UnmarshalJSON(data []byte) error {
	type embed SlimDeploymentRead
	var unmarshaler = struct {
		embed
		Created        *core.DateTime `json:"created"`
		LastDeployedOn *core.DateTime `json:"last_deployed_on"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SlimDeploymentRead(unmarshaler.embed)
	s.Created = unmarshaler.Created.Time()
	s.LastDeployedOn = unmarshaler.LastDeployedOn.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SlimDeploymentRead) MarshalJSON() ([]byte, error) {
	type embed SlimDeploymentRead
	var marshaler = struct {
		embed
		Created        *core.DateTime `json:"created"`
		LastDeployedOn *core.DateTime `json:"last_deployed_on"`
	}{
		embed:          embed(*s),
		Created:        core.NewDateTime(s.Created),
		LastDeployedOn: core.NewDateTime(s.LastDeployedOn),
	}
	return json.Marshal(marshaler)
}

func (s *SlimDeploymentRead) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type DeploymentsListRequestStatus string

const (
	DeploymentsListRequestStatusActive   DeploymentsListRequestStatus = "ACTIVE"
	DeploymentsListRequestStatusArchived DeploymentsListRequestStatus = "ARCHIVED"
)

func NewDeploymentsListRequestStatusFromString(s string) (DeploymentsListRequestStatus, error) {
	switch s {
	case "ACTIVE":
		return DeploymentsListRequestStatusActive, nil
	case "ARCHIVED":
		return DeploymentsListRequestStatusArchived, nil
	}
	var t DeploymentsListRequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DeploymentsListRequestStatus) Ptr() *DeploymentsListRequestStatus {
	return &d
}

type ListDeploymentReleaseTagsRequestSource string

const (
	ListDeploymentReleaseTagsRequestSourceSystem ListDeploymentReleaseTagsRequestSource = "SYSTEM"
	ListDeploymentReleaseTagsRequestSourceUser   ListDeploymentReleaseTagsRequestSource = "USER"
)

func NewListDeploymentReleaseTagsRequestSourceFromString(s string) (ListDeploymentReleaseTagsRequestSource, error) {
	switch s {
	case "SYSTEM":
		return ListDeploymentReleaseTagsRequestSourceSystem, nil
	case "USER":
		return ListDeploymentReleaseTagsRequestSourceUser, nil
	}
	var t ListDeploymentReleaseTagsRequestSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListDeploymentReleaseTagsRequestSource) Ptr() *ListDeploymentReleaseTagsRequestSource {
	return &l
}

type PatchedDeploymentReleaseTagUpdateRequest struct {
	// The ID of the Deployment History Item to tag
	HistoryItemId *string `json:"history_item_id,omitempty" url:"-"`
}
