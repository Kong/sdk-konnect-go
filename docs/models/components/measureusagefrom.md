# MeasureUsageFrom

The time from which usage is measured. Defaults to the entitlement creation
time.


## Supported Types

### BillingEntitlementMeasureUsageFromPreset

```go
measureUsageFrom := components.CreateMeasureUsageFromBillingEntitlementMeasureUsageFromPreset(components.BillingEntitlementMeasureUsageFromPreset{/* values here */})
```

### 

```go
measureUsageFrom := components.CreateMeasureUsageFromDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch measureUsageFrom.Type {
	case components.MeasureUsageFromTypeBillingEntitlementMeasureUsageFromPreset:
		// measureUsageFrom.BillingEntitlementMeasureUsageFromPreset is populated
	case components.MeasureUsageFromTypeDateTime:
		// measureUsageFrom.DateTime is populated
}
```
