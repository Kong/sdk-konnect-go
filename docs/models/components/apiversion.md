# APIVersion

Cohere API version. `v1` uses the legacy `/v1/chat` endpoint; `v2` (default)
uses `/v2/chat` and supports tool calling.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIVersionV1

// Open enum: custom values can be created with a direct type cast
custom := components.APIVersion("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `APIVersionV1` | v1             |
| `APIVersionV2` | v2             |