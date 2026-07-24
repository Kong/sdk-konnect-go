# AIGatewayModelProviderVertexAuth


## Supported Types

### AIGatewayModelProviderConfigAuthBasic

```go
aiGatewayModelProviderVertexAuth := components.CreateAIGatewayModelProviderVertexAuthBasic(components.AIGatewayModelProviderConfigAuthBasic{/* values here */})
```

### AIGatewayModelProviderConfigAuthVertex

```go
aiGatewayModelProviderVertexAuth := components.CreateAIGatewayModelProviderVertexAuthVertex(components.AIGatewayModelProviderConfigAuthVertex{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderVertexAuth.Type {
	case components.AIGatewayModelProviderVertexAuthTypeBasic:
		// aiGatewayModelProviderVertexAuth.AIGatewayModelProviderConfigAuthBasic is populated
	case components.AIGatewayModelProviderVertexAuthTypeVertex:
		// aiGatewayModelProviderVertexAuth.AIGatewayModelProviderConfigAuthVertex is populated
}
```
