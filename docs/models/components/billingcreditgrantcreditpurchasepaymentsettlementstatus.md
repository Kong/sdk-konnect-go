# BillingCreditGrantCreditPurchasePaymentSettlementStatus

Current payment settlement status.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingCreditGrantCreditPurchasePaymentSettlementStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.BillingCreditGrantCreditPurchasePaymentSettlementStatus("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `BillingCreditGrantCreditPurchasePaymentSettlementStatusPending`    | pending                                                             |
| `BillingCreditGrantCreditPurchasePaymentSettlementStatusAuthorized` | authorized                                                          |
| `BillingCreditGrantCreditPurchasePaymentSettlementStatusSettled`    | settled                                                             |