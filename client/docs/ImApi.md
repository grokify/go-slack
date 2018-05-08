# \ImApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImClose**](ImApi.md#ImClose) | **Post** /im.close | 
[**ImHistory**](ImApi.md#ImHistory) | **Get** /im.history | 
[**ImList**](ImApi.md#ImList) | **Get** /im.list | 
[**ImMark**](ImApi.md#ImMark) | **Post** /im.mark | 
[**ImOpen**](ImApi.md#ImOpen) | **Post** /im.open | 
[**ImReplies**](ImApi.md#ImReplies) | **Get** /im.replies | 


# **ImClose**
> DefaultSuccessTemplate ImClose(ctx, optional)


Close a direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;im:write&#x60; | 
 **channel** | **string**| Direct message channel to close. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImHistory**
> ImHistorySuccessSchema ImHistory(ctx, optional)


Fetches history of messages and events from direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32**| Number of messages to return, between 1 and 1000. | 
 **unreads** | **bool**| Include &#x60;unread_count_display&#x60; in the output? | 
 **inclusive** | **bool**| Include messages with latest or oldest timestamp in results. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;im:history&#x60; | 
 **oldest** | **float32**| Start of time range of messages to include in results. | 
 **channel** | **string**| Direct message channel to fetch history for. | 
 **latest** | **float32**| End of time range of messages to include in results. | 

### Return type

[**ImHistorySuccessSchema**](im.history success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImList**
> ImListSuccessSchema ImList(ctx, optional)


Lists direct message channels for the calling user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;im:read&#x60; | 
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 

### Return type

[**ImListSuccessSchema**](im.list success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImMark**
> ImMarkSuccessSchema ImMark(ctx, optional)


Sets the read cursor in a direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;im:write&#x60; | 
 **ts** | **float32**| Timestamp of the most recently seen message. | 
 **channel** | **string**| Direct message channel to set reading cursor in. | 

### Return type

[**ImMarkSuccessSchema**](im.mark success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImOpen**
> ImOpenSuccessSchema ImOpen(ctx, optional)


Opens a direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;im:write&#x60; | 
 **returnIm** | **bool**| Boolean, indicates you want the full IM channel definition in the response. | 
 **user** | **string**| User to open a direct message channel with. | 
 **includeLocale** | **bool**| Set this to &#x60;true&#x60; to receive the locale for this im. Defaults to &#x60;false&#x60; | 

### Return type

[**ImOpenSuccessSchema**](im.open success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImReplies**
> DefaultSuccessTemplate ImReplies(ctx, optional)


Retrieve a thread of messages posted to a direct message conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **float32**| Unique identifier of a thread&#39;s parent message | 
 **token** | **string**| Authentication token. Requires scope: &#x60;im:history&#x60; | 
 **channel** | **string**| Direct message channel to fetch thread from | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

