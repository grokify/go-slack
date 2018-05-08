# \RtmApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RtmConnect**](RtmApi.md#RtmConnect) | **Get** /rtm.connect | 
[**RtmStart**](RtmApi.md#RtmStart) | **Get** /rtm.start | 


# **RtmConnect**
> DefaultSuccessTemplate RtmConnect(ctx, optional)


Starts a Real Time Messaging session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presenceSub** | **bool**| Only deliver presence events when requested by subscription. See [presence subscriptions](/docs/presence-and-status#subscriptions). | 
 **token** | **string**| Authentication token. Requires scope: &#x60;rtm:stream&#x60; | 
 **batchPresenceAware** | **bool**| Batch presence deliveries via subscription. Enabling changes the shape of &#x60;presence_change&#x60; events. See [batch presence](/docs/presence-and-status#batching). | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RtmStart**
> DefaultSuccessTemplate RtmStart(ctx, optional)


Starts a Real Time Messaging session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noLatest** | **bool**| Exclude latest timestamps for channels, groups, mpims, and ims. Automatically sets &#x60;no_unreads&#x60; to &#x60;1&#x60; | 
 **simpleLatest** | **bool**| Return timestamp only for latest message object of each channel (improves performance). | 
 **includeLocale** | **bool**| Set this to &#x60;true&#x60; to receive the locale for users and channels. Defaults to &#x60;false&#x60; | 
 **presenceSub** | **bool**| Only deliver presence events when requested by subscription. See [presence subscriptions](/docs/presence-and-status#subscriptions). | 
 **noUnreads** | **bool**| Skip unread counts for each channel (improves performance). | 
 **batchPresenceAware** | **bool**| Batch presence deliveries via subscription. Enabling changes the shape of &#x60;presence_change&#x60; events. See [batch presence](/docs/presence-and-status#batching). | 
 **mpimAware** | **bool**| Returns MPIMs to the client in the API response. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;rtm:stream&#x60; | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

