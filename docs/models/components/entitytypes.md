# EntityTypes

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EntityTypesBillingInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.EntityTypes("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `EntityTypesBillingInvoice` | billing-invoice             |
| `EntityTypesAccessToken`    | access-token                |
| `EntityTypesWebhook`        | webhook                     |
| `EntityTypesDevPortal`      | dev-portal                  |
| `EntityTypesDataplaneGroup` | dataplane-group             |
| `EntityTypesDataplane`      | dataplane                   |
| `EntityTypesKaiEnablement`  | kai-enablement              |