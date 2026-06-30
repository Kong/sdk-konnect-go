# PlatformGranularity

Force time grouping into buckets of the specified duration. Only has an effect if "time" is in the "dimensions" list.

If granularity is not specified and "time" is in the dimensions list, a default will be chosen based on the time range requested.

Different relative times support different granularities:


  - 24h                                          => daily
  - 7d, current_week, previous_week              => daily, weekly
  - 30d, current_month, previous_month           => daily, weekly, monthly
  - 90d, 180d, 365d, current_quarter,
    previous_quarter                             => daily, weekly, monthly

For absolute time ranges, daily will be used.

Granularity values:


  - `daily`: Groups data into 24-hour buckets.
  - `weekly`: Groups data into 7-day buckets.
  - `monthly`: Groups data into calendar month buckets.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PlatformGranularityDaily

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformGranularity("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `PlatformGranularityDaily`   | daily                        |
| `PlatformGranularityWeekly`  | weekly                       |
| `PlatformGranularityMonthly` | monthly                      |