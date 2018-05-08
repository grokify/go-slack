/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slack

type ObjsChannel struct {
	IsGeneral bool `json:"is_general,omitempty"`

	NameNormalized string `json:"name_normalized"`

	LastRead string `json:"last_read,omitempty"`

	Creator string `json:"creator"`

	IsMember bool `json:"is_member,omitempty"`

	IsArchived bool `json:"is_archived,omitempty"`

	Topic *ObjsChannelTopic `json:"topic"`

	UnreadCountDisplay int32 `json:"unread_count_display,omitempty"`

	Id string `json:"id"`

	IsOrgShared bool `json:"is_org_shared"`

	IsChannel bool `json:"is_channel"`

	Name string `json:"name"`

	Priority int32 `json:"priority,omitempty"`

	IsMoved int32 `json:"is_moved,omitempty"`

	AcceptedUser string `json:"accepted_user,omitempty"`

	IsPendingExtShared bool `json:"is_pending_ext_shared,omitempty"`

	IsMpim bool `json:"is_mpim"`

	IsReadOnly bool `json:"is_read_only,omitempty"`

	Purpose *ObjsChannelTopic `json:"purpose"`

	Members []string `json:"members"`

	IsPrivate bool `json:"is_private"`

	PreviousNames []string `json:"previous_names,omitempty"`

	NumMembers int32 `json:"num_members,omitempty"`

	IsShared bool `json:"is_shared"`

	Created int32 `json:"created"`

	PendingShared []string `json:"pending_shared,omitempty"`

	UnreadCount int32 `json:"unread_count,omitempty"`

	Unlinked int32 `json:"unlinked,omitempty"`
}