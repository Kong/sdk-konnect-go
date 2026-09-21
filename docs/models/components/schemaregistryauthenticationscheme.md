# SchemaRegistryAuthenticationScheme

The authentication configuration for the schema registry.


## Supported Types

### SchemaRegistryAuthenticationBasic

```go
schemaRegistryAuthenticationScheme := components.CreateSchemaRegistryAuthenticationSchemeBasic(components.SchemaRegistryAuthenticationBasic{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch schemaRegistryAuthenticationScheme.Type {
	case components.SchemaRegistryAuthenticationSchemeTypeBasic:
		// schemaRegistryAuthenticationScheme.SchemaRegistryAuthenticationBasic is populated
}
```
