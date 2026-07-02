# SuggestedRuleActionEvaluated


## Supported Types

### ArchiveActionPayload

```go
suggestedRuleActionEvaluated := components.CreateSuggestedRuleActionEvaluatedArchiveActionPayload(components.ArchiveActionPayload{/* values here */})
```

### IgnoreActionPayload

```go
suggestedRuleActionEvaluated := components.CreateSuggestedRuleActionEvaluatedIgnoreActionPayload(components.IgnoreActionPayload{/* values here */})
```

### MapActionPayload

```go
suggestedRuleActionEvaluated := components.CreateSuggestedRuleActionEvaluatedMapActionPayload(components.MapActionPayload{/* values here */})
```

### CreateActionEvaluated

```go
suggestedRuleActionEvaluated := components.CreateSuggestedRuleActionEvaluatedCreateActionEvaluated(components.CreateActionEvaluated{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestedRuleActionEvaluated.Type {
	case components.SuggestedRuleActionEvaluatedTypeArchiveActionPayload:
		// suggestedRuleActionEvaluated.ArchiveActionPayload is populated
	case components.SuggestedRuleActionEvaluatedTypeIgnoreActionPayload:
		// suggestedRuleActionEvaluated.IgnoreActionPayload is populated
	case components.SuggestedRuleActionEvaluatedTypeMapActionPayload:
		// suggestedRuleActionEvaluated.MapActionPayload is populated
	case components.SuggestedRuleActionEvaluatedTypeCreateActionEvaluated:
		// suggestedRuleActionEvaluated.CreateActionEvaluated is populated
}
```
