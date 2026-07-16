# CreditPurchasePaymentSettlementStatus

The new payment settlement status.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreditPurchasePaymentSettlementStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.CreditPurchasePaymentSettlementStatus("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CreditPurchasePaymentSettlementStatusPending`    | pending                                           |
| `CreditPurchasePaymentSettlementStatusAuthorized` | authorized                                        |
| `CreditPurchasePaymentSettlementStatusSettled`    | settled                                           |