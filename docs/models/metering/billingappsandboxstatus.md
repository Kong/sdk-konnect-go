# BillingAppSandboxStatus

Status of the app connection.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingAppSandboxStatusReady

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingAppSandboxStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `BillingAppSandboxStatusReady`        | ready                                 |
| `BillingAppSandboxStatusUnauthorized` | unauthorized                          |