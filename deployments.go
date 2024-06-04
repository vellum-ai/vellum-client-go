// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
)

type DeploySandboxWorkflowRequest struct {
	// The Vellum-generated ID of the Workflow Deployment you'd like to update. Cannot specify both this and workflow_deployment_name. Leave null to create a new Workflow Deployment.
	WorkflowDeploymentId *string `json:"workflow_deployment_id,omitempty"`
	// The unique name of the Workflow Deployment you'd like to either create or update. Cannot specify both this and workflow_deployment_id. If provided and matches an existing Workflow Deployment, that Workflow Deployment will be updated. Otherwise, a new Prompt Deployment will be created.
	WorkflowDeploymentName *string `json:"workflow_deployment_name,omitempty"`
	// In the event that a new Workflow Deployment is created, this will be the label it's given.
	Label *string `json:"label,omitempty"`
	// Optionally provide the release tags that you'd like to be associated with the latest release of the created/updated Prompt Deployment.
	ReleaseTags []string `json:"release_tags,omitempty"`
}

type DeploymentsListRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-"`
	// The initial index from which to return the results.
	Offset *int `json:"-"`
	// Which field to use when ordering the results.
	Ordering *string `json:"-"`
	// status
	Status *DeploymentsListRequestStatus `json:"-"`
}

type DeploymentProviderPayloadRequest struct {
	// The ID of the deployment. Must provide either this or deployment_name.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// The name of the deployment. Must provide either this or deployment_id.
	DeploymentName *string `json:"deployment_name,omitempty"`
	// The list of inputs defined in the Prompt's deployment with their corresponding values.
	Inputs []*PromptDeploymentInputRequest `json:"inputs,omitempty"`
	// Optionally specify a release tag if you want to pin to a specific release of the Workflow Deployment
	ReleaseTag *string `json:"release_tag,omitempty"`
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

type PatchedDeploymentReleaseTagUpdateRequest struct {
	// The ID of the Deployment History Item to tag
	HistoryItemId *string `json:"history_item_id,omitempty"`
}
