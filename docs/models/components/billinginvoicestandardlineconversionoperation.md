# BillingInvoiceStandardLineConversionOperation

The arithmetic operation to apply to the raw metered quantity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingInvoiceStandardLineConversionOperationDivide

// Open enum: custom values can be created with a direct type cast
custom := components.BillingInvoiceStandardLineConversionOperation("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `BillingInvoiceStandardLineConversionOperationDivide`   | divide                                                  |
| `BillingInvoiceStandardLineConversionOperationMultiply` | multiply                                                |