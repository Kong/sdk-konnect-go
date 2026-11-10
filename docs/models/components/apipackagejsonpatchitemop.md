# APIPackageJSONPatchItemOp

The operation to perform.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIPackageJSONPatchItemOpAdd

// Open enum: custom values can be created with a direct type cast
custom := components.APIPackageJSONPatchItemOp("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `APIPackageJSONPatchItemOpAdd`     | add                                |
| `APIPackageJSONPatchItemOpRemove`  | remove                             |
| `APIPackageJSONPatchItemOpReplace` | replace                            |