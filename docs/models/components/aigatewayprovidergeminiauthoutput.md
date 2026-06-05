# AIGatewayProviderGeminiAuthOutput


## Supported Types

### AIGatewayProviderConfigAuthBasicOutput

```go
aiGatewayProviderGeminiAuthOutput := components.CreateAIGatewayProviderGeminiAuthOutputBasic(components.AIGatewayProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayProviderConfigAuthGCPOutput

```go
aiGatewayProviderGeminiAuthOutput := components.CreateAIGatewayProviderGeminiAuthOutputGcp(components.AIGatewayProviderConfigAuthGCPOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderGeminiAuthOutput.Type {
	case components.AIGatewayProviderGeminiAuthOutputTypeBasic:
		// aiGatewayProviderGeminiAuthOutput.AIGatewayProviderConfigAuthBasicOutput is populated
	case components.AIGatewayProviderGeminiAuthOutputTypeGcp:
		// aiGatewayProviderGeminiAuthOutput.AIGatewayProviderConfigAuthGCPOutput is populated
}
```
