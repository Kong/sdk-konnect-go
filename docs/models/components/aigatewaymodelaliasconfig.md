# AIGatewayModelAliasConfig

Configuration for routing to this model using an alias.


## Supported Types

### AIGatewayModelAliasConfigBody

```go
aiGatewayModelAliasConfig := components.CreateAIGatewayModelAliasConfigAIGatewayModelAliasConfigBody(components.AIGatewayModelAliasConfigBody{/* values here */})
```

### AIGatewayModelAliasConfigHeaders

```go
aiGatewayModelAliasConfig := components.CreateAIGatewayModelAliasConfigAIGatewayModelAliasConfigHeaders(components.AIGatewayModelAliasConfigHeaders{/* values here */})
```

### AIGatewayModelAliasConfigPath

```go
aiGatewayModelAliasConfig := components.CreateAIGatewayModelAliasConfigAIGatewayModelAliasConfigPath(components.AIGatewayModelAliasConfigPath{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelAliasConfig.Type {
	case components.AIGatewayModelAliasConfigTypeAIGatewayModelAliasConfigBody:
		// aiGatewayModelAliasConfig.AIGatewayModelAliasConfigBody is populated
	case components.AIGatewayModelAliasConfigTypeAIGatewayModelAliasConfigHeaders:
		// aiGatewayModelAliasConfig.AIGatewayModelAliasConfigHeaders is populated
	case components.AIGatewayModelAliasConfigTypeAIGatewayModelAliasConfigPath:
		// aiGatewayModelAliasConfig.AIGatewayModelAliasConfigPath is populated
}
```
