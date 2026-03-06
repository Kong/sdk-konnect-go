# Required

Whether tax ID collection is required.

Defaults to "never".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RequiredIfSupported

// Open enum: custom values can be created with a direct type cast
custom := components.Required("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `RequiredIfSupported` | if_supported          |
| `RequiredNever`       | never                 |