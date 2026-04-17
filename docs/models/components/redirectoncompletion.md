# RedirectOnCompletion

Redirect behavior for embedded checkout sessions.

Controls when to redirect users after completion.
See: https://docs.stripe.com/payments/checkout/custom-success-page?payment-ui=embedded-form

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RedirectOnCompletionAlways

// Open enum: custom values can be created with a direct type cast
custom := components.RedirectOnCompletion("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `RedirectOnCompletionAlways`     | always                           |
| `RedirectOnCompletionIfRequired` | if_required                      |
| `RedirectOnCompletionNever`      | never                            |