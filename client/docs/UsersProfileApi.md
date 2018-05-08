# \UsersProfileApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersProfileGet**](UsersProfileApi.md#UsersProfileGet) | **Get** /users.profile.get | 
[**UsersProfileSet**](UsersProfileApi.md#UsersProfileSet) | **Post** /users.profile.set | 


# **UsersProfileGet**
> DefaultSuccessTemplate UsersProfileGet(ctx, optional)


Retrieves a user's profile information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;users.profile:read&#x60; | 
 **includeLabels** | **bool**| Include labels for each ID in custom profile fields | 
 **user** | **string**| User to retrieve profile info for | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersProfileSet**
> DefaultSuccessTemplate UsersProfileSet(ctx, optional)


Set the profile information for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | **string**| Collection of key:value pairs presented as a URL-encoded JSON hash. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;users.profile:write&#x60; | 
 **user** | **string**| ID of user to change. This argument may only be specified by team admins on paid teams. | 
 **value** | **string**| Value to set a single key to. Usable only if &#x60;profile&#x60; is not passed. | 
 **name** | **string**| Name of a single key to set. Usable only if &#x60;profile&#x60; is not passed. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

