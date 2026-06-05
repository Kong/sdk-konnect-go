# CatalogResourceRefFilterParametersIntegrationDisplayName


## Supported Types

### StringFieldFilter

```go
catalogResourceRefFilterParametersIntegrationDisplayName := components.CreateCatalogResourceRefFilterParametersIntegrationDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceRefFilterParametersIntegrationDisplayName.Type {
	case components.CatalogResourceRefFilterParametersIntegrationDisplayNameTypeStringFieldFilter:
		// catalogResourceRefFilterParametersIntegrationDisplayName.StringFieldFilter is populated
}
```
