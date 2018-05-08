/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slack

// Schema for successful response from im.open method
type ImOpenSuccessSchema struct {
	NoOp bool `json:"no_op,omitempty"`

	AlreadyOpen bool `json:"already_open,omitempty"`

	Ok bool `json:"ok"`

	Channel *ImOpenSuccessSchemaChannel `json:"channel"`
}