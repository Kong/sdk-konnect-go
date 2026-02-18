# EventGatewayConsumePolicyUpdate

The typed schema of the consume policy to modify it.


## Supported Types

### EventGatewayModifyHeadersPolicy

```go
eventGatewayConsumePolicyUpdate := components.CreateEventGatewayConsumePolicyUpdateModifyHeaders(components.EventGatewayModifyHeadersPolicy{/* values here */})
```

### EventGatewayConsumeSchemaValidationPolicy

```go
eventGatewayConsumePolicyUpdate := components.CreateEventGatewayConsumePolicyUpdateSchemaValidation(components.EventGatewayConsumeSchemaValidationPolicy{/* values here */})
```

### EventGatewayDecryptPolicy

```go
eventGatewayConsumePolicyUpdate := components.CreateEventGatewayConsumePolicyUpdateDecrypt(components.EventGatewayDecryptPolicy{/* values here */})
```

### EventGatewaySkipRecordPolicy

```go
eventGatewayConsumePolicyUpdate := components.CreateEventGatewayConsumePolicyUpdateSkipRecord(components.EventGatewaySkipRecordPolicy{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayConsumePolicyUpdate.Type {
	case components.EventGatewayConsumePolicyUpdateTypeModifyHeaders:
		// eventGatewayConsumePolicyUpdate.EventGatewayModifyHeadersPolicy is populated
	case components.EventGatewayConsumePolicyUpdateTypeSchemaValidation:
		// eventGatewayConsumePolicyUpdate.EventGatewayConsumeSchemaValidationPolicy is populated
	case components.EventGatewayConsumePolicyUpdateTypeDecrypt:
		// eventGatewayConsumePolicyUpdate.EventGatewayDecryptPolicy is populated
	case components.EventGatewayConsumePolicyUpdateTypeSkipRecord:
		// eventGatewayConsumePolicyUpdate.EventGatewaySkipRecordPolicy is populated
}
```
