# \GroupsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsArchive**](GroupsApi.md#GroupsArchive) | **Post** /groups.archive | 
[**GroupsCreate**](GroupsApi.md#GroupsCreate) | **Post** /groups.create | 
[**GroupsCreateChild**](GroupsApi.md#GroupsCreateChild) | **Post** /groups.createChild | 
[**GroupsHistory**](GroupsApi.md#GroupsHistory) | **Get** /groups.history | 
[**GroupsInfo**](GroupsApi.md#GroupsInfo) | **Get** /groups.info | 
[**GroupsInvite**](GroupsApi.md#GroupsInvite) | **Post** /groups.invite | 
[**GroupsKick**](GroupsApi.md#GroupsKick) | **Post** /groups.kick | 
[**GroupsLeave**](GroupsApi.md#GroupsLeave) | **Post** /groups.leave | 
[**GroupsList**](GroupsApi.md#GroupsList) | **Get** /groups.list | 
[**GroupsMark**](GroupsApi.md#GroupsMark) | **Post** /groups.mark | 
[**GroupsOpen**](GroupsApi.md#GroupsOpen) | **Post** /groups.open | 
[**GroupsRename**](GroupsApi.md#GroupsRename) | **Post** /groups.rename | 
[**GroupsReplies**](GroupsApi.md#GroupsReplies) | **Get** /groups.replies | 
[**GroupsSetPurpose**](GroupsApi.md#GroupsSetPurpose) | **Post** /groups.setPurpose | 
[**GroupsSetTopic**](GroupsApi.md#GroupsSetTopic) | **Post** /groups.setTopic | 
[**GroupsUnarchive**](GroupsApi.md#GroupsUnarchive) | **Post** /groups.unarchive | 


# **GroupsArchive**
> DefaultSuccessTemplate GroupsArchive(ctx, optional)


Archives a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **string**| Private channel to archive | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsCreate**
> GroupsCreateSuccessSchema GroupsCreate(ctx, optional)


Creates a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool**| Whether to return errors on invalid channel name instead of modifying it to meet the specified criteria. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **name** | **string**| Name of private channel to create | 

### Return type

[**GroupsCreateSuccessSchema**](groups.create success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsCreateChild**
> DefaultSuccessTemplate GroupsCreateChild(ctx, optional)


Clones and archives a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **string**| Private channel to clone and archive. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsHistory**
> GroupsHistorySuccessSchema GroupsHistory(ctx, optional)


Fetches history of messages and events from a private channel.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:history&#x60; | 
 **oldest** | **float32**| Start of time range of messages to include in results. | 
 **channel** | **string**| Private channel to fetch history for. | 
 **latest** | **float32**| End of time range of messages to include in results. | 

### Return type

[**GroupsHistorySuccessSchema**](groups.history success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsInfo**
> GroupsInfoSuccessSchema GroupsInfo(ctx, optional)


Gets information about a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:read&#x60; | 
 **includeLocale** | **bool**| Set this to &#x60;true&#x60; to receive the locale for this group. Defaults to &#x60;false&#x60; | 
 **channel** | **string**| Private channel to get info on | 

### Return type

[**GroupsInfoSuccessSchema**](groups.info success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsInvite**
> GroupsInviteSuccessSchema GroupsInvite(ctx, optional)


Invites a user to a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **user** | **string**| User to invite. | 
 **channel** | **string**| Private channel to invite user to. | 

### Return type

[**GroupsInviteSuccessSchema**](groups.invite success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsKick**
> DefaultSuccessTemplate GroupsKick(ctx, optional)


Removes a user from a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **user** | **string**| User to remove from private channel. | 
 **channel** | **string**| Private channel to remove user from. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsLeave**
> DefaultSuccessTemplate GroupsLeave(ctx, optional)


Leaves a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **string**| Private channel to leave | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsList**
> GroupsListSuccessSchema GroupsList(ctx, optional)


Lists private channels that the calling user has access to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeMembers** | **bool**| Exclude the &#x60;members&#x60; from each &#x60;group&#x60; | 
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:read&#x60; | 
 **excludeArchived** | **bool**| Don&#39;t return archived private channels. | 

### Return type

[**GroupsListSuccessSchema**](groups.list success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsMark**
> GroupsMarkSuccessSchema GroupsMark(ctx, optional)


Sets the read cursor in a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **ts** | **float32**| Timestamp of the most recently seen message. | 
 **channel** | **string**| Private channel to set reading cursor in. | 

### Return type

[**GroupsMarkSuccessSchema**](groups.mark success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsOpen**
> DefaultSuccessTemplate GroupsOpen(ctx, optional)


Opens a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **string**| Private channel to open. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsRename**
> DefaultSuccessTemplate GroupsRename(ctx, optional)


Renames a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool**| Whether to return errors on invalid channel name instead of modifying it to meet the specified criteria. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **name** | **string**| New name for private channel. | 
 **channel** | **string**| Private channel to rename | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsReplies**
> DefaultSuccessTemplate GroupsReplies(ctx, optional)


Retrieve a thread of messages posted to a private channel

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
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:history&#x60; | 
 **channel** | **string**| Private channel to fetch thread from | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsSetPurpose**
> DefaultSuccessTemplate GroupsSetPurpose(ctx, optional)


Sets the purpose for a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **purpose** | **string**| The new purpose | 
 **channel** | **string**| Private channel to set the purpose of | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsSetTopic**
> DefaultSuccessTemplate GroupsSetTopic(ctx, optional)


Sets the topic for a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **string**| The new topic | 
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **string**| Private channel to set the topic of | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsUnarchive**
> DefaultSuccessTemplate GroupsUnarchive(ctx, optional)


Unarchives a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **string**| Private channel to unarchive | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

