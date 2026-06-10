# Required

Whether tax ID collection is required.

Defaults to "never".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.RequiredIfSupported

// Open enum: custom values can be created with a direct type cast
custom := metering.Required("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `RequiredIfSupported` | if_supported          |
| `RequiredNever`       | never                 |