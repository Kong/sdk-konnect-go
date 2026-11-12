# ManagedCacheUsageFilters


## Supported Types

### MetricsControlPlaneFilterByField

```go
managedCacheUsageFilters := components.CreateManagedCacheUsageFiltersControlPlane(components.MetricsControlPlaneFilterByField{/* values here */})
```

### MetricsDataPlaneGroupFilterByField

```go
managedCacheUsageFilters := components.CreateManagedCacheUsageFiltersDataPlaneGroup(components.MetricsDataPlaneGroupFilterByField{/* values here */})
```

### MetricsManagedCacheFilterByField

```go
managedCacheUsageFilters := components.CreateManagedCacheUsageFiltersManagedCache(components.MetricsManagedCacheFilterByField{/* values here */})
```

### MetricsNetworkFilterByField

```go
managedCacheUsageFilters := components.CreateManagedCacheUsageFiltersNetwork(components.MetricsNetworkFilterByField{/* values here */})
```

### MetricsProviderFilterByField

```go
managedCacheUsageFilters := components.CreateManagedCacheUsageFiltersProvider(components.MetricsProviderFilterByField{/* values here */})
```

### MetricsProviderRegionFilterByField

```go
managedCacheUsageFilters := components.CreateManagedCacheUsageFiltersProviderRegion(components.MetricsProviderRegionFilterByField{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch managedCacheUsageFilters.Type {
	case components.ManagedCacheUsageFiltersTypeControlPlane:
		// managedCacheUsageFilters.MetricsControlPlaneFilterByField is populated
	case components.ManagedCacheUsageFiltersTypeDataPlaneGroup:
		// managedCacheUsageFilters.MetricsDataPlaneGroupFilterByField is populated
	case components.ManagedCacheUsageFiltersTypeManagedCache:
		// managedCacheUsageFilters.MetricsManagedCacheFilterByField is populated
	case components.ManagedCacheUsageFiltersTypeNetwork:
		// managedCacheUsageFilters.MetricsNetworkFilterByField is populated
	case components.ManagedCacheUsageFiltersTypeProvider:
		// managedCacheUsageFilters.MetricsProviderFilterByField is populated
	case components.ManagedCacheUsageFiltersTypeProviderRegion:
		// managedCacheUsageFilters.MetricsProviderRegionFilterByField is populated
}
```
