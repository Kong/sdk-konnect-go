# UnitCost

Optional per-unit cost configuration. Use "manual" for a fixed per-unit cost, or
"llm" to look up cost from the LLM cost database based on meter group-by
properties.


## Supported Types

### BillingFeatureManualUnitCost

```go
unitCost := components.CreateUnitCostManual(metering.BillingFeatureManualUnitCost{/* values here */})
```

### BillingFeatureLLMUnitCost

```go
unitCost := components.CreateUnitCostLlm(metering.BillingFeatureLLMUnitCost{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unitCost.Type {
	case components.UnitCostUnionTypeManual:
		// unitCost.BillingFeatureManualUnitCost is populated
	case components.UnitCostUnionTypeLlm:
		// unitCost.BillingFeatureLLMUnitCost is populated
}
```
