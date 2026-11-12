# BillingChargeUsageBasedSubscription

The subscription that originated the charge, when the charge was created from a
subscription item.

By default, only the `id`, `phase.id`, `phase.item.id` of the subscription is
returned. For more details use the `subscription` expand.


## Supported Types

### SubscriptionBillingSubscription

```go
billingChargeUsageBasedSubscription := components.CreateBillingChargeUsageBasedSubscriptionSubscriptionBillingSubscription(components.SubscriptionBillingSubscription{/* values here */})
```

### SubscriptionBillingSubscriptionReference

```go
billingChargeUsageBasedSubscription := components.CreateBillingChargeUsageBasedSubscriptionSubscriptionBillingSubscriptionReference(components.SubscriptionBillingSubscriptionReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeUsageBasedSubscription.Type {
	case components.BillingChargeUsageBasedSubscriptionTypeSubscriptionBillingSubscription:
		// billingChargeUsageBasedSubscription.SubscriptionBillingSubscription is populated
	case components.BillingChargeUsageBasedSubscriptionTypeSubscriptionBillingSubscriptionReference:
		// billingChargeUsageBasedSubscription.SubscriptionBillingSubscriptionReference is populated
}
```
