# Location

Whether the credential travels in a header, the request body, or the query string. Defaults to `header`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LocationHeader

// Open enum: custom values can be created with a direct type cast
custom := components.Location("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `LocationHeader` | header           |
| `LocationBody`   | body             |
| `LocationQuery`  | query            |