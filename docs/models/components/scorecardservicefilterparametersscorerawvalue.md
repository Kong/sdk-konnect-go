# ScorecardServiceFilterParametersScoreRawValue


## Supported Types

### NumericFieldFilter

```go
scorecardServiceFilterParametersScoreRawValue := components.CreateScorecardServiceFilterParametersScoreRawValueNumericFieldFilter(components.NumericFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scorecardServiceFilterParametersScoreRawValue.Type {
	case components.ScorecardServiceFilterParametersScoreRawValueTypeNumericFieldFilter:
		// scorecardServiceFilterParametersScoreRawValue.NumericFieldFilter is populated
}
```
