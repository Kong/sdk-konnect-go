# BillingFlatFeeChargePaymentTerm

Payment term of the flat fee charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingFlatFeeChargePaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := components.BillingFlatFeeChargePaymentTerm("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BillingFlatFeeChargePaymentTermInAdvance` | in_advance                                 |
| `BillingFlatFeeChargePaymentTermInArrears` | in_arrears                                 |