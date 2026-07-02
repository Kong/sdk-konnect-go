# ResourceActionFilterParametersIntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
resourceActionFilterParametersIntegrationInstanceName := components.CreateResourceActionFilterParametersIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceActionFilterParametersIntegrationInstanceName.Type {
	case components.ResourceActionFilterParametersIntegrationInstanceNameTypeStringFieldFilter:
		// resourceActionFilterParametersIntegrationInstanceName.StringFieldFilter is populated
}
```
