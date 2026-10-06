# BillingCharge

Customer charge.


## Supported Types

### BillingChargeFlatFee

```go
billingCharge := components.CreateBillingChargeFlatFee(components.BillingChargeFlatFee{/* values here */})
```

### BillingChargeUsageBased

```go
billingCharge := components.CreateBillingChargeUsageBased(components.BillingChargeUsageBased{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingCharge.Type {
	case components.BillingChargeTypeFlatFee:
		// billingCharge.BillingChargeFlatFee is populated
	case components.BillingChargeTypeUsageBased:
		// billingCharge.BillingChargeUsageBased is populated
}
```
