# AIGatewayModelProviderBedrockAuthOutput


## Supported Types

### AIGatewayModelProviderConfigAuthBasicOutput

```go
aiGatewayModelProviderBedrockAuthOutput := components.CreateAIGatewayModelProviderBedrockAuthOutputBasic(components.AIGatewayModelProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayModelProviderConfigAuthAWSOutput

```go
aiGatewayModelProviderBedrockAuthOutput := components.CreateAIGatewayModelProviderBedrockAuthOutputAws(components.AIGatewayModelProviderConfigAuthAWSOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderBedrockAuthOutput.Type {
	case components.AIGatewayModelProviderBedrockAuthOutputTypeBasic:
		// aiGatewayModelProviderBedrockAuthOutput.AIGatewayModelProviderConfigAuthBasicOutput is populated
	case components.AIGatewayModelProviderBedrockAuthOutputTypeAws:
		// aiGatewayModelProviderBedrockAuthOutput.AIGatewayModelProviderConfigAuthAWSOutput is populated
}
```
