# PlatformRelativeTimeRangeTimeRange

Platform usage data is not available at sub-day granularity. A sub-day duration
(`15m`, `1h`, `6h`, `12h`) is accepted but silently coerced up to `24h` rather
than rejected.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PlatformRelativeTimeRangeTimeRangeFifteenm

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformRelativeTimeRangeTimeRange("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `PlatformRelativeTimeRangeTimeRangeFifteenm`                  | 15m                                                           |
| `PlatformRelativeTimeRangeTimeRangeOneh`                      | 1h                                                            |
| `PlatformRelativeTimeRangeTimeRangeSixh`                      | 6h                                                            |
| `PlatformRelativeTimeRangeTimeRangeTwelveh`                   | 12h                                                           |
| `PlatformRelativeTimeRangeTimeRangeTwentyFourh`               | 24h                                                           |
| `PlatformRelativeTimeRangeTimeRangeSevend`                    | 7d                                                            |
| `PlatformRelativeTimeRangeTimeRangeThirtyd`                   | 30d                                                           |
| `PlatformRelativeTimeRangeTimeRangeNinetyd`                   | 90d                                                           |
| `PlatformRelativeTimeRangeTimeRangeOneHundredAndEightyd`      | 180d                                                          |
| `PlatformRelativeTimeRangeTimeRangeThreeHundredAndSixtyFived` | 365d                                                          |
| `PlatformRelativeTimeRangeTimeRangeCurrentWeek`               | current_week                                                  |
| `PlatformRelativeTimeRangeTimeRangeCurrentMonth`              | current_month                                                 |
| `PlatformRelativeTimeRangeTimeRangeCurrentQuarter`            | current_quarter                                               |
| `PlatformRelativeTimeRangeTimeRangePreviousWeek`              | previous_week                                                 |
| `PlatformRelativeTimeRangeTimeRangePreviousMonth`             | previous_month                                                |
| `PlatformRelativeTimeRangeTimeRangePreviousQuarter`           | previous_quarter                                              |