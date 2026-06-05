# CatalogServiceIntegrationEventsFilterParametersActorName


## Supported Types

### StringFieldFilter

```go
catalogServiceIntegrationEventsFilterParametersActorName := components.CreateCatalogServiceIntegrationEventsFilterParametersActorNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogServiceIntegrationEventsFilterParametersActorName.Type {
	case components.CatalogServiceIntegrationEventsFilterParametersActorNameTypeStringFieldFilter:
		// catalogServiceIntegrationEventsFilterParametersActorName.StringFieldFilter is populated
}
```
