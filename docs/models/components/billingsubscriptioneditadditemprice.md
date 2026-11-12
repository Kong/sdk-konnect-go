# BillingSubscriptionEditAddItemPrice

The price of the rate card.


## Supported Types

### BillingPriceFree

```go
billingSubscriptionEditAddItemPrice := components.CreateBillingSubscriptionEditAddItemPriceFree(components.BillingPriceFree{/* values here */})
```

### BillingPriceFlat

```go
billingSubscriptionEditAddItemPrice := components.CreateBillingSubscriptionEditAddItemPriceFlat(components.BillingPriceFlat{/* values here */})
```

### BillingPriceUnit

```go
billingSubscriptionEditAddItemPrice := components.CreateBillingSubscriptionEditAddItemPriceUnit(components.BillingPriceUnit{/* values here */})
```

### BillingPriceGraduated

```go
billingSubscriptionEditAddItemPrice := components.CreateBillingSubscriptionEditAddItemPriceGraduated(components.BillingPriceGraduated{/* values here */})
```

### BillingPriceVolume

```go
billingSubscriptionEditAddItemPrice := components.CreateBillingSubscriptionEditAddItemPriceVolume(components.BillingPriceVolume{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionEditAddItemPrice.Type {
	case components.BillingSubscriptionEditAddItemPriceTypeFree:
		// billingSubscriptionEditAddItemPrice.BillingPriceFree is populated
	case components.BillingSubscriptionEditAddItemPriceTypeFlat:
		// billingSubscriptionEditAddItemPrice.BillingPriceFlat is populated
	case components.BillingSubscriptionEditAddItemPriceTypeUnit:
		// billingSubscriptionEditAddItemPrice.BillingPriceUnit is populated
	case components.BillingSubscriptionEditAddItemPriceTypeGraduated:
		// billingSubscriptionEditAddItemPrice.BillingPriceGraduated is populated
	case components.BillingSubscriptionEditAddItemPriceTypeVolume:
		// billingSubscriptionEditAddItemPrice.BillingPriceVolume is populated
}
```
