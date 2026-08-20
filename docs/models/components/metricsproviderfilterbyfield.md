# MetricsProviderFilterByField


## Supported Types

### MetricsProviderFilterByFieldMultiselectFilters

```go
metricsProviderFilterByField := components.CreateMetricsProviderFilterByFieldMetricsProviderFilterByFieldMultiselectFilters(components.MetricsProviderFilterByFieldMultiselectFilters{/* values here */})
```

### MetricsProviderFilterByFieldEmptyFilters

```go
metricsProviderFilterByField := components.CreateMetricsProviderFilterByFieldMetricsProviderFilterByFieldEmptyFilters(components.MetricsProviderFilterByFieldEmptyFilters{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsProviderFilterByField.Type {
	case components.MetricsProviderFilterByFieldTypeMetricsProviderFilterByFieldMultiselectFilters:
		// metricsProviderFilterByField.MetricsProviderFilterByFieldMultiselectFilters is populated
	case components.MetricsProviderFilterByFieldTypeMetricsProviderFilterByFieldEmptyFilters:
		// metricsProviderFilterByField.MetricsProviderFilterByFieldEmptyFilters is populated
}
```
