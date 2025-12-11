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

