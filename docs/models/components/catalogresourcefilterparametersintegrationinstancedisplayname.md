# CatalogResourceFilterParametersIntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
catalogResourceFilterParametersIntegrationInstanceDisplayName := components.CreateCatalogResourceFilterParametersIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceFilterParametersIntegrationInstanceDisplayName.Type {
	case components.CatalogResourceFilterParametersIntegrationInstanceDisplayNameTypeStringFieldFilter:
		// catalogResourceFilterParametersIntegrationInstanceDisplayName.StringFieldFilter is populated
}
```
