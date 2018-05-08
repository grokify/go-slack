# \MpimApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpimClose**](MpimApi.md#MpimClose) | **Post** /mpim.close | 
[**MpimHistory**](MpimApi.md#MpimHistory) | **Get** /mpim.history | 
[**MpimList**](MpimApi.md#MpimList) | **Get** /mpim.list | 
[**MpimMark**](MpimApi.md#MpimMark) | **Post** /mpim.mark | 
[**MpimOpen**](MpimApi.md#MpimOpen) | **Post** /mpim.open | 
[**MpimReplies**](MpimApi.md#MpimReplies) | **Get** /mpim.replies | 


# **MpimClose**
> DefaultSuccessTemplate MpimClose(ctx, optional)


Closes a multiparty direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;mpim:write&#x60; | 
 **channel** | **string**| MPIM to close. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimHistory**
> DefaultSuccessTemplate MpimHistory(ctx, optional)


Fetches history of messages and events from a multiparty direct message.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;mpim:history&#x60; | 
 **oldest** | **float32**| Start of time range of messages to include in results. | 
 **channel** | **string**| Multiparty direct message to fetch history for. | 
 **latest** | **float32**| End of time range of messages to include in results. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimList**
> DefaultSuccessTemplate MpimList(ctx, optional)


Lists multiparty direct message channels for the calling user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;mpim:read&#x60; | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimMark**
> DefaultSuccessTemplate MpimMark(ctx, optional)


Sets the read cursor in a multiparty direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;mpim:write&#x60; | 
 **ts** | **float32**| Timestamp of the most recently seen message. | 
 **channel** | **string**| multiparty direct message channel to set reading cursor in. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimOpen**
> MpimOpenSuccessSchema MpimOpen(ctx, optional)


This method opens a multiparty direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;mpim:write&#x60; | 
 **users** | **string**| Comma separated lists of users.  The ordering of the users is preserved whenever a MPIM group is returned. | 

### Return type

[**MpimOpenSuccessSchema**](mpim.open success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimReplies**
> DefaultSuccessTemplate MpimReplies(ctx, optional)


Retrieve a thread of messages posted to a direct message conversation from a multiparty direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **float32**| Unique identifier of a thread&#39;s parent message. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;mpim:history&#x60; | 
 **channel** | **string**| Multiparty direct message channel to fetch thread from. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

