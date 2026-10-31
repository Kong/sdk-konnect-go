# BillingSubscriptionCreateTiming

When the subscription should start. If not provided, the subscription starts
immediately. Provide a future timestamp to schedule the subscription to start
later — this creates a not-yet-active, scheduled subscription. A timestamp in
the past is rejected.


## Supported Types

### BillingSubscriptionCreateTimingEnum

```go
billingSubscriptionCreateTiming := components.CreateBillingSubscriptionCreateTimingBillingSubscriptionCreateTimingEnum(components.BillingSubscriptionCreateTimingEnum{/* values here */})
```

### 

```go
billingSubscriptionCreateTiming := components.CreateBillingSubscriptionCreateTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionCreateTiming.Type {
	case components.BillingSubscriptionCreateTimingTypeBillingSubscriptionCreateTimingEnum:
		// billingSubscriptionCreateTiming.BillingSubscriptionCreateTimingEnum is populated
	case components.BillingSubscriptionCreateTimingTypeDateTime:
		// billingSubscriptionCreateTiming.DateTime is populated
}
```
