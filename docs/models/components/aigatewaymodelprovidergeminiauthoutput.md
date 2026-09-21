# AIGatewayModelProviderGeminiAuthOutput


## Supported Types

### AIGatewayModelProviderConfigAuthBasicOutput

```go
aiGatewayModelProviderGeminiAuthOutput := components.CreateAIGatewayModelProviderGeminiAuthOutputBasic(components.AIGatewayModelProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayModelProviderConfigAuthGCPOutput

```go
aiGatewayModelProviderGeminiAuthOutput := components.CreateAIGatewayModelProviderGeminiAuthOutputGcp(components.AIGatewayModelProviderConfigAuthGCPOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderGeminiAuthOutput.Type {
	case components.AIGatewayModelProviderGeminiAuthOutputTypeBasic:
		// aiGatewayModelProviderGeminiAuthOutput.AIGatewayModelProviderConfigAuthBasicOutput is populated
	case components.AIGatewayModelProviderGeminiAuthOutputTypeGcp:
		// aiGatewayModelProviderGeminiAuthOutput.AIGatewayModelProviderConfigAuthGCPOutput is populated
}
```
