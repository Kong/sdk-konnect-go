# AICostBudget

A budget sets a spending limit, or declares an exemption from limits, for part of the organization's AI usage. `scope` selects the usage it covers.

`effect` tells the two kinds apart. A limit budget carries its settings in `limit`, `window`, and `posture`, and an exemption carries the period it covers in `validity`. The fields of the other kind are `null`.


## Supported Types

### AICostLimitBudget

```go
aiCostBudget := components.CreateAICostBudgetLimit(components.AICostLimitBudget{/* values here */})
```

### AICostExemptBudget

```go
aiCostBudget := components.CreateAICostBudgetExempt(components.AICostExemptBudget{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostBudget.Type {
	case components.AICostBudgetTypeLimit:
		// aiCostBudget.AICostLimitBudget is populated
	case components.AICostBudgetTypeExempt:
		// aiCostBudget.AICostExemptBudget is populated
}
```
