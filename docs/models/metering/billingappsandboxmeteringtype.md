# BillingAppSandboxMeteringType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingAppSandboxMeteringTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingAppSandboxMeteringType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `BillingAppSandboxMeteringTypeSandbox`           | sandbox                                          |
| `BillingAppSandboxMeteringTypeStripe`            | stripe                                           |
| `BillingAppSandboxMeteringTypeExternalInvoicing` | external_invoicing                               |