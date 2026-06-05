# BillingCreditTransactionType

The type of credit transaction.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingCreditTransactionTypeFunded

// Open enum: custom values can be created with a direct type cast
custom := components.BillingCreditTransactionType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `BillingCreditTransactionTypeFunded`   | funded                                 |
| `BillingCreditTransactionTypeConsumed` | consumed                               |
| `BillingCreditTransactionTypeExpired`  | expired                                |