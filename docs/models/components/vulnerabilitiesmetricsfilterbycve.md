# VulnerabilitiesMetricsFilterByCVE

Filters a metrics query by `cve_id`.


## Supported Types

### VulnerabilitiesMetricsFilterByCVEEmptyValueMetricsFilter

```go
vulnerabilitiesMetricsFilterByCVE := components.CreateVulnerabilitiesMetricsFilterByCVEVulnerabilitiesMetricsFilterByCVEEmptyValueMetricsFilter(components.VulnerabilitiesMetricsFilterByCVEEmptyValueMetricsFilter{/* values here */})
```

### VulnerabilitiesMetricsFilterByCVE2

```go
vulnerabilitiesMetricsFilterByCVE := components.CreateVulnerabilitiesMetricsFilterByCVEVulnerabilitiesMetricsFilterByCVE2(components.VulnerabilitiesMetricsFilterByCVE2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vulnerabilitiesMetricsFilterByCVE.Type {
	case components.VulnerabilitiesMetricsFilterByCVETypeVulnerabilitiesMetricsFilterByCVEEmptyValueMetricsFilter:
		// vulnerabilitiesMetricsFilterByCVE.VulnerabilitiesMetricsFilterByCVEEmptyValueMetricsFilter is populated
	case components.VulnerabilitiesMetricsFilterByCVETypeVulnerabilitiesMetricsFilterByCVE2:
		// vulnerabilitiesMetricsFilterByCVE.VulnerabilitiesMetricsFilterByCVE2 is populated
}
```
