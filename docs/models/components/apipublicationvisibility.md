# APIPublicationVisibility

The visibility of the API in the portal.
Public API publications do not require authentication to view and retrieve information about them.
Private API publications require authentication to retrieve information about them.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIPublicationVisibilityPublic

// Open enum: custom values can be created with a direct type cast
custom := components.APIPublicationVisibility("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `APIPublicationVisibilityPublic`  | public                            |
| `APIPublicationVisibilityPrivate` | private                           |