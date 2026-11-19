# UpsertCostBudgetRequestBody


## Supported Types

### AICostLimitBudgetInput

```go
upsertCostBudgetRequestBody := operations.CreateUpsertCostBudgetRequestBodyLimit(components.AICostLimitBudgetInput{/* values here */})
```

### AICostExemptBudgetInput

```go
upsertCostBudgetRequestBody := operations.CreateUpsertCostBudgetRequestBodyExempt(components.AICostExemptBudgetInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch upsertCostBudgetRequestBody.Type {
	case operations.UpsertCostBudgetRequestBodyTypeLimit:
		// upsertCostBudgetRequestBody.AICostLimitBudgetInput is populated
	case operations.UpsertCostBudgetRequestBodyTypeExempt:
		// upsertCostBudgetRequestBody.AICostExemptBudgetInput is populated
}
```
