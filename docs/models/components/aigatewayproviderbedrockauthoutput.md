# AIGatewayProviderBedrockAuthOutput


## Supported Types

### AIGatewayProviderConfigAuthBasicOutput

```go
aiGatewayProviderBedrockAuthOutput := components.CreateAIGatewayProviderBedrockAuthOutputBasic(components.AIGatewayProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayProviderConfigAuthAWSOutput

```go
aiGatewayProviderBedrockAuthOutput := components.CreateAIGatewayProviderBedrockAuthOutputAws(components.AIGatewayProviderConfigAuthAWSOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderBedrockAuthOutput.Type {
	case components.AIGatewayProviderBedrockAuthOutputTypeBasic:
		// aiGatewayProviderBedrockAuthOutput.AIGatewayProviderConfigAuthBasicOutput is populated
	case components.AIGatewayProviderBedrockAuthOutputTypeAws:
		// aiGatewayProviderBedrockAuthOutput.AIGatewayProviderConfigAuthAWSOutput is populated
}
```
