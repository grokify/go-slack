/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slack

type ConversationsRepliesSuccessSchemaMessages struct {
	ThreadTs string `json:"thread_ts"`

	Subscribed bool `json:"subscribed,omitempty"`

	SourceTeam string `json:"source_team,omitempty"`

	LastRead string `json:"last_read,omitempty"`

	UserProfile *ObjsUserProfileShort `json:"user_profile,omitempty"`

	Text string `json:"text"`

	Team string `json:"team,omitempty"`

	Ts string `json:"ts"`

	UnreadCount int32 `json:"unread_count,omitempty"`

	ReplyCount int32 `json:"reply_count,omitempty"`

	User string `json:"user"`

	Replies []ConversationsRepliesSuccessSchemaReplies `json:"replies,omitempty"`

	Type_ string `json:"type"`

	UserTeam string `json:"user_team,omitempty"`

	ParentUserId string `json:"parent_user_id,omitempty"`

	IsStarred bool `json:"is_starred,omitempty"`
}
