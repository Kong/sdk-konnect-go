# CatalogResourceFilterParametersIntegrationName


## Supported Types

### StringFieldFilter

```go
catalogResourceFilterParametersIntegrationName := components.CreateCatalogResourceFilterParametersIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceFilterParametersIntegrationName.Type {
	case components.CatalogResourceFilterParametersIntegrationNameTypeStringFieldFilter:
		// catalogResourceFilterParametersIntegrationName.StringFieldFilter is populated
}
```
