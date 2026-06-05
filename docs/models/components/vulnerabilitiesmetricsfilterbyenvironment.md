# VulnerabilitiesMetricsFilterByEnvironment

Filters a metrics query by `environment`.


## Supported Types

### EmptyValueMetricsFilter

```go
vulnerabilitiesMetricsFilterByEnvironment := components.CreateVulnerabilitiesMetricsFilterByEnvironmentEmptyValueMetricsFilter(components.EmptyValueMetricsFilter{/* values here */})
```

### VulnerabilitiesMetricsFilterByEnvironment2

```go
vulnerabilitiesMetricsFilterByEnvironment := components.CreateVulnerabilitiesMetricsFilterByEnvironmentVulnerabilitiesMetricsFilterByEnvironment2(components.VulnerabilitiesMetricsFilterByEnvironment2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vulnerabilitiesMetricsFilterByEnvironment.Type {
	case components.VulnerabilitiesMetricsFilterByEnvironmentTypeEmptyValueMetricsFilter:
		// vulnerabilitiesMetricsFilterByEnvironment.EmptyValueMetricsFilter is populated
	case components.VulnerabilitiesMetricsFilterByEnvironmentTypeVulnerabilitiesMetricsFilterByEnvironment2:
		// vulnerabilitiesMetricsFilterByEnvironment.VulnerabilitiesMetricsFilterByEnvironment2 is populated
}
```
