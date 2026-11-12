# BillingChargeUsageBasedPrice

The price of the charge.

`free` prices are rejected on create: a usage-based charge always carries a
concrete price.


## Supported Types

### BillingPriceUnit

```go
billingChargeUsageBasedPrice := components.CreateBillingChargeUsageBasedPriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
billingChargeUsageBasedPrice := components.CreateBillingChargeUsageBasedPriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
billingChargeUsageBasedPrice := components.CreateBillingChargeUsageBasedPriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeUsageBasedPrice.Type {
	case components.BillingChargeUsageBasedPriceTypeUnit:
		// billingChargeUsageBasedPrice.BillingPriceUnit is populated
	case components.BillingChargeUsageBasedPriceTypeGraduated:
		// billingChargeUsageBasedPrice.BillingPriceGraduated is populated
	case components.BillingChargeUsageBasedPriceTypeVolume:
		// billingChargeUsageBasedPrice.BillingPriceVolume is populated
}
```
