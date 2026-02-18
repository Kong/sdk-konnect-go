# ScorecardEntitySelector

Selector used to dynamically target catalog entities that will be
included in the given scorecard's evaluated score.



## Supported Types

### ServiceSelector

```go
scorecardEntitySelector := components.CreateScorecardEntitySelectorServiceSelector(components.ServiceSelector{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scorecardEntitySelector.Type {
	case components.ScorecardEntitySelectorTypeServiceSelector:
		// scorecardEntitySelector.ServiceSelector is populated
}
```
