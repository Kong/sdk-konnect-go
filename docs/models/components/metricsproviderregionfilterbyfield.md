# MetricsProviderRegionFilterByField


## Supported Types

### MetricsProviderRegionFilterByFieldMultiselectFilters

```go
metricsProviderRegionFilterByField := components.CreateMetricsProviderRegionFilterByFieldMetricsProviderRegionFilterByFieldMultiselectFilters(components.MetricsProviderRegionFilterByFieldMultiselectFilters{/* values here */})
```

### MetricsProviderRegionFilterByFieldEmptyFilters

```go
metricsProviderRegionFilterByField := components.CreateMetricsProviderRegionFilterByFieldMetricsProviderRegionFilterByFieldEmptyFilters(components.MetricsProviderRegionFilterByFieldEmptyFilters{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsProviderRegionFilterByField.Type {
	case components.MetricsProviderRegionFilterByFieldTypeMetricsProviderRegionFilterByFieldMultiselectFilters:
		// metricsProviderRegionFilterByField.MetricsProviderRegionFilterByFieldMultiselectFilters is populated
	case components.MetricsProviderRegionFilterByFieldTypeMetricsProviderRegionFilterByFieldEmptyFilters:
		// metricsProviderRegionFilterByField.MetricsProviderRegionFilterByFieldEmptyFilters is populated
}
```
