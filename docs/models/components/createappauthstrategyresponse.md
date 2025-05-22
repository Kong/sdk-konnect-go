# CreateAppAuthStrategyResponse

A set of plugin configurations that represent how the gateway will perform authentication and authorization for a Product Version. Called “Auth Strategy” for short in the context of portals/applications. The plugins are synced to any Gateway Service that is currently linked or becomes linked to the Product Version.


## Supported Types

### AppAuthStrategyKeyAuthResponse

```go
createAppAuthStrategyResponse := components.CreateCreateAppAuthStrategyResponseKeyAuth(components.AppAuthStrategyKeyAuthResponse{/* values here */})
```

### AppAuthStrategyOpenIDConnectResponse

```go
createAppAuthStrategyResponse := components.CreateCreateAppAuthStrategyResponseOpenidConnect(components.AppAuthStrategyOpenIDConnectResponse{/* values here */})
```

