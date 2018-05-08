# \ChannelsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsArchive**](ChannelsApi.md#ChannelsArchive) | **Post** /channels.archive | 
[**ChannelsCreate**](ChannelsApi.md#ChannelsCreate) | **Post** /channels.create | 
[**ChannelsHistory**](ChannelsApi.md#ChannelsHistory) | **Get** /channels.history | 
[**ChannelsInfo**](ChannelsApi.md#ChannelsInfo) | **Get** /channels.info | 
[**ChannelsInvite**](ChannelsApi.md#ChannelsInvite) | **Post** /channels.invite | 
[**ChannelsJoin**](ChannelsApi.md#ChannelsJoin) | **Post** /channels.join | 
[**ChannelsKick**](ChannelsApi.md#ChannelsKick) | **Post** /channels.kick | 
[**ChannelsLeave**](ChannelsApi.md#ChannelsLeave) | **Post** /channels.leave | 
[**ChannelsList**](ChannelsApi.md#ChannelsList) | **Get** /channels.list | 
[**ChannelsMark**](ChannelsApi.md#ChannelsMark) | **Post** /channels.mark | 
[**ChannelsRename**](ChannelsApi.md#ChannelsRename) | **Post** /channels.rename | 
[**ChannelsReplies**](ChannelsApi.md#ChannelsReplies) | **Get** /channels.replies | 
[**ChannelsSetPurpose**](ChannelsApi.md#ChannelsSetPurpose) | **Post** /channels.setPurpose | 
[**ChannelsSetTopic**](ChannelsApi.md#ChannelsSetTopic) | **Post** /channels.setTopic | 
[**ChannelsUnarchive**](ChannelsApi.md#ChannelsUnarchive) | **Post** /channels.unarchive | 


# **ChannelsArchive**
> ChannelsArchiveSuccessSchema ChannelsArchive(ctx, optional)


Archives a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **string**| Channel to archive | 

### Return type

[**ChannelsArchiveSuccessSchema**](channels.archive success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsCreate**
> ChannelsCreateErrorSchema ChannelsCreate(ctx, optional)


Creates a channel.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **name** | **string**| Name of channel to create | 

### Return type

[**ChannelsCreateErrorSchema**](channels.create error schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsHistory**
> ChannelsHistorySuccessSchema ChannelsHistory(ctx, optional)


Fetches history of messages and events from a channel.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:history&#x60; | 
 **oldest** | **float32**| Start of time range of messages to include in results. | 
 **channel** | **string**| Channel to fetch history for. | 
 **latest** | **float32**| End of time range of messages to include in results. | 

### Return type

[**ChannelsHistorySuccessSchema**](channels.history success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsInfo**
> ChannelsInfoSuccessSchema ChannelsInfo(ctx, optional)


Gets information about a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:read&#x60; | 
 **includeLocale** | **bool**| Set this to &#x60;true&#x60; to receive the locale for this channel. Defaults to &#x60;false&#x60; | 
 **channel** | **string**| Channel to get info on | 

### Return type

[**ChannelsInfoSuccessSchema**](channels.info success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsInvite**
> ChannelsInviteErrorSchema ChannelsInvite(ctx, optional)


Invites a user to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **user** | **string**| User to invite to channel. | 
 **channel** | **string**| Channel to invite user to. | 

### Return type

[**ChannelsInviteErrorSchema**](channels.invite error schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsJoin**
> DefaultSuccessTemplate ChannelsJoin(ctx, optional)


Joins a channel, creating it if needed.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **name** | **string**| Name of channel to join | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsKick**
> DefaultSuccessTemplate ChannelsKick(ctx, optional)


Removes a user from a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **user** | **string**| User to remove from channel. | 
 **channel** | **string**| Channel to remove user from. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsLeave**
> DefaultSuccessTemplate ChannelsLeave(ctx, optional)


Leaves a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **string**| Channel to leave | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsList**
> ChannelsListSuccessSchema ChannelsList(ctx, optional)


Lists all channels in a Slack team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeMembers** | **bool**| Exclude the &#x60;members&#x60; collection from each &#x60;channel&#x60; | 
 **cursor** | **string**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:read&#x60; | 
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **excludeArchived** | **bool**| Exclude archived channels from the list | 

### Return type

[**ChannelsListSuccessSchema**](channels.list success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsMark**
> ChannelsMarkSuccessSchema ChannelsMark(ctx, optional)


Sets the read cursor in a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **ts** | **float32**| Timestamp of the most recently seen message. | 
 **channel** | **string**| Channel to set reading cursor in. | 

### Return type

[**ChannelsMarkSuccessSchema**](channels.mark success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsRename**
> DefaultSuccessTemplate ChannelsRename(ctx, optional)


Renames a channel.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **name** | **string**| New name for channel. | 
 **channel** | **string**| Channel to rename | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsReplies**
> DefaultSuccessTemplate ChannelsReplies(ctx, optional)


Retrieve a thread of messages posted to a channel

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
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:history&#x60; | 
 **channel** | **string**| Channel to fetch thread from | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsSetPurpose**
> DefaultSuccessTemplate ChannelsSetPurpose(ctx, optional)


Sets the purpose for a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **purpose** | **string**| The new purpose | 
 **channel** | **string**| Channel to set the purpose of | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsSetTopic**
> DefaultSuccessTemplate ChannelsSetTopic(ctx, optional)


Sets the topic for a channel.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **string**| Channel to set the topic of | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsUnarchive**
> DefaultSuccessTemplate ChannelsUnarchive(ctx, optional)


Unarchives a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **string**| Channel to unarchive | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

