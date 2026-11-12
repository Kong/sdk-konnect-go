# BillingSubscriptionEditTiming

When the requested changes should take effect. Defaults to immediate.


## Supported Types

### BillingSubscriptionEditTimingEnum

```go
billingSubscriptionEditTiming := components.CreateBillingSubscriptionEditTimingBillingSubscriptionEditTimingEnum(components.BillingSubscriptionEditTimingEnum{/* values here */})
```

### 

```go
billingSubscriptionEditTiming := components.CreateBillingSubscriptionEditTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionEditTiming.Type {
	case components.BillingSubscriptionEditTimingTypeBillingSubscriptionEditTimingEnum:
		// billingSubscriptionEditTiming.BillingSubscriptionEditTimingEnum is populated
	case components.BillingSubscriptionEditTimingTypeDateTime:
		// billingSubscriptionEditTiming.DateTime is populated
}
```
