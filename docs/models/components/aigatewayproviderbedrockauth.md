# AIGatewayProviderBedrockAuth


## Supported Types

### AIGatewayProviderConfigAuthBasic

```go
aiGatewayProviderBedrockAuth := components.CreateAIGatewayProviderBedrockAuthBasic(components.AIGatewayProviderConfigAuthBasic{/* values here */})
```

### AIGatewayProviderConfigAuthAWS

```go
aiGatewayProviderBedrockAuth := components.CreateAIGatewayProviderBedrockAuthAws(components.AIGatewayProviderConfigAuthAWS{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderBedrockAuth.Type {
	case components.AIGatewayProviderBedrockAuthTypeBasic:
		// aiGatewayProviderBedrockAuth.AIGatewayProviderConfigAuthBasic is populated
	case components.AIGatewayProviderBedrockAuthTypeAws:
		// aiGatewayProviderBedrockAuth.AIGatewayProviderConfigAuthAWS is populated
}
```
