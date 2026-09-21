# BillingAppCatalogItemType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppCatalogItemTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppCatalogItemType("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `BillingAppCatalogItemTypeSandbox`           | sandbox                                      |
| `BillingAppCatalogItemTypeStripe`            | stripe                                       |
| `BillingAppCatalogItemTypeExternalInvoicing` | external_invoicing                           |