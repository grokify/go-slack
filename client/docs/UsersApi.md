# \UsersApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConversations**](UsersApi.md#UsersConversations) | **Get** /users.conversations | 
[**UsersDeletePhoto**](UsersApi.md#UsersDeletePhoto) | **Post** /users.deletePhoto | 
[**UsersGetPresence**](UsersApi.md#UsersGetPresence) | **Get** /users.getPresence | 
[**UsersIdentity**](UsersApi.md#UsersIdentity) | **Get** /users.identity | 
[**UsersInfo**](UsersApi.md#UsersInfo) | **Get** /users.info | 
[**UsersList**](UsersApi.md#UsersList) | **Get** /users.list | 
[**UsersLookupByEmail**](UsersApi.md#UsersLookupByEmail) | **Get** /users.lookupByEmail | 
[**UsersProfileGet**](UsersApi.md#UsersProfileGet) | **Get** /users.profile.get | 
[**UsersProfileSet**](UsersApi.md#UsersProfileSet) | **Post** /users.profile.set | 
[**UsersSetActive**](UsersApi.md#UsersSetActive) | **Post** /users.setActive | 
[**UsersSetPhoto**](UsersApi.md#UsersSetPhoto) | **Post** /users.setPhoto | 
[**UsersSetPresence**](UsersApi.md#UsersSetPresence) | **Post** /users.setPresence | 


# **UsersConversations**
> ConversationsListSuccessSchema1 UsersConversations(ctx, optional)


List conversations the calling user may access.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:read&#x60; | 
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. Must be an integer no larger than 1000. | 
 **user** | **string**| Browse conversations by a specific user ID&#39;s membership. Non-public channels are restricted to those where the calling user shares membership. | 
 **excludeArchived** | **bool**| Set to &#x60;true&#x60; to exclude archived channels from the list | 
 **types** | **string**| Mix and match channel types by providing a comma-separated list of any combination of &#x60;public_channel&#x60;, &#x60;private_channel&#x60;, &#x60;mpim&#x60;, &#x60;im&#x60; | 

### Return type

[**ConversationsListSuccessSchema1**](conversations.list success schema_1.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersDeletePhoto**
> DefaultSuccessTemplate UsersDeletePhoto(ctx, optional)


Delete the user profile photo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;users.profile:write&#x60; | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGetPresence**
> ApiMethodUsersGetPresence UsersGetPresence(ctx, optional)


Gets user presence information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;users:read&#x60; | 
 **user** | **string**| User to get presence info on. Defaults to the authed user. | 

### Return type

[**ApiMethodUsersGetPresence**](API method users.getPresence.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdentity**
> DefaultSuccessTemplate UsersIdentity(ctx, optional)


Get a user's identity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;identity.basic&#x60; | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersInfo**
> UsersInfoSuccessSchema UsersInfo(ctx, optional)


Gets information about a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeLocale** | **bool**| Set this to &#x60;true&#x60; to receive the locale for this user. Defaults to &#x60;false&#x60; | 
 **token** | **string**| Authentication token. Requires scope: &#x60;users:read&#x60; | 
 **user** | **string**| User to get info on | 

### Return type

[**UsersInfoSuccessSchema**](users.info success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersList**
> UsersListSchema UsersList(ctx, optional)


Lists all users in a Slack team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presence** | **bool**| Deprecated. Whether to include presence data in the output. Defaults to &#x60;false&#x60;. Setting this to &#x60;true&#x60; reduces performance, especially with large teams. | 
 **cursor** | **string**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;users:read&#x60; | 
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **includeLocale** | **bool**| Set this to &#x60;true&#x60; to receive the locale for users. Defaults to &#x60;false&#x60; | 

### Return type

[**UsersListSchema**](users.list schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersLookupByEmail**
> UsersLookupByEmailSuccessSchema UsersLookupByEmail(ctx, optional)


Find a user with an email address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;users:read.email&#x60; | 
 **email** | **string**| An email address belonging to a user in the workspace | 

### Return type

[**UsersLookupByEmailSuccessSchema**](users.lookupByEmail success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **UsersSetActive**
> DefaultSuccessTemplate UsersSetActive(ctx, optional)


Marked a user as active. Deprecated and non-functional.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;users:write&#x60; | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSetPhoto**
> DefaultSuccessTemplate UsersSetPhoto(ctx, optional)


Set the user profile photo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | **string**| File contents via &#x60;multipart/form-data&#x60;. | 
 **cropW** | **int32**| Width/height of crop box (always square) | 
 **token** | **string**| Authentication token. Requires scope: &#x60;users.profile:write&#x60; | 
 **cropY** | **int32**| Y coordinate of top-left corner of crop box | 
 **cropX** | **int32**| X coordinate of top-left corner of crop box | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSetPresence**
> DefaultSuccessTemplate UsersSetPresence(ctx, optional)


Manually sets user presence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;users:write&#x60; | 
 **presence** | **string**| Either &#x60;auto&#x60; or &#x60;away&#x60; | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

