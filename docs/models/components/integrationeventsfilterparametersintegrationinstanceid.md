# IntegrationEventsFilterParametersIntegrationInstanceID


## Supported Types

### StringFieldFilterExact

```go
integrationEventsFilterParametersIntegrationInstanceID := components.CreateIntegrationEventsFilterParametersIntegrationInstanceIDStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationEventsFilterParametersIntegrationInstanceID.Type {
	case components.IntegrationEventsFilterParametersIntegrationInstanceIDTypeStringFieldFilterExact:
		// integrationEventsFilterParametersIntegrationInstanceID.StringFieldFilterExact is populated
}
```
