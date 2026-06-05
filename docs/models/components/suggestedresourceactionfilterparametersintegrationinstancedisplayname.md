# SuggestedResourceActionFilterParametersIntegrationInstanceDisplayName


## Supported Types

### StringFieldFilter

```go
suggestedResourceActionFilterParametersIntegrationInstanceDisplayName := components.CreateSuggestedResourceActionFilterParametersIntegrationInstanceDisplayNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedResourceActionFilterParametersIntegrationInstanceDisplayName.Type {
	case components.SuggestedResourceActionFilterParametersIntegrationInstanceDisplayNameTypeStringFieldFilter:
		// suggestedResourceActionFilterParametersIntegrationInstanceDisplayName.StringFieldFilter is populated
}
```
