# \FilesApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesCommentsAdd**](FilesApi.md#FilesCommentsAdd) | **Post** /files.comments.add | 
[**FilesCommentsDelete**](FilesApi.md#FilesCommentsDelete) | **Post** /files.comments.delete | 
[**FilesCommentsEdit**](FilesApi.md#FilesCommentsEdit) | **Post** /files.comments.edit | 
[**FilesDelete**](FilesApi.md#FilesDelete) | **Post** /files.delete | 
[**FilesInfo**](FilesApi.md#FilesInfo) | **Get** /files.info | 
[**FilesList**](FilesApi.md#FilesList) | **Get** /files.list | 
[**FilesRevokePublicURL**](FilesApi.md#FilesRevokePublicURL) | **Post** /files.revokePublicURL | 
[**FilesSharedPublicURL**](FilesApi.md#FilesSharedPublicURL) | **Post** /files.sharedPublicURL | 
[**FilesUpload**](FilesApi.md#FilesUpload) | **Post** /files.upload | 


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

# **FilesDelete**
> FilesDeleteSchema FilesDelete(ctx, optional)


Deletes a file.

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
 **file** | **string**| ID of file to delete. | 

### Return type

[**FilesDeleteSchema**](files.delete schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesInfo**
> FilesInfoSchema FilesInfo(ctx, optional)


Gets information about a team file.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;files:read&#x60; | 
 **file** | **string**| Specify a file by providing its ID. | 
 **page** | **string**|  | 

### Return type

[**FilesInfoSchema**](files.info schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesList**
> FilesListSchema FilesList(ctx, optional)


Lists & filters team files.

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
 **channel** | **string**| Filter files appearing in a specific channel, indicated by its ID. | 
 **tsTo** | **float32**| Filter files created before this timestamp (inclusive). | 
 **tsFrom** | **float32**| Filter files created after this timestamp (inclusive). | 
 **token** | **string**| Authentication token. Requires scope: &#x60;files:read&#x60; | 
 **user** | **string**| Filter files created by a single user. | 
 **page** | **string**|  | 
 **types** | **string**| Filter files by type:  * &#x60;all&#x60; - All files * &#x60;spaces&#x60; - Posts * &#x60;snippets&#x60; - Snippets * &#x60;images&#x60; - Image files * &#x60;gdocs&#x60; - Google docs * &#x60;zips&#x60; - Zip files * &#x60;pdfs&#x60; - PDF files  You can pass multiple values in the types argument, like &#x60;types&#x3D;spaces,snippets&#x60;.The default value is &#x60;all&#x60;, which does not filter the list. | 

### Return type

[**FilesListSchema**](files.list schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesRevokePublicURL**
> DefaultSuccessTemplate FilesRevokePublicURL(ctx, optional)


Revokes public/external sharing access for a file

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
 **file** | **string**| File to revoke | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesSharedPublicURL**
> DefaultSuccessTemplate FilesSharedPublicURL(ctx, optional)


Enables a file for public/external sharing.

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
 **file** | **string**| File to share | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesUpload**
> FilesUploadSchema FilesUpload(ctx, optional)


Uploads or creates a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channels** | **string**| Comma-separated list of channel names or IDs where the file will be shared. | 
 **title** | **string**| Title of file. | 
 **initialComment** | **string**| Initial comment to add to file. | 
 **filetype** | **string**| A [file type](/types/file#file_types) identifier. | 
 **filename** | **string**| Filename of file. | 
 **content** | **string**| File contents via a POST variable. If omitting this parameter, you must provide a &#x60;file&#x60;. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **file** | **string**| File contents via &#x60;multipart/form-data&#x60;. If omitting this parameter, you must submit &#x60;content&#x60;. | 

### Return type

[**FilesUploadSchema**](files.upload schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

