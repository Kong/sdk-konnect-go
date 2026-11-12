# CreateChargeFlatFeeRequestPaymentTerm

Payment term of the flat fee charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateChargeFlatFeeRequestPaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := components.CreateChargeFlatFeeRequestPaymentTerm("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateChargeFlatFeeRequestPaymentTermInAdvance` | in_advance                                       |
| `CreateChargeFlatFeeRequestPaymentTermInArrears` | in_arrears                                       |