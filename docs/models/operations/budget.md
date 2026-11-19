# Budget

The limit budget the balance belongs to: a reference by default, or the full budget when the request sends `expand=budget`, so a list can show its name and scope without reading it.


## Supported Types

### AICostLimitBudget

```go
budget := operations.CreateBudgetAICostLimitBudget(components.AICostLimitBudget{/* values here */})
```

### AICostBudgetReference

```go
budget := operations.CreateBudgetAICostBudgetReference(components.AICostBudgetReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch budget.Type {
	case operations.BudgetTypeAICostLimitBudget:
		// budget.AICostLimitBudget is populated
	case operations.BudgetTypeAICostBudgetReference:
		// budget.AICostBudgetReference is populated
}
```
