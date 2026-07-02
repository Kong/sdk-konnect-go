# AIGatewayProviderVertexAuthOutput


## Supported Types

### AIGatewayProviderConfigAuthBasicOutput

```go
aiGatewayProviderVertexAuthOutput := components.CreateAIGatewayProviderVertexAuthOutputBasic(components.AIGatewayProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayProviderConfigAuthGCPOutput

```go
aiGatewayProviderVertexAuthOutput := components.CreateAIGatewayProviderVertexAuthOutputGcp(components.AIGatewayProviderConfigAuthGCPOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderVertexAuthOutput.Type {
	case components.AIGatewayProviderVertexAuthOutputTypeBasic:
		// aiGatewayProviderVertexAuthOutput.AIGatewayProviderConfigAuthBasicOutput is populated
	case components.AIGatewayProviderVertexAuthOutputTypeGcp:
		// aiGatewayProviderVertexAuthOutput.AIGatewayProviderConfigAuthGCPOutput is populated
}
```
