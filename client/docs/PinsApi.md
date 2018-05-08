# \PinsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PinsAdd**](PinsApi.md#PinsAdd) | **Post** /pins.add | 
[**PinsList**](PinsApi.md#PinsList) | **Get** /pins.list | 
[**PinsRemove**](PinsApi.md#PinsRemove) | **Post** /pins.remove | 


# **PinsAdd**
> PinsAddSchema PinsAdd(ctx, optional)


Pins an item to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **string**| File comment to pin. | 
 **timestamp** | **float32**| Timestamp of the message to pin. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;pins:write&#x60; | 
 **file** | **string**| File to pin. | 
 **channel** | **string**| Channel to pin the item in. | 

### Return type

[**PinsAddSchema**](pins.add schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PinsList**
> PinsList(ctx, optional)


Lists items pinned to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;pins:read&#x60; | 
 **channel** | **string**| Channel to get pinned items for. | 

### Return type

 (empty response body)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PinsRemove**
> PinsRemoveSchema PinsRemove(ctx, optional)


Un-pins an item from a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **string**| File comment to un-pin. | 
 **timestamp** | **float32**| Timestamp of the message to un-pin. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;pins:write&#x60; | 
 **file** | **string**| File to un-pin. | 
 **channel** | **string**| Channel where the item is pinned to. | 

### Return type

[**PinsRemoveSchema**](pins.remove schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

