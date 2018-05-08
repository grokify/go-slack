# \ReactionsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReactionsAdd**](ReactionsApi.md#ReactionsAdd) | **Post** /reactions.add | 
[**ReactionsGet**](ReactionsApi.md#ReactionsGet) | **Get** /reactions.get | 
[**ReactionsList**](ReactionsApi.md#ReactionsList) | **Get** /reactions.list | 
[**ReactionsRemove**](ReactionsApi.md#ReactionsRemove) | **Post** /reactions.remove | 


# **ReactionsAdd**
> ReactionsAddSchema ReactionsAdd(ctx, optional)


Adds a reaction to an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Reaction (emoji) name. | 
 **fileComment** | **string**| File comment to add reaction to. | 
 **timestamp** | **float32**| Timestamp of the message to add reaction to. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;reactions:write&#x60; | 
 **file** | **string**| File to add reaction to. | 
 **channel** | **string**| Channel where the message to add reaction to was posted. | 

### Return type

[**ReactionsAddSchema**](reactions.add schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsGet**
> interface{} ReactionsGet(ctx, optional)


Gets reactions for an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **bool**| If true always return the complete reaction list. | 
 **fileComment** | **string**| File comment to get reactions for. | 
 **timestamp** | **float32**| Timestamp of the message to get reactions for. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;reactions:read&#x60; | 
 **file** | **string**| File to get reactions for. | 
 **channel** | **string**| Channel where the message to get reactions for was posted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsList**
> ReactionsListSchema ReactionsList(ctx, optional)


Lists reactions made by a user.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;reactions:read&#x60; | 
 **full** | **bool**| If true always return the complete reaction list. | 
 **user** | **string**| Show reactions made by this user. Defaults to the authed user. | 
 **page** | **string**|  | 

### Return type

[**ReactionsListSchema**](reactions.list schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsRemove**
> ReactionsRemoveSchema ReactionsRemove(ctx, optional)


Removes a reaction from an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Reaction (emoji) name. | 
 **fileComment** | **string**| File comment to remove reaction from. | 
 **timestamp** | **float32**| Timestamp of the message to remove reaction from. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;reactions:write&#x60; | 
 **file** | **string**| File to remove reaction from. | 
 **channel** | **string**| Channel where the message to remove reaction from was posted. | 

### Return type

[**ReactionsRemoveSchema**](reactions.remove schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

