# AIGatewayProviderAzureAuthOutput


## Supported Types

### AIGatewayProviderConfigAuthBasicOutput

```go
aiGatewayProviderAzureAuthOutput := components.CreateAIGatewayProviderAzureAuthOutputBasic(components.AIGatewayProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayProviderConfigAuthAzureOutput

```go
aiGatewayProviderAzureAuthOutput := components.CreateAIGatewayProviderAzureAuthOutputAzure(components.AIGatewayProviderConfigAuthAzureOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderAzureAuthOutput.Type {
	case components.AIGatewayProviderAzureAuthOutputTypeBasic:
		// aiGatewayProviderAzureAuthOutput.AIGatewayProviderConfigAuthBasicOutput is populated
	case components.AIGatewayProviderAzureAuthOutputTypeAzure:
		// aiGatewayProviderAzureAuthOutput.AIGatewayProviderConfigAuthAzureOutput is populated
}
```
