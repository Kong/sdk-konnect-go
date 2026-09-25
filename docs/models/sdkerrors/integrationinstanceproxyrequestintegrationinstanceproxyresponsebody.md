# IntegrationInstanceProxyRequestIntegrationInstanceProxyResponseBody

Not Found


## Supported Types

### NotFoundError

```go
integrationInstanceProxyRequestIntegrationInstanceProxyResponseBody := sdkerrors.CreateIntegrationInstanceProxyRequestIntegrationInstanceProxyResponseBodyNotFoundError(sdkerrors.NotFoundError{/* values here */})
```

### IntegrationNotInstalledError

```go
integrationInstanceProxyRequestIntegrationInstanceProxyResponseBody := sdkerrors.CreateIntegrationInstanceProxyRequestIntegrationInstanceProxyResponseBodyIntegrationNotInstalledError(sdkerrors.IntegrationNotInstalledError{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceProxyRequestIntegrationInstanceProxyResponseBody.Type {
	case sdkerrors.IntegrationInstanceProxyRequestIntegrationInstanceProxyResponseBodyTypeNotFoundError:
		// integrationInstanceProxyRequestIntegrationInstanceProxyResponseBody.NotFoundError is populated
	case sdkerrors.IntegrationInstanceProxyRequestIntegrationInstanceProxyResponseBodyTypeIntegrationNotInstalledError:
		// integrationInstanceProxyRequestIntegrationInstanceProxyResponseBody.IntegrationNotInstalledError is populated
}
```
