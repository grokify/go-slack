# \FilesCommentsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesCommentsAdd**](FilesCommentsApi.md#FilesCommentsAdd) | **Post** /files.comments.add | 
[**FilesCommentsDelete**](FilesCommentsApi.md#FilesCommentsDelete) | **Post** /files.comments.delete | 
[**FilesCommentsEdit**](FilesCommentsApi.md#FilesCommentsEdit) | **Post** /files.comments.edit | 


# **FilesCommentsAdd**
> FilesCommentsAddSchema FilesCommentsAdd(ctx, optional)


Add a comment to an existing file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | **string**| Text of the comment to add. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **file** | **string**| File to add a comment to. | 

### Return type

[**FilesCommentsAddSchema**](files.comments.add schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesCommentsDelete**
> FilesCommentsDeleteSchema FilesCommentsDelete(ctx, optional)


Deletes an existing comment on a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **id** | **string**| The comment to delete. | 
 **file** | **string**| File to delete a comment from. | 

### Return type

[**FilesCommentsDeleteSchema**](files.comments.delete schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesCommentsEdit**
> FilesCommentsEditSchema FilesCommentsEdit(ctx, optional)


Edit an existing file comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **comment** | **string**| Text of the comment to edit. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **id** | **string**| The comment to edit. | 
 **file** | **string**| File containing the comment to edit. | 

### Return type

[**FilesCommentsEditSchema**](files.comments.edit schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

