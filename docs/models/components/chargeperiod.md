# ChargePeriod

The length of each row's charge period: a calendar day or an hour in `time_zone`, which `ChargePeriodStart` and `ChargePeriodEnd` bound. The range selects the events and doesn't shorten the first or last period, so those periods can extend beyond `starts_at` and `ends_at`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ChargePeriodDay

// Open enum: custom values can be created with a direct type cast
custom := components.ChargePeriod("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `ChargePeriodDay`  | day                |
| `ChargePeriodHour` | hour               |