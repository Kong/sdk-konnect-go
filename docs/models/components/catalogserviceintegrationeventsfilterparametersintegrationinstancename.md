# CatalogServiceIntegrationEventsFilterParametersIntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
catalogServiceIntegrationEventsFilterParametersIntegrationInstanceName := components.CreateCatalogServiceIntegrationEventsFilterParametersIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogServiceIntegrationEventsFilterParametersIntegrationInstanceName.Type {
	case components.CatalogServiceIntegrationEventsFilterParametersIntegrationInstanceNameTypeStringFieldFilter:
		// catalogServiceIntegrationEventsFilterParametersIntegrationInstanceName.StringFieldFilter is populated
}
```
