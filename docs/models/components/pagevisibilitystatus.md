# PageVisibilityStatus

Whether a page is publicly accessible to non-authenticated users.
If not provided, the default_page_visibility value of the portal will be used.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PageVisibilityStatusPublic

// Open enum: custom values can be created with a direct type cast
custom := components.PageVisibilityStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PageVisibilityStatusPublic`  | public                        |
| `PageVisibilityStatusPrivate` | private                       |