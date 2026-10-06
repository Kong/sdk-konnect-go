# CreateChargeUsageBasedRequestCostBasis

Defines how a custom-currency charge is converted into its fiat invoice
currency; the resolved rate is exposed through `resolved_cost_basis`. Required
when `currency` is custom and the charge settles as `credit_then_invoice`; must
be omitted otherwise.


## Supported Types

### BillingChargeCostBasisDynamic

```go
createChargeUsageBasedRequestCostBasis := components.CreateCreateChargeUsageBasedRequestCostBasisDynamic(components.BillingChargeCostBasisDynamic{/* values here */})
```

### BillingChargeCostBasisPinned

```go
createChargeUsageBasedRequestCostBasis := components.CreateCreateChargeUsageBasedRequestCostBasisPinned(components.BillingChargeCostBasisPinned{/* values here */})
```

### BillingChargeCostBasisManual

```go
createChargeUsageBasedRequestCostBasis := components.CreateCreateChargeUsageBasedRequestCostBasisManual(components.BillingChargeCostBasisManual{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createChargeUsageBasedRequestCostBasis.Type {
	case components.CreateChargeUsageBasedRequestCostBasisTypeDynamic:
		// createChargeUsageBasedRequestCostBasis.BillingChargeCostBasisDynamic is populated
	case components.CreateChargeUsageBasedRequestCostBasisTypePinned:
		// createChargeUsageBasedRequestCostBasis.BillingChargeCostBasisPinned is populated
	case components.CreateChargeUsageBasedRequestCostBasisTypeManual:
		// createChargeUsageBasedRequestCostBasis.BillingChargeCostBasisManual is populated
}
```
