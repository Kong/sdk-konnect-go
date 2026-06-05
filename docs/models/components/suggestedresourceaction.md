# SuggestedResourceAction


## Supported Types

### ArchiveActionPayload

```go
suggestedResourceAction := components.CreateSuggestedResourceActionArchiveActionPayload(components.ArchiveActionPayload{/* values here */})
```

### IgnoreActionPayload

```go
suggestedResourceAction := components.CreateSuggestedResourceActionIgnoreActionPayload(components.IgnoreActionPayload{/* values here */})
```

### MapServiceAction

```go
suggestedResourceAction := components.CreateSuggestedResourceActionMapServiceAction(components.MapServiceAction{/* values here */})
```

### CreateAndMapServiceActionPayload

```go
suggestedResourceAction := components.CreateSuggestedResourceActionCreateAndMapServiceActionPayload(components.CreateAndMapServiceActionPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedResourceAction.Type {
	case components.SuggestedResourceActionTypeArchiveActionPayload:
		// suggestedResourceAction.ArchiveActionPayload is populated
	case components.SuggestedResourceActionTypeIgnoreActionPayload:
		// suggestedResourceAction.IgnoreActionPayload is populated
	case components.SuggestedResourceActionTypeMapServiceAction:
		// suggestedResourceAction.MapServiceAction is populated
	case components.SuggestedResourceActionTypeCreateAndMapServiceActionPayload:
		// suggestedResourceAction.CreateAndMapServiceActionPayload is populated
}
```
