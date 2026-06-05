# Price

The price of the rate card.


## Supported Types

### BillingPriceFree

```go
price := components.CreatePriceFree(components.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
price := components.CreatePriceFlat(components.BillingPriceFlat{/* values here */})
```

### BillingPriceUnit

```go
price := components.CreatePriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
price := components.CreatePriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
price := components.CreatePriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch price.Type {
	case components.PriceUnionTypeFree:
		// price.BillingPriceFree is populated
	case components.PriceUnionTypeFlat:
		// price.BillingPriceFlat is populated
	case components.PriceUnionTypeUnit:
		// price.BillingPriceUnit is populated
	case components.PriceUnionTypeGraduated:
		// price.BillingPriceGraduated is populated
	case components.PriceUnionTypeVolume:
		// price.BillingPriceVolume is populated
}
```
