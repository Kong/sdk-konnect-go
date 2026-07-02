# CreateSubscriptionAddonRequestTiming

The timing of the operation. After the create or update, a new entry will be
created in the timeline.


## Supported Types

### BillingSubscriptionEditTimingEnum

```go
createSubscriptionAddonRequestTiming := components.CreateCreateSubscriptionAddonRequestTimingBillingSubscriptionEditTimingEnum(components.BillingSubscriptionEditTimingEnum{/* values here */})
```

### 

```go
createSubscriptionAddonRequestTiming := components.CreateCreateSubscriptionAddonRequestTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createSubscriptionAddonRequestTiming.Type {
	case components.CreateSubscriptionAddonRequestTimingTypeBillingSubscriptionEditTimingEnum:
		// createSubscriptionAddonRequestTiming.BillingSubscriptionEditTimingEnum is populated
	case components.CreateSubscriptionAddonRequestTimingTypeDateTime:
		// createSubscriptionAddonRequestTiming.DateTime is populated
}
```
