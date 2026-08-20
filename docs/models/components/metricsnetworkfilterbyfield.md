# MetricsNetworkFilterByField


## Supported Types

### MetricsNetworkFilterByFieldMultiselectFilters

```go
metricsNetworkFilterByField := components.CreateMetricsNetworkFilterByFieldMetricsNetworkFilterByFieldMultiselectFilters(components.MetricsNetworkFilterByFieldMultiselectFilters{/* values here */})
```

### MetricsNetworkFilterByFieldEmptyFilters

```go
metricsNetworkFilterByField := components.CreateMetricsNetworkFilterByFieldMetricsNetworkFilterByFieldEmptyFilters(components.MetricsNetworkFilterByFieldEmptyFilters{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsNetworkFilterByField.Type {
	case components.MetricsNetworkFilterByFieldTypeMetricsNetworkFilterByFieldMultiselectFilters:
		// metricsNetworkFilterByField.MetricsNetworkFilterByFieldMultiselectFilters is populated
	case components.MetricsNetworkFilterByFieldTypeMetricsNetworkFilterByFieldEmptyFilters:
		// metricsNetworkFilterByField.MetricsNetworkFilterByFieldEmptyFilters is populated
}
```
