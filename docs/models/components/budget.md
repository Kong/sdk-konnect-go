# Budget

The limit budget whose threshold was crossed: a reference by default, or the full budget when the request sends `expand=budget`.


## Supported Types

### AICostLimitBudget

```go
budget := components.CreateBudgetAICostLimitBudget(components.AICostLimitBudget{/* values here */})
```

### AICostBudgetReference

```go
budget := components.CreateBudgetAICostBudgetReference(components.AICostBudgetReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch budget.Type {
	case components.BudgetTypeAICostLimitBudget:
		// budget.AICostLimitBudget is populated
	case components.BudgetTypeAICostBudgetReference:
		// budget.AICostBudgetReference is populated
}
```
