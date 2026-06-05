# BillingCharge

Customer charge.


## Supported Types

### BillingFlatFeeCharge

```go
billingCharge := components.CreateBillingChargeFlatFee(components.BillingFlatFeeCharge{/* values here */})
```

### BillingUsageBasedCharge

```go
billingCharge := components.CreateBillingChargeUsageBased(components.BillingUsageBasedCharge{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingCharge.Type {
	case components.BillingChargeTypeFlatFee:
		// billingCharge.BillingFlatFeeCharge is populated
	case components.BillingChargeTypeUsageBased:
		// billingCharge.BillingUsageBasedCharge is populated
}
```
