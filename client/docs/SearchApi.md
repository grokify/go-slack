# \SearchApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAll**](SearchApi.md#SearchAll) | **Get** /search.all | 
[**SearchFiles**](SearchApi.md#SearchFiles) | **Get** /search.files | 
[**SearchMessages**](SearchApi.md#SearchMessages) | **Get** /search.messages | 


# **SearchAll**
> DefaultSuccessTemplate SearchAll(ctx, optional)


Searches for messages and files matching a query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortDir** | **string**| Change sort direction to ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **query** | **string**| Search query. May contains booleans, etc. | 
 **sort** | **string**| Return matches sorted by either &#x60;score&#x60; or &#x60;timestamp&#x60;. | 
 **count** | **string**|  | 
 **token** | **string**| Authentication token. Requires scope: &#x60;search:read&#x60; | 
 **highlight** | **bool**| Pass a value of &#x60;true&#x60; to enable query highlight markers (see below). | 
 **page** | **string**|  | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchFiles**
> DefaultSuccessTemplate SearchFiles(ctx, optional)


Searches for files matching a query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortDir** | **string**| Change sort direction to ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **query** | **string**| Search query. May contain booleans, etc. | 
 **sort** | **string**| Return matches sorted by either &#x60;score&#x60; or &#x60;timestamp&#x60;. | 
 **highlight** | **bool**| Pass a value of &#x60;true&#x60; to enable query highlight markers (see below). | 
 **count** | **string**|  | 
 **token** | **string**| Authentication token. Requires scope: &#x60;search:read&#x60; | 
 **page** | **string**|  | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchMessages**
> DefaultSuccessTemplate SearchMessages(ctx, optional)


Searches for messages matching a query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortDir** | **string**| Change sort direction to ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **query** | **string**| Search query. May contains booleans, etc. | 
 **sort** | **string**| Return matches sorted by either &#x60;score&#x60; or &#x60;timestamp&#x60;. | 
 **count** | **string**| Pass the number of results you want per \&quot;page\&quot;. Maximum of &#x60;100&#x60;. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;search:read&#x60; | 
 **highlight** | **bool**| Pass a value of &#x60;true&#x60; to enable query highlight markers (see below). | 
 **page** | **string**|  | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

