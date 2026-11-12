# SchemaRegistryAuthenticationSensitiveDataAwareScheme

The authentication configuration for the schema registry.
The Event Gateway only requires read-only access to the schema registry to fetch schemas for validation, so we recommend using credentials with read-only scope.



## Supported Types

### SchemaRegistryAuthenticationBasicSensitiveDataAware

```go
schemaRegistryAuthenticationSensitiveDataAwareScheme := components.CreateSchemaRegistryAuthenticationSensitiveDataAwareSchemeBasic(components.SchemaRegistryAuthenticationBasicSensitiveDataAware{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch schemaRegistryAuthenticationSensitiveDataAwareScheme.Type {
	case components.SchemaRegistryAuthenticationSensitiveDataAwareSchemeTypeBasic:
		// schemaRegistryAuthenticationSensitiveDataAwareScheme.SchemaRegistryAuthenticationBasicSensitiveDataAware is populated
}
```
