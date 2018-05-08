# \ChatApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDelete**](ChatApi.md#ChatDelete) | **Post** /chat.delete | 
[**ChatGetPermalink**](ChatApi.md#ChatGetPermalink) | **Get** /chat.getPermalink | 
[**ChatMeMessage**](ChatApi.md#ChatMeMessage) | **Post** /chat.meMessage | 
[**ChatPostEphemeral**](ChatApi.md#ChatPostEphemeral) | **Post** /chat.postEphemeral | 
[**ChatPostMessage**](ChatApi.md#ChatPostMessage) | **Post** /chat.postMessage | 
[**ChatUnfurl**](ChatApi.md#ChatUnfurl) | **Post** /chat.unfurl | 
[**ChatUpdate**](ChatApi.md#ChatUpdate) | **Post** /chat.update | 


# **ChatDelete**
> ChatDeleteSuccessSchema ChatDelete(ctx, optional)


Deletes a message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asUser** | **bool**| Pass true to delete the message as the authed user with &#x60;chat:write:user&#x60; scope. [Bot users](/bot-users) in this context are considered authed users. If unused or false, the message will be deleted with &#x60;chat:write:bot&#x60; scope. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **ts** | **float32**| Timestamp of the message to be deleted. | 
 **channel** | **string**| Channel containing the message to be deleted. | 

### Return type

[**ChatDeleteSuccessSchema**](chat.delete success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatGetPermalink**
> ChatGetPermalinkSuccessSchema ChatGetPermalink(ctx, optional)


Retrieve a permalink URL for a specific extant message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **messageTs** | **float32**| A message&#39;s &#x60;ts&#x60; value, uniquely identifying it within a channel | 
 **channel** | **string**| The ID of the conversation or channel containing the message | 

### Return type

[**ChatGetPermalinkSuccessSchema**](chat.getPermalink success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatMeMessage**
> DefaultSuccessTemplate ChatMeMessage(ctx, optional)


Share a me message into a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string**| Text of the message to send. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;chat:write:user&#x60; | 
 **channel** | **string**| Channel to send message to. Can be a public channel, private group or IM channel. Can be an encoded ID, or a name. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatPostEphemeral**
> ChatPostEphemeralSuccessSchema ChatPostEphemeral(ctx, optional)


Sends an ephemeral message to a user in a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachments** | **string**| A JSON-based array of structured attachments, presented as a URL-encoded string. | 
 **text** | **string**| Text of the message to send. See below for an explanation of [formatting](#formatting). This field is usually required, unless you&#39;re providing only &#x60;attachments&#x60; instead. | 
 **linkNames** | **bool**| Find and link channel names and usernames. | 
 **parse** | **string**| Change how messages are treated. Defaults to &#x60;none&#x60;. See [below](#formatting). | 
 **token** | **string**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **user** | **string**| &#x60;id&#x60; of the user who will receive the ephemeral message. The user should be in the channel specified by the &#x60;channel&#x60; argument. | 
 **asUser** | **bool**| Pass true to post the message as the authed bot. Defaults to false. | 
 **channel** | **string**| Channel, private group, or IM channel to send message to. Can be an encoded ID, or a name. | 

### Return type

[**ChatPostEphemeralSuccessSchema**](chat.postEphemeral success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatPostMessage**
> ChatPostMessageSuccessSchema ChatPostMessage(ctx, optional)


Sends a message to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Set your bot&#39;s user name. Must be used in conjunction with &#x60;as_user&#x60; set to false, otherwise ignored. See [authorship](#authorship) below. | 
 **threadTs** | **float32**| Provide another message&#39;s &#x60;ts&#x60; value to make this message a reply. Avoid using a reply&#39;s &#x60;ts&#x60; value; use its parent instead. | 
 **attachments** | **string**| A JSON-based array of structured attachments, presented as a URL-encoded string. | 
 **unfurlLinks** | **bool**| Pass true to enable unfurling of primarily text-based content. | 
 **text** | **string**| Text of the message to send. See below for an explanation of [formatting](#formatting). This field is usually required, unless you&#39;re providing only &#x60;attachments&#x60; instead. | 
 **unfurlMedia** | **bool**| Pass false to disable unfurling of media content. | 
 **parse** | **string**| Change how messages are treated. Defaults to &#x60;none&#x60;. See [below](#formatting). | 
 **asUser** | **bool**| Pass true to post the message as the authed user, instead of as a bot. Defaults to false. See [authorship](#authorship) below. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **mrkdwn** | **bool**| Disable Slack markup parsing by setting to &#x60;false&#x60;. Enabled by default. | 
 **iconEmoji** | **string**| Emoji to use as the icon for this message. Overrides &#x60;icon_url&#x60;. Must be used in conjunction with &#x60;as_user&#x60; set to &#x60;false&#x60;, otherwise ignored. See [authorship](#authorship) below. | 
 **linkNames** | **bool**| Find and link channel names and usernames. | 
 **iconUrl** | **string**| URL to an image to use as the icon for this message. Must be used in conjunction with &#x60;as_user&#x60; set to false, otherwise ignored. See [authorship](#authorship) below. | 
 **channel** | **string**| Channel, private group, or IM channel to send message to. Can be an encoded ID, or a name. See [below](#channels) for more details. | 
 **replyBroadcast** | **bool**| Used in conjunction with &#x60;thread_ts&#x60; and indicates whether reply should be made visible to everyone in the channel or conversation. Defaults to &#x60;false&#x60;. | 

### Return type

[**ChatPostMessageSuccessSchema**](chat.postMessage success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatUnfurl**
> ChatUnfurlSuccessSchema ChatUnfurl(ctx, optional)


Provide custom unfurl behavior for user-posted URLs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAuthMessage** | **string**| Provide a simply-formatted string to send as an ephemeral message to the user as invitation to authenticate further and enable full unfurling behavior | 
 **userAuthRequired** | **bool**| Set to &#x60;true&#x60; or &#x60;1&#x60; to indicate the user must install your Slack app to trigger unfurls for this domain | 
 **unfurls** | **string**| URL-encoded JSON map with keys set to URLs featured in the the message, pointing to their unfurl message attachments. | 
 **ts** | **string**| Timestamp of the message to add unfurl behavior to. | 
 **userAuthUrl** | **string**| Send users to this custom URL where they will complete authentication in your app to fully trigger unfurling. Value should be properly URL-encoded. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;links:write&#x60; | 
 **channel** | **string**| Channel ID of the message | 

### Return type

[**ChatUnfurlSuccessSchema**](chat.unfurl success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatUpdate**
> ChatUpdateSuccessSchema ChatUpdate(ctx, optional)


Updates a message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachments** | **string**| A JSON-based array of structured attachments, presented as a URL-encoded string. This field is required when not presenting &#x60;text&#x60;. | 
 **text** | **string**| New text for the message, using the [default formatting rules](/docs/formatting). It&#39;s not required when presenting &#x60;attachments&#x60;. | 
 **ts** | **float32**| Timestamp of the message to be updated. | 
 **parse** | **string**| Change how messages are treated. Defaults to &#x60;client&#x60;, unlike &#x60;chat.postMessage&#x60;. See [below](#formatting). | 
 **asUser** | **bool**| Pass true to update the message as the authed user. [Bot users](/bot-users) in this context are considered authed users. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **linkNames** | **bool**| Find and link channel names and usernames. Defaults to &#x60;none&#x60;. This parameter should be used in conjunction with &#x60;parse&#x60;. To set &#x60;link_names&#x60; to &#x60;1&#x60;, specify a &#x60;parse&#x60; mode of &#x60;full&#x60;. | 
 **channel** | **string**| Channel containing the message to be updated. | 

### Return type

[**ChatUpdateSuccessSchema**](chat.update success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

