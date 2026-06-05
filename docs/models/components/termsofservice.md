# TermsOfService

Requires customers to accept terms of service before payment.

Requires a valid terms of service URL in your Stripe Dashboard settings.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TermsOfServiceNone

// Open enum: custom values can be created with a direct type cast
custom := components.TermsOfService("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `TermsOfServiceNone`     | none                     |
| `TermsOfServiceRequired` | required                 |