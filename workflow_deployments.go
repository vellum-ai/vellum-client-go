// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
)

type WorkflowDeploymentsListRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-"`
	// The initial index from which to return the results.
	Offset *int `json:"-"`
	// Which field to use when ordering the results.
	Ordering *string `json:"-"`
	// The current status of the workflow deployment
	//
	// - `ACTIVE` - Active
	// - `ARCHIVED` - Archived
	Status *WorkflowDeploymentsListRequestStatus `json:"-"`
}

type WorkflowDeploymentsListRequestStatus string

const (
	WorkflowDeploymentsListRequestStatusActive   WorkflowDeploymentsListRequestStatus = "ACTIVE"
	WorkflowDeploymentsListRequestStatusArchived WorkflowDeploymentsListRequestStatus = "ARCHIVED"
)

func NewWorkflowDeploymentsListRequestStatusFromString(s string) (WorkflowDeploymentsListRequestStatus, error) {
	switch s {
	case "ACTIVE":
		return WorkflowDeploymentsListRequestStatusActive, nil
	case "ARCHIVED":
		return WorkflowDeploymentsListRequestStatusArchived, nil
	}
	var t WorkflowDeploymentsListRequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (w WorkflowDeploymentsListRequestStatus) Ptr() *WorkflowDeploymentsListRequestStatus {
	return &w
}