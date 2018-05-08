# \ConversationsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversationsArchive**](ConversationsApi.md#ConversationsArchive) | **Post** /conversations.archive | 
[**ConversationsClose**](ConversationsApi.md#ConversationsClose) | **Post** /conversations.close | 
[**ConversationsCreate**](ConversationsApi.md#ConversationsCreate) | **Post** /conversations.create | 
[**ConversationsHistory**](ConversationsApi.md#ConversationsHistory) | **Get** /conversations.history | 
[**ConversationsInfo**](ConversationsApi.md#ConversationsInfo) | **Get** /conversations.info | 
[**ConversationsInvite**](ConversationsApi.md#ConversationsInvite) | **Post** /conversations.invite | 
[**ConversationsJoin**](ConversationsApi.md#ConversationsJoin) | **Post** /conversations.join | 
[**ConversationsKick**](ConversationsApi.md#ConversationsKick) | **Post** /conversations.kick | 
[**ConversationsLeave**](ConversationsApi.md#ConversationsLeave) | **Post** /conversations.leave | 
[**ConversationsList**](ConversationsApi.md#ConversationsList) | **Get** /conversations.list | 
[**ConversationsMembers**](ConversationsApi.md#ConversationsMembers) | **Get** /conversations.members | 
[**ConversationsOpen**](ConversationsApi.md#ConversationsOpen) | **Post** /conversations.open | 
[**ConversationsRename**](ConversationsApi.md#ConversationsRename) | **Post** /conversations.rename | 
[**ConversationsReplies**](ConversationsApi.md#ConversationsReplies) | **Get** /conversations.replies | 
[**ConversationsSetPurpose**](ConversationsApi.md#ConversationsSetPurpose) | **Post** /conversations.setPurpose | 
[**ConversationsSetTopic**](ConversationsApi.md#ConversationsSetTopic) | **Post** /conversations.setTopic | 
[**ConversationsUnarchive**](ConversationsApi.md#ConversationsUnarchive) | **Post** /conversations.unarchive | 


# **ConversationsArchive**
> ConversationsArchiveSuccessSchema ConversationsArchive(ctx, optional)


Archives a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **string**| ID of conversation to archive | 

### Return type

[**ConversationsArchiveSuccessSchema**](conversations.archive success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsClose**
> ConversationsCloseSuccessSchema ConversationsClose(ctx, optional)


Closes a direct message or multi-person direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **string**| Conversation to close. | 

### Return type

[**ConversationsCloseSuccessSchema**](conversations.close success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsCreate**
> ConversationsCreateSuccessSchema ConversationsCreate(ctx, optional)


Initiates a public or private channel-based conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **name** | **string**| Name of the public or private channel to create | 
 **isPrivate** | **bool**| Create a private channel instead of a public one | 

### Return type

[**ConversationsCreateSuccessSchema**](conversations.create success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsHistory**
> ConversationsHistorySuccessSchema ConversationsHistory(ctx, optional)


Fetches a conversation's history of messages and events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inclusive** | **bool**| Include messages with latest or oldest timestamp in results only when either timestamp is specified. | 
 **cursor** | **string**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:history&#x60; | 
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **oldest** | **float32**| Start of time range of messages to include in results. | 
 **channel** | **string**| Conversation ID to fetch history for. | 
 **latest** | **float32**| End of time range of messages to include in results. | 

### Return type

[**ConversationsHistorySuccessSchema**](conversations.history success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsInfo**
> ConversationsInfoSuccessSchema ConversationsInfo(ctx, optional)


Retrieve information about a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:read&#x60; | 
 **channel** | **string**| Conversation ID to learn more about | 
 **includeLocale** | **bool**| Set this to &#x60;true&#x60; to receive the locale for this conversation. Defaults to &#x60;false&#x60; | 

### Return type

[**ConversationsInfoSuccessSchema**](conversations.info success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsInvite**
> ConversationsInviteErrorSchema ConversationsInvite(ctx, optional)


Invites users to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **users** | **string**| A comma separated list of user IDs. Up to 30 users may be listed. | 
 **channel** | **string**| The ID of the public or private channel to invite user(s) to. | 

### Return type

[**ConversationsInviteErrorSchema**](conversations.invite error schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsJoin**
> ConversationsJoinSuccessSchema ConversationsJoin(ctx, optional)


Joins an existing conversation.

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
 **channel** | **string**| ID of conversation to join | 

### Return type

[**ConversationsJoinSuccessSchema**](conversations.join success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsKick**
> ConversationsKickSuccessSchema ConversationsKick(ctx, optional)


Removes a user from a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **user** | **string**| User ID to be removed. | 
 **channel** | **string**| ID of conversation to remove user from. | 

### Return type

[**ConversationsKickSuccessSchema**](conversations.kick success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsLeave**
> ConversationsLeaveSuccessSchema ConversationsLeave(ctx, optional)


Leaves a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **string**| Conversation to leave | 

### Return type

[**ConversationsLeaveSuccessSchema**](conversations.leave success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsList**
> ConversationsListSuccessSchema ConversationsList(ctx, optional)


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
 **cursor** | **string**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:read&#x60; | 
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. Must be an integer no larger than 1000. | 
 **excludeArchived** | **bool**| Set to &#x60;true&#x60; to exclude archived channels from the list | 
 **types** | **string**| Mix and match channel types by providing a comma-separated list of any combination of &#x60;public_channel&#x60;, &#x60;private_channel&#x60;, &#x60;mpim&#x60;, &#x60;im&#x60; | 

### Return type

[**ConversationsListSuccessSchema**](conversations.list success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsMembers**
> ConversationsMembersSuccessSchema ConversationsMembers(ctx, optional)


Retrieve members of a conversation.

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
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **channel** | **string**| ID of the conversation to retrieve members for | 

### Return type

[**ConversationsMembersSuccessSchema**](conversations.members success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsOpen**
> ConversationsOpenSuccessSchema ConversationsOpen(ctx, optional)


Opens or resumes a direct message or multi-person direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **returnIm** | **bool**| Boolean, indicates you want the full IM channel definition in the response. | 
 **users** | **string**| Comma separated lists of users. If only one user is included, this creates a 1:1 DM.  The ordering of the users is preserved whenever a multi-person direct message is returned. Supply a &#x60;channel&#x60; when not supplying &#x60;users&#x60;. | 
 **channel** | **string**| Resume a conversation by supplying an &#x60;im&#x60; or &#x60;mpim&#x60;&#39;s ID. Or provide the &#x60;users&#x60; field instead. | 

### Return type

[**ConversationsOpenSuccessSchema**](conversations.open success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsRename**
> ConversationsRenameSuccessSchema ConversationsRename(ctx, optional)


Renames a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **name** | **string**| New name for conversation. | 
 **channel** | **string**| ID of conversation to rename | 

### Return type

[**ConversationsRenameSuccessSchema**](conversations.rename success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsReplies**
> ConversationsRepliesSuccessSchema ConversationsReplies(ctx, optional)


Retrieve a thread of messages posted to a conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inclusive** | **bool**| Include messages with latest or oldest timestamp in results only when either timestamp is specified. | 
 **ts** | **float32**| Unique identifier of a thread&#39;s parent message. | 
 **cursor** | **string**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:history&#x60; | 
 **limit** | **int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **oldest** | **float32**| Start of time range of messages to include in results. | 
 **channel** | **string**| Conversation ID to fetch thread from. | 
 **latest** | **float32**| End of time range of messages to include in results. | 

### Return type

[**ConversationsRepliesSuccessSchema**](conversations.replies success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsSetPurpose**
> ConversationsSetPurposeSuccessSchema ConversationsSetPurpose(ctx, optional)


Sets the purpose for a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **purpose** | **string**| A new, specialer purpose | 
 **channel** | **string**| Conversation to set the purpose of | 

### Return type

[**ConversationsSetPurposeSuccessSchema**](conversations.setPurpose success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsSetTopic**
> ConversationsSetTopicSuccessSchema ConversationsSetTopic(ctx, optional)


Sets the topic for a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **string**| The new topic string. Does not support formatting or linkification. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **string**| Conversation to set the topic of | 

### Return type

[**ConversationsSetTopicSuccessSchema**](conversations.setTopic success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsUnarchive**
> ConversationsUnarchiveSuccessSchema ConversationsUnarchive(ctx, optional)


Reverses conversation archival.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **string**| ID of conversation to unarchive | 

### Return type

[**ConversationsUnarchiveSuccessSchema**](conversations.unarchive success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

