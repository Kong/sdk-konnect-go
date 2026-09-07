# AIGatewayModelVectorDBConfigOutput

Configuration for the vector database used by the model.


## Supported Types

### AIGatewayModelVectorDBConfigPgVectorOutput

```go
aiGatewayModelVectorDBConfigOutput := components.CreateAIGatewayModelVectorDBConfigOutputPgvector(components.AIGatewayModelVectorDBConfigPgVectorOutput{/* values here */})
```

### AIGatewayModelVectorDBConfigRedisOutput

```go
aiGatewayModelVectorDBConfigOutput := components.CreateAIGatewayModelVectorDBConfigOutputRedis(components.AIGatewayModelVectorDBConfigRedisOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelVectorDBConfigOutput.Type {
	case components.AIGatewayModelVectorDBConfigOutputTypePgvector:
		// aiGatewayModelVectorDBConfigOutput.AIGatewayModelVectorDBConfigPgVectorOutput is populated
	case components.AIGatewayModelVectorDBConfigOutputTypeRedis:
		// aiGatewayModelVectorDBConfigOutput.AIGatewayModelVectorDBConfigRedisOutput is populated
}
```
