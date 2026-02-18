# EventGatewayProducePolicyUpdate

The typed schema of the produce policy to modify it.


## Supported Types

### EventGatewayModifyHeadersPolicy

```go
eventGatewayProducePolicyUpdate := components.CreateEventGatewayProducePolicyUpdateModifyHeaders(components.EventGatewayModifyHeadersPolicy{/* values here */})
```

### EventGatewayProduceSchemaValidationPolicy

```go
eventGatewayProducePolicyUpdate := components.CreateEventGatewayProducePolicyUpdateSchemaValidation(components.EventGatewayProduceSchemaValidationPolicy{/* values here */})
```

### EventGatewayEncryptPolicy

```go
eventGatewayProducePolicyUpdate := components.CreateEventGatewayProducePolicyUpdateEncrypt(components.EventGatewayEncryptPolicy{/* values here */})
```

### EventGatewaySkipRecordPolicy

```go
eventGatewayProducePolicyUpdate := components.CreateEventGatewayProducePolicyUpdateSkipRecord(components.EventGatewaySkipRecordPolicy{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayProducePolicyUpdate.Type {
	case components.EventGatewayProducePolicyUpdateTypeModifyHeaders:
		// eventGatewayProducePolicyUpdate.EventGatewayModifyHeadersPolicy is populated
	case components.EventGatewayProducePolicyUpdateTypeSchemaValidation:
		// eventGatewayProducePolicyUpdate.EventGatewayProduceSchemaValidationPolicy is populated
	case components.EventGatewayProducePolicyUpdateTypeEncrypt:
		// eventGatewayProducePolicyUpdate.EventGatewayEncryptPolicy is populated
	case components.EventGatewayProducePolicyUpdateTypeSkipRecord:
		// eventGatewayProducePolicyUpdate.EventGatewaySkipRecordPolicy is populated
}
```
