# ScorecardServiceFilterParametersScoreValue


## Supported Types

### StringFieldFilter

```go
scorecardServiceFilterParametersScoreValue := components.CreateScorecardServiceFilterParametersScoreValueStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scorecardServiceFilterParametersScoreValue.Type {
	case components.ScorecardServiceFilterParametersScoreValueTypeStringFieldFilter:
		// scorecardServiceFilterParametersScoreValue.StringFieldFilter is populated
}
```
