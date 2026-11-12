# ControlPlaneGeoFieldFilter


## Supported Types

### ControlPlaneGeoFieldEqualsFilter

```go
controlPlaneGeoFieldFilter := components.CreateControlPlaneGeoFieldFilterControlPlaneGeoFieldEqualsFilter(components.ControlPlaneGeoFieldEqualsFilter{/* values here */})
```

### ControlPlaneGeoFieldNotEqualsFilter

```go
controlPlaneGeoFieldFilter := components.CreateControlPlaneGeoFieldFilterControlPlaneGeoFieldNotEqualsFilter(components.ControlPlaneGeoFieldNotEqualsFilter{/* values here */})
```

### ControlPlaneGeoFieldOrEqualityFilter

```go
controlPlaneGeoFieldFilter := components.CreateControlPlaneGeoFieldFilterControlPlaneGeoFieldOrEqualityFilter(components.ControlPlaneGeoFieldOrEqualityFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch controlPlaneGeoFieldFilter.Type {
	case components.ControlPlaneGeoFieldFilterTypeControlPlaneGeoFieldEqualsFilter:
		// controlPlaneGeoFieldFilter.ControlPlaneGeoFieldEqualsFilter is populated
	case components.ControlPlaneGeoFieldFilterTypeControlPlaneGeoFieldNotEqualsFilter:
		// controlPlaneGeoFieldFilter.ControlPlaneGeoFieldNotEqualsFilter is populated
	case components.ControlPlaneGeoFieldFilterTypeControlPlaneGeoFieldOrEqualityFilter:
		// controlPlaneGeoFieldFilter.ControlPlaneGeoFieldOrEqualityFilter is populated
}
```
