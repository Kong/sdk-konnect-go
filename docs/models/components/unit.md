# Unit

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UnitMilliseconds

// Open enum: custom values can be created with a direct type cast
custom := components.Unit("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `UnitMilliseconds` | milliseconds       |
| `UnitSeconds`      | seconds            |
| `UnitMinutes`      | minutes            |
| `UnitHours`        | hours              |
| `UnitDays`         | days               |
| `UnitMonths`       | months             |
| `UnitYears`        | years              |