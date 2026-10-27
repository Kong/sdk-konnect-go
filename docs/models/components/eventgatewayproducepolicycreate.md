# EventGatewayProducePolicyCreate

The typed schema of the produce policy to modify it.


## Supported Types

### EventGatewayModifyHeadersPolicyCreate

```go
eventGatewayProducePolicyCreate := components.CreateEventGatewayProducePolicyCreateModifyHeaders(components.EventGatewayModifyHeadersPolicyCreate{/* values here */})
```

### EventGatewayProduceSchemaValidationPolicy

```go
eventGatewayProducePolicyCreate := components.CreateEventGatewayProducePolicyCreateSchemaValidation(components.EventGatewayProduceSchemaValidationPolicy{/* values here */})
```

### EventGatewayEncryptPolicy

```go
eventGatewayProducePolicyCreate := components.CreateEventGatewayProducePolicyCreateEncrypt(components.EventGatewayEncryptPolicy{/* values here */})
```

### EventGatewayParsedRecordEncryptFieldsPolicyCreate

```go
eventGatewayProducePolicyCreate := components.CreateEventGatewayProducePolicyCreateEncryptFields(components.EventGatewayParsedRecordEncryptFieldsPolicyCreate{/* values here */})
```

### EventGatewayParsedRecordTranscodeProducePolicyCreate

```go
eventGatewayProducePolicyCreate := components.CreateEventGatewayProducePolicyCreateTranscode(components.EventGatewayParsedRecordTranscodeProducePolicyCreate{/* values here */})
```

### EventGatewayParsedRecordMaskFieldsProducePolicyCreate

```go
eventGatewayProducePolicyCreate := components.CreateEventGatewayProducePolicyCreateMaskFields(components.EventGatewayParsedRecordMaskFieldsProducePolicyCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayProducePolicyCreate.Type {
	case components.EventGatewayProducePolicyCreateTypeModifyHeaders:
		// eventGatewayProducePolicyCreate.EventGatewayModifyHeadersPolicyCreate is populated
	case components.EventGatewayProducePolicyCreateTypeSchemaValidation:
		// eventGatewayProducePolicyCreate.EventGatewayProduceSchemaValidationPolicy is populated
	case components.EventGatewayProducePolicyCreateTypeEncrypt:
		// eventGatewayProducePolicyCreate.EventGatewayEncryptPolicy is populated
	case components.EventGatewayProducePolicyCreateTypeEncryptFields:
		// eventGatewayProducePolicyCreate.EventGatewayParsedRecordEncryptFieldsPolicyCreate is populated
	case components.EventGatewayProducePolicyCreateTypeTranscode:
		// eventGatewayProducePolicyCreate.EventGatewayParsedRecordTranscodeProducePolicyCreate is populated
	case components.EventGatewayProducePolicyCreateTypeMaskFields:
		// eventGatewayProducePolicyCreate.EventGatewayParsedRecordMaskFieldsProducePolicyCreate is populated
}
```
