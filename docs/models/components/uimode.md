# UIMode

The UI mode for the checkout session.

"hosted" displays a Stripe-hosted page. "embedded" integrates directly into your app.
Defaults to "hosted".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UIModeEmbedded

// Open enum: custom values can be created with a direct type cast
custom := components.UIMode("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `UIModeEmbedded` | embedded         |
| `UIModeHosted`   | hosted           |