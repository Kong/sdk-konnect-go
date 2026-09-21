# BillingAppType

The type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `BillingAppTypeSandbox`           | sandbox                           |
| `BillingAppTypeStripe`            | stripe                            |
| `BillingAppTypeExternalInvoicing` | external_invoicing                |