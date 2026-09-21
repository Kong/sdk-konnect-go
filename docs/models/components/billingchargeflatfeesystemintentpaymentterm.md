# BillingChargeFlatFeeSystemIntentPaymentTerm

Payment term of the flat fee charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeFlatFeeSystemIntentPaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeFlatFeeSystemIntentPaymentTerm("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `BillingChargeFlatFeeSystemIntentPaymentTermInAdvance` | in_advance                                             |
| `BillingChargeFlatFeeSystemIntentPaymentTermInArrears` | in_arrears                                             |