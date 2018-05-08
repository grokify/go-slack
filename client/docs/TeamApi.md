# \TeamApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamAccessLogs**](TeamApi.md#TeamAccessLogs) | **Get** /team.accessLogs | 
[**TeamBillableInfo**](TeamApi.md#TeamBillableInfo) | **Get** /team.billableInfo | 
[**TeamInfo**](TeamApi.md#TeamInfo) | **Get** /team.info | 
[**TeamIntegrationLogs**](TeamApi.md#TeamIntegrationLogs) | **Get** /team.integrationLogs | 
[**TeamProfileGet**](TeamApi.md#TeamProfileGet) | **Get** /team.profile.get | 


# **TeamAccessLogs**
> DefaultSuccessTemplate TeamAccessLogs(ctx, optional)


Gets the access logs for the current team.

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
 **token** | **string**| Authentication token. Requires scope: &#x60;admin&#x60; | 
 **page** | **string**|  | 
 **before** | **int32**| End of time range of logs to include in results (inclusive). | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamBillableInfo**
> DefaultSuccessTemplate TeamBillableInfo(ctx, optional)


Gets billable users information for the current team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;admin&#x60; | 
 **user** | **string**| A user to retrieve the billable information for. Defaults to all users. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamInfo**
> TeamInfoSchema TeamInfo(ctx, optional)


Gets information about the current team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Authentication token. Requires scope: &#x60;team:read&#x60; | 

### Return type

[**TeamInfoSchema**](team.info schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamIntegrationLogs**
> DefaultSuccessTemplate TeamIntegrationLogs(ctx, optional)


Gets the integration logs for the current team.

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
 **changeType** | **string**| Filter logs with this change type. Defaults to all logs. | 
 **appId** | **int32**| Filter logs to this Slack app. Defaults to all logs. | 
 **token** | **string**| Authentication token. Requires scope: &#x60;admin&#x60; | 
 **user** | **string**| Filter logs generated by this user’s actions. Defaults to all logs. | 
 **serviceId** | **int32**| Filter logs to this service. Defaults to all logs. | 
 **page** | **string**|  | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamProfileGet**
> TeamProfileGetSuccessSchema TeamProfileGet(ctx, optional)


Retrieve a team's profile.

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
 **visibility** | **string**| Filter by visibility. | 

### Return type

[**TeamProfileGetSuccessSchema**](team.profile.get success schema.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

