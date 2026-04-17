# ControlPlaneGeoFieldEqualsFilter

Filter a control-plane geo by exact match.


## Supported Types

### ControlPlaneGeo

```go
controlPlaneGeoFieldEqualsFilter := components.CreateControlPlaneGeoFieldEqualsFilterControlPlaneGeo(components.ControlPlaneGeo{/* values here */})
```

### ControlPlaneGeoFieldEqualsComparison

```go
controlPlaneGeoFieldEqualsFilter := components.CreateControlPlaneGeoFieldEqualsFilterControlPlaneGeoFieldEqualsComparison(components.ControlPlaneGeoFieldEqualsComparison{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch controlPlaneGeoFieldEqualsFilter.Type {
	case components.ControlPlaneGeoFieldEqualsFilterTypeControlPlaneGeo:
		// controlPlaneGeoFieldEqualsFilter.ControlPlaneGeo is populated
	case components.ControlPlaneGeoFieldEqualsFilterTypeControlPlaneGeoFieldEqualsComparison:
		// controlPlaneGeoFieldEqualsFilter.ControlPlaneGeoFieldEqualsComparison is populated
}
```
