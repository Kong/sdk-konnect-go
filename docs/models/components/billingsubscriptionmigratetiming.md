# BillingSubscriptionMigrateTiming

When the migration takes effect: immediately (default), at the next billing
cycle, or at an explicit billing-aligned timestamp. In-place migrations may
target a later phase when the phase timelines match. The target plan reference
is saved now; changed items take effect at the requested time.


## Supported Types

### BillingSubscriptionEditTimingEnum

```go
billingSubscriptionMigrateTiming := components.CreateBillingSubscriptionMigrateTimingBillingSubscriptionEditTimingEnum(components.BillingSubscriptionEditTimingEnum{/* values here */})
```

### 

```go
billingSubscriptionMigrateTiming := components.CreateBillingSubscriptionMigrateTimingDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingSubscriptionMigrateTiming.Type {
	case components.BillingSubscriptionMigrateTimingTypeBillingSubscriptionEditTimingEnum:
		// billingSubscriptionMigrateTiming.BillingSubscriptionEditTimingEnum is populated
	case components.BillingSubscriptionMigrateTimingTypeDateTime:
		// billingSubscriptionMigrateTiming.DateTime is populated
}
```
