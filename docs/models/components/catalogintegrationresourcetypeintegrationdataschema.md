# CatalogIntegrationResourceTypeIntegrationDataSchema


## Supported Types

### EmptySchema

```go
catalogIntegrationResourceTypeIntegrationDataSchema := components.CreateCatalogIntegrationResourceTypeIntegrationDataSchemaEmptySchema(components.EmptySchema{/* values here */})
```

### CatalogIntegrationResourceTypeIntegrationDataSimpleSchema

```go
catalogIntegrationResourceTypeIntegrationDataSchema := components.CreateCatalogIntegrationResourceTypeIntegrationDataSchemaCatalogIntegrationResourceTypeIntegrationDataSimpleSchema(components.CatalogIntegrationResourceTypeIntegrationDataSimpleSchema{/* values here */})
```

### CatalogIntegrationResourceTypeIntegrationDataJSONSchema

```go
catalogIntegrationResourceTypeIntegrationDataSchema := components.CreateCatalogIntegrationResourceTypeIntegrationDataSchemaCatalogIntegrationResourceTypeIntegrationDataJSONSchema(components.CatalogIntegrationResourceTypeIntegrationDataJSONSchema{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogIntegrationResourceTypeIntegrationDataSchema.Type {
	case components.CatalogIntegrationResourceTypeIntegrationDataSchemaUnionTypeEmptySchema:
		// catalogIntegrationResourceTypeIntegrationDataSchema.EmptySchema is populated
	case components.CatalogIntegrationResourceTypeIntegrationDataSchemaUnionTypeCatalogIntegrationResourceTypeIntegrationDataSimpleSchema:
		// catalogIntegrationResourceTypeIntegrationDataSchema.CatalogIntegrationResourceTypeIntegrationDataSimpleSchema is populated
	case components.CatalogIntegrationResourceTypeIntegrationDataSchemaUnionTypeCatalogIntegrationResourceTypeIntegrationDataJSONSchema:
		// catalogIntegrationResourceTypeIntegrationDataSchema.CatalogIntegrationResourceTypeIntegrationDataJSONSchema is populated
}
```
