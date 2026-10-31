# CreateCreditGrantRequestCostBasis

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
createCreditGrantRequestCostBasis := components.CreateCreateCreditGrantRequestCostBasisDynamic(components.CreateChargeCostBasisDynamic{/* values here */})
```

### CreateChargeCostBasisPinned

```go
createCreditGrantRequestCostBasis := components.CreateCreateCreditGrantRequestCostBasisPinned(components.CreateChargeCostBasisPinned{/* values here */})
```

### CreateChargeCostBasisManual

```go
createCreditGrantRequestCostBasis := components.CreateCreateCreditGrantRequestCostBasisManual(components.CreateChargeCostBasisManual{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createCreditGrantRequestCostBasis.Type {
	case components.CreateCreditGrantRequestCostBasisTypeDynamic:
		// createCreditGrantRequestCostBasis.CreateChargeCostBasisDynamic is populated
	case components.CreateCreditGrantRequestCostBasisTypePinned:
		// createCreditGrantRequestCostBasis.CreateChargeCostBasisPinned is populated
	case components.CreateCreditGrantRequestCostBasisTypeManual:
		// createCreditGrantRequestCostBasis.CreateChargeCostBasisManual is populated
}
```
