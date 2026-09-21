# FeatureUnitCost

Optional per-unit cost configuration. Use "manual" for a fixed per-unit cost, or
"llm" to look up cost from the LLM cost database based on meter group-by
properties.


## Supported Types

### BillingFeatureManualUnitCost

```go
featureUnitCost := components.CreateFeatureUnitCostManual(components.BillingFeatureManualUnitCost{/* values here */})
```

### BillingFeatureLLMUnitCost

```go
featureUnitCost := components.CreateFeatureUnitCostLlm(components.BillingFeatureLLMUnitCost{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch featureUnitCost.Type {
	case components.FeatureUnitCostTypeManual:
		// featureUnitCost.BillingFeatureManualUnitCost is populated
	case components.FeatureUnitCostTypeLlm:
		// featureUnitCost.BillingFeatureLLMUnitCost is populated
}
```
