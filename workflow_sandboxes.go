// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/vellum-ai/vellum-client-go/core"
)

type DeploySandboxWorkflowRequest struct {
	// The Vellum-generated ID of the Workflow Deployment you'd like to update. Cannot specify both this and workflow_deployment_name. Leave null to create a new Workflow Deployment.
	WorkflowDeploymentId *string `json:"workflow_deployment_id,omitempty" url:"-"`
	// The unique name of the Workflow Deployment you'd like to either create or update. Cannot specify both this and workflow_deployment_id. If provided and matches an existing Workflow Deployment, that Workflow Deployment will be updated. Otherwise, a new Prompt Deployment will be created.
	WorkflowDeploymentName *string `json:"workflow_deployment_name,omitempty" url:"-"`
	// In the event that a new Workflow Deployment is created, this will be the label it's given.
	Label *string `json:"label,omitempty" url:"-"`
	// Optionally provide the release tags that you'd like to be associated with the latest release of the created/updated Prompt Deployment.
	ReleaseTags []string `json:"release_tags,omitempty" url:"-"`
	// Optionally provide a description that details what's new in this Release.
	ReleaseDescription *string `json:"release_description,omitempty" url:"-"`
}

type ListWorkflowSandboxExamplesRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The initial index from which to return the results.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Which field to use when ordering the results.
	Ordering *string                                `json:"-" url:"ordering,omitempty"`
	Tag      *ListWorkflowSandboxExamplesRequestTag `json:"-" url:"tag,omitempty"`
}

type PaginatedWorkflowSandboxExampleList struct {
	Count    *int                      `json:"count,omitempty" url:"count,omitempty"`
	Next     *string                   `json:"next,omitempty" url:"next,omitempty"`
	Previous *string                   `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*WorkflowSandboxExample `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PaginatedWorkflowSandboxExampleList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedWorkflowSandboxExampleList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedWorkflowSandboxExampleList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedWorkflowSandboxExampleList(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedWorkflowSandboxExampleList) String() string {
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

type WorkflowSandboxExample struct {
	Id           string  `json:"id" url:"id"`
	Label        string  `json:"label" url:"label"`
	Description  *string `json:"description,omitempty" url:"description,omitempty"`
	IconName     *string `json:"icon_name,omitempty" url:"icon_name,omitempty"`
	UiImageUrl   *string `json:"ui_image_url,omitempty" url:"ui_image_url,omitempty"`
	CodeImageUrl *string `json:"code_image_url,omitempty" url:"code_image_url,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (w *WorkflowSandboxExample) GetExtraProperties() map[string]interface{} {
	return w.extraProperties
}

func (w *WorkflowSandboxExample) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkflowSandboxExample
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkflowSandboxExample(value)

	extraProperties, err := core.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties

	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WorkflowSandboxExample) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type ListWorkflowSandboxExamplesRequestTag string

const (
	ListWorkflowSandboxExamplesRequestTagOnboarding ListWorkflowSandboxExamplesRequestTag = "ONBOARDING"
	ListWorkflowSandboxExamplesRequestTagTemplates  ListWorkflowSandboxExamplesRequestTag = "TEMPLATES"
)

func NewListWorkflowSandboxExamplesRequestTagFromString(s string) (ListWorkflowSandboxExamplesRequestTag, error) {
	switch s {
	case "ONBOARDING":
		return ListWorkflowSandboxExamplesRequestTagOnboarding, nil
	case "TEMPLATES":
		return ListWorkflowSandboxExamplesRequestTagTemplates, nil
	}
	var t ListWorkflowSandboxExamplesRequestTag
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListWorkflowSandboxExamplesRequestTag) Ptr() *ListWorkflowSandboxExamplesRequestTag {
	return &l
}
