# BillingSubscriptionChangeResponseMeteringStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingSubscriptionChangeResponseMeteringStatusActive

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingSubscriptionChangeResponseMeteringStatus("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `BillingSubscriptionChangeResponseMeteringStatusActive`    | active                                                     |
| `BillingSubscriptionChangeResponseMeteringStatusInactive`  | inactive                                                   |
| `BillingSubscriptionChangeResponseMeteringStatusCanceled`  | canceled                                                   |
| `BillingSubscriptionChangeResponseMeteringStatusScheduled` | scheduled                                                  |