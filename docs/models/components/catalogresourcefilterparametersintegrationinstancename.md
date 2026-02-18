# CatalogResourceFilterParametersIntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
catalogResourceFilterParametersIntegrationInstanceName := components.CreateCatalogResourceFilterParametersIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceFilterParametersIntegrationInstanceName.Type {
	case components.CatalogResourceFilterParametersIntegrationInstanceNameTypeStringFieldFilter:
		// catalogResourceFilterParametersIntegrationInstanceName.StringFieldFilter is populated
}
```
