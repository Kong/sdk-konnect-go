# SchemaRegistryReference

A reference to a schema Registry.


## Supported Types

### SchemaRegistryReferenceByID

```go
schemaRegistryReference := components.CreateSchemaRegistryReferenceSchemaRegistryReferenceByID(components.SchemaRegistryReferenceByID{/* values here */})
```

### SchemaRegistryReferenceByName

```go
schemaRegistryReference := components.CreateSchemaRegistryReferenceSchemaRegistryReferenceByName(components.SchemaRegistryReferenceByName{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch schemaRegistryReference.Type {
	case components.SchemaRegistryReferenceTypeSchemaRegistryReferenceByID:
		// schemaRegistryReference.SchemaRegistryReferenceByID is populated
	case components.SchemaRegistryReferenceTypeSchemaRegistryReferenceByName:
		// schemaRegistryReference.SchemaRegistryReferenceByName is populated
}
```
