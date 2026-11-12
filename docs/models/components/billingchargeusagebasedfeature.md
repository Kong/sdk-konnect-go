# BillingChargeUsageBasedFeature

The feature associated with the charge.


## Supported Types

### BillingChargeUsageBasedFeatureFeature

```go
billingChargeUsageBasedFeature := components.CreateBillingChargeUsageBasedFeatureBillingChargeUsageBasedFeatureFeature(components.BillingChargeUsageBasedFeatureFeature{/* values here */})
```

### BillingChargeUsageBasedFeatureFeatureReference

```go
billingChargeUsageBasedFeature := components.CreateBillingChargeUsageBasedFeatureBillingChargeUsageBasedFeatureFeatureReference(components.BillingChargeUsageBasedFeatureFeatureReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeUsageBasedFeature.Type {
	case components.BillingChargeUsageBasedFeatureTypeBillingChargeUsageBasedFeatureFeature:
		// billingChargeUsageBasedFeature.BillingChargeUsageBasedFeatureFeature is populated
	case components.BillingChargeUsageBasedFeatureTypeBillingChargeUsageBasedFeatureFeatureReference:
		// billingChargeUsageBasedFeature.BillingChargeUsageBasedFeatureFeatureReference is populated
}
```
