# CostBasis

Defines how custom-currency credits are priced in the purchase `currency`; the
resolved rate is exposed through `resolved_cost_basis`.

Fiat grants accept only a `manual` cost basis without `fiat_currency`, where
`rate` is the fiat cost per credit unit. Custom-currency grants require a cost
basis of any type with `fiat_currency` set and equal to the purchase `currency`.
A `dynamic` cost basis is resolved at the grant's effective time, so the
currency cost basis must be effective by then. Cannot be combined with
`per_unit_cost_basis`.


## Supported Types

### CreateChargeCostBasisDynamic

```go
costBasis := components.CreateCostBasisDynamic(components.CreateChargeCostBasisDynamic{/* values here */})
```

### CreateChargeCostBasisPinned

```go
costBasis := components.CreateCostBasisPinned(components.CreateChargeCostBasisPinned{/* values here */})
```

### CreateChargeCostBasisManual

```go
costBasis := components.CreateCostBasisManual(components.CreateChargeCostBasisManual{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch costBasis.Type {
	case components.CostBasisTypeDynamic:
		// costBasis.CreateChargeCostBasisDynamic is populated
	case components.CostBasisTypePinned:
		// costBasis.CreateChargeCostBasisPinned is populated
	case components.CostBasisTypeManual:
		// costBasis.CreateChargeCostBasisManual is populated
}
```
