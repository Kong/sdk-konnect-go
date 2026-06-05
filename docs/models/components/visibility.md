# Visibility

Whether a menu item is public or private. Private menu items are only accessible to authenticated users.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VisibilityPublic

// Open enum: custom values can be created with a direct type cast
custom := components.Visibility("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `VisibilityPublic`  | public              |
| `VisibilityPrivate` | private             |