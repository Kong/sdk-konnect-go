# EvaluationPassingServicesCount


## Supported Types

### NumericFieldFilter

```go
evaluationPassingServicesCount := components.CreateEvaluationPassingServicesCountNumericFieldFilter(components.NumericFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch evaluationPassingServicesCount.Type {
	case components.EvaluationPassingServicesCountTypeNumericFieldFilter:
		// evaluationPassingServicesCount.NumericFieldFilter is populated
}
```
