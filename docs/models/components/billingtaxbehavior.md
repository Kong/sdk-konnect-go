# BillingTaxBehavior

Tax behavior.

This enum is used to specify whether tax is included in the price or excluded
from the price.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingTaxBehaviorInclusive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingTaxBehavior("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `BillingTaxBehaviorInclusive` | inclusive                     |
| `BillingTaxBehaviorExclusive` | exclusive                     |