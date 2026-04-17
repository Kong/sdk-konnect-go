# BillingSubscriptionChangeTiming

Timing configuration for the change, when the change should take effect.
For changing a subscription, the accepted values depend on the subscription configuration.


## Supported Types

### BillingSubscriptionEditTimingEnum

```go
billingSubscriptionChangeTiming := components.CreateBillingSubscriptionChangeTimingBillingSubscriptionEditTimingEnum(components.BillingSubscriptionEditTimingEnum{/* values here */})
```

### 

```go
billingSubscriptionChangeTiming := components.CreateBillingSubscriptionChangeTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionChangeTiming.Type {
	case components.BillingSubscriptionChangeTimingTypeBillingSubscriptionEditTimingEnum:
		// billingSubscriptionChangeTiming.BillingSubscriptionEditTimingEnum is populated
	case components.BillingSubscriptionChangeTimingTypeDateTime:
		// billingSubscriptionChangeTiming.DateTime is populated
}
```
