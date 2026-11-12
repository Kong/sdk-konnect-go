# SchemaRegistryAuthenticationScheme

The authentication configuration for the schema registry.
The Event Gateway only requires read-only access to the schema registry to fetch schemas for validation, so we recommend using credentials with read-only scope.



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
