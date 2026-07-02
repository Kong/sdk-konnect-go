# AIGatewayProviderAzureAuth


## Supported Types

### AIGatewayProviderConfigAuthBasic

```go
aiGatewayProviderAzureAuth := components.CreateAIGatewayProviderAzureAuthBasic(components.AIGatewayProviderConfigAuthBasic{/* values here */})
```

### AIGatewayProviderConfigAuthAzure

```go
aiGatewayProviderAzureAuth := components.CreateAIGatewayProviderAzureAuthAzure(components.AIGatewayProviderConfigAuthAzure{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderAzureAuth.Type {
	case components.AIGatewayProviderAzureAuthTypeBasic:
		// aiGatewayProviderAzureAuth.AIGatewayProviderConfigAuthBasic is populated
	case components.AIGatewayProviderAzureAuthTypeAzure:
		// aiGatewayProviderAzureAuth.AIGatewayProviderConfigAuthAzure is populated
}
```
