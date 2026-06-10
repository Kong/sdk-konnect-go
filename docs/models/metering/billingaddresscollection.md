# BillingAddressCollection

Whether to collect the customer's billing address.

Defaults to auto, which only collects the address when necessary for tax
calculation.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingAddressCollectionAuto

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingAddressCollection("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `BillingAddressCollectionAuto`     | auto                               |
| `BillingAddressCollectionRequired` | required                           |