# Method

The prediction method:

- `auto`: `weekly_seasonal` once 14 completed training days exist, and `exponential_smoothing` before that.
- `weekly_seasonal`: the mean daily spend for each weekday. Needs 14 completed training days.
- `linear_regression`: a straight-line fit of daily spend, extended to the end of the month. Needs 2 completed training days.
- `exponential_smoothing`: a weighted average of recent daily spend, carried forward. Needs 2 completed training days.

A method named in the request never falls back to another when the history is too short. The response's `status` is then `insufficient_history`. Defaults to `auto`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.MethodAuto

// Open enum: custom values can be created with a direct type cast
custom := operations.Method("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `MethodAuto`                 | auto                         |
| `MethodWeeklySeasonal`       | weekly_seasonal              |
| `MethodLinearRegression`     | linear_regression            |
| `MethodExponentialSmoothing` | exponential_smoothing        |