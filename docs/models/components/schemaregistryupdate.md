# SchemaRegistryUpdate

The typed schema of the schema registry to modify it.


## Supported Types

### SchemaRegistryConfluentSensitiveDataAware

```go
schemaRegistryUpdate := components.CreateSchemaRegistryUpdateConfluent(components.SchemaRegistryConfluentSensitiveDataAware{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch schemaRegistryUpdate.Type {
	case components.SchemaRegistryUpdateTypeConfluent:
		// schemaRegistryUpdate.SchemaRegistryConfluentSensitiveDataAware is populated
}
```
