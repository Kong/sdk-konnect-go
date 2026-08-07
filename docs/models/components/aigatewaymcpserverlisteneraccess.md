# AIGatewayMCPServerListenerAccess

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayMCPServerListenerConsumer

```go
aiGatewayMCPServerListenerAccess := components.CreateAIGatewayMCPServerListenerAccessConsumer(components.AIGatewayMCPServerListenerConsumer{/* values here */})
```

### AIGatewayMCPServerListenerOauth

```go
aiGatewayMCPServerListenerAccess := components.CreateAIGatewayMCPServerListenerAccessOauthAccessToken(components.AIGatewayMCPServerListenerOauth{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerListenerAccess.Type {
	case components.AIGatewayMCPServerListenerAccessTypeConsumer:
		// aiGatewayMCPServerListenerAccess.AIGatewayMCPServerListenerConsumer is populated
	case components.AIGatewayMCPServerListenerAccessTypeOauthAccessToken:
		// aiGatewayMCPServerListenerAccess.AIGatewayMCPServerListenerOauth is populated
}
```
