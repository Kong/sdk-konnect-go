# BillingChargeFlatFeeFeature

The feature associated with the charge, when applicable.


## Supported Types

### FeatureFeature

```go
billingChargeFlatFeeFeature := components.CreateBillingChargeFlatFeeFeatureFeatureFeature(components.FeatureFeature{/* values here */})
```

### FeatureFeatureReference

```go
billingChargeFlatFeeFeature := components.CreateBillingChargeFlatFeeFeatureFeatureFeatureReference(components.FeatureFeatureReference{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingChargeFlatFeeFeature.Type {
	case components.BillingChargeFlatFeeFeatureTypeFeatureFeature:
		// billingChargeFlatFeeFeature.FeatureFeature is populated
	case components.BillingChargeFlatFeeFeatureTypeFeatureFeatureReference:
		// billingChargeFlatFeeFeature.FeatureFeatureReference is populated
}
```
