// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/vellum-ai/vellum-client-go/core"
)

type RegisterPromptRequestRequest struct {
	// A human-friendly label for corresponding entities created in Vellum.
	Label string `json:"label"`
	// A uniquely-identifying name for corresponding entities created in Vellum.
	Name string `json:"name"`
	// Information about how to execute the prompt template.
	Prompt *RegisterPromptPromptInfoRequest `json:"prompt,omitempty"`
	// The initial LLM provider to use for this prompt
	//
	// * `ANTHROPIC` - Anthropic
	// * `COHERE` - Cohere
	// * `GOOGLE` - Google
	// * `HOSTED` - Hosted
	// * `MOSAICML` - MosaicML
	// * `OPENAI` - OpenAI
	// * `HUGGINGFACE` - HuggingFace
	// * `MYSTIC` - Mystic
	// * `PYQ` - Pyq
	Provider ProviderEnum `json:"provider,omitempty"`
	// The initial model to use for this prompt
	Model string `json:"model"`
	// The initial model parameters to use for  this prompt
	Parameters *RegisterPromptModelParametersRequest `json:"parameters,omitempty"`
	// Optionally include additional metadata to store along with the prompt.
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// * `ANTHROPIC` - Anthropic
// * `COHERE` - Cohere
// * `GOOGLE` - Google
// * `HOSTED` - Hosted
// * `MOSAICML` - MosaicML
// * `OPENAI` - OpenAI
// * `HUGGINGFACE` - HuggingFace
// * `MYSTIC` - Mystic
// * `PYQ` - Pyq
type ProviderEnum string

const (
	ProviderEnumAnthropic   ProviderEnum = "ANTHROPIC"
	ProviderEnumCohere      ProviderEnum = "COHERE"
	ProviderEnumGoogle      ProviderEnum = "GOOGLE"
	ProviderEnumHosted      ProviderEnum = "HOSTED"
	ProviderEnumMosaicml    ProviderEnum = "MOSAICML"
	ProviderEnumOpenai      ProviderEnum = "OPENAI"
	ProviderEnumHuggingface ProviderEnum = "HUGGINGFACE"
	ProviderEnumMystic      ProviderEnum = "MYSTIC"
	ProviderEnumPyq         ProviderEnum = "PYQ"
)

func NewProviderEnumFromString(s string) (ProviderEnum, error) {
	switch s {
	case "ANTHROPIC":
		return ProviderEnumAnthropic, nil
	case "COHERE":
		return ProviderEnumCohere, nil
	case "GOOGLE":
		return ProviderEnumGoogle, nil
	case "HOSTED":
		return ProviderEnumHosted, nil
	case "MOSAICML":
		return ProviderEnumMosaicml, nil
	case "OPENAI":
		return ProviderEnumOpenai, nil
	case "HUGGINGFACE":
		return ProviderEnumHuggingface, nil
	case "MYSTIC":
		return ProviderEnumMystic, nil
	case "PYQ":
		return ProviderEnumPyq, nil
	}
	var t ProviderEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p ProviderEnum) Ptr() *ProviderEnum {
	return &p
}

type RegisterPromptModelParametersRequest struct {
	Temperature      float64             `json:"temperature"`
	MaxTokens        int                 `json:"max_tokens"`
	Stop             []string            `json:"stop,omitempty"`
	TopP             float64             `json:"top_p"`
	TopK             *float64            `json:"top_k,omitempty"`
	FrequencyPenalty float64             `json:"frequency_penalty"`
	PresencePenalty  float64             `json:"presence_penalty"`
	LogitBias        map[string]*float64 `json:"logit_bias,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RegisterPromptModelParametersRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler RegisterPromptModelParametersRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RegisterPromptModelParametersRequest(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RegisterPromptModelParametersRequest) String() string {
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

type RegisterPromptPromptInfoRequest struct {
	PromptSyntaxVersion *int                            `json:"prompt_syntax_version,omitempty"`
	PromptBlockData     *PromptTemplateBlockDataRequest `json:"prompt_block_data,omitempty"`
	// The input variables specified in the prompt template.
	InputVariables []*RegisteredPromptInputVariableRequest `json:"input_variables,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RegisterPromptPromptInfoRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler RegisterPromptPromptInfoRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RegisterPromptPromptInfoRequest(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RegisterPromptPromptInfoRequest) String() string {
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

type RegisterPromptResponse struct {
	// Information about the generated prompt
	Prompt *RegisterPromptPrompt `json:"prompt,omitempty"`
	// Information about the generated sandbox snapshot
	SandboxSnapshot *RegisteredPromptSandboxSnapshot `json:"sandbox_snapshot,omitempty"`
	// Information about the generated sandbox
	Sandbox *RegisteredPromptSandbox `json:"sandbox,omitempty"`
	// Information about the generated model version
	ModelVersion *RegisteredPromptModelVersion `json:"model_version,omitempty"`
	// Information about the generated deployment
	Deployment *RegisteredPromptDeployment `json:"deployment,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RegisterPromptResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RegisterPromptResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RegisterPromptResponse(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RegisterPromptResponse) String() string {
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
