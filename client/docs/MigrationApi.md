# \MigrationApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MigrationExchange**](MigrationApi.md#MigrationExchange) | **Get** /migration.exchange | 


# **MigrationExchange**
> Http200ResponseSchemaForMigrationExchange MigrationExchange(ctx, optional)


For Enterprise Grid workspaces, map local user IDs to global user IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;tokens.basic&#x60; | 
 **toOld** | **bool**| Specify &#x60;true&#x60; to convert &#x60;W&#x60; global user IDs to workspace-specific &#x60;U&#x60; IDs. Defaults to &#x60;false&#x60;. | 
 **users** | **string**| A comma-separated list of user ids, up to 400 per request | 

### Return type

[**Http200ResponseSchemaForMigrationExchange**](HTTP 200 response schema for migration.exchange.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

