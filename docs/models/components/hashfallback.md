# HashFallback

What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.HashFallbackConsumer

// Open enum: custom values can be created with a direct type cast
custom := components.HashFallback("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `HashFallbackConsumer`   | consumer                 |
| `HashFallbackCookie`     | cookie                   |
| `HashFallbackHeader`     | header                   |
| `HashFallbackIP`         | ip                       |
| `HashFallbackNone`       | none                     |
| `HashFallbackPath`       | path                     |
| `HashFallbackQueryArg`   | query_arg                |
| `HashFallbackURICapture` | uri_capture              |