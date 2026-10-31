# BillingChargeUsageBasedSystemIntentPrice

The price of the charge.

`free` prices are rejected on create: a usage-based charge always carries a
concrete price.


## Supported Types

### BillingPriceUnit

```go
billingChargeUsageBasedSystemIntentPrice := components.CreateBillingChargeUsageBasedSystemIntentPriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
billingChargeUsageBasedSystemIntentPrice := components.CreateBillingChargeUsageBasedSystemIntentPriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
billingChargeUsageBasedSystemIntentPrice := components.CreateBillingChargeUsageBasedSystemIntentPriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeUsageBasedSystemIntentPrice.Type {
	case components.BillingChargeUsageBasedSystemIntentPriceTypeUnit:
		// billingChargeUsageBasedSystemIntentPrice.BillingPriceUnit is populated
	case components.BillingChargeUsageBasedSystemIntentPriceTypeGraduated:
		// billingChargeUsageBasedSystemIntentPrice.BillingPriceGraduated is populated
	case components.BillingChargeUsageBasedSystemIntentPriceTypeVolume:
		// billingChargeUsageBasedSystemIntentPrice.BillingPriceVolume is populated
}
```
