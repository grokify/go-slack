# \UsergroupsUsersApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupsUsersList**](UsergroupsUsersApi.md#UsergroupsUsersList) | **Get** /usergroups.users.list | 
[**UsergroupsUsersUpdate**](UsergroupsUsersApi.md#UsergroupsUsersUpdate) | **Post** /usergroups.users.update | 


# **UsergroupsUsersList**
> DefaultSuccessTemplate UsergroupsUsersList(ctx, optional)


List all users in a User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;usergroups:read&#x60; | 
 **includeDisabled** | **bool**| Allow results that involve disabled User Groups. | 
 **usergroup** | **string**| The encoded ID of the User Group to update. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsUsersUpdate**
> DefaultSuccessTemplate UsergroupsUsersUpdate(ctx, optional)


Update the list of users for a User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCount** | **bool**| Include the number of users in the User Group. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **users** | **string**| A comma separated string of encoded user IDs that represent the entire list of users for the User Group. | 
 **usergroup** | **string**| The encoded ID of the User Group to update. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

