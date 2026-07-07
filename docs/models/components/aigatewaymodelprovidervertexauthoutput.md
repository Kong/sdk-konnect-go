# AIGatewayModelProviderVertexAuthOutput


## Supported Types

### AIGatewayModelProviderConfigAuthBasicOutput

```go
aiGatewayModelProviderVertexAuthOutput := components.CreateAIGatewayModelProviderVertexAuthOutputBasic(components.AIGatewayModelProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayModelProviderConfigAuthGCPOutput

```go
aiGatewayModelProviderVertexAuthOutput := components.CreateAIGatewayModelProviderVertexAuthOutputGcp(components.AIGatewayModelProviderConfigAuthGCPOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderVertexAuthOutput.Type {
	case components.AIGatewayModelProviderVertexAuthOutputTypeBasic:
		// aiGatewayModelProviderVertexAuthOutput.AIGatewayModelProviderConfigAuthBasicOutput is populated
	case components.AIGatewayModelProviderVertexAuthOutputTypeGcp:
		// aiGatewayModelProviderVertexAuthOutput.AIGatewayModelProviderConfigAuthGCPOutput is populated
}
```
