# SuggestedResourceActionFilterParametersActionType


## Supported Types

### StringFieldFilterExact

```go
suggestedResourceActionFilterParametersActionType := components.CreateSuggestedResourceActionFilterParametersActionTypeStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedResourceActionFilterParametersActionType.Type {
	case components.SuggestedResourceActionFilterParametersActionTypeTypeStringFieldFilterExact:
		// suggestedResourceActionFilterParametersActionType.StringFieldFilterExact is populated
}
```
