# CacheScope

Whether the result may be cached across authorization contexts. `public` is rejected when
the server's tool list is filtered per subject by `default_tool_acls` or a tool's own
`access.acls`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CacheScopePublic

// Open enum: custom values can be created with a direct type cast
custom := components.CacheScope("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `CacheScopePublic`  | public              |
| `CacheScopePrivate` | private             |