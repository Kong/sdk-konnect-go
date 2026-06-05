# CatalogIntegrationConfigSchema


## Supported Types

### CatalogIntegrationConfigSchema1

```go
catalogIntegrationConfigSchema := components.CreateCatalogIntegrationConfigSchemaCatalogIntegrationConfigSchema1(components.CatalogIntegrationConfigSchema1{/* values here */})
```

### StringConfigFieldSchema

```go
catalogIntegrationConfigSchema := components.CreateCatalogIntegrationConfigSchemaStringConfigFieldSchema(components.StringConfigFieldSchema{/* values here */})
```

### EnumConfigFieldSchema

```go
catalogIntegrationConfigSchema := components.CreateCatalogIntegrationConfigSchemaEnumConfigFieldSchema(components.EnumConfigFieldSchema{/* values here */})
```

### BooleanConfigFieldSchema

```go
catalogIntegrationConfigSchema := components.CreateCatalogIntegrationConfigSchemaBooleanConfigFieldSchema(components.BooleanConfigFieldSchema{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogIntegrationConfigSchema.Type {
	case components.CatalogIntegrationConfigSchemaTypeCatalogIntegrationConfigSchema1:
		// catalogIntegrationConfigSchema.CatalogIntegrationConfigSchema1 is populated
	case components.CatalogIntegrationConfigSchemaTypeStringConfigFieldSchema:
		// catalogIntegrationConfigSchema.StringConfigFieldSchema is populated
	case components.CatalogIntegrationConfigSchemaTypeEnumConfigFieldSchema:
		// catalogIntegrationConfigSchema.EnumConfigFieldSchema is populated
	case components.CatalogIntegrationConfigSchemaTypeBooleanConfigFieldSchema:
		// catalogIntegrationConfigSchema.BooleanConfigFieldSchema is populated
}
```
