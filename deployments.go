// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
)

type DeploymentsListRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-"`
	// The initial index from which to return the results.
	Offset *int `json:"-"`
	// Which field to use when ordering the results.
	Ordering *string `json:"-"`
	// The current status of the deployment
	//
	// - `ACTIVE` - Active
	// - `ARCHIVED` - Archived
	Status *DeploymentsListRequestStatus `json:"-"`
}

type DeploymentProviderPayloadRequest struct {
	// The ID of the deployment. Must provide either this or deployment_name.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// The name of the deployment. Must provide either this or deployment_id.
	DeploymentName *string `json:"deployment_name,omitempty"`
	// The list of inputs defined in the Prompt's deployment with their corresponding values.
	Inputs []*PromptDeploymentInputRequest `json:"inputs,omitempty"`
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
