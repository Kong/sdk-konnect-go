# CreateAIGatewayDatastoreRequest

Configuration for an AI Gateway Datastore.


## Supported Types

### RedisCEDatastore

```go
createAIGatewayDatastoreRequest := components.CreateCreateAIGatewayDatastoreRequestRedisCe(components.RedisCEDatastore{/* values here */})
```

### RedisEEDatastore

```go
createAIGatewayDatastoreRequest := components.CreateCreateAIGatewayDatastoreRequestRedisEe(components.RedisEEDatastore{/* values here */})
```

### VectorDBDatastore

```go
createAIGatewayDatastoreRequest := components.CreateCreateAIGatewayDatastoreRequestVectordb(components.VectorDBDatastore{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayDatastoreRequest.Type {
	case components.CreateAIGatewayDatastoreRequestTypeRedisCe:
		// createAIGatewayDatastoreRequest.RedisCEDatastore is populated
	case components.CreateAIGatewayDatastoreRequestTypeRedisEe:
		// createAIGatewayDatastoreRequest.RedisEEDatastore is populated
	case components.CreateAIGatewayDatastoreRequestTypeVectordb:
		// createAIGatewayDatastoreRequest.VectorDBDatastore is populated
}
```
