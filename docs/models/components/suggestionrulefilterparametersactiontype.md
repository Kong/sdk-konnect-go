# SuggestionRuleFilterParametersActionType


## Supported Types

### StringFieldFilterExact

```go
suggestionRuleFilterParametersActionType := components.CreateSuggestionRuleFilterParametersActionTypeStringFieldFilterExact(components.StringFieldFilterExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch suggestionRuleFilterParametersActionType.Type {
	case components.SuggestionRuleFilterParametersActionTypeTypeStringFieldFilterExact:
		// suggestionRuleFilterParametersActionType.StringFieldFilterExact is populated
}
```
