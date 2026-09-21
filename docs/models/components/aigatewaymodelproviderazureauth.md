# AIGatewayModelProviderAzureAuth


## Supported Types

### AIGatewayModelProviderConfigAuthBasic

```go
aiGatewayModelProviderAzureAuth := components.CreateAIGatewayModelProviderAzureAuthBasic(components.AIGatewayModelProviderConfigAuthBasic{/* values here */})
```

### AIGatewayModelProviderConfigAuthAzure

```go
aiGatewayModelProviderAzureAuth := components.CreateAIGatewayModelProviderAzureAuthAzure(components.AIGatewayModelProviderConfigAuthAzure{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderAzureAuth.Type {
	case components.AIGatewayModelProviderAzureAuthTypeBasic:
		// aiGatewayModelProviderAzureAuth.AIGatewayModelProviderConfigAuthBasic is populated
	case components.AIGatewayModelProviderAzureAuthTypeAzure:
		// aiGatewayModelProviderAzureAuth.AIGatewayModelProviderConfigAuthAzure is populated
}
```
