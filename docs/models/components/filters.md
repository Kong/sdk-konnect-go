# Filters


## Supported Types

### VulnerabilitiesMetricsFilterByService

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByService(components.VulnerabilitiesMetricsFilterByService{/* values here */})
```

### VulnerabilitiesMetricsFilterBySeverity

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterBySeverity(components.VulnerabilitiesMetricsFilterBySeverity{/* values here */})
```

### VulnerabilitiesMetricsFilterByState

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByState(components.VulnerabilitiesMetricsFilterByState{/* values here */})
```

### VulnerabilitiesMetricsFilterByType

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByType(components.VulnerabilitiesMetricsFilterByType{/* values here */})
```

### VulnerabilitiesMetricsFilterBySource

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterBySource(components.VulnerabilitiesMetricsFilterBySource{/* values here */})
```

### VulnerabilitiesMetricsFilterByEnvironment

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByEnvironment(components.VulnerabilitiesMetricsFilterByEnvironment{/* values here */})
```

### VulnerabilitiesMetricsFilterByRegion

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByRegion(components.VulnerabilitiesMetricsFilterByRegion{/* values here */})
```

### VulnerabilitiesMetricsFilterByCVE

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByCVE(components.VulnerabilitiesMetricsFilterByCVE{/* values here */})
```

### VulnerabilitiesMetricsFilterByPackageName

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByPackageName(components.VulnerabilitiesMetricsFilterByPackageName{/* values here */})
```

### VulnerabilitiesMetricsFilterByScanAttributes

```go
filters := components.CreateFiltersVulnerabilitiesMetricsFilterByScanAttributes(components.VulnerabilitiesMetricsFilterByScanAttributes{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch filters.Type {
	case components.FiltersTypeVulnerabilitiesMetricsFilterByService:
		// filters.VulnerabilitiesMetricsFilterByService is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterBySeverity:
		// filters.VulnerabilitiesMetricsFilterBySeverity is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterByState:
		// filters.VulnerabilitiesMetricsFilterByState is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterByType:
		// filters.VulnerabilitiesMetricsFilterByType is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterBySource:
		// filters.VulnerabilitiesMetricsFilterBySource is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterByEnvironment:
		// filters.VulnerabilitiesMetricsFilterByEnvironment is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterByRegion:
		// filters.VulnerabilitiesMetricsFilterByRegion is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterByCVE:
		// filters.VulnerabilitiesMetricsFilterByCVE is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterByPackageName:
		// filters.VulnerabilitiesMetricsFilterByPackageName is populated
	case components.FiltersTypeVulnerabilitiesMetricsFilterByScanAttributes:
		// filters.VulnerabilitiesMetricsFilterByScanAttributes is populated
}
```
