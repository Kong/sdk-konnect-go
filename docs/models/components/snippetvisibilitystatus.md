# SnippetVisibilityStatus

Whether a snippet is publicly accessible to non-authenticated users.
If not provided, the default_page_visibility value of the portal will be used.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SnippetVisibilityStatusPublic

// Open enum: custom values can be created with a direct type cast
custom := components.SnippetVisibilityStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `SnippetVisibilityStatusPublic`  | public                           |
| `SnippetVisibilityStatusPrivate` | private                          |