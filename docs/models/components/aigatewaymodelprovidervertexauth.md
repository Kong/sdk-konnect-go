# AIGatewayModelProviderVertexAuth


## Supported Types

### AIGatewayModelProviderConfigAuthBasic

```go
aiGatewayModelProviderVertexAuth := components.CreateAIGatewayModelProviderVertexAuthBasic(components.AIGatewayModelProviderConfigAuthBasic{/* values here */})
```

### AIGatewayModelProviderConfigAuthGCP

```go
aiGatewayModelProviderVertexAuth := components.CreateAIGatewayModelProviderVertexAuthGcp(components.AIGatewayModelProviderConfigAuthGCP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProviderVertexAuth.Type {
	case components.AIGatewayModelProviderVertexAuthTypeBasic:
		// aiGatewayModelProviderVertexAuth.AIGatewayModelProviderConfigAuthBasic is populated
	case components.AIGatewayModelProviderVertexAuthTypeGcp:
		// aiGatewayModelProviderVertexAuth.AIGatewayModelProviderConfigAuthGCP is populated
}
```
