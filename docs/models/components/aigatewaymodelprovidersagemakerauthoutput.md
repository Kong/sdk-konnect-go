# AIGatewayModelProviderSagemakerAuthOutput


## Supported Types

### AIGatewayModelProviderConfigAuthBasicOutput

```go
aiGatewayModelProviderSagemakerAuthOutput := components.CreateAIGatewayModelProviderSagemakerAuthOutputBasic(components.AIGatewayModelProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayModelProviderConfigAuthSagemakerOutput

```go
aiGatewayModelProviderSagemakerAuthOutput := components.CreateAIGatewayModelProviderSagemakerAuthOutputSagemaker(components.AIGatewayModelProviderConfigAuthSagemakerOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderSagemakerAuthOutput.Type {
	case components.AIGatewayModelProviderSagemakerAuthOutputTypeBasic:
		// aiGatewayModelProviderSagemakerAuthOutput.AIGatewayModelProviderConfigAuthBasicOutput is populated
	case components.AIGatewayModelProviderSagemakerAuthOutputTypeSagemaker:
		// aiGatewayModelProviderSagemakerAuthOutput.AIGatewayModelProviderConfigAuthSagemakerOutput is populated
}
```
