# Timing

If not provided the subscription is canceled immediately.


## Supported Types

### BillingSubscriptionEditTimingEnum

```go
timing := components.CreateTimingBillingSubscriptionEditTimingEnum(components.BillingSubscriptionEditTimingEnum{/* values here */})
```

### 

```go
timing := components.CreateTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch timing.Type {
	case components.TimingTypeBillingSubscriptionEditTimingEnum:
		// timing.BillingSubscriptionEditTimingEnum is populated
	case components.TimingTypeDateTime:
		// timing.DateTime is populated
}
```
