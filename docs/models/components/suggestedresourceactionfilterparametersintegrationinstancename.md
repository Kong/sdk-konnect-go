# SuggestedResourceActionFilterParametersIntegrationInstanceName


## Supported Types

### StringFieldFilter

```go
suggestedResourceActionFilterParametersIntegrationInstanceName := components.CreateSuggestedResourceActionFilterParametersIntegrationInstanceNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedResourceActionFilterParametersIntegrationInstanceName.Type {
	case components.SuggestedResourceActionFilterParametersIntegrationInstanceNameTypeStringFieldFilter:
		// suggestedResourceActionFilterParametersIntegrationInstanceName.StringFieldFilter is populated
}
```
