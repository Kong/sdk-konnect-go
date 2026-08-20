# MetricsDataPlaneGroupFilterByField


## Supported Types

### MetricsDataPlaneGroupFilterByFieldMultiselectFilters

```go
metricsDataPlaneGroupFilterByField := components.CreateMetricsDataPlaneGroupFilterByFieldMetricsDataPlaneGroupFilterByFieldMultiselectFilters(components.MetricsDataPlaneGroupFilterByFieldMultiselectFilters{/* values here */})
```

### MetricsDataPlaneGroupFilterByFieldEmptyFilters

```go
metricsDataPlaneGroupFilterByField := components.CreateMetricsDataPlaneGroupFilterByFieldMetricsDataPlaneGroupFilterByFieldEmptyFilters(components.MetricsDataPlaneGroupFilterByFieldEmptyFilters{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsDataPlaneGroupFilterByField.Type {
	case components.MetricsDataPlaneGroupFilterByFieldTypeMetricsDataPlaneGroupFilterByFieldMultiselectFilters:
		// metricsDataPlaneGroupFilterByField.MetricsDataPlaneGroupFilterByFieldMultiselectFilters is populated
	case components.MetricsDataPlaneGroupFilterByFieldTypeMetricsDataPlaneGroupFilterByFieldEmptyFilters:
		// metricsDataPlaneGroupFilterByField.MetricsDataPlaneGroupFilterByFieldEmptyFilters is populated
}
```
