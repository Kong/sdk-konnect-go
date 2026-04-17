# EventGatewayConsumePolicyCreate

The typed schema of the consume policy to modify it.


## Supported Types

### EventGatewayModifyHeadersPolicyCreate

```go
eventGatewayConsumePolicyCreate := components.CreateEventGatewayConsumePolicyCreateModifyHeaders(components.EventGatewayModifyHeadersPolicyCreate{/* values here */})
```

### EventGatewayConsumeSchemaValidationPolicy

```go
eventGatewayConsumePolicyCreate := components.CreateEventGatewayConsumePolicyCreateSchemaValidation(components.EventGatewayConsumeSchemaValidationPolicy{/* values here */})
```

### EventGatewayDecryptPolicy

```go
eventGatewayConsumePolicyCreate := components.CreateEventGatewayConsumePolicyCreateDecrypt(components.EventGatewayDecryptPolicy{/* values here */})
```

### EventGatewaySkipRecordPolicyCreate

```go
eventGatewayConsumePolicyCreate := components.CreateEventGatewayConsumePolicyCreateSkipRecord(components.EventGatewaySkipRecordPolicyCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayConsumePolicyCreate.Type {
	case components.EventGatewayConsumePolicyCreateTypeModifyHeaders:
		// eventGatewayConsumePolicyCreate.EventGatewayModifyHeadersPolicyCreate is populated
	case components.EventGatewayConsumePolicyCreateTypeSchemaValidation:
		// eventGatewayConsumePolicyCreate.EventGatewayConsumeSchemaValidationPolicy is populated
	case components.EventGatewayConsumePolicyCreateTypeDecrypt:
		// eventGatewayConsumePolicyCreate.EventGatewayDecryptPolicy is populated
	case components.EventGatewayConsumePolicyCreateTypeSkipRecord:
		// eventGatewayConsumePolicyCreate.EventGatewaySkipRecordPolicyCreate is populated
}
```
