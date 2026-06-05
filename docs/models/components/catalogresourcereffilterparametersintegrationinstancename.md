# CatalogResourceRefFilterParametersIntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
catalogResourceRefFilterParametersIntegrationInstanceName := components.CreateCatalogResourceRefFilterParametersIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogResourceRefFilterParametersIntegrationInstanceName.Type {
	case components.CatalogResourceRefFilterParametersIntegrationInstanceNameTypeStringFieldFilter:
		// catalogResourceRefFilterParametersIntegrationInstanceName.StringFieldFilter is populated
}
```
