# SchemaRegistryCreate

The typed schema of the schema registry to create it.


## Supported Types

### SchemaRegistryConfluent

```go
schemaRegistryCreate := components.CreateSchemaRegistryCreateConfluent(components.SchemaRegistryConfluent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch schemaRegistryCreate.Type {
	case components.SchemaRegistryCreateTypeConfluent:
		// schemaRegistryCreate.SchemaRegistryConfluent is populated
}
```
