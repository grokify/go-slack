# \UsergroupsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupsCreate**](UsergroupsApi.md#UsergroupsCreate) | **Post** /usergroups.create | 
[**UsergroupsDisable**](UsergroupsApi.md#UsergroupsDisable) | **Post** /usergroups.disable | 
[**UsergroupsEnable**](UsergroupsApi.md#UsergroupsEnable) | **Post** /usergroups.enable | 
[**UsergroupsList**](UsergroupsApi.md#UsergroupsList) | **Get** /usergroups.list | 
[**UsergroupsUpdate**](UsergroupsApi.md#UsergroupsUpdate) | **Post** /usergroups.update | 
[**UsergroupsUsersList**](UsergroupsApi.md#UsergroupsUsersList) | **Get** /usergroups.users.list | 
[**UsergroupsUsersUpdate**](UsergroupsApi.md#UsergroupsUsersUpdate) | **Post** /usergroups.users.update | 


# **UsergroupsCreate**
> DefaultSuccessTemplate UsergroupsCreate(ctx, optional)


Create a User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handle** | **string**| A mention handle. Must be unique among channels, users and User Groups. | 
 **name** | **string**| A name for the User Group. Must be unique among User Groups. | 
 **channels** | **string**| A comma separated string of encoded channel IDs for which the User Group uses as a default. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **bool**| Include the number of users in each User Group. | 
 **description** | **string**| A short description of the User Group. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsDisable**
> DefaultSuccessTemplate UsergroupsDisable(ctx, optional)


Disable an existing User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **bool**| Include the number of users in the User Group. | 
 **usergroup** | **string**| The encoded ID of the User Group to disable. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsEnable**
> DefaultSuccessTemplate UsergroupsEnable(ctx, optional)


Enable a User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **bool**| Include the number of users in the User Group. | 
 **usergroup** | **string**| The encoded ID of the User Group to enable. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsList**
> DefaultSuccessTemplate UsergroupsList(ctx, optional)


List all User Groups for a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUsers** | **bool**| Include the list of users for each User Group. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;usergroups:read&#x60; | 
 **includeCount** | **bool**| Include the number of users in each User Group. | 
 **includeDisabled** | **bool**| Include disabled User Groups. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsUpdate**
> DefaultSuccessTemplate UsergroupsUpdate(ctx, optional)


Update an existing User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handle** | **string**| A mention handle. Must be unique among channels, users and User Groups. | 
 **description** | **string**| A short description of the User Group. | 
 **channels** | **string**| A comma separated string of encoded channel IDs for which the User Group uses as a default. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **bool**| Include the number of users in the User Group. | 
 **usergroup** | **string**| The encoded ID of the User Group to update. | 
 **name** | **string**| A name for the User Group. Must be unique among User Groups. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

