# PortalMenuItemVisibility

Whether a menu item is public or private. Private menu items are only accessible to authenticated users.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PortalMenuItemVisibilityPublic

// Open enum: custom values can be created with a direct type cast
custom := components.PortalMenuItemVisibility("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PortalMenuItemVisibilityPublic`  | public                            |
| `PortalMenuItemVisibilityPrivate` | private                           |