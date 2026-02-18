# EvaluationAttemptedAt


## Supported Types

### DateTimeFieldFilter

```go
evaluationAttemptedAt := components.CreateEvaluationAttemptedAtDateTimeFieldFilter(components.DateTimeFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch evaluationAttemptedAt.Type {
	case components.EvaluationAttemptedAtTypeDateTimeFieldFilter:
		// evaluationAttemptedAt.DateTimeFieldFilter is populated
}
```
