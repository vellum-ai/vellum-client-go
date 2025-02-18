// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/vellum-ai/vellum-client-go/core"
	time "time"
)

type DocumentsListRequest struct {
	// Filter down to only those documents that are included in the specified index. You may provide either the Vellum-generated ID or the unique name of the index specified upon initial creation.
	DocumentIndexId *string `json:"-" url:"document_index_id,omitempty"`
	// Number of results to return per page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The initial index from which to return the results.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Which field to use when ordering the results.
	Ordering *string `json:"-" url:"ordering,omitempty"`
	// A search term.
	Search *string `json:"-" url:"search,omitempty"`
}

type PatchedDocumentUpdateRequest struct {
	// A human-readable label for the document. Defaults to the originally uploaded file's file name.
	Label *string `json:"label,omitempty" url:"-"`
	// The current status of the document
	//
	// * `ACTIVE` - Active
	Status *DocumentStatus `json:"status,omitempty" url:"-"`
	// A JSON object containing any metadata associated with the document that you'd like to filter upon later.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
}

// A detailed representation of the link between a Document and a Document Index it's a member of.
type DocumentDocumentToDocumentIndex struct {
	// Vellum-generated ID that uniquely identifies this link.
	Id string `json:"id" url:"id"`
	// Vellum-generated ID that uniquely identifies the index this document is included in.
	DocumentIndexId string `json:"document_index_id" url:"document_index_id"`
	// An enum value representing where this document is along its indexing lifecycle for this index.
	//
	// * `AWAITING_PROCESSING` - Awaiting Processing
	// * `QUEUED` - Queued
	// * `INDEXING` - Indexing
	// * `INDEXED` - Indexed
	// * `FAILED` - Failed
	IndexingState        *IndexingStateEnum `json:"indexing_state,omitempty" url:"indexing_state,omitempty"`
	ExtractedTextFileUrl *string            `json:"extracted_text_file_url,omitempty" url:"extracted_text_file_url,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DocumentDocumentToDocumentIndex) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DocumentDocumentToDocumentIndex) UnmarshalJSON(data []byte) error {
	type unmarshaler DocumentDocumentToDocumentIndex
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DocumentDocumentToDocumentIndex(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DocumentDocumentToDocumentIndex) String() string {
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

// * `QUEUED` - Queued
// * `PROCESSING` - Processing
// * `PROCESSED` - Processed
// * `FAILED` - Failed
// * `UNKNOWN` - Unknown
type DocumentProcessingState string

const (
	DocumentProcessingStateQueued     DocumentProcessingState = "QUEUED"
	DocumentProcessingStateProcessing DocumentProcessingState = "PROCESSING"
	DocumentProcessingStateProcessed  DocumentProcessingState = "PROCESSED"
	DocumentProcessingStateFailed     DocumentProcessingState = "FAILED"
	DocumentProcessingStateUnknown    DocumentProcessingState = "UNKNOWN"
)

func NewDocumentProcessingStateFromString(s string) (DocumentProcessingState, error) {
	switch s {
	case "QUEUED":
		return DocumentProcessingStateQueued, nil
	case "PROCESSING":
		return DocumentProcessingStateProcessing, nil
	case "PROCESSED":
		return DocumentProcessingStateProcessed, nil
	case "FAILED":
		return DocumentProcessingStateFailed, nil
	case "UNKNOWN":
		return DocumentProcessingStateUnknown, nil
	}
	var t DocumentProcessingState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DocumentProcessingState) Ptr() *DocumentProcessingState {
	return &d
}

type DocumentRead struct {
	Id string `json:"id" url:"id"`
	// The unique id of this document as it exists in the user's system.
	ExternalId     *string   `json:"external_id,omitempty" url:"external_id,omitempty"`
	LastUploadedAt time.Time `json:"last_uploaded_at" url:"last_uploaded_at"`
	// A human-readable label for the document. Defaults to the originally uploaded file's file name.
	Label           string                  `json:"label" url:"label"`
	ProcessingState DocumentProcessingState `json:"processing_state" url:"processing_state"`
	// The current status of the document
	//
	// * `ACTIVE` - Active
	Status                    *DocumentStatus                    `json:"status,omitempty" url:"status,omitempty"`
	OriginalFileUrl           *string                            `json:"original_file_url,omitempty" url:"original_file_url,omitempty"`
	ProcessedFileUrl          *string                            `json:"processed_file_url,omitempty" url:"processed_file_url,omitempty"`
	DocumentToDocumentIndexes []*DocumentDocumentToDocumentIndex `json:"document_to_document_indexes" url:"document_to_document_indexes"`
	// A previously supplied JSON object containing metadata that can be filtered on when searching.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DocumentRead) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DocumentRead) UnmarshalJSON(data []byte) error {
	type embed DocumentRead
	var unmarshaler = struct {
		embed
		LastUploadedAt *core.DateTime `json:"last_uploaded_at"`
	}{
		embed: embed(*d),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*d = DocumentRead(unmarshaler.embed)
	d.LastUploadedAt = unmarshaler.LastUploadedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DocumentRead) MarshalJSON() ([]byte, error) {
	type embed DocumentRead
	var marshaler = struct {
		embed
		LastUploadedAt *core.DateTime `json:"last_uploaded_at"`
	}{
		embed:          embed(*d),
		LastUploadedAt: core.NewDateTime(d.LastUploadedAt),
	}
	return json.Marshal(marshaler)
}

func (d *DocumentRead) String() string {
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

// * `ACTIVE` - Active
type DocumentStatus = string

// * `AWAITING_PROCESSING` - Awaiting Processing
// * `QUEUED` - Queued
// * `INDEXING` - Indexing
// * `INDEXED` - Indexed
// * `FAILED` - Failed
type IndexingStateEnum string

const (
	IndexingStateEnumAwaitingProcessing IndexingStateEnum = "AWAITING_PROCESSING"
	IndexingStateEnumQueued             IndexingStateEnum = "QUEUED"
	IndexingStateEnumIndexing           IndexingStateEnum = "INDEXING"
	IndexingStateEnumIndexed            IndexingStateEnum = "INDEXED"
	IndexingStateEnumFailed             IndexingStateEnum = "FAILED"
)

func NewIndexingStateEnumFromString(s string) (IndexingStateEnum, error) {
	switch s {
	case "AWAITING_PROCESSING":
		return IndexingStateEnumAwaitingProcessing, nil
	case "QUEUED":
		return IndexingStateEnumQueued, nil
	case "INDEXING":
		return IndexingStateEnumIndexing, nil
	case "INDEXED":
		return IndexingStateEnumIndexed, nil
	case "FAILED":
		return IndexingStateEnumFailed, nil
	}
	var t IndexingStateEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i IndexingStateEnum) Ptr() *IndexingStateEnum {
	return &i
}

type PaginatedSlimDocumentList struct {
	Count    *int            `json:"count,omitempty" url:"count,omitempty"`
	Next     *string         `json:"next,omitempty" url:"next,omitempty"`
	Previous *string         `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*SlimDocument `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PaginatedSlimDocumentList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedSlimDocumentList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedSlimDocumentList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedSlimDocumentList(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedSlimDocumentList) String() string {
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

// * `EXCEEDED_CHARACTER_LIMIT` - Exceeded Character Limit
// * `INVALID_FILE` - Invalid File
type ProcessingFailureReasonEnum string

const (
	ProcessingFailureReasonEnumExceededCharacterLimit ProcessingFailureReasonEnum = "EXCEEDED_CHARACTER_LIMIT"
	ProcessingFailureReasonEnumInvalidFile            ProcessingFailureReasonEnum = "INVALID_FILE"
)

func NewProcessingFailureReasonEnumFromString(s string) (ProcessingFailureReasonEnum, error) {
	switch s {
	case "EXCEEDED_CHARACTER_LIMIT":
		return ProcessingFailureReasonEnumExceededCharacterLimit, nil
	case "INVALID_FILE":
		return ProcessingFailureReasonEnumInvalidFile, nil
	}
	var t ProcessingFailureReasonEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p ProcessingFailureReasonEnum) Ptr() *ProcessingFailureReasonEnum {
	return &p
}

type SlimDocument struct {
	// Vellum-generated ID that uniquely identifies this document.
	Id string `json:"id" url:"id"`
	// The external ID that was originally provided when uploading the document.
	ExternalId *string `json:"external_id,omitempty" url:"external_id,omitempty"`
	// A timestamp representing when this document was most recently uploaded.
	LastUploadedAt time.Time `json:"last_uploaded_at" url:"last_uploaded_at"`
	// Human-friendly name for this document.
	Label           string                  `json:"label" url:"label"`
	ProcessingState DocumentProcessingState `json:"processing_state" url:"processing_state"`
	// An enum value representing why the document could not be processed. Is null unless processing_state is FAILED.
	//
	// * `EXCEEDED_CHARACTER_LIMIT` - Exceeded Character Limit
	// * `INVALID_FILE` - Invalid File
	ProcessingFailureReason *ProcessingFailureReasonEnum `json:"processing_failure_reason,omitempty" url:"processing_failure_reason,omitempty"`
	// The document's current status.
	//
	// * `ACTIVE` - Active
	Status *DocumentStatus `json:"status,omitempty" url:"status,omitempty"`
	// A list of keywords associated with this document. Originally provided when uploading the document.
	Keywords []string `json:"keywords,omitempty" url:"keywords,omitempty"`
	// A previously supplied JSON object containing metadata that can be filtered on when searching.
	Metadata                  map[string]interface{}                 `json:"metadata,omitempty" url:"metadata,omitempty"`
	DocumentToDocumentIndexes []*SlimDocumentDocumentToDocumentIndex `json:"document_to_document_indexes" url:"document_to_document_indexes"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SlimDocument) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SlimDocument) UnmarshalJSON(data []byte) error {
	type embed SlimDocument
	var unmarshaler = struct {
		embed
		LastUploadedAt *core.DateTime `json:"last_uploaded_at"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SlimDocument(unmarshaler.embed)
	s.LastUploadedAt = unmarshaler.LastUploadedAt.Time()

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SlimDocument) MarshalJSON() ([]byte, error) {
	type embed SlimDocument
	var marshaler = struct {
		embed
		LastUploadedAt *core.DateTime `json:"last_uploaded_at"`
	}{
		embed:          embed(*s),
		LastUploadedAt: core.NewDateTime(s.LastUploadedAt),
	}
	return json.Marshal(marshaler)
}

func (s *SlimDocument) String() string {
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

// A slim representation of the link between a Document and a Document Index it's a member of.
type SlimDocumentDocumentToDocumentIndex struct {
	// Vellum-generated ID that uniquely identifies this link.
	Id string `json:"id" url:"id"`
	// Vellum-generated ID that uniquely identifies the index this document is included in.
	DocumentIndexId string `json:"document_index_id" url:"document_index_id"`
	// An enum value representing where this document is along its indexing lifecycle for this index.
	//
	// * `AWAITING_PROCESSING` - Awaiting Processing
	// * `QUEUED` - Queued
	// * `INDEXING` - Indexing
	// * `INDEXED` - Indexed
	// * `FAILED` - Failed
	IndexingState *IndexingStateEnum `json:"indexing_state,omitempty" url:"indexing_state,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SlimDocumentDocumentToDocumentIndex) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SlimDocumentDocumentToDocumentIndex) UnmarshalJSON(data []byte) error {
	type unmarshaler SlimDocumentDocumentToDocumentIndex
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SlimDocumentDocumentToDocumentIndex(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SlimDocumentDocumentToDocumentIndex) String() string {
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

type UploadDocumentResponse struct {
	// The ID of the newly created document.
	DocumentId string `json:"document_id" url:"document_id"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UploadDocumentResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UploadDocumentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UploadDocumentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UploadDocumentResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UploadDocumentResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UploadDocumentBodyRequest struct {
	// Optionally include the names of all indexes that you'd like this document to be included in
	AddToIndexNames []string `json:"add_to_index_names,omitempty" url:"-"`
	// Optionally include an external ID for this document. This is useful if you want to re-upload the same document later when its contents change and would like it to be re-indexed.
	ExternalId *string `json:"external_id,omitempty" url:"-"`
	// A human-friendly name for this document. Typically the filename.
	Label string `json:"label" url:"-"`
	// Optionally include a list of keywords that'll be associated with this document. Used when performing keyword searches.
	Keywords []string `json:"keywords,omitempty" url:"-"`
	// A stringified JSON object containing any metadata associated with the document that you'd like to filter upon later.
	Metadata *string `json:"metadata,omitempty" url:"-"`
}
