# CatalogServiceIntegrationEventsFilterParametersIntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
catalogServiceIntegrationEventsFilterParametersIntegrationInstanceDisplayName := components.CreateCatalogServiceIntegrationEventsFilterParametersIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogServiceIntegrationEventsFilterParametersIntegrationInstanceDisplayName.Type {
	case components.CatalogServiceIntegrationEventsFilterParametersIntegrationInstanceDisplayNameTypeStringFieldFilter:
		// catalogServiceIntegrationEventsFilterParametersIntegrationInstanceDisplayName.StringFieldFilter is populated
}
```
