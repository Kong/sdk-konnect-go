# CreateBillingProfileRequestSubscriptionEndProrationMode

Controls how subscription-ending shortened service periods are billed.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateBillingProfileRequestSubscriptionEndProrationModeBillFullPeriod

// Open enum: custom values can be created with a direct type cast
custom := components.CreateBillingProfileRequestSubscriptionEndProrationMode("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `CreateBillingProfileRequestSubscriptionEndProrationModeBillFullPeriod`   | bill_full_period                                                          |
| `CreateBillingProfileRequestSubscriptionEndProrationModeBillActualPeriod` | bill_actual_period                                                        |