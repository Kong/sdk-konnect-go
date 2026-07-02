# CatalogResourceRefFilterParametersIntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
catalogResourceRefFilterParametersIntegrationInstanceDisplayName := components.CreateCatalogResourceRefFilterParametersIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceRefFilterParametersIntegrationInstanceDisplayName.Type {
	case components.CatalogResourceRefFilterParametersIntegrationInstanceDisplayNameTypeStringFieldFilter:
		// catalogResourceRefFilterParametersIntegrationInstanceDisplayName.StringFieldFilter is populated
}
```
