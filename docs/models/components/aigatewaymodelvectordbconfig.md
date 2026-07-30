# AIGatewayModelVectorDBConfig

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for the vector database used by the model.


## Supported Types

### AIGatewayModelVectorDBConfigPgVector

```go
aiGatewayModelVectorDBConfig := components.CreateAIGatewayModelVectorDBConfigPgvector(components.AIGatewayModelVectorDBConfigPgVector{/* values here */})
```

### AIGatewayModelVectorDBConfigRedis

```go
aiGatewayModelVectorDBConfig := components.CreateAIGatewayModelVectorDBConfigRedis(components.AIGatewayModelVectorDBConfigRedis{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelVectorDBConfig.Type {
	case components.AIGatewayModelVectorDBConfigTypePgvector:
		// aiGatewayModelVectorDBConfig.AIGatewayModelVectorDBConfigPgVector is populated
	case components.AIGatewayModelVectorDBConfigTypeRedis:
		// aiGatewayModelVectorDBConfig.AIGatewayModelVectorDBConfigRedis is populated
}
```
