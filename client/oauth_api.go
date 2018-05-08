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

type OauthApiService service

/* OauthApiService
Exchanges a temporary OAuth code for an API token.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "clientSecret" (string) Issued when you created your application.
    @param "code" (string) The &#x60;code&#x60; param returned via the OAuth callback.
    @param "clientId" (string) Issued when you created your application.
    @param "redirectUri" (string) This must match the originally submitted URI (if one was sent).
@return DefaultSuccessTemplate*/
func (a *OauthApiService) OauthAccess(ctx context.Context, localVarOptionals map[string]interface{}) (DefaultSuccessTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     DefaultSuccessTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth.access"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["clientSecret"], "string", "clientSecret"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["code"], "string", "code"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["clientId"], "string", "clientId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["redirectUri"], "string", "redirectUri"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["clientSecret"].(string); localVarOk {
		localVarQueryParams.Add("client_secret", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["code"].(string); localVarOk {
		localVarQueryParams.Add("code", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["clientId"].(string); localVarOk {
		localVarQueryParams.Add("client_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["redirectUri"].(string); localVarOk {
		localVarQueryParams.Add("redirect_uri", parameterToString(localVarTempParam, ""))
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

/* OauthApiService
Exchanges a temporary OAuth verifier code for a workspace token.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "clientSecret" (string) Issued when you created your application.
    @param "code" (string) The &#x60;code&#x60; param returned via the OAuth callback.
    @param "singleChannel" (bool) Request the user to add your app only to a single channel.
    @param "clientId" (string) Issued when you created your application.
    @param "redirectUri" (string) This must match the originally submitted URI (if one was sent).
@return DefaultSuccessTemplate*/
func (a *OauthApiService) OauthToken(ctx context.Context, localVarOptionals map[string]interface{}) (DefaultSuccessTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     DefaultSuccessTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth.token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["clientSecret"], "string", "clientSecret"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["code"], "string", "code"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["singleChannel"], "bool", "singleChannel"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["clientId"], "string", "clientId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["redirectUri"], "string", "redirectUri"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["clientSecret"].(string); localVarOk {
		localVarQueryParams.Add("client_secret", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["code"].(string); localVarOk {
		localVarQueryParams.Add("code", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["singleChannel"].(bool); localVarOk {
		localVarQueryParams.Add("single_channel", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["clientId"].(string); localVarOk {
		localVarQueryParams.Add("client_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["redirectUri"].(string); localVarOk {
		localVarQueryParams.Add("redirect_uri", parameterToString(localVarTempParam, ""))
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
