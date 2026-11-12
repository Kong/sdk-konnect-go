# EventGatewayProduceSchemaValidationPolicyConfig

The configuration of the produce schema validation policy.


## Supported Types

### EventGatewayProduceSchemaValidationPolicySchemaRegistryConfig

```go
eventGatewayProduceSchemaValidationPolicyConfig := components.CreateEventGatewayProduceSchemaValidationPolicyConfigConfluentSchemaRegistry(components.EventGatewayProduceSchemaValidationPolicySchemaRegistryConfig{/* values here */})
```

### EventGatewayProduceSchemaValidationPolicyJSONConfig

```go
eventGatewayProduceSchemaValidationPolicyConfig := components.CreateEventGatewayProduceSchemaValidationPolicyConfigJSON(components.EventGatewayProduceSchemaValidationPolicyJSONConfig{/* values here */})
```

### EventGatewayProduceSchemaValidationPolicyInlineSchemaConfig

```go
eventGatewayProduceSchemaValidationPolicyConfig := components.CreateEventGatewayProduceSchemaValidationPolicyConfigInlineSchema(components.EventGatewayProduceSchemaValidationPolicyInlineSchemaConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayProduceSchemaValidationPolicyConfig.Type {
	case components.EventGatewayProduceSchemaValidationPolicyConfigTypeConfluentSchemaRegistry:
		// eventGatewayProduceSchemaValidationPolicyConfig.EventGatewayProduceSchemaValidationPolicySchemaRegistryConfig is populated
	case components.EventGatewayProduceSchemaValidationPolicyConfigTypeJSON:
		// eventGatewayProduceSchemaValidationPolicyConfig.EventGatewayProduceSchemaValidationPolicyJSONConfig is populated
	case components.EventGatewayProduceSchemaValidationPolicyConfigTypeInlineSchema:
		// eventGatewayProduceSchemaValidationPolicyConfig.EventGatewayProduceSchemaValidationPolicyInlineSchemaConfig is populated
}
```
