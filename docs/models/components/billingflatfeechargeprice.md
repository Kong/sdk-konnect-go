# BillingFlatFeeChargePrice

The price of the charge.


## Supported Types

### PriceBillingPriceFree

```go
billingFlatFeeChargePrice := components.CreateBillingFlatFeeChargePriceFree(components.PriceBillingPriceFree{/* values here */})
```

### PriceBillingPriceFlat

```go
billingFlatFeeChargePrice := components.CreateBillingFlatFeeChargePriceFlat(components.PriceBillingPriceFlat{/* values here */})
```

### PriceBillingPriceUnit

```go
billingFlatFeeChargePrice := components.CreateBillingFlatFeeChargePriceUnit(components.PriceBillingPriceUnit{/* values here */})
```

### PriceBillingPriceGraduated

```go
billingFlatFeeChargePrice := components.CreateBillingFlatFeeChargePriceGraduated(components.PriceBillingPriceGraduated{/* values here */})
```

### PriceBillingPriceVolume

```go
billingFlatFeeChargePrice := components.CreateBillingFlatFeeChargePriceVolume(components.PriceBillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingFlatFeeChargePrice.Type {
	case components.BillingFlatFeeChargePriceTypeFree:
		// billingFlatFeeChargePrice.PriceBillingPriceFree is populated
	case components.BillingFlatFeeChargePriceTypeFlat:
		// billingFlatFeeChargePrice.PriceBillingPriceFlat is populated
	case components.BillingFlatFeeChargePriceTypeUnit:
		// billingFlatFeeChargePrice.PriceBillingPriceUnit is populated
	case components.BillingFlatFeeChargePriceTypeGraduated:
		// billingFlatFeeChargePrice.PriceBillingPriceGraduated is populated
	case components.BillingFlatFeeChargePriceTypeVolume:
		// billingFlatFeeChargePrice.PriceBillingPriceVolume is populated
}
```
