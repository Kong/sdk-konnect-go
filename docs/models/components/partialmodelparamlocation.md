# PartialModelParamLocation

Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialModelParamLocationBody

// Open enum: custom values can be created with a direct type cast
custom := components.PartialModelParamLocation("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PartialModelParamLocationBody`  | body                             |
| `PartialModelParamLocationQuery` | query                            |