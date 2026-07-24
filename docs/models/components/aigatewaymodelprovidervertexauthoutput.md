# AIGatewayModelProviderVertexAuthOutput


## Supported Types

### AIGatewayModelProviderConfigAuthBasicOutput

```go
aiGatewayModelProviderVertexAuthOutput := components.CreateAIGatewayModelProviderVertexAuthOutputBasic(components.AIGatewayModelProviderConfigAuthBasicOutput{/* values here */})
```

### AIGatewayModelProviderConfigAuthVertexOutput

```go
aiGatewayModelProviderVertexAuthOutput := components.CreateAIGatewayModelProviderVertexAuthOutputVertex(components.AIGatewayModelProviderConfigAuthVertexOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderVertexAuthOutput.Type {
	case components.AIGatewayModelProviderVertexAuthOutputTypeBasic:
		// aiGatewayModelProviderVertexAuthOutput.AIGatewayModelProviderConfigAuthBasicOutput is populated
	case components.AIGatewayModelProviderVertexAuthOutputTypeVertex:
		// aiGatewayModelProviderVertexAuthOutput.AIGatewayModelProviderConfigAuthVertexOutput is populated
}
```
