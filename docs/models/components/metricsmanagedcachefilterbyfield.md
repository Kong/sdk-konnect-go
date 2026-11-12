# MetricsManagedCacheFilterByField


## Supported Types

### MetricsManagedCacheFilterByFieldMultiselectFilters

```go
metricsManagedCacheFilterByField := components.CreateMetricsManagedCacheFilterByFieldMetricsManagedCacheFilterByFieldMultiselectFilters(components.MetricsManagedCacheFilterByFieldMultiselectFilters{/* values here */})
```

### MetricsManagedCacheFilterByFieldEmptyFilters

```go
metricsManagedCacheFilterByField := components.CreateMetricsManagedCacheFilterByFieldMetricsManagedCacheFilterByFieldEmptyFilters(components.MetricsManagedCacheFilterByFieldEmptyFilters{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsManagedCacheFilterByField.Type {
	case components.MetricsManagedCacheFilterByFieldTypeMetricsManagedCacheFilterByFieldMultiselectFilters:
		// metricsManagedCacheFilterByField.MetricsManagedCacheFilterByFieldMultiselectFilters is populated
	case components.MetricsManagedCacheFilterByFieldTypeMetricsManagedCacheFilterByFieldEmptyFilters:
		// metricsManagedCacheFilterByField.MetricsManagedCacheFilterByFieldEmptyFilters is populated
}
```
