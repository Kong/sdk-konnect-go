# VisibilityStatus

Whether the resource is publicly accessible to non-authenticated users.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VisibilityStatusPublic

// Open enum: custom values can be created with a direct type cast
custom := components.VisibilityStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `VisibilityStatusPublic`  | public                    |
| `VisibilityStatusPrivate` | private                   |