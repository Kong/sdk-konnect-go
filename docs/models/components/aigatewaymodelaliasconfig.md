# AIGatewayModelAliasConfig

Configuration for routing to this model using an alias.


## Supported Types

### AIGatewayModelAliasConfigPath

```go
aiGatewayModelAliasConfig := components.CreateAIGatewayModelAliasConfigAIGatewayModelAliasConfigPath(components.AIGatewayModelAliasConfigPath{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelAliasConfig.Type {
	case components.AIGatewayModelAliasConfigTypeAIGatewayModelAliasConfigPath:
		// aiGatewayModelAliasConfig.AIGatewayModelAliasConfigPath is populated
}
```
