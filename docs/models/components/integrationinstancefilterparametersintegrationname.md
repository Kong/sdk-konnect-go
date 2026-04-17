# IntegrationInstanceFilterParametersIntegrationName


## Supported Types

### StringFieldFilter

```go
integrationInstanceFilterParametersIntegrationName := components.CreateIntegrationInstanceFilterParametersIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceFilterParametersIntegrationName.Type {
	case components.IntegrationInstanceFilterParametersIntegrationNameTypeStringFieldFilter:
		// integrationInstanceFilterParametersIntegrationName.StringFieldFilter is populated
}
```
