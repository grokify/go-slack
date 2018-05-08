# \StarsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StarsAdd**](StarsApi.md#StarsAdd) | **Post** /stars.add | 
[**StarsList**](StarsApi.md#StarsList) | **Get** /stars.list | 
[**StarsRemove**](StarsApi.md#StarsRemove) | **Post** /stars.remove | 


# **StarsAdd**
> DefaultSuccessTemplate StarsAdd(ctx, optional)


Adds a star to an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **string**| File comment to add star to. | 
 **timestamp** | **float32**| Timestamp of the message to add star to. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;stars:write&#x60; | 
 **channel** | **string**| Channel to add star to, or channel where the message to add star to was posted (used with &#x60;timestamp&#x60;). | 
 **file** | **string**| File to add star to. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StarsList**
> DefaultSuccessTemplate StarsList(ctx, optional)


Lists stars for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **string**|  | 
 **token** | **string**| Authentication token. Requires scope: &#x60;stars:read&#x60; | 
 **page** | **string**|  | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StarsRemove**
> DefaultSuccessTemplate StarsRemove(ctx, optional)


Removes a star from an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **string**| File comment to remove star from. | 
 **timestamp** | **float32**| Timestamp of the message to remove star from. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;stars:write&#x60; | 
 **channel** | **string**| Channel to remove star from, or channel where the message to remove star from was posted (used with &#x60;timestamp&#x60;). | 
 **file** | **string**| File to remove star from. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

