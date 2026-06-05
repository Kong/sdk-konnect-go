# CatalogServiceIntegrationEventsFilterParametersActorID


## Supported Types

### StringFieldFilter

```go
catalogServiceIntegrationEventsFilterParametersActorID := components.CreateCatalogServiceIntegrationEventsFilterParametersActorIDStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogServiceIntegrationEventsFilterParametersActorID.Type {
	case components.CatalogServiceIntegrationEventsFilterParametersActorIDTypeStringFieldFilter:
		// catalogServiceIntegrationEventsFilterParametersActorID.StringFieldFilter is populated
}
```
