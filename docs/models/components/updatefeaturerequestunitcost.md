# UpdateFeatureRequestUnitCost

Optional per-unit cost configuration. Use "manual" for a fixed per-unit cost, or
"llm" to look up cost from the LLM cost database based on meter group-by
properties. Set to `null` to clear the existing unit cost; omit to leave it
unchanged.


## Supported Types

### UnitCostBillingFeatureManualUnitCost

```go
updateFeatureRequestUnitCost := components.CreateUpdateFeatureRequestUnitCostManual(components.UnitCostBillingFeatureManualUnitCost{/* values here */})
```

### UnitCostBillingFeatureLLMUnitCost

```go
updateFeatureRequestUnitCost := components.CreateUpdateFeatureRequestUnitCostLlm(components.UnitCostBillingFeatureLLMUnitCost{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateFeatureRequestUnitCost.Type {
	case components.UpdateFeatureRequestUnitCostTypeManual:
		// updateFeatureRequestUnitCost.UnitCostBillingFeatureManualUnitCost is populated
	case components.UpdateFeatureRequestUnitCostTypeLlm:
		// updateFeatureRequestUnitCost.UnitCostBillingFeatureLLMUnitCost is populated
}
```
