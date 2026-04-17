# BillingAppSandboxDefinitionType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppSandboxDefinitionTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppSandboxDefinitionType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `BillingAppSandboxDefinitionTypeSandbox`           | sandbox                                            |
| `BillingAppSandboxDefinitionTypeStripe`            | stripe                                             |
| `BillingAppSandboxDefinitionTypeExternalInvoicing` | external_invoicing                                 |