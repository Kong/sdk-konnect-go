# UpdatePortalDefaultPageVisibility

The default visibility of pages in the portal. If set to `public`, newly created pages are visible to unauthenticated developers. If set to `private`, newly created pages are hidden from unauthenticated developers.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpdatePortalDefaultPageVisibilityPublic

// Open enum: custom values can be created with a direct type cast
custom := components.UpdatePortalDefaultPageVisibility("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `UpdatePortalDefaultPageVisibilityPublic`  | public                                     |
| `UpdatePortalDefaultPageVisibilityPrivate` | private                                    |