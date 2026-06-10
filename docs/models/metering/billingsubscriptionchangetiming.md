# BillingSubscriptionChangeTiming

Timing configuration for the change, when the change should take effect. For
changing a subscription, the accepted values depend on the subscription
configuration.


## Supported Types

### TimingBillingSubscriptionEditTimingEnum

```go
billingSubscriptionChangeTiming := components.CreateBillingSubscriptionChangeTimingTimingBillingSubscriptionEditTimingEnum(metering.TimingBillingSubscriptionEditTimingEnum{/* values here */})
```

### 

```go
billingSubscriptionChangeTiming := components.CreateBillingSubscriptionChangeTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionChangeTiming.Type {
	case components.BillingSubscriptionChangeTimingTypeTimingBillingSubscriptionEditTimingEnum:
		// billingSubscriptionChangeTiming.TimingBillingSubscriptionEditTimingEnum is populated
	case components.BillingSubscriptionChangeTimingTypeDateTime:
		// billingSubscriptionChangeTiming.DateTime is populated
}
```
