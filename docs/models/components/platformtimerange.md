# PlatformTimeRange

The time range to query for platform data.


## Supported Types

### PlatformRelativeTimeRange

```go
platformTimeRange := components.CreatePlatformTimeRangeRelative(components.PlatformRelativeTimeRange{/* values here */})
```

### MetricsAbsoluteTimeRangeDtoV2

```go
platformTimeRange := components.CreatePlatformTimeRangeAbsolute(components.MetricsAbsoluteTimeRangeDtoV2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformTimeRange.Type {
	case components.PlatformTimeRangeTypeRelative:
		// platformTimeRange.PlatformRelativeTimeRange is populated
	case components.PlatformTimeRangeTypeAbsolute:
		// platformTimeRange.MetricsAbsoluteTimeRangeDtoV2 is populated
}
```
