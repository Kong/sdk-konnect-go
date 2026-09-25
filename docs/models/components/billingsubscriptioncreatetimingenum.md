# BillingSubscriptionCreateTimingEnum

Subscription create timing. Only immediate is supported as an enum value —
unlike edit timing, next_billing_cycle is not accepted on create, since a new
subscription has no current billing cycle to schedule against.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionCreateTimingEnumImmediate
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `BillingSubscriptionCreateTimingEnumImmediate` | immediate                                      |