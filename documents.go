// This file was auto-generated by Fern from our API Definition.

package api

type DocumentsListRequest struct {
	// Filter down to only those documents that are included in the specified index. You may provide either the Vellum-generated ID or the unique name of the index specified upon initial creation.
	DocumentIndexId *string `json:"-"`
	// Number of results to return per page.
	Limit *int `json:"-"`
	// The initial index from which to return the results.
	Offset *int `json:"-"`
	// Which field to use when ordering the results.
	Ordering *string `json:"-"`
}

type PatchedDocumentUpdateRequest struct {
	// A human-readable label for the document. Defaults to the originally uploaded file's file name.
	Label *string `json:"label,omitempty"`
	// The current status of the document
	//
	// * `ACTIVE` - Active
	Status *DocumentStatus `json:"status,omitempty"`
	// A JSON object containing any metadata associated with the document that you'd like to filter upon later.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type UploadDocumentBodyRequest struct {
	// Optionally include the names of all indexes that you'd like this document to be included in
	AddToIndexNames []string `json:"add_to_index_names,omitempty"`
	// Optionally include an external ID for this document. This is useful if you want to re-upload the same document later when its contents change and would like it to be re-indexed.
	ExternalId *string `json:"external_id,omitempty"`
	// A human-friendly name for this document. Typically the filename.
	Label string `json:"label"`
	// Optionally include a list of keywords that'll be associated with this document. Used when performing keyword searches.
	Keywords []string `json:"keywords,omitempty"`
	// A stringified JSON object containing any metadata associated with the document that you'd like to filter upon later.
	Metadata *string `json:"metadata,omitempty"`
}
