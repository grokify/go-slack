/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slack

// Schema for successful response from files.info method
type FilesInfoSchema struct {
	Paging *ObjsPaging `json:"paging"`

	Ok bool `json:"ok"`

	File *ObjsFile `json:"file"`

	Comments *ObjsComments `json:"comments"`
}
