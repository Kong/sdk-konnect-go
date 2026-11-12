# APIPackageImageStatus

The status of the image.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIPackageImageStatusUploading

// Open enum: custom values can be created with a direct type cast
custom := components.APIPackageImageStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `APIPackageImageStatusUploading`  | uploading                         |
| `APIPackageImageStatusValidating` | validating                        |
| `APIPackageImageStatusValid`      | valid                             |
| `APIPackageImageStatusInvalid`    | invalid                           |