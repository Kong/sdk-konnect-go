# AICostObservedSubjectTypeFilter

Filters a field by its value with `eq`, `neq`, or `oeq`. Accepts `principal`, `consumer`, and `oidc`. A bare value means `eq`.


## Supported Types

### AICostObservedSubjectType

```go
aiCostObservedSubjectTypeFilter := components.CreateAICostObservedSubjectTypeFilterAICostObservedSubjectType(components.AICostObservedSubjectType{/* values here */})
```

### AICostObservedSubjectTypeEqualsFilter

```go
aiCostObservedSubjectTypeFilter := components.CreateAICostObservedSubjectTypeFilterAICostObservedSubjectTypeEqualsFilter(components.AICostObservedSubjectTypeEqualsFilter{/* values here */})
```

### AICostObservedSubjectTypeNEQFilter

```go
aiCostObservedSubjectTypeFilter := components.CreateAICostObservedSubjectTypeFilterAICostObservedSubjectTypeNEQFilter(components.AICostObservedSubjectTypeNEQFilter{/* values here */})
```

### StringFieldOEQFilter

```go
aiCostObservedSubjectTypeFilter := components.CreateAICostObservedSubjectTypeFilterStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostObservedSubjectTypeFilter.Type {
	case components.AICostObservedSubjectTypeFilterTypeAICostObservedSubjectType:
		// aiCostObservedSubjectTypeFilter.AICostObservedSubjectType is populated
	case components.AICostObservedSubjectTypeFilterTypeAICostObservedSubjectTypeEqualsFilter:
		// aiCostObservedSubjectTypeFilter.AICostObservedSubjectTypeEqualsFilter is populated
	case components.AICostObservedSubjectTypeFilterTypeAICostObservedSubjectTypeNEQFilter:
		// aiCostObservedSubjectTypeFilter.AICostObservedSubjectTypeNEQFilter is populated
	case components.AICostObservedSubjectTypeFilterTypeStringFieldOEQFilter:
		// aiCostObservedSubjectTypeFilter.StringFieldOEQFilter is populated
}
```
