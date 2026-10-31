# CatalogResourceFilterParametersIntegrationData


## Supported Types

### CatalogResourceIntegrationDataFieldFilter

```go
catalogResourceFilterParametersIntegrationData := components.CreateCatalogResourceFilterParametersIntegrationDataCatalogResourceIntegrationDataFieldFilter(components.CatalogResourceIntegrationDataFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceFilterParametersIntegrationData.Type {
	case components.CatalogResourceFilterParametersIntegrationDataTypeCatalogResourceIntegrationDataFieldFilter:
		// catalogResourceFilterParametersIntegrationData.CatalogResourceIntegrationDataFieldFilter is populated
}
```
