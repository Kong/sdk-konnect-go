# CreateAppAuthStrategyRequest

Request body for creating an Application Auth Strategy


## Supported Types

### AppAuthStrategyKeyAuthRequest

```go
createAppAuthStrategyRequest := components.CreateCreateAppAuthStrategyRequestKeyAuth(components.AppAuthStrategyKeyAuthRequest{/* values here */})
```

### AppAuthStrategyOpenIDConnectRequest

```go
createAppAuthStrategyRequest := components.CreateCreateAppAuthStrategyRequestOpenidConnect(components.AppAuthStrategyOpenIDConnectRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAppAuthStrategyRequest.Type {
	case components.CreateAppAuthStrategyRequestTypeKeyAuth:
		// createAppAuthStrategyRequest.AppAuthStrategyKeyAuthRequest is populated
	case components.CreateAppAuthStrategyRequestTypeOpenidConnect:
		// createAppAuthStrategyRequest.AppAuthStrategyOpenIDConnectRequest is populated
}
```
