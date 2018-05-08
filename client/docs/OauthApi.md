# \OauthApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthAccess**](OauthApi.md#OauthAccess) | **Get** /oauth.access | 
[**OauthToken**](OauthApi.md#OauthToken) | **Get** /oauth.token | 


# **OauthAccess**
> DefaultSuccessTemplate OauthAccess(ctx, optional)


Exchanges a temporary OAuth code for an API token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientSecret** | **string**| Issued when you created your application. | 
 **code** | **string**| The &#x60;code&#x60; param returned via the OAuth callback. | 
 **clientId** | **string**| Issued when you created your application. | 
 **redirectUri** | **string**| This must match the originally submitted URI (if one was sent). | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OauthToken**
> DefaultSuccessTemplate OauthToken(ctx, optional)


Exchanges a temporary OAuth verifier code for a workspace token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientSecret** | **string**| Issued when you created your application. | 
 **code** | **string**| The &#x60;code&#x60; param returned via the OAuth callback. | 
 **singleChannel** | **bool**| Request the user to add your app only to a single channel. | 
 **clientId** | **string**| Issued when you created your application. | 
 **redirectUri** | **string**| This must match the originally submitted URI (if one was sent). | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

