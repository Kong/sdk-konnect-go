# ResourceActionFilterParametersIntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
resourceActionFilterParametersIntegrationInstanceDisplayName := components.CreateResourceActionFilterParametersIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resourceActionFilterParametersIntegrationInstanceDisplayName.Type {
	case components.ResourceActionFilterParametersIntegrationInstanceDisplayNameTypeStringFieldFilter:
		// resourceActionFilterParametersIntegrationInstanceDisplayName.StringFieldFilter is populated
}
```
