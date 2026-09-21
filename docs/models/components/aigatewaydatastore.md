# AIGatewayDatastore

Configuration for an AI Gateway Datastore.


## Supported Types

### AIGatewayDatastoreRedisCEDatastore

```go
aiGatewayDatastore := components.CreateAIGatewayDatastoreRedisCe(components.AIGatewayDatastoreRedisCEDatastore{/* values here */})
```

### AIGatewayDatastoreRedisEEDatastore

```go
aiGatewayDatastore := components.CreateAIGatewayDatastoreRedisEe(components.AIGatewayDatastoreRedisEEDatastore{/* values here */})
```

### AIGatewayDatastoreVectorDBDatastore

```go
aiGatewayDatastore := components.CreateAIGatewayDatastoreVectordb(components.AIGatewayDatastoreVectorDBDatastore{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayDatastore.Type {
	case components.AIGatewayDatastoreTypeRedisCe:
		// aiGatewayDatastore.AIGatewayDatastoreRedisCEDatastore is populated
	case components.AIGatewayDatastoreTypeRedisEe:
		// aiGatewayDatastore.AIGatewayDatastoreRedisEEDatastore is populated
	case components.AIGatewayDatastoreTypeVectordb:
		// aiGatewayDatastore.AIGatewayDatastoreVectorDBDatastore is populated
}
```
