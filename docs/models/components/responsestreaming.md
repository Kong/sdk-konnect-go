# ResponseStreaming

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ResponseStreamingAllow

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseStreaming("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ResponseStreamingAllow`  | allow                     |
| `ResponseStreamingAlways` | always                    |
| `ResponseStreamingDeny`   | deny                      |