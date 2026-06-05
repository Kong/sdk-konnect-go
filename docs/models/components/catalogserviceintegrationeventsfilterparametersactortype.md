# CatalogServiceIntegrationEventsFilterParametersActorType


## Supported Types

### StringFieldFilter

```go
catalogServiceIntegrationEventsFilterParametersActorType := components.CreateCatalogServiceIntegrationEventsFilterParametersActorTypeStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogServiceIntegrationEventsFilterParametersActorType.Type {
	case components.CatalogServiceIntegrationEventsFilterParametersActorTypeTypeStringFieldFilter:
		// catalogServiceIntegrationEventsFilterParametersActorType.StringFieldFilter is populated
}
```
