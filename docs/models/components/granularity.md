# Granularity

Force time grouping into buckets of the specified duration.  Only has an effect if "time" is in the "dimensions" list.

The granularity of the result may be coarser than requested.  The finest allowed granularity depends on the query's time range: data farther in the past may have coarser granularity.  The exact result granularity will be reported in the response `meta.granularity_ms` field.

If granularity is not specified and "time" is in the dimensions list, a default will be chosen based on the time range requested.

Different relative times support different granularities:
  - 15m => tenSecondly, thirtySecondly, minutely
  - 1h  => tenSecondly, thirtySecondly, minutely, fiveMinutely, tenMinutely
  - 6h  => thirtySecondly, minutely, fiveMinutely, tenMinutely, thirtyMinutely, hourly
  - 12h => minutely, fiveMinutely, tenMinutely, thirtyMinutely, hourly
  - 24h => fiveMinutely, tenMinutely, thirtyMinutely, hourly
  - 7d  => thirtyMinutely, hourly, twoHourly, twelveHourly, daily
  - 30d => hourly, twoHourly, twelveHourly, daily, weekly

For special time ranges:
  - current_week, previous_week   => thirtyMinutely, hourly, twoHourly, twelveHourly, daily
  - current_month, previous_month => hourly, twoHourly, twelveHourly, daily, weekly

For absolute time ranges, daily will be used.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.GranularityTenSecondly

// Open enum: custom values can be created with a direct type cast
custom := components.Granularity("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `GranularityTenSecondly`    | tenSecondly                 |
| `GranularityThirtySecondly` | thirtySecondly              |
| `GranularityMinutely`       | minutely                    |
| `GranularityFiveMinutely`   | fiveMinutely                |
| `GranularityTenMinutely`    | tenMinutely                 |
| `GranularityThirtyMinutely` | thirtyMinutely              |
| `GranularityHourly`         | hourly                      |
| `GranularityTwoHourly`      | twoHourly                   |
| `GranularityTwelveHourly`   | twelveHourly                |
| `GranularityDaily`          | daily                       |
| `GranularityWeekly`         | weekly                      |