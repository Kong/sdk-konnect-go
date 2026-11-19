# CreateCostBudgetRequestBody


## Supported Types

### AICostLimitBudgetInput

```go
createCostBudgetRequestBody := operations.CreateCreateCostBudgetRequestBodyLimit(components.AICostLimitBudgetInput{/* values here */})
```

### AICostExemptBudgetInput

```go
createCostBudgetRequestBody := operations.CreateCreateCostBudgetRequestBodyExempt(components.AICostExemptBudgetInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createCostBudgetRequestBody.Type {
	case operations.CreateCostBudgetRequestBodyTypeLimit:
		// createCostBudgetRequestBody.AICostLimitBudgetInput is populated
	case operations.CreateCostBudgetRequestBodyTypeExempt:
		// createCostBudgetRequestBody.AICostExemptBudgetInput is populated
}
```
