# ForecastCostUsageEventsMethod

The method the service used. For a request with `auto`, it's the method that `auto` chose, so it's never `auto`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.ForecastCostUsageEventsMethodWeeklySeasonal

// Open enum: custom values can be created with a direct type cast
custom := operations.ForecastCostUsageEventsMethod("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `ForecastCostUsageEventsMethodWeeklySeasonal`       | weekly_seasonal                                     |
| `ForecastCostUsageEventsMethodLinearRegression`     | linear_regression                                   |
| `ForecastCostUsageEventsMethodExponentialSmoothing` | exponential_smoothing                               |