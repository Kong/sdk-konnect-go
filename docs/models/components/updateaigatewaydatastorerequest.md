# UpdateAIGatewayDatastoreRequest

Configuration for an AI Gateway Datastore.


## Supported Types

### RedisCEDatastore

```go
updateAIGatewayDatastoreRequest := components.CreateUpdateAIGatewayDatastoreRequestRedisCe(components.RedisCEDatastore{/* values here */})
```

### RedisEEDatastore

```go
updateAIGatewayDatastoreRequest := components.CreateUpdateAIGatewayDatastoreRequestRedisEe(components.RedisEEDatastore{/* values here */})
```

### VectorDBDatastore

```go
updateAIGatewayDatastoreRequest := components.CreateUpdateAIGatewayDatastoreRequestVectordb(components.VectorDBDatastore{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayDatastoreRequest.Type {
	case components.UpdateAIGatewayDatastoreRequestTypeRedisCe:
		// updateAIGatewayDatastoreRequest.RedisCEDatastore is populated
	case components.UpdateAIGatewayDatastoreRequestTypeRedisEe:
		// updateAIGatewayDatastoreRequest.RedisEEDatastore is populated
	case components.UpdateAIGatewayDatastoreRequestTypeVectordb:
		// updateAIGatewayDatastoreRequest.VectorDBDatastore is populated
}
```
