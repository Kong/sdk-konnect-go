# PlatformRelativeTimeRangeTimeRange

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PlatformRelativeTimeRangeTimeRangeTwentyFourh

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformRelativeTimeRangeTimeRange("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
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