# CreateAIGatewayAuthStrategyRequest


## Supported Types

### AIGatewayAuthStrategyKeyAuth

```go
createAIGatewayAuthStrategyRequest := components.CreateCreateAIGatewayAuthStrategyRequestKeyAuth(components.AIGatewayAuthStrategyKeyAuth{/* values here */})
```

### AIGatewayAuthStrategyOpenIDConnect

```go
createAIGatewayAuthStrategyRequest := components.CreateCreateAIGatewayAuthStrategyRequestOpenidConnect(components.AIGatewayAuthStrategyOpenIDConnect{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayAuthStrategyRequest.Type {
	case components.CreateAIGatewayAuthStrategyRequestTypeKeyAuth:
		// createAIGatewayAuthStrategyRequest.AIGatewayAuthStrategyKeyAuth is populated
	case components.CreateAIGatewayAuthStrategyRequestTypeOpenidConnect:
		// createAIGatewayAuthStrategyRequest.AIGatewayAuthStrategyOpenIDConnect is populated
}
```
