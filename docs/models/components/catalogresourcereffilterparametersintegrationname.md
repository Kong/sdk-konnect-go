# CatalogResourceRefFilterParametersIntegrationName


## Supported Types

### StringFieldFilter

```go
catalogResourceRefFilterParametersIntegrationName := components.CreateCatalogResourceRefFilterParametersIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceRefFilterParametersIntegrationName.Type {
	case components.CatalogResourceRefFilterParametersIntegrationNameTypeStringFieldFilter:
		// catalogResourceRefFilterParametersIntegrationName.StringFieldFilter is populated
}
```
