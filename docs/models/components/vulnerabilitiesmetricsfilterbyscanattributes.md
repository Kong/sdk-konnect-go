# VulnerabilitiesMetricsFilterByScanAttributes

Filters a metrics query by `scan_attributes` properties.
Dot-notation off of `scan_attributes` is used to specify the property to filter by.

Example: `scan_attributes.image_tag`.



## Supported Types

### VulnerabilitiesMetricsFilterByScanAttributesEmptyValueMetricsFilter

```go
vulnerabilitiesMetricsFilterByScanAttributes := components.CreateVulnerabilitiesMetricsFilterByScanAttributesVulnerabilitiesMetricsFilterByScanAttributesEmptyValueMetricsFilter(components.VulnerabilitiesMetricsFilterByScanAttributesEmptyValueMetricsFilter{/* values here */})
```

### StringValueMetricsFilter

```go
vulnerabilitiesMetricsFilterByScanAttributes := components.CreateVulnerabilitiesMetricsFilterByScanAttributesStringValueMetricsFilter(components.StringValueMetricsFilter{/* values here */})
```

### NumericValueMetricsFilter

```go
vulnerabilitiesMetricsFilterByScanAttributes := components.CreateVulnerabilitiesMetricsFilterByScanAttributesNumericValueMetricsFilter(components.NumericValueMetricsFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vulnerabilitiesMetricsFilterByScanAttributes.Type {
	case components.VulnerabilitiesMetricsFilterByScanAttributesTypeVulnerabilitiesMetricsFilterByScanAttributesEmptyValueMetricsFilter:
		// vulnerabilitiesMetricsFilterByScanAttributes.VulnerabilitiesMetricsFilterByScanAttributesEmptyValueMetricsFilter is populated
	case components.VulnerabilitiesMetricsFilterByScanAttributesTypeStringValueMetricsFilter:
		// vulnerabilitiesMetricsFilterByScanAttributes.StringValueMetricsFilter is populated
	case components.VulnerabilitiesMetricsFilterByScanAttributesTypeNumericValueMetricsFilter:
		// vulnerabilitiesMetricsFilterByScanAttributes.NumericValueMetricsFilter is populated
}
```
