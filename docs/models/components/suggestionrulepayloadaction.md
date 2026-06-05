# SuggestionRulePayloadAction


## Supported Types

### ArchiveActionPayload

```go
suggestionRulePayloadAction := components.CreateSuggestionRulePayloadActionArchiveActionPayload(components.ArchiveActionPayload{/* values here */})
```

### MapActionPayload

```go
suggestionRulePayloadAction := components.CreateSuggestionRulePayloadActionMapActionPayload(components.MapActionPayload{/* values here */})
```

### CreateOrMapAction

```go
suggestionRulePayloadAction := components.CreateSuggestionRulePayloadActionCreateOrMapAction(components.CreateOrMapAction{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestionRulePayloadAction.Type {
	case components.SuggestionRulePayloadActionTypeArchiveActionPayload:
		// suggestionRulePayloadAction.ArchiveActionPayload is populated
	case components.SuggestionRulePayloadActionTypeMapActionPayload:
		// suggestionRulePayloadAction.MapActionPayload is populated
	case components.SuggestionRulePayloadActionTypeCreateOrMapAction:
		// suggestionRulePayloadAction.CreateOrMapAction is populated
}
```
