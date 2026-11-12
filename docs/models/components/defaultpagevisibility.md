# DefaultPageVisibility

The default visibility of pages in the portal. If set to `public`, newly created pages are visible to unauthenticated developers. If set to `private`, newly created pages are hidden from unauthenticated developers.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DefaultPageVisibilityPublic

// Open enum: custom values can be created with a direct type cast
custom := components.DefaultPageVisibility("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `DefaultPageVisibilityPublic`  | public                         |
| `DefaultPageVisibilityPrivate` | private                        |