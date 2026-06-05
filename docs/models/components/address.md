# Address

Whether to save the billing address to customer.address.

Defaults to "never".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AddressAuto

// Open enum: custom values can be created with a direct type cast
custom := components.Address("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `AddressAuto`  | auto           |
| `AddressNever` | never          |