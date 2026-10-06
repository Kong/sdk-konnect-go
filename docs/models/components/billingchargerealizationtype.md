# BillingChargeRealizationType

The type of the realization run.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeRealizationTypeFinalRealization

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeRealizationType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `BillingChargeRealizationTypeFinalRealization` | final_realization                              |
| `BillingChargeRealizationTypePartialInvoice`   | partial_invoice                                |
| `BillingChargeRealizationTypeOutstanding`      | outstanding                                    |
| `BillingChargeRealizationTypeVoided`           | voided                                         |