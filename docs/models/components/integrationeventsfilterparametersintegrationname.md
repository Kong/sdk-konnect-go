# IntegrationEventsFilterParametersIntegrationName


## Supported Types

### StringFieldFilter

```go
integrationEventsFilterParametersIntegrationName := components.CreateIntegrationEventsFilterParametersIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationEventsFilterParametersIntegrationName.Type {
	case components.IntegrationEventsFilterParametersIntegrationNameTypeStringFieldFilter:
		// integrationEventsFilterParametersIntegrationName.StringFieldFilter is populated
}
```
