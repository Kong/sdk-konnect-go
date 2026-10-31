# BillingSubscriptionItemPrice

The price of the rate card.


## Supported Types

### BillingPriceFree

```go
billingSubscriptionItemPrice := components.CreateBillingSubscriptionItemPriceFree(components.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
billingSubscriptionItemPrice := components.CreateBillingSubscriptionItemPriceFlat(components.BillingPriceFlat{/* values here */})
```

### BillingPriceUnit

```go
billingSubscriptionItemPrice := components.CreateBillingSubscriptionItemPriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
billingSubscriptionItemPrice := components.CreateBillingSubscriptionItemPriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
billingSubscriptionItemPrice := components.CreateBillingSubscriptionItemPriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionItemPrice.Type {
	case components.BillingSubscriptionItemPriceTypeFree:
		// billingSubscriptionItemPrice.BillingPriceFree is populated
	case components.BillingSubscriptionItemPriceTypeFlat:
		// billingSubscriptionItemPrice.BillingPriceFlat is populated
	case components.BillingSubscriptionItemPriceTypeUnit:
		// billingSubscriptionItemPrice.BillingPriceUnit is populated
	case components.BillingSubscriptionItemPriceTypeGraduated:
		// billingSubscriptionItemPrice.BillingPriceGraduated is populated
	case components.BillingSubscriptionItemPriceTypeVolume:
		// billingSubscriptionItemPrice.BillingPriceVolume is populated
}
```
