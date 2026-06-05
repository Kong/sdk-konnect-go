# BillingUsageBasedChargePrice

The price of the charge.


## Supported Types

### BillingUsageBasedChargePriceBillingPriceFree

```go
billingUsageBasedChargePrice := components.CreateBillingUsageBasedChargePriceFree(components.BillingUsageBasedChargePriceBillingPriceFree{/* values here */})
```

### BillingUsageBasedChargePriceBillingPriceFlat

```go
billingUsageBasedChargePrice := components.CreateBillingUsageBasedChargePriceFlat(components.BillingUsageBasedChargePriceBillingPriceFlat{/* values here */})
```

### BillingUsageBasedChargePriceBillingPriceUnit

```go
billingUsageBasedChargePrice := components.CreateBillingUsageBasedChargePriceUnit(components.BillingUsageBasedChargePriceBillingPriceUnit{/* values here */})
```

### BillingUsageBasedChargePriceBillingPriceGraduated

```go
billingUsageBasedChargePrice := components.CreateBillingUsageBasedChargePriceGraduated(components.BillingUsageBasedChargePriceBillingPriceGraduated{/* values here */})
```

### BillingUsageBasedChargePriceBillingPriceVolume

```go
billingUsageBasedChargePrice := components.CreateBillingUsageBasedChargePriceVolume(components.BillingUsageBasedChargePriceBillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingUsageBasedChargePrice.Type {
	case components.BillingUsageBasedChargePriceTypeFree:
		// billingUsageBasedChargePrice.BillingUsageBasedChargePriceBillingPriceFree is populated
	case components.BillingUsageBasedChargePriceTypeFlat:
		// billingUsageBasedChargePrice.BillingUsageBasedChargePriceBillingPriceFlat is populated
	case components.BillingUsageBasedChargePriceTypeUnit:
		// billingUsageBasedChargePrice.BillingUsageBasedChargePriceBillingPriceUnit is populated
	case components.BillingUsageBasedChargePriceTypeGraduated:
		// billingUsageBasedChargePrice.BillingUsageBasedChargePriceBillingPriceGraduated is populated
	case components.BillingUsageBasedChargePriceTypeVolume:
		// billingUsageBasedChargePrice.BillingUsageBasedChargePriceBillingPriceVolume is populated
}
```
