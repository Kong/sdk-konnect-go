# AIGatewayModelProviderAzureAuthOutput


## Supported Types

### AIGatewayModelProviderConfigAuthBasicOutput

```go
aiGatewayModelProviderAzureAuthOutput := components.CreateAIGatewayModelProviderAzureAuthOutputBasic(components.AIGatewayModelProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayModelProviderConfigAuthAzureOutput

```go
aiGatewayModelProviderAzureAuthOutput := components.CreateAIGatewayModelProviderAzureAuthOutputAzure(components.AIGatewayModelProviderConfigAuthAzureOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderAzureAuthOutput.Type {
	case components.AIGatewayModelProviderAzureAuthOutputTypeBasic:
		// aiGatewayModelProviderAzureAuthOutput.AIGatewayModelProviderConfigAuthBasicOutput is populated
	case components.AIGatewayModelProviderAzureAuthOutputTypeAzure:
		// aiGatewayModelProviderAzureAuthOutput.AIGatewayModelProviderConfigAuthAzureOutput is populated
}
```
