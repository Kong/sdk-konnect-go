# MetricsControlPlaneFilterByField


## Supported Types

### MultiselectFilters

```go
metricsControlPlaneFilterByField := components.CreateMetricsControlPlaneFilterByFieldMultiselectFilters(components.MultiselectFilters{/* values here */})
```

### EmptyFilters

```go
metricsControlPlaneFilterByField := components.CreateMetricsControlPlaneFilterByFieldEmptyFilters(components.EmptyFilters{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsControlPlaneFilterByField.Type {
	case components.MetricsControlPlaneFilterByFieldTypeMultiselectFilters:
		// metricsControlPlaneFilterByField.MultiselectFilters is populated
	case components.MetricsControlPlaneFilterByFieldTypeEmptyFilters:
		// metricsControlPlaneFilterByField.EmptyFilters is populated
}
```
