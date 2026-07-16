# CreateChargeUsageBasedRequestPrice

The price of the charge.


## Supported Types

### BillingPriceFree

```go
createChargeUsageBasedRequestPrice := components.CreateCreateChargeUsageBasedRequestPriceFree(components.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
createChargeUsageBasedRequestPrice := components.CreateCreateChargeUsageBasedRequestPriceFlat(components.BillingPriceFlat{/* values here */})
```

### BillingPriceUnit

```go
createChargeUsageBasedRequestPrice := components.CreateCreateChargeUsageBasedRequestPriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
createChargeUsageBasedRequestPrice := components.CreateCreateChargeUsageBasedRequestPriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
createChargeUsageBasedRequestPrice := components.CreateCreateChargeUsageBasedRequestPriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createChargeUsageBasedRequestPrice.Type {
	case components.CreateChargeUsageBasedRequestPriceTypeFree:
		// createChargeUsageBasedRequestPrice.BillingPriceFree is populated
	case components.CreateChargeUsageBasedRequestPriceTypeFlat:
		// createChargeUsageBasedRequestPrice.BillingPriceFlat is populated
	case components.CreateChargeUsageBasedRequestPriceTypeUnit:
		// createChargeUsageBasedRequestPrice.BillingPriceUnit is populated
	case components.CreateChargeUsageBasedRequestPriceTypeGraduated:
		// createChargeUsageBasedRequestPrice.BillingPriceGraduated is populated
	case components.CreateChargeUsageBasedRequestPriceTypeVolume:
		// createChargeUsageBasedRequestPrice.BillingPriceVolume is populated
}
```
