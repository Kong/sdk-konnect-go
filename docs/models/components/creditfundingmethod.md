# CreditFundingMethod

Funding method of the grant.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreditFundingMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.CreditFundingMethod("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CreditFundingMethodNone`     | none                          |
| `CreditFundingMethodInvoice`  | invoice                       |
| `CreditFundingMethodExternal` | external                      |