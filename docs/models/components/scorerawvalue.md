# ScoreRawValue


## Supported Types

### NumericFieldFilter

```go
scoreRawValue := components.CreateScoreRawValueNumericFieldFilter(components.NumericFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scoreRawValue.Type {
	case components.ScoreRawValueTypeNumericFieldFilter:
		// scoreRawValue.NumericFieldFilter is populated
}
```
