# BillingAppExternalInvoicingStatus

Status of the app connection.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingAppExternalInvoicingStatusReady

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingAppExternalInvoicingStatus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `BillingAppExternalInvoicingStatusReady`        | ready                                           |
| `BillingAppExternalInvoicingStatusUnauthorized` | unauthorized                                    |