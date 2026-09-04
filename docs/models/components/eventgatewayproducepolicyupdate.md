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

### EventGatewayParsedRecordEncryptFieldsPolicy

```go
eventGatewayProducePolicyUpdate := components.CreateEventGatewayProducePolicyUpdateEncryptFields(components.EventGatewayParsedRecordEncryptFieldsPolicy{/* values here */})
```

### EventGatewayParsedRecordTranscodeProducePolicy

```go
eventGatewayProducePolicyUpdate := components.CreateEventGatewayProducePolicyUpdateTranscode(components.EventGatewayParsedRecordTranscodeProducePolicy{/* values here */})
```

### EventGatewayParsedRecordMaskFieldsProducePolicy

```go
eventGatewayProducePolicyUpdate := components.CreateEventGatewayProducePolicyUpdateMaskFields(components.EventGatewayParsedRecordMaskFieldsProducePolicy{/* values here */})
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
	case components.EventGatewayProducePolicyUpdateTypeEncryptFields:
		// eventGatewayProducePolicyUpdate.EventGatewayParsedRecordEncryptFieldsPolicy is populated
	case components.EventGatewayProducePolicyUpdateTypeTranscode:
		// eventGatewayProducePolicyUpdate.EventGatewayParsedRecordTranscodeProducePolicy is populated
	case components.EventGatewayProducePolicyUpdateTypeMaskFields:
		// eventGatewayProducePolicyUpdate.EventGatewayParsedRecordMaskFieldsProducePolicy is populated
}
```
