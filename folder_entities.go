// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
)

type AddEntityToFolderRequest struct {
	// The ID of the entity you would like to move.
	EntityId string `json:"entity_id" url:"-"`
}

type FolderEntitiesListRequest struct {
	// Filter down to only those objects whose entities have a status matching the status specified.
	//
	// - `ACTIVE` - Active
	// - `ARCHIVED` - Archived
	EntityStatus *FolderEntitiesListRequestEntityStatus `json:"-" url:"entity_status,omitempty"`
	// Number of results to return per page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The initial index from which to return the results.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Which field to use when ordering the results.
	Ordering *string `json:"-" url:"ordering,omitempty"`
	// Filter down to only those entities whose parent folder has the specified ID.
	//
	// To filter by an entity's parent folder, provide the ID of the parent folder. To filter by the root directory, provide
	// a string representing the entity type of the root directory (e.g. "PROMPT_SANDBOX" or "TEST_SUITE").
	ParentFolderId string `json:"-" url:"parent_folder_id"`
}

type FolderEntitiesListRequestEntityStatus string

const (
	FolderEntitiesListRequestEntityStatusActive   FolderEntitiesListRequestEntityStatus = "ACTIVE"
	FolderEntitiesListRequestEntityStatusArchived FolderEntitiesListRequestEntityStatus = "ARCHIVED"
)

func NewFolderEntitiesListRequestEntityStatusFromString(s string) (FolderEntitiesListRequestEntityStatus, error) {
	switch s {
	case "ACTIVE":
		return FolderEntitiesListRequestEntityStatusActive, nil
	case "ARCHIVED":
		return FolderEntitiesListRequestEntityStatusArchived, nil
	}
	var t FolderEntitiesListRequestEntityStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (f FolderEntitiesListRequestEntityStatus) Ptr() *FolderEntitiesListRequestEntityStatus {
	return &f
}
