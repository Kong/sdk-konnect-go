# IntegrationEventsFilterParametersIntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
integrationEventsFilterParametersIntegrationInstanceName := components.CreateIntegrationEventsFilterParametersIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationEventsFilterParametersIntegrationInstanceName.Type {
	case components.IntegrationEventsFilterParametersIntegrationInstanceNameTypeStringFieldFilter:
		// integrationEventsFilterParametersIntegrationInstanceName.StringFieldFilter is populated
}
```
