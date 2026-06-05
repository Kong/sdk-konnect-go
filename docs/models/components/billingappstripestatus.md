# BillingAppStripeStatus

Status of the app connection.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppStripeStatusReady

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppStripeStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `BillingAppStripeStatusReady`        | ready                                |
| `BillingAppStripeStatusUnauthorized` | unauthorized                         |