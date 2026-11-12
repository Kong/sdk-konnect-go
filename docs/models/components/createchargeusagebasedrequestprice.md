# CreateChargeUsageBasedRequestPrice

The price of the charge.

`free` prices are rejected on create: a usage-based charge always carries a
concrete price.


## Supported Types

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
	case components.CreateChargeUsageBasedRequestPriceTypeUnit:
		// createChargeUsageBasedRequestPrice.BillingPriceUnit is populated
	case components.CreateChargeUsageBasedRequestPriceTypeGraduated:
		// createChargeUsageBasedRequestPrice.BillingPriceGraduated is populated
	case components.CreateChargeUsageBasedRequestPriceTypeVolume:
		// createChargeUsageBasedRequestPrice.BillingPriceVolume is populated
}
```
