# SuggestedRuleAction


## Supported Types

### ArchiveActionPayload

```go
suggestedRuleAction := components.CreateSuggestedRuleActionArchiveActionPayload(components.ArchiveActionPayload{/* values here */})
```

### IgnoreActionPayload

```go
suggestedRuleAction := components.CreateSuggestedRuleActionIgnoreActionPayload(components.IgnoreActionPayload{/* values here */})
```

### MapActionPayload

```go
suggestedRuleAction := components.CreateSuggestedRuleActionMapActionPayload(components.MapActionPayload{/* values here */})
```

### CreateOrMapAction

```go
suggestedRuleAction := components.CreateSuggestedRuleActionCreateOrMapAction(components.CreateOrMapAction{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedRuleAction.Type {
	case components.SuggestedRuleActionTypeArchiveActionPayload:
		// suggestedRuleAction.ArchiveActionPayload is populated
	case components.SuggestedRuleActionTypeIgnoreActionPayload:
		// suggestedRuleAction.IgnoreActionPayload is populated
	case components.SuggestedRuleActionTypeMapActionPayload:
		// suggestedRuleAction.MapActionPayload is populated
	case components.SuggestedRuleActionTypeCreateOrMapAction:
		// suggestedRuleAction.CreateOrMapAction is populated
}
```
