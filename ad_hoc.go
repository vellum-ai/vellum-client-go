// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/vellum-ai/vellum-client-go/core"
)

type AdHocExecutePrompt struct {
	MlModel        string                `json:"ml_model" url:"-"`
	InputValues    []*PromptRequestInput `json:"input_values,omitempty" url:"-"`
	InputVariables []*VellumVariable     `json:"input_variables,omitempty" url:"-"`
	Parameters     *PromptParameters     `json:"parameters,omitempty" url:"-"`
	Settings       *PromptSettings       `json:"settings,omitempty" url:"-"`
	Blocks         []*PromptBlock        `json:"blocks,omitempty" url:"-"`
	Functions      []*FunctionDefinition `json:"functions,omitempty" url:"-"`
	ExpandMeta     *AdHocExpandMeta      `json:"expand_meta,omitempty" url:"-"`
}

type AdHocExecutePromptStream struct {
	MlModel        string                `json:"ml_model" url:"-"`
	InputValues    []*PromptRequestInput `json:"input_values,omitempty" url:"-"`
	InputVariables []*VellumVariable     `json:"input_variables,omitempty" url:"-"`
	Parameters     *PromptParameters     `json:"parameters,omitempty" url:"-"`
	Settings       *PromptSettings       `json:"settings,omitempty" url:"-"`
	Blocks         []*PromptBlock        `json:"blocks,omitempty" url:"-"`
	Functions      []*FunctionDefinition `json:"functions,omitempty" url:"-"`
	ExpandMeta     *AdHocExpandMeta      `json:"expand_meta,omitempty" url:"-"`
}

type AdHocExecutePromptEvent struct {
	InitiatedAdHocExecutePromptEvent *InitiatedAdHocExecutePromptEvent
	StreamingAdHocExecutePromptEvent *StreamingAdHocExecutePromptEvent
	FulfilledAdHocExecutePromptEvent *FulfilledAdHocExecutePromptEvent
	RejectedAdHocExecutePromptEvent  *RejectedAdHocExecutePromptEvent
}

func (a *AdHocExecutePromptEvent) UnmarshalJSON(data []byte) error {
	valueInitiatedAdHocExecutePromptEvent := new(InitiatedAdHocExecutePromptEvent)
	if err := json.Unmarshal(data, &valueInitiatedAdHocExecutePromptEvent); err == nil {
		a.InitiatedAdHocExecutePromptEvent = valueInitiatedAdHocExecutePromptEvent
		return nil
	}
	valueStreamingAdHocExecutePromptEvent := new(StreamingAdHocExecutePromptEvent)
	if err := json.Unmarshal(data, &valueStreamingAdHocExecutePromptEvent); err == nil {
		a.StreamingAdHocExecutePromptEvent = valueStreamingAdHocExecutePromptEvent
		return nil
	}
	valueFulfilledAdHocExecutePromptEvent := new(FulfilledAdHocExecutePromptEvent)
	if err := json.Unmarshal(data, &valueFulfilledAdHocExecutePromptEvent); err == nil {
		a.FulfilledAdHocExecutePromptEvent = valueFulfilledAdHocExecutePromptEvent
		return nil
	}
	valueRejectedAdHocExecutePromptEvent := new(RejectedAdHocExecutePromptEvent)
	if err := json.Unmarshal(data, &valueRejectedAdHocExecutePromptEvent); err == nil {
		a.RejectedAdHocExecutePromptEvent = valueRejectedAdHocExecutePromptEvent
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AdHocExecutePromptEvent) MarshalJSON() ([]byte, error) {
	if a.InitiatedAdHocExecutePromptEvent != nil {
		return json.Marshal(a.InitiatedAdHocExecutePromptEvent)
	}
	if a.StreamingAdHocExecutePromptEvent != nil {
		return json.Marshal(a.StreamingAdHocExecutePromptEvent)
	}
	if a.FulfilledAdHocExecutePromptEvent != nil {
		return json.Marshal(a.FulfilledAdHocExecutePromptEvent)
	}
	if a.RejectedAdHocExecutePromptEvent != nil {
		return json.Marshal(a.RejectedAdHocExecutePromptEvent)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AdHocExecutePromptEventVisitor interface {
	VisitInitiatedAdHocExecutePromptEvent(*InitiatedAdHocExecutePromptEvent) error
	VisitStreamingAdHocExecutePromptEvent(*StreamingAdHocExecutePromptEvent) error
	VisitFulfilledAdHocExecutePromptEvent(*FulfilledAdHocExecutePromptEvent) error
	VisitRejectedAdHocExecutePromptEvent(*RejectedAdHocExecutePromptEvent) error
}

func (a *AdHocExecutePromptEvent) Accept(visitor AdHocExecutePromptEventVisitor) error {
	if a.InitiatedAdHocExecutePromptEvent != nil {
		return visitor.VisitInitiatedAdHocExecutePromptEvent(a.InitiatedAdHocExecutePromptEvent)
	}
	if a.StreamingAdHocExecutePromptEvent != nil {
		return visitor.VisitStreamingAdHocExecutePromptEvent(a.StreamingAdHocExecutePromptEvent)
	}
	if a.FulfilledAdHocExecutePromptEvent != nil {
		return visitor.VisitFulfilledAdHocExecutePromptEvent(a.FulfilledAdHocExecutePromptEvent)
	}
	if a.RejectedAdHocExecutePromptEvent != nil {
		return visitor.VisitRejectedAdHocExecutePromptEvent(a.RejectedAdHocExecutePromptEvent)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AdHocExpandMeta struct {
	// If enabled, the response will include model host cost tracking. This may increase latency for some model hosts.
	Cost *bool `json:"cost,omitempty" url:"cost,omitempty"`
	// If enabled, the response will include the model identifier representing the ML Model invoked by the Prompt.
	ModelName *bool `json:"model_name,omitempty" url:"model_name,omitempty"`
	// If enabled, the response will include model host usage tracking. This may increase latency for some model hosts.
	Usage *bool `json:"usage,omitempty" url:"usage,omitempty"`
	// If enabled, the response will include the reason provided by the model for why the execution finished.
	FinishReason *bool `json:"finish_reason,omitempty" url:"finish_reason,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AdHocExpandMeta) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AdHocExpandMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler AdHocExpandMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AdHocExpandMeta(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AdHocExpandMeta) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The subset of the metadata tracked by Vellum during prompt execution that the request opted into with `expand_meta`.
type AdHocFulfilledPromptExecutionMeta struct {
	Latency      *int              `json:"latency,omitempty" url:"latency,omitempty"`
	FinishReason *FinishReasonEnum `json:"finish_reason,omitempty" url:"finish_reason,omitempty"`
	Usage        *MlModelUsage     `json:"usage,omitempty" url:"usage,omitempty"`
	Cost         *Price            `json:"cost,omitempty" url:"cost,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AdHocFulfilledPromptExecutionMeta) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AdHocFulfilledPromptExecutionMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler AdHocFulfilledPromptExecutionMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AdHocFulfilledPromptExecutionMeta(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AdHocFulfilledPromptExecutionMeta) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The subset of the metadata tracked by Vellum during prompt execution that the request opted into with `expand_meta`.
type AdHocInitiatedPromptExecutionMeta struct {
	ModelName *string `json:"model_name,omitempty" url:"model_name,omitempty"`
	Latency   *int    `json:"latency,omitempty" url:"latency,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AdHocInitiatedPromptExecutionMeta) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AdHocInitiatedPromptExecutionMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler AdHocInitiatedPromptExecutionMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AdHocInitiatedPromptExecutionMeta(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AdHocInitiatedPromptExecutionMeta) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The subset of the metadata tracked by Vellum during prompt execution that the request opted into with `expand_meta`.
type AdHocRejectedPromptExecutionMeta struct {
	Latency      *int              `json:"latency,omitempty" url:"latency,omitempty"`
	FinishReason *FinishReasonEnum `json:"finish_reason,omitempty" url:"finish_reason,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AdHocRejectedPromptExecutionMeta) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AdHocRejectedPromptExecutionMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler AdHocRejectedPromptExecutionMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AdHocRejectedPromptExecutionMeta(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AdHocRejectedPromptExecutionMeta) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The subset of the metadata tracked by Vellum during prompt execution that the request opted into with `expand_meta`.
type AdHocStreamingPromptExecutionMeta struct {
	Latency *int `json:"latency,omitempty" url:"latency,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AdHocStreamingPromptExecutionMeta) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AdHocStreamingPromptExecutionMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler AdHocStreamingPromptExecutionMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AdHocStreamingPromptExecutionMeta(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AdHocStreamingPromptExecutionMeta) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The final data event returned indicating that the stream has ended and all final resolved values from the model can be found.
type FulfilledAdHocExecutePromptEvent struct {
	Outputs     []*PromptOutput                    `json:"outputs" url:"outputs"`
	ExecutionId string                             `json:"execution_id" url:"execution_id"`
	Meta        *AdHocFulfilledPromptExecutionMeta `json:"meta,omitempty" url:"meta,omitempty"`
	state       string

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (f *FulfilledAdHocExecutePromptEvent) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FulfilledAdHocExecutePromptEvent) State() string {
	return f.state
}

func (f *FulfilledAdHocExecutePromptEvent) UnmarshalJSON(data []byte) error {
	type embed FulfilledAdHocExecutePromptEvent
	var unmarshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*f),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*f = FulfilledAdHocExecutePromptEvent(unmarshaler.embed)
	if unmarshaler.State != "FULFILLED" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", f, "FULFILLED", unmarshaler.State)
	}
	f.state = unmarshaler.State

	extraProperties, err := core.ExtractExtraProperties(data, *f, "state")
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties

	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FulfilledAdHocExecutePromptEvent) MarshalJSON() ([]byte, error) {
	type embed FulfilledAdHocExecutePromptEvent
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*f),
		State: "FULFILLED",
	}
	return json.Marshal(marshaler)
}

func (f *FulfilledAdHocExecutePromptEvent) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

// The initial data returned indicating that the response from the model has returned and begun streaming.
type InitiatedAdHocExecutePromptEvent struct {
	Meta        *AdHocInitiatedPromptExecutionMeta `json:"meta,omitempty" url:"meta,omitempty"`
	ExecutionId string                             `json:"execution_id" url:"execution_id"`
	state       string

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *InitiatedAdHocExecutePromptEvent) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *InitiatedAdHocExecutePromptEvent) State() string {
	return i.state
}

func (i *InitiatedAdHocExecutePromptEvent) UnmarshalJSON(data []byte) error {
	type embed InitiatedAdHocExecutePromptEvent
	var unmarshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*i),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*i = InitiatedAdHocExecutePromptEvent(unmarshaler.embed)
	if unmarshaler.State != "INITIATED" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", i, "INITIATED", unmarshaler.State)
	}
	i.state = unmarshaler.State

	extraProperties, err := core.ExtractExtraProperties(data, *i, "state")
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *InitiatedAdHocExecutePromptEvent) MarshalJSON() ([]byte, error) {
	type embed InitiatedAdHocExecutePromptEvent
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*i),
		State: "INITIATED",
	}
	return json.Marshal(marshaler)
}

func (i *InitiatedAdHocExecutePromptEvent) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type PromptRequestChatHistoryInput struct {
	// The variable's name, as defined in the Prompt.
	Key   string         `json:"key" url:"key"`
	Value []*ChatMessage `json:"value" url:"value"`
	type_ string

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PromptRequestChatHistoryInput) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PromptRequestChatHistoryInput) Type() string {
	return p.type_
}

func (p *PromptRequestChatHistoryInput) UnmarshalJSON(data []byte) error {
	type embed PromptRequestChatHistoryInput
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PromptRequestChatHistoryInput(unmarshaler.embed)
	if unmarshaler.Type != "CHAT_HISTORY" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "CHAT_HISTORY", unmarshaler.Type)
	}
	p.type_ = unmarshaler.Type

	extraProperties, err := core.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PromptRequestChatHistoryInput) MarshalJSON() ([]byte, error) {
	type embed PromptRequestChatHistoryInput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "CHAT_HISTORY",
	}
	return json.Marshal(marshaler)
}

func (p *PromptRequestChatHistoryInput) String() string {
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

type PromptRequestInput struct {
	PromptRequestStringInput      *PromptRequestStringInput
	PromptRequestJsonInput        *PromptRequestJsonInput
	PromptRequestChatHistoryInput *PromptRequestChatHistoryInput
}

func (p *PromptRequestInput) UnmarshalJSON(data []byte) error {
	valuePromptRequestStringInput := new(PromptRequestStringInput)
	if err := json.Unmarshal(data, &valuePromptRequestStringInput); err == nil {
		p.PromptRequestStringInput = valuePromptRequestStringInput
		return nil
	}
	valuePromptRequestJsonInput := new(PromptRequestJsonInput)
	if err := json.Unmarshal(data, &valuePromptRequestJsonInput); err == nil {
		p.PromptRequestJsonInput = valuePromptRequestJsonInput
		return nil
	}
	valuePromptRequestChatHistoryInput := new(PromptRequestChatHistoryInput)
	if err := json.Unmarshal(data, &valuePromptRequestChatHistoryInput); err == nil {
		p.PromptRequestChatHistoryInput = valuePromptRequestChatHistoryInput
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, p)
}

func (p PromptRequestInput) MarshalJSON() ([]byte, error) {
	if p.PromptRequestStringInput != nil {
		return json.Marshal(p.PromptRequestStringInput)
	}
	if p.PromptRequestJsonInput != nil {
		return json.Marshal(p.PromptRequestJsonInput)
	}
	if p.PromptRequestChatHistoryInput != nil {
		return json.Marshal(p.PromptRequestChatHistoryInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", p)
}

type PromptRequestInputVisitor interface {
	VisitPromptRequestStringInput(*PromptRequestStringInput) error
	VisitPromptRequestJsonInput(*PromptRequestJsonInput) error
	VisitPromptRequestChatHistoryInput(*PromptRequestChatHistoryInput) error
}

func (p *PromptRequestInput) Accept(visitor PromptRequestInputVisitor) error {
	if p.PromptRequestStringInput != nil {
		return visitor.VisitPromptRequestStringInput(p.PromptRequestStringInput)
	}
	if p.PromptRequestJsonInput != nil {
		return visitor.VisitPromptRequestJsonInput(p.PromptRequestJsonInput)
	}
	if p.PromptRequestChatHistoryInput != nil {
		return visitor.VisitPromptRequestChatHistoryInput(p.PromptRequestChatHistoryInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", p)
}

type PromptRequestJsonInput struct {
	// The variable's name, as defined in the Prompt.
	Key   string      `json:"key" url:"key"`
	Value interface{} `json:"value" url:"value"`
	type_ string

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PromptRequestJsonInput) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PromptRequestJsonInput) Type() string {
	return p.type_
}

func (p *PromptRequestJsonInput) UnmarshalJSON(data []byte) error {
	type embed PromptRequestJsonInput
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PromptRequestJsonInput(unmarshaler.embed)
	if unmarshaler.Type != "JSON" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "JSON", unmarshaler.Type)
	}
	p.type_ = unmarshaler.Type

	extraProperties, err := core.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PromptRequestJsonInput) MarshalJSON() ([]byte, error) {
	type embed PromptRequestJsonInput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "JSON",
	}
	return json.Marshal(marshaler)
}

func (p *PromptRequestJsonInput) String() string {
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

type PromptRequestStringInput struct {
	// The variable's name, as defined in the Prompt.
	Key   string `json:"key" url:"key"`
	Value string `json:"value" url:"value"`
	type_ string

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PromptRequestStringInput) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PromptRequestStringInput) Type() string {
	return p.type_
}

func (p *PromptRequestStringInput) UnmarshalJSON(data []byte) error {
	type embed PromptRequestStringInput
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PromptRequestStringInput(unmarshaler.embed)
	if unmarshaler.Type != "STRING" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "STRING", unmarshaler.Type)
	}
	p.type_ = unmarshaler.Type

	extraProperties, err := core.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PromptRequestStringInput) MarshalJSON() ([]byte, error) {
	type embed PromptRequestStringInput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "STRING",
	}
	return json.Marshal(marshaler)
}

func (p *PromptRequestStringInput) String() string {
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

// The final data returned indicating an error occurred during the stream.
type RejectedAdHocExecutePromptEvent struct {
	Error       *VellumError                      `json:"error" url:"error"`
	ExecutionId string                            `json:"execution_id" url:"execution_id"`
	Meta        *AdHocRejectedPromptExecutionMeta `json:"meta,omitempty" url:"meta,omitempty"`
	state       string

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RejectedAdHocExecutePromptEvent) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RejectedAdHocExecutePromptEvent) State() string {
	return r.state
}

func (r *RejectedAdHocExecutePromptEvent) UnmarshalJSON(data []byte) error {
	type embed RejectedAdHocExecutePromptEvent
	var unmarshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*r),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*r = RejectedAdHocExecutePromptEvent(unmarshaler.embed)
	if unmarshaler.State != "REJECTED" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", r, "REJECTED", unmarshaler.State)
	}
	r.state = unmarshaler.State

	extraProperties, err := core.ExtractExtraProperties(data, *r, "state")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RejectedAdHocExecutePromptEvent) MarshalJSON() ([]byte, error) {
	type embed RejectedAdHocExecutePromptEvent
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*r),
		State: "REJECTED",
	}
	return json.Marshal(marshaler)
}

func (r *RejectedAdHocExecutePromptEvent) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// The data returned for each delta during the prompt execution stream.
type StreamingAdHocExecutePromptEvent struct {
	Output      *PromptOutput                      `json:"output" url:"output"`
	OutputIndex int                                `json:"output_index" url:"output_index"`
	ExecutionId string                             `json:"execution_id" url:"execution_id"`
	Meta        *AdHocStreamingPromptExecutionMeta `json:"meta,omitempty" url:"meta,omitempty"`
	// The subset of the raw response from the model that the request opted into with `expand_raw`.
	Raw   map[string]interface{} `json:"raw,omitempty" url:"raw,omitempty"`
	state string

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *StreamingAdHocExecutePromptEvent) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *StreamingAdHocExecutePromptEvent) State() string {
	return s.state
}

func (s *StreamingAdHocExecutePromptEvent) UnmarshalJSON(data []byte) error {
	type embed StreamingAdHocExecutePromptEvent
	var unmarshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = StreamingAdHocExecutePromptEvent(unmarshaler.embed)
	if unmarshaler.State != "STREAMING" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "STREAMING", unmarshaler.State)
	}
	s.state = unmarshaler.State

	extraProperties, err := core.ExtractExtraProperties(data, *s, "state")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *StreamingAdHocExecutePromptEvent) MarshalJSON() ([]byte, error) {
	type embed StreamingAdHocExecutePromptEvent
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*s),
		State: "STREAMING",
	}
	return json.Marshal(marshaler)
}

func (s *StreamingAdHocExecutePromptEvent) String() string {
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
