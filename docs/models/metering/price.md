# Price

The price of the rate card.


## Supported Types

### BillingPriceFree

```go
price := components.CreatePriceFree(metering.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
price := components.CreatePriceFlat(metering.BillingPriceFlat{/* values here */})
```

### BillingPriceUnit

```go
price := components.CreatePriceUnit(metering.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
price := components.CreatePriceGraduated(metering.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
price := components.CreatePriceVolume(metering.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch price.Type {
	case components.PriceTypeFree:
		// price.BillingPriceFree is populated
	case components.PriceTypeFlat:
		// price.BillingPriceFlat is populated
	case components.PriceTypeUnit:
		// price.BillingPriceUnit is populated
	case components.PriceTypeGraduated:
		// price.BillingPriceGraduated is populated
	case components.PriceTypeVolume:
		// price.BillingPriceVolume is populated
}
```
