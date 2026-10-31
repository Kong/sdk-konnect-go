# Criteria


## Supported Types

### UpdateScorecardCriteria

```go
criteria := components.CreateCriteriaUpdateScorecardCriteria(components.UpdateScorecardCriteria{/* values here */})
```

### CreateScorecardCriteria

```go
criteria := components.CreateCriteriaCreateScorecardCriteria(components.CreateScorecardCriteria{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch criteria.Type {
	case components.CriteriaTypeUpdateScorecardCriteria:
		// criteria.UpdateScorecardCriteria is populated
	case components.CriteriaTypeCreateScorecardCriteria:
		// criteria.CreateScorecardCriteria is populated
}
```
