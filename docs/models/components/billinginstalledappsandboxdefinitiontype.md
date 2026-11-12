# BillingInstalledAppSandboxDefinitionType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingInstalledAppSandboxDefinitionTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingInstalledAppSandboxDefinitionType("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `BillingInstalledAppSandboxDefinitionTypeSandbox`           | sandbox                                                     |
| `BillingInstalledAppSandboxDefinitionTypeStripe`            | stripe                                                      |
| `BillingInstalledAppSandboxDefinitionTypeExternalInvoicing` | external_invoicing                                          |