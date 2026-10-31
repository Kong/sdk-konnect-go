# SubscriptionEndProrationMode

Controls how subscription-ending shortened service periods are billed.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SubscriptionEndProrationModeBillFullPeriod

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionEndProrationMode("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `SubscriptionEndProrationModeBillFullPeriod`   | bill_full_period                               |
| `SubscriptionEndProrationModeBillActualPeriod` | bill_actual_period                             |