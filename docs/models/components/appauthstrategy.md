# AppAuthStrategy

A set of plugin configurations that represent how the gateway will perform authentication and authorization for a Product Version. Called “Auth Strategy” for short in the context of portals/applications. The plugins are synced to any Gateway Service that is currently linked or becomes linked to the Product Version.


## Supported Types

### AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse

```go
appAuthStrategy := components.CreateAppAuthStrategyKeyAuth(components.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse{/* values here */})
```

### AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse

```go
appAuthStrategy := components.CreateAppAuthStrategyOpenidConnect(components.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch appAuthStrategy.Type {
	case components.AppAuthStrategyTypeKeyAuth:
		// appAuthStrategy.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse is populated
	case components.AppAuthStrategyTypeOpenidConnect:
		// appAuthStrategy.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse is populated
}
```
