# CreditGrantVoidPaymentAdjustment

How voiding adjusts payment state related to the grant.

Currently only `none` is supported: voiding does not adjust invoices, payment
authorization, settlement, payment intents, or external collection state. If
payment later completes, the original invoiced amount may still be collected.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreditGrantVoidPaymentAdjustmentNone
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `CreditGrantVoidPaymentAdjustmentNone` | none                                   |