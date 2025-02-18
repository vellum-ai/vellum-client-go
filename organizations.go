// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/vellum-ai/vellum-client-go/core"
)

// * `AUTO_ACCEPT_FROM_SHARED_DOMAIN` - Auto-Accept from Shared Domain
// * `ALLOW_REQUESTS_FROM_SHARED_DOMAIN` - Allows Requests from Shared Domains
// * `REQUIRE_EXPLICIT_INVITE` - Require Explicit Invite
type NewMemberJoinBehaviorEnum string

const (
	NewMemberJoinBehaviorEnumAutoAcceptFromSharedDomain    NewMemberJoinBehaviorEnum = "AUTO_ACCEPT_FROM_SHARED_DOMAIN"
	NewMemberJoinBehaviorEnumAllowRequestsFromSharedDomain NewMemberJoinBehaviorEnum = "ALLOW_REQUESTS_FROM_SHARED_DOMAIN"
	NewMemberJoinBehaviorEnumRequireExplicitInvite         NewMemberJoinBehaviorEnum = "REQUIRE_EXPLICIT_INVITE"
)

func NewNewMemberJoinBehaviorEnumFromString(s string) (NewMemberJoinBehaviorEnum, error) {
	switch s {
	case "AUTO_ACCEPT_FROM_SHARED_DOMAIN":
		return NewMemberJoinBehaviorEnumAutoAcceptFromSharedDomain, nil
	case "ALLOW_REQUESTS_FROM_SHARED_DOMAIN":
		return NewMemberJoinBehaviorEnumAllowRequestsFromSharedDomain, nil
	case "REQUIRE_EXPLICIT_INVITE":
		return NewMemberJoinBehaviorEnumRequireExplicitInvite, nil
	}
	var t NewMemberJoinBehaviorEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (n NewMemberJoinBehaviorEnum) Ptr() *NewMemberJoinBehaviorEnum {
	return &n
}

type OrganizationRead struct {
	Id                    string                    `json:"id" url:"id"`
	Name                  string                    `json:"name" url:"name"`
	AllowStaffAccess      *bool                     `json:"allow_staff_access,omitempty" url:"allow_staff_access,omitempty"`
	NewMemberJoinBehavior NewMemberJoinBehaviorEnum `json:"new_member_join_behavior" url:"new_member_join_behavior"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (o *OrganizationRead) GetExtraProperties() map[string]interface{} {
	return o.extraProperties
}

func (o *OrganizationRead) UnmarshalJSON(data []byte) error {
	type unmarshaler OrganizationRead
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrganizationRead(value)

	extraProperties, err := core.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties

	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *OrganizationRead) String() string {
	if len(o._rawJSON) > 0 {
		if value, err := core.StringifyJSON(o._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}
