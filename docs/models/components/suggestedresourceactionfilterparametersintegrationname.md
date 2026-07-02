# SuggestedResourceActionFilterParametersIntegrationName


## Supported Types

### StringFieldFilter

```go
suggestedResourceActionFilterParametersIntegrationName := components.CreateSuggestedResourceActionFilterParametersIntegrationNameStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedResourceActionFilterParametersIntegrationName.Type {
	case components.SuggestedResourceActionFilterParametersIntegrationNameTypeStringFieldFilter:
		// suggestedResourceActionFilterParametersIntegrationName.StringFieldFilter is populated
}
```
