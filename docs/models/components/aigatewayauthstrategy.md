# AIGatewayAuthStrategy

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayAuthStrategyKeyAuthResponse

```go
aiGatewayAuthStrategy := components.CreateAIGatewayAuthStrategyKeyAuth(components.AIGatewayAuthStrategyKeyAuthResponse{/* values here */})
```

### AIGatewayAuthStrategyOpenIDConnectResponse

```go
aiGatewayAuthStrategy := components.CreateAIGatewayAuthStrategyOpenidConnect(components.AIGatewayAuthStrategyOpenIDConnectResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayAuthStrategy.Type {
	case components.AIGatewayAuthStrategyTypeKeyAuth:
		// aiGatewayAuthStrategy.AIGatewayAuthStrategyKeyAuthResponse is populated
	case components.AIGatewayAuthStrategyTypeOpenidConnect:
		// aiGatewayAuthStrategy.AIGatewayAuthStrategyOpenIDConnectResponse is populated
}
```
