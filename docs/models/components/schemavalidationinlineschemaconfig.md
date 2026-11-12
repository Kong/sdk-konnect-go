# SchemaValidationInlineSchemaConfig

Defines the schema for a record key or value, inline.

**Requires a minimum runtime version of `1.3`**.


## Supported Types

### SchemaValidationInlineSchemaConfigJSON

```go
schemaValidationInlineSchemaConfig := components.CreateSchemaValidationInlineSchemaConfigJSON(components.SchemaValidationInlineSchemaConfigJSON{/* values here */})
```

### SchemaValidationInlineSchemaConfigAvro

```go
schemaValidationInlineSchemaConfig := components.CreateSchemaValidationInlineSchemaConfigAvro(components.SchemaValidationInlineSchemaConfigAvro{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch schemaValidationInlineSchemaConfig.Type {
	case components.SchemaValidationInlineSchemaConfigTypeJSON:
		// schemaValidationInlineSchemaConfig.SchemaValidationInlineSchemaConfigJSON is populated
	case components.SchemaValidationInlineSchemaConfigTypeAvro:
		// schemaValidationInlineSchemaConfig.SchemaValidationInlineSchemaConfigAvro is populated
}
```
