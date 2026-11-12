# IntegrationInstanceFilterParametersIntegrationDisplayName


## Supported Types

### StringFieldFilter

```go
integrationInstanceFilterParametersIntegrationDisplayName := components.CreateIntegrationInstanceFilterParametersIntegrationDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceFilterParametersIntegrationDisplayName.Type {
	case components.IntegrationInstanceFilterParametersIntegrationDisplayNameTypeStringFieldFilter:
		// integrationInstanceFilterParametersIntegrationDisplayName.StringFieldFilter is populated
}
```
