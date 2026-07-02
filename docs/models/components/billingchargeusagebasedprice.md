# BillingChargeUsageBasedPrice

The price of the charge.


## Supported Types

### BillingPriceFree

```go
billingChargeUsageBasedPrice := components.CreateBillingChargeUsageBasedPriceFree(components.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
billingChargeUsageBasedPrice := components.CreateBillingChargeUsageBasedPriceFlat(components.BillingPriceFlat{/* values here */})
```

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
	case components.BillingChargeUsageBasedPriceTypeFree:
		// billingChargeUsageBasedPrice.BillingPriceFree is populated
	case components.BillingChargeUsageBasedPriceTypeFlat:
		// billingChargeUsageBasedPrice.BillingPriceFlat is populated
	case components.BillingChargeUsageBasedPriceTypeUnit:
		// billingChargeUsageBasedPrice.BillingPriceUnit is populated
	case components.BillingChargeUsageBasedPriceTypeGraduated:
		// billingChargeUsageBasedPrice.BillingPriceGraduated is populated
	case components.BillingChargeUsageBasedPriceTypeVolume:
		// billingChargeUsageBasedPrice.BillingPriceVolume is populated
}
```
