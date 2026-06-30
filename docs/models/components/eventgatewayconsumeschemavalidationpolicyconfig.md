# EventGatewayConsumeSchemaValidationPolicyConfig

The configuration of the consume schema validation policy.


## Supported Types

### EventGatewayConsumeSchemaValidationPolicySchemaRegistryConfig

```go
eventGatewayConsumeSchemaValidationPolicyConfig := components.CreateEventGatewayConsumeSchemaValidationPolicyConfigConfluentSchemaRegistry(components.EventGatewayConsumeSchemaValidationPolicySchemaRegistryConfig{/* values here */})
```

### EventGatewayConsumeSchemaValidationPolicyJSONConfig

```go
eventGatewayConsumeSchemaValidationPolicyConfig := components.CreateEventGatewayConsumeSchemaValidationPolicyConfigJSON(components.EventGatewayConsumeSchemaValidationPolicyJSONConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayConsumeSchemaValidationPolicyConfig.Type {
	case components.EventGatewayConsumeSchemaValidationPolicyConfigTypeConfluentSchemaRegistry:
		// eventGatewayConsumeSchemaValidationPolicyConfig.EventGatewayConsumeSchemaValidationPolicySchemaRegistryConfig is populated
	case components.EventGatewayConsumeSchemaValidationPolicyConfigTypeJSON:
		// eventGatewayConsumeSchemaValidationPolicyConfig.EventGatewayConsumeSchemaValidationPolicyJSONConfig is populated
}
```
