# EvaluationSuccessfullyEvaluatedAt


## Supported Types

### DateTimeFieldFilter

```go
evaluationSuccessfullyEvaluatedAt := components.CreateEvaluationSuccessfullyEvaluatedAtDateTimeFieldFilter(components.DateTimeFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch evaluationSuccessfullyEvaluatedAt.Type {
	case components.EvaluationSuccessfullyEvaluatedAtTypeDateTimeFieldFilter:
		// evaluationSuccessfullyEvaluatedAt.DateTimeFieldFilter is populated
}
```
