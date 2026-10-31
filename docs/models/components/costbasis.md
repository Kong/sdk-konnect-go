# CostBasis

Defines how a custom-currency charge is converted into its fiat invoice
currency; the resolved rate is exposed through `resolved_cost_basis`. Required
when `currency` is custom and the charge settles as `credit_then_invoice`; must
be omitted otherwise.


## Supported Types

### BillingChargeCostBasisDynamic

```go
costBasis := components.CreateCostBasisDynamic(components.BillingChargeCostBasisDynamic{/* values here */})
```

### BillingChargeCostBasisPinned

```go
costBasis := components.CreateCostBasisPinned(components.BillingChargeCostBasisPinned{/* values here */})
```

### BillingChargeCostBasisManual

```go
costBasis := components.CreateCostBasisManual(components.BillingChargeCostBasisManual{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch costBasis.Type {
	case components.CostBasisTypeDynamic:
		// costBasis.BillingChargeCostBasisDynamic is populated
	case components.CostBasisTypePinned:
		// costBasis.BillingChargeCostBasisPinned is populated
	case components.CostBasisTypeManual:
		// costBasis.BillingChargeCostBasisManual is populated
}
```
