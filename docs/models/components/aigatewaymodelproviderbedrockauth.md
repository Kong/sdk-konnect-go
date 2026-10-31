# AIGatewayModelProviderBedrockAuth


## Supported Types

### AIGatewayModelProviderConfigAuthBasic

```go
aiGatewayModelProviderBedrockAuth := components.CreateAIGatewayModelProviderBedrockAuthBasic(components.AIGatewayModelProviderConfigAuthBasic{/* values here */})
```

### AIGatewayModelProviderConfigAuthAWS

```go
aiGatewayModelProviderBedrockAuth := components.CreateAIGatewayModelProviderBedrockAuthAws(components.AIGatewayModelProviderConfigAuthAWS{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderBedrockAuth.Type {
	case components.AIGatewayModelProviderBedrockAuthTypeBasic:
		// aiGatewayModelProviderBedrockAuth.AIGatewayModelProviderConfigAuthBasic is populated
	case components.AIGatewayModelProviderBedrockAuthTypeAws:
		// aiGatewayModelProviderBedrockAuth.AIGatewayModelProviderConfigAuthAWS is populated
}
```
