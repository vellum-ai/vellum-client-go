// This file was auto-generated by Fern from our API Definition.

package api

type DeploySandboxPromptRequest struct {
	// The Vellum-generated ID of the Prompt Deployment you'd like to update. Cannot specify both this and prompt_deployment_name. Leave null to create a new Prompt Deployment.
	PromptDeploymentId *string `json:"prompt_deployment_id,omitempty"`
	// The unique name of the Prompt Deployment you'd like to either create or update. Cannot specify both this and prompt_deployment_id. If provided and matches an existing Prompt Deployment, that Prompt Deployment will be updated. Otherwise, a new Prompt Deployment will be created.
	PromptDeploymentName *string `json:"prompt_deployment_name,omitempty"`
	// In the event that a new Prompt Deployment is created, this will be the label it's given.
	Label *string `json:"label,omitempty"`
	// Optionally provide the release tags that you'd like to be associated with the latest release of the created/updated Prompt Deployment.
	ReleaseTags []string `json:"release_tags,omitempty"`
}
