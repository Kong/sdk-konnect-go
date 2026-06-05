# IntegrationEventsFilterParametersIntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
integrationEventsFilterParametersIntegrationInstanceDisplayName := components.CreateIntegrationEventsFilterParametersIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationEventsFilterParametersIntegrationInstanceDisplayName.Type {
	case components.IntegrationEventsFilterParametersIntegrationInstanceDisplayNameTypeStringFieldFilter:
		// integrationEventsFilterParametersIntegrationInstanceDisplayName.StringFieldFilter is populated
}
```
