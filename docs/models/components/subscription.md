# Subscription

The subscription that originated the charge, when the charge was created from a
subscription item.

By default, only the `id`, `phase.id`, `phase.item.id` of the subscription is
returned. For more details use the `subscription` expand.


## Supported Types

### BillingChargeFlatFeeSubscriptionBillingSubscription

```go
subscription := components.CreateSubscriptionBillingChargeFlatFeeSubscriptionBillingSubscription(components.BillingChargeFlatFeeSubscriptionBillingSubscription{/* values here */})
```

### BillingSubscriptionReference

```go
subscription := components.CreateSubscriptionBillingSubscriptionReference(components.BillingSubscriptionReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscription.Type {
	case components.SubscriptionTypeBillingChargeFlatFeeSubscriptionBillingSubscription:
		// subscription.BillingChargeFlatFeeSubscriptionBillingSubscription is populated
	case components.SubscriptionTypeBillingSubscriptionReference:
		// subscription.BillingSubscriptionReference is populated
}
```
