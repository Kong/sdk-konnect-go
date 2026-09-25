# BillingSubscriptionAddonUpdateTiming

The timing of the update. A new entry is appended to the add-on's timeline at
this point.


## Supported Types

### BillingSubscriptionEditTimingEnum

```go
billingSubscriptionAddonUpdateTiming := components.CreateBillingSubscriptionAddonUpdateTimingBillingSubscriptionEditTimingEnum(components.BillingSubscriptionEditTimingEnum{/* values here */})
```

### 

```go
billingSubscriptionAddonUpdateTiming := components.CreateBillingSubscriptionAddonUpdateTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionAddonUpdateTiming.Type {
	case components.BillingSubscriptionAddonUpdateTimingTypeBillingSubscriptionEditTimingEnum:
		// billingSubscriptionAddonUpdateTiming.BillingSubscriptionEditTimingEnum is populated
	case components.BillingSubscriptionAddonUpdateTimingTypeDateTime:
		// billingSubscriptionAddonUpdateTiming.DateTime is populated
}
```
