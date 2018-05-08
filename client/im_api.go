/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slack

import (
	"encoding/json"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ImApiService service

/* ImApiService
Close a direct message channel.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "token" (string) Authentication token. Requires scope: &#x60;im:write&#x60;
    @param "channel" (string) Direct message channel to close.
@return DefaultSuccessTemplate*/
func (a *ImApiService) ImClose(ctx context.Context, localVarOptionals map[string]interface{}) (DefaultSuccessTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     DefaultSuccessTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/im.close"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["channel"], "string", "channel"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarHeaderParams["token"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["channel"].(string); localVarOk {
		localVarFormParams.Add("channel", parameterToString(localVarTempParam, ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ImApiService
Fetches history of messages and events from direct message channel.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "count" (int32) Number of messages to return, between 1 and 1000.
    @param "unreads" (bool) Include &#x60;unread_count_display&#x60; in the output?
    @param "inclusive" (bool) Include messages with latest or oldest timestamp in results.
    @param "token" (string) Authentication token. Requires scope: &#x60;im:history&#x60;
    @param "oldest" (float32) Start of time range of messages to include in results.
    @param "channel" (string) Direct message channel to fetch history for.
    @param "latest" (float32) End of time range of messages to include in results.
@return ImHistorySuccessSchema*/
func (a *ImApiService) ImHistory(ctx context.Context, localVarOptionals map[string]interface{}) (ImHistorySuccessSchema, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ImHistorySuccessSchema
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/im.history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["count"], "int32", "count"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["unreads"], "bool", "unreads"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["inclusive"], "bool", "inclusive"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["oldest"], "float32", "oldest"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["channel"], "string", "channel"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["latest"], "float32", "latest"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["count"].(int32); localVarOk {
		localVarQueryParams.Add("count", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["unreads"].(bool); localVarOk {
		localVarQueryParams.Add("unreads", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["inclusive"].(bool); localVarOk {
		localVarQueryParams.Add("inclusive", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["oldest"].(float32); localVarOk {
		localVarQueryParams.Add("oldest", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["channel"].(string); localVarOk {
		localVarQueryParams.Add("channel", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["latest"].(float32); localVarOk {
		localVarQueryParams.Add("latest", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ImApiService
Lists direct message channels for the calling user.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "cursor" (string) Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail.
    @param "token" (string) Authentication token. Requires scope: &#x60;im:read&#x60;
    @param "limit" (int32) The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached.
@return ImListSuccessSchema*/
func (a *ImApiService) ImList(ctx context.Context, localVarOptionals map[string]interface{}) (ImListSuccessSchema, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ImListSuccessSchema
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/im.list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["cursor"], "string", "cursor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["cursor"].(string); localVarOk {
		localVarQueryParams.Add("cursor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ImApiService
Sets the read cursor in a direct message channel.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "token" (string) Authentication token. Requires scope: &#x60;im:write&#x60;
    @param "ts" (float32) Timestamp of the most recently seen message.
    @param "channel" (string) Direct message channel to set reading cursor in.
@return ImMarkSuccessSchema*/
func (a *ImApiService) ImMark(ctx context.Context, localVarOptionals map[string]interface{}) (ImMarkSuccessSchema, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ImMarkSuccessSchema
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/im.mark"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["ts"], "float32", "ts"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["channel"], "string", "channel"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarHeaderParams["token"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["ts"].(float32); localVarOk {
		localVarFormParams.Add("ts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["channel"].(string); localVarOk {
		localVarFormParams.Add("channel", parameterToString(localVarTempParam, ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ImApiService
Opens a direct message channel.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "token" (string) Authentication token. Requires scope: &#x60;im:write&#x60;
    @param "returnIm" (bool) Boolean, indicates you want the full IM channel definition in the response.
    @param "user" (string) User to open a direct message channel with.
    @param "includeLocale" (bool) Set this to &#x60;true&#x60; to receive the locale for this im. Defaults to &#x60;false&#x60;
@return ImOpenSuccessSchema*/
func (a *ImApiService) ImOpen(ctx context.Context, localVarOptionals map[string]interface{}) (ImOpenSuccessSchema, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ImOpenSuccessSchema
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/im.open"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["returnIm"], "bool", "returnIm"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["user"], "string", "user"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeLocale"], "bool", "includeLocale"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarHeaderParams["token"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["returnIm"].(bool); localVarOk {
		localVarFormParams.Add("return_im", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["user"].(string); localVarOk {
		localVarFormParams.Add("user", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeLocale"].(bool); localVarOk {
		localVarFormParams.Add("include_locale", parameterToString(localVarTempParam, ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ImApiService
Retrieve a thread of messages posted to a direct message conversation
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "threadTs" (float32) Unique identifier of a thread&#39;s parent message
    @param "token" (string) Authentication token. Requires scope: &#x60;im:history&#x60;
    @param "channel" (string) Direct message channel to fetch thread from
@return DefaultSuccessTemplate*/
func (a *ImApiService) ImReplies(ctx context.Context, localVarOptionals map[string]interface{}) (DefaultSuccessTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     DefaultSuccessTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/im.replies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["threadTs"], "float32", "threadTs"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["channel"], "string", "channel"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["threadTs"].(float32); localVarOk {
		localVarQueryParams.Add("thread_ts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["channel"].(string); localVarOk {
		localVarQueryParams.Add("channel", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
