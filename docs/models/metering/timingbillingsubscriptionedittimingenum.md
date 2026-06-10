# TimingBillingSubscriptionEditTimingEnum

Subscription edit timing. When immediate, the requested changes take effect
immediately. When next_billing_cycle, the requested changes take effect at the
next billing cycle.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.TimingBillingSubscriptionEditTimingEnumImmediate

// Open enum: custom values can be created with a direct type cast
custom := metering.TimingBillingSubscriptionEditTimingEnum("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `TimingBillingSubscriptionEditTimingEnumImmediate`        | immediate                                                 |
| `TimingBillingSubscriptionEditTimingEnumNextBillingCycle` | next_billing_cycle                                        |