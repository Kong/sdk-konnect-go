# FeatureFeatureUnitCost

Optional per-unit cost configuration. Use "manual" for a fixed per-unit cost, or
"llm" to look up cost from the LLM cost database based on meter group-by
properties.


## Supported Types

### BillingFeatureManualUnitCost

```go
featureFeatureUnitCost := components.CreateFeatureFeatureUnitCostManual(components.BillingFeatureManualUnitCost{/* values here */})
```

### BillingFeatureLLMUnitCost

```go
featureFeatureUnitCost := components.CreateFeatureFeatureUnitCostLlm(components.BillingFeatureLLMUnitCost{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch featureFeatureUnitCost.Type {
	case components.FeatureFeatureUnitCostTypeManual:
		// featureFeatureUnitCost.BillingFeatureManualUnitCost is populated
	case components.FeatureFeatureUnitCostTypeLlm:
		// featureFeatureUnitCost.BillingFeatureLLMUnitCost is populated
}
```
