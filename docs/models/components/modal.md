# Modal

The modal type this price applies to.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ModalText

// Open enum: custom values can be created with a direct type cast
custom := components.Modal("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `ModalText`  | text         |
| `ModalAudio` | audio        |
| `ModalImage` | image        |
| `ModalVideo` | video        |