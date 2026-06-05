# VulnerabilitiesMetricsFilterByPackageName

Filters a metrics query by `package.name`.


## Supported Types

### VulnerabilitiesMetricsFilterByPackageNameEmptyValueMetricsFilter

```go
vulnerabilitiesMetricsFilterByPackageName := components.CreateVulnerabilitiesMetricsFilterByPackageNameVulnerabilitiesMetricsFilterByPackageNameEmptyValueMetricsFilter(components.VulnerabilitiesMetricsFilterByPackageNameEmptyValueMetricsFilter{/* values here */})
```

### VulnerabilitiesMetricsFilterByPackageName2

```go
vulnerabilitiesMetricsFilterByPackageName := components.CreateVulnerabilitiesMetricsFilterByPackageNameVulnerabilitiesMetricsFilterByPackageName2(components.VulnerabilitiesMetricsFilterByPackageName2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vulnerabilitiesMetricsFilterByPackageName.Type {
	case components.VulnerabilitiesMetricsFilterByPackageNameTypeVulnerabilitiesMetricsFilterByPackageNameEmptyValueMetricsFilter:
		// vulnerabilitiesMetricsFilterByPackageName.VulnerabilitiesMetricsFilterByPackageNameEmptyValueMetricsFilter is populated
	case components.VulnerabilitiesMetricsFilterByPackageNameTypeVulnerabilitiesMetricsFilterByPackageName2:
		// vulnerabilitiesMetricsFilterByPackageName.VulnerabilitiesMetricsFilterByPackageName2 is populated
}
```
