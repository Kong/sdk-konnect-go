# NavigationLinkVisibility

Who sees the link. `public` for every visitor, `private` for signed-in developers only.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NavigationLinkVisibilityPublic

// Open enum: custom values can be created with a direct type cast
custom := components.NavigationLinkVisibility("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `NavigationLinkVisibilityPublic`  | public                            |
| `NavigationLinkVisibilityPrivate` | private                           |