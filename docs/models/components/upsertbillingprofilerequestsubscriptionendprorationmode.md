# UpsertBillingProfileRequestSubscriptionEndProrationMode

Controls how subscription-ending shortened service periods are billed.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpsertBillingProfileRequestSubscriptionEndProrationModeBillFullPeriod

// Open enum: custom values can be created with a direct type cast
custom := components.UpsertBillingProfileRequestSubscriptionEndProrationMode("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `UpsertBillingProfileRequestSubscriptionEndProrationModeBillFullPeriod`   | bill_full_period                                                          |
| `UpsertBillingProfileRequestSubscriptionEndProrationModeBillActualPeriod` | bill_actual_period                                                        |