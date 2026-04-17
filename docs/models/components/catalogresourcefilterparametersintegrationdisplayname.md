# CatalogResourceFilterParametersIntegrationDisplayName


## Supported Types

### StringFieldFilter

```go
catalogResourceFilterParametersIntegrationDisplayName := components.CreateCatalogResourceFilterParametersIntegrationDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceFilterParametersIntegrationDisplayName.Type {
	case components.CatalogResourceFilterParametersIntegrationDisplayNameTypeStringFieldFilter:
		// catalogResourceFilterParametersIntegrationDisplayName.StringFieldFilter is populated
}
```
