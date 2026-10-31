# CreateFeatureRequestUnitCost

Optional per-unit cost configuration. Use "manual" for a fixed per-unit cost, or
"llm" to look up cost from the LLM cost database based on meter group-by
properties.


## Supported Types

### BillingFeatureManualUnitCost

```go
createFeatureRequestUnitCost := components.CreateCreateFeatureRequestUnitCostManual(components.BillingFeatureManualUnitCost{/* values here */})
```

### BillingFeatureLLMUnitCostInput

```go
createFeatureRequestUnitCost := components.CreateCreateFeatureRequestUnitCostLlm(components.BillingFeatureLLMUnitCostInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createFeatureRequestUnitCost.Type {
	case components.CreateFeatureRequestUnitCostTypeManual:
		// createFeatureRequestUnitCost.BillingFeatureManualUnitCost is populated
	case components.CreateFeatureRequestUnitCostTypeLlm:
		// createFeatureRequestUnitCost.BillingFeatureLLMUnitCostInput is populated
}
```
