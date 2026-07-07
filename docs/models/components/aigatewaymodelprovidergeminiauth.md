# AIGatewayModelProviderGeminiAuth


## Supported Types

### AIGatewayModelProviderConfigAuthBasic

```go
aiGatewayModelProviderGeminiAuth := components.CreateAIGatewayModelProviderGeminiAuthBasic(components.AIGatewayModelProviderConfigAuthBasic{/* values here */})
```

### AIGatewayModelProviderConfigAuthGCP

```go
aiGatewayModelProviderGeminiAuth := components.CreateAIGatewayModelProviderGeminiAuthGcp(components.AIGatewayModelProviderConfigAuthGCP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderGeminiAuth.Type {
	case components.AIGatewayModelProviderGeminiAuthTypeBasic:
		// aiGatewayModelProviderGeminiAuth.AIGatewayModelProviderConfigAuthBasic is populated
	case components.AIGatewayModelProviderGeminiAuthTypeGcp:
		// aiGatewayModelProviderGeminiAuth.AIGatewayModelProviderConfigAuthGCP is populated
}
```
