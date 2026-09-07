# UpdateAIGatewayAuthStrategyRequest


## Supported Types

### AIGatewayAuthStrategyKeyAuth

```go
updateAIGatewayAuthStrategyRequest := components.CreateUpdateAIGatewayAuthStrategyRequestKeyAuth(components.AIGatewayAuthStrategyKeyAuth{/* values here */})
```

### AIGatewayAuthStrategyOpenIDConnect

```go
updateAIGatewayAuthStrategyRequest := components.CreateUpdateAIGatewayAuthStrategyRequestOpenidConnect(components.AIGatewayAuthStrategyOpenIDConnect{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayAuthStrategyRequest.Type {
	case components.UpdateAIGatewayAuthStrategyRequestTypeKeyAuth:
		// updateAIGatewayAuthStrategyRequest.AIGatewayAuthStrategyKeyAuth is populated
	case components.UpdateAIGatewayAuthStrategyRequestTypeOpenidConnect:
		// updateAIGatewayAuthStrategyRequest.AIGatewayAuthStrategyOpenIDConnect is populated
}
```
