# APIDocumentStatus

If `status=published` the document will be visible in your live portal

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIDocumentStatusPublished

// Open enum: custom values can be created with a direct type cast
custom := components.APIDocumentStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `APIDocumentStatusPublished`   | published                      |
| `APIDocumentStatusUnpublished` | unpublished                    |