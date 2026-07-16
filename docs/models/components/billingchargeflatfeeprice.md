# BillingChargeFlatFeePrice

The price of the charge.


## Supported Types

### PriceBillingPriceFree

```go
billingChargeFlatFeePrice := components.CreateBillingChargeFlatFeePriceFree(components.PriceBillingPriceFree{/* values here */})
```

### PriceBillingPriceFlat

```go
billingChargeFlatFeePrice := components.CreateBillingChargeFlatFeePriceFlat(components.PriceBillingPriceFlat{/* values here */})
```

### PriceBillingPriceUnit

```go
billingChargeFlatFeePrice := components.CreateBillingChargeFlatFeePriceUnit(components.PriceBillingPriceUnit{/* values here */})
```

### PriceBillingPriceGraduated

```go
billingChargeFlatFeePrice := components.CreateBillingChargeFlatFeePriceGraduated(components.PriceBillingPriceGraduated{/* values here */})
```

### PriceBillingPriceVolume

```go
billingChargeFlatFeePrice := components.CreateBillingChargeFlatFeePriceVolume(components.PriceBillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeFlatFeePrice.Type {
	case components.BillingChargeFlatFeePriceTypeFree:
		// billingChargeFlatFeePrice.PriceBillingPriceFree is populated
	case components.BillingChargeFlatFeePriceTypeFlat:
		// billingChargeFlatFeePrice.PriceBillingPriceFlat is populated
	case components.BillingChargeFlatFeePriceTypeUnit:
		// billingChargeFlatFeePrice.PriceBillingPriceUnit is populated
	case components.BillingChargeFlatFeePriceTypeGraduated:
		// billingChargeFlatFeePrice.PriceBillingPriceGraduated is populated
	case components.BillingChargeFlatFeePriceTypeVolume:
		// billingChargeFlatFeePrice.PriceBillingPriceVolume is populated
}
```
