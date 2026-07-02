# CatalogResourceRefFilterParametersIntegrationData


## Supported Types

### CatalogResourceIntegrationDataFieldFilter

```go
catalogResourceRefFilterParametersIntegrationData := components.CreateCatalogResourceRefFilterParametersIntegrationDataCatalogResourceIntegrationDataFieldFilter(components.CatalogResourceIntegrationDataFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceRefFilterParametersIntegrationData.Type {
	case components.CatalogResourceRefFilterParametersIntegrationDataTypeCatalogResourceIntegrationDataFieldFilter:
		// catalogResourceRefFilterParametersIntegrationData.CatalogResourceIntegrationDataFieldFilter is populated
}
```
