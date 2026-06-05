# VulnerabilitiesMetricsFilterByRegion

Filters a metrics query by `region`.


## Supported Types

### VulnerabilitiesMetricsFilterByRegionEmptyValueMetricsFilter

```go
vulnerabilitiesMetricsFilterByRegion := components.CreateVulnerabilitiesMetricsFilterByRegionVulnerabilitiesMetricsFilterByRegionEmptyValueMetricsFilter(components.VulnerabilitiesMetricsFilterByRegionEmptyValueMetricsFilter{/* values here */})
```

### VulnerabilitiesMetricsFilterByRegion2

```go
vulnerabilitiesMetricsFilterByRegion := components.CreateVulnerabilitiesMetricsFilterByRegionVulnerabilitiesMetricsFilterByRegion2(components.VulnerabilitiesMetricsFilterByRegion2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vulnerabilitiesMetricsFilterByRegion.Type {
	case components.VulnerabilitiesMetricsFilterByRegionTypeVulnerabilitiesMetricsFilterByRegionEmptyValueMetricsFilter:
		// vulnerabilitiesMetricsFilterByRegion.VulnerabilitiesMetricsFilterByRegionEmptyValueMetricsFilter is populated
	case components.VulnerabilitiesMetricsFilterByRegionTypeVulnerabilitiesMetricsFilterByRegion2:
		// vulnerabilitiesMetricsFilterByRegion.VulnerabilitiesMetricsFilterByRegion2 is populated
}
```
