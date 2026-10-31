# APIImageStatus

The status of the image.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIImageStatusUploading

// Open enum: custom values can be created with a direct type cast
custom := components.APIImageStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `APIImageStatusUploading`  | uploading                  |
| `APIImageStatusValidating` | validating                 |
| `APIImageStatusValid`      | valid                      |
| `APIImageStatusInvalid`    | invalid                    |