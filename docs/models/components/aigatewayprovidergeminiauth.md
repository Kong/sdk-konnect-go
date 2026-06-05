# AIGatewayProviderGeminiAuth


## Supported Types

### AIGatewayProviderConfigAuthBasic

```go
aiGatewayProviderGeminiAuth := components.CreateAIGatewayProviderGeminiAuthBasic(components.AIGatewayProviderConfigAuthBasic{/* values here */})
```

### AIGatewayProviderConfigAuthGCP

```go
aiGatewayProviderGeminiAuth := components.CreateAIGatewayProviderGeminiAuthGcp(components.AIGatewayProviderConfigAuthGCP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderGeminiAuth.Type {
	case components.AIGatewayProviderGeminiAuthTypeBasic:
		// aiGatewayProviderGeminiAuth.AIGatewayProviderConfigAuthBasic is populated
	case components.AIGatewayProviderGeminiAuthTypeGcp:
		// aiGatewayProviderGeminiAuth.AIGatewayProviderConfigAuthGCP is populated
}
```
