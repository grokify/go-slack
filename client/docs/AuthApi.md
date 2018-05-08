# \AuthApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthRevoke**](AuthApi.md#AuthRevoke) | **Get** /auth.revoke | 
[**AuthTest**](AuthApi.md#AuthTest) | **Get** /auth.test | 


# **AuthRevoke**
> DefaultSuccessTemplate AuthRevoke(ctx, optional)


Revokes a token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **test** | **bool**| Setting this parameter to &#x60;1&#x60; triggers a _testing mode_ where the specified token will not actually be revoked. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;none&#x60; | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthTest**
> AuthTestSuccessSchema AuthTest(ctx, optional)


Checks authentication & identity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;none&#x60; | 

### Return type

[**AuthTestSuccessSchema**](auth.test success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

