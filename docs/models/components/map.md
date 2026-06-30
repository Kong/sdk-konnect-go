# Map


## Supported Types

### CriteriaEvaluationRelationResult

```go
map := components.CreateMapCriteriaEvaluationRelationResult(components.CriteriaEvaluationRelationResult{/* values here */})
```

### CriteriaEvaluationErrorResult

```go
map := components.CreateMapCriteriaEvaluationErrorResult(components.CriteriaEvaluationErrorResult{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch map.Type {
	case components.MapTypeCriteriaEvaluationRelationResult:
		// map.CriteriaEvaluationRelationResult is populated
	case components.MapTypeCriteriaEvaluationErrorResult:
		// map.CriteriaEvaluationErrorResult is populated
}
```
