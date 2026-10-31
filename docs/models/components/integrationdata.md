# IntegrationData


## Supported Types

### CatalogResourceIntegrationDataFieldFilter

```go
integrationData := components.CreateIntegrationDataCatalogResourceIntegrationDataFieldFilter(components.CatalogResourceIntegrationDataFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationData.Type {
	case components.IntegrationDataTypeCatalogResourceIntegrationDataFieldFilter:
		// integrationData.CatalogResourceIntegrationDataFieldFilter is populated
}
```
