# SuggestedResourceActionFilterParametersIntegrationInstanceID


## Supported Types

### StringFieldFilterExact

```go
suggestedResourceActionFilterParametersIntegrationInstanceID := components.CreateSuggestedResourceActionFilterParametersIntegrationInstanceIDStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedResourceActionFilterParametersIntegrationInstanceID.Type {
	case components.SuggestedResourceActionFilterParametersIntegrationInstanceIDTypeStringFieldFilterExact:
		// suggestedResourceActionFilterParametersIntegrationInstanceID.StringFieldFilterExact is populated
}
```
