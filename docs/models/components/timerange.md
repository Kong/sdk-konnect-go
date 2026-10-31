# TimeRange

The time range to query.


## Supported Types

### MetricsRelativeTimeRangeDtoV2

```go
timeRange := components.CreateTimeRangeRelative(components.MetricsRelativeTimeRangeDtoV2{/* values here */})
```

### MetricsAbsoluteTimeRangeDtoV2

```go
timeRange := components.CreateTimeRangeAbsolute(components.MetricsAbsoluteTimeRangeDtoV2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch timeRange.Type {
	case components.TimeRangeTypeRelative:
		// timeRange.MetricsRelativeTimeRangeDtoV2 is populated
	case components.TimeRangeTypeAbsolute:
		// timeRange.MetricsAbsoluteTimeRangeDtoV2 is populated
}
```
