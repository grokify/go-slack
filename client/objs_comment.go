/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slack

type ObjsComment struct {
	Comment string `json:"comment,omitempty"`

	Reactions []ObjsReaction `json:"reactions,omitempty"`

	Created int32 `json:"created,omitempty"`

	Timestamp int32 `json:"timestamp,omitempty"`

	PinnedTo []string `json:"pinned_to,omitempty"`

	IsIntro bool `json:"is_intro,omitempty"`

	User string `json:"user,omitempty"`

	Id string `json:"id,omitempty"`
}
