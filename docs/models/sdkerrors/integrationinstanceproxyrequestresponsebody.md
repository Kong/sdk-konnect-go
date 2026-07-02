# IntegrationInstanceProxyRequestResponseBody

Error response indicating the proxy request could not be sent due to invalid state of the request body or integration instance



## Supported Types

### BadRequestError

```go
integrationInstanceProxyRequestResponseBody := sdkerrors.CreateIntegrationInstanceProxyRequestResponseBodyBadRequestError(components.BadRequestError{/* values here */})
```

### IntegrationUnauthorizedError

```go
integrationInstanceProxyRequestResponseBody := sdkerrors.CreateIntegrationInstanceProxyRequestResponseBodyIntegrationUnauthorizedError(components.IntegrationUnauthorizedError{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceProxyRequestResponseBody.Type {
	case sdkerrors.IntegrationInstanceProxyRequestResponseBodyTypeBadRequestError:
		// integrationInstanceProxyRequestResponseBody.BadRequestError is populated
	case sdkerrors.IntegrationInstanceProxyRequestResponseBodyTypeIntegrationUnauthorizedError:
		// integrationInstanceProxyRequestResponseBody.IntegrationUnauthorizedError is populated
}
```
