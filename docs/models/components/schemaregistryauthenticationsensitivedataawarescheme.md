# SchemaRegistryAuthenticationSensitiveDataAwareScheme

The authentication configuration for the schema registry.


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
