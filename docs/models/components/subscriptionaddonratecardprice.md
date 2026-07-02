# SubscriptionAddonRateCardPrice

The price of the rate card.


## Supported Types

### BillingPriceFree

```go
subscriptionAddonRateCardPrice := components.CreateSubscriptionAddonRateCardPriceFree(components.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
subscriptionAddonRateCardPrice := components.CreateSubscriptionAddonRateCardPriceFlat(components.BillingPriceFlat{/* values here */})
```

### BillingPriceUnit

```go
subscriptionAddonRateCardPrice := components.CreateSubscriptionAddonRateCardPriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
subscriptionAddonRateCardPrice := components.CreateSubscriptionAddonRateCardPriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
subscriptionAddonRateCardPrice := components.CreateSubscriptionAddonRateCardPriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionAddonRateCardPrice.Type {
	case components.SubscriptionAddonRateCardPriceTypeFree:
		// subscriptionAddonRateCardPrice.BillingPriceFree is populated
	case components.SubscriptionAddonRateCardPriceTypeFlat:
		// subscriptionAddonRateCardPrice.BillingPriceFlat is populated
	case components.SubscriptionAddonRateCardPriceTypeUnit:
		// subscriptionAddonRateCardPrice.BillingPriceUnit is populated
	case components.SubscriptionAddonRateCardPriceTypeGraduated:
		// subscriptionAddonRateCardPrice.BillingPriceGraduated is populated
	case components.SubscriptionAddonRateCardPriceTypeVolume:
		// subscriptionAddonRateCardPrice.BillingPriceVolume is populated
}
```
