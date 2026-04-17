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
}
```
