# AIGatewayModelProviderSagemakerAuth


## Supported Types

### AIGatewayModelProviderConfigAuthBasic

```go
aiGatewayModelProviderSagemakerAuth := components.CreateAIGatewayModelProviderSagemakerAuthBasic(components.AIGatewayModelProviderConfigAuthBasic{/* values here */})
```

### AIGatewayModelProviderConfigAuthSagemaker

```go
aiGatewayModelProviderSagemakerAuth := components.CreateAIGatewayModelProviderSagemakerAuthSagemaker(components.AIGatewayModelProviderConfigAuthSagemaker{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderSagemakerAuth.Type {
	case components.AIGatewayModelProviderSagemakerAuthTypeBasic:
		// aiGatewayModelProviderSagemakerAuth.AIGatewayModelProviderConfigAuthBasic is populated
	case components.AIGatewayModelProviderSagemakerAuthTypeSagemaker:
		// aiGatewayModelProviderSagemakerAuth.AIGatewayModelProviderConfigAuthSagemaker is populated
}
```
