# CatalogServiceIntegrationEventsFilterParametersIntegrationInstanceID


## Supported Types

### StringFieldFilterExact

```go
catalogServiceIntegrationEventsFilterParametersIntegrationInstanceID := components.CreateCatalogServiceIntegrationEventsFilterParametersIntegrationInstanceIDStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogServiceIntegrationEventsFilterParametersIntegrationInstanceID.Type {
	case components.CatalogServiceIntegrationEventsFilterParametersIntegrationInstanceIDTypeStringFieldFilterExact:
		// catalogServiceIntegrationEventsFilterParametersIntegrationInstanceID.StringFieldFilterExact is populated
}
```
