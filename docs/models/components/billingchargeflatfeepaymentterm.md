# BillingChargeFlatFeePaymentTerm

Payment term of the flat fee charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeFlatFeePaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeFlatFeePaymentTerm("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BillingChargeFlatFeePaymentTermInAdvance` | in_advance                                 |
| `BillingChargeFlatFeePaymentTermInArrears` | in_arrears                                 |