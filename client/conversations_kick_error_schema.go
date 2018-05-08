/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slack

// Schema for error response conversations.kick method
type ConversationsKickErrorSchema struct {
	Needed string `json:"needed,omitempty"`

	Error_ string `json:"error"`

	Ok bool `json:"ok"`

	Provided string `json:"provided,omitempty"`
}
