# BillingChargeRealizationStatus

The settlement status of the payment.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeRealizationStatusAuthorized

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeRealizationStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BillingChargeRealizationStatusAuthorized` | authorized                                 |
| `BillingChargeRealizationStatusSettled`    | settled                                    |