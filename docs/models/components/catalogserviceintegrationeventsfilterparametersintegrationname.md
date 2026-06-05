# CatalogServiceIntegrationEventsFilterParametersIntegrationName


## Supported Types

### StringFieldFilter

```go
catalogServiceIntegrationEventsFilterParametersIntegrationName := components.CreateCatalogServiceIntegrationEventsFilterParametersIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogServiceIntegrationEventsFilterParametersIntegrationName.Type {
	case components.CatalogServiceIntegrationEventsFilterParametersIntegrationNameTypeStringFieldFilter:
		// catalogServiceIntegrationEventsFilterParametersIntegrationName.StringFieldFilter is populated
}
```
