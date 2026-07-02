# VulnerabilitiesMetricsQueryTimeRange

The time range to query.


## Supported Types

### VulnerabilityMetricsRelativeTimeRange

```go
vulnerabilitiesMetricsQueryTimeRange := components.CreateVulnerabilitiesMetricsQueryTimeRangeRelative(components.VulnerabilityMetricsRelativeTimeRange{/* values here */})
```

### VulnerabilityMetricsAbsoluteTimeRange

```go
vulnerabilitiesMetricsQueryTimeRange := components.CreateVulnerabilitiesMetricsQueryTimeRangeAbsolute(components.VulnerabilityMetricsAbsoluteTimeRange{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vulnerabilitiesMetricsQueryTimeRange.Type {
	case components.VulnerabilitiesMetricsQueryTimeRangeTypeRelative:
		// vulnerabilitiesMetricsQueryTimeRange.VulnerabilityMetricsRelativeTimeRange is populated
	case components.VulnerabilitiesMetricsQueryTimeRangeTypeAbsolute:
		// vulnerabilitiesMetricsQueryTimeRange.VulnerabilityMetricsAbsoluteTimeRange is populated
}
```
