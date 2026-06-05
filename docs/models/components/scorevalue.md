# ScoreValue


## Supported Types

### StringFieldFilter

```go
scoreValue := components.CreateScoreValueStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scoreValue.Type {
	case components.ScoreValueTypeStringFieldFilter:
		// scoreValue.StringFieldFilter is populated
}
```
