# Address

Whether to save the billing address to customer.address.

Defaults to "never".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.AddressAuto

// Open enum: custom values can be created with a direct type cast
custom := metering.Address("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `AddressAuto`  | auto           |
| `AddressNever` | never          |