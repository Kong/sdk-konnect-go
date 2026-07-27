# ListPortalsResponseDefaultAPIVisibility

The default visibility of APIs in the portal. If set to `public`, newly published APIs are visible to unauthenticated developers. If set to `private`, newly published APIs are hidden from unauthenticated developers.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ListPortalsResponseDefaultAPIVisibilityPublic

// Open enum: custom values can be created with a direct type cast
custom := components.ListPortalsResponseDefaultAPIVisibility("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `ListPortalsResponseDefaultAPIVisibilityPublic`  | public                                           |
| `ListPortalsResponseDefaultAPIVisibilityPrivate` | private                                          |