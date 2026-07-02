# AIGatewayProviderVertexAuth


## Supported Types

### AIGatewayProviderConfigAuthBasic

```go
aiGatewayProviderVertexAuth := components.CreateAIGatewayProviderVertexAuthBasic(components.AIGatewayProviderConfigAuthBasic{/* values here */})
```

### AIGatewayProviderConfigAuthGCP

```go
aiGatewayProviderVertexAuth := components.CreateAIGatewayProviderVertexAuthGcp(components.AIGatewayProviderConfigAuthGCP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProviderVertexAuth.Type {
	case components.AIGatewayProviderVertexAuthTypeBasic:
		// aiGatewayProviderVertexAuth.AIGatewayProviderConfigAuthBasic is populated
	case components.AIGatewayProviderVertexAuthTypeGcp:
		// aiGatewayProviderVertexAuth.AIGatewayProviderConfigAuthGCP is populated
}
```
