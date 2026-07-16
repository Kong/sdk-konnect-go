# Location

Specify whether the param name and value options go in a query string, or the POST form/JSON body.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LocationBody

// Open enum: custom values can be created with a direct type cast
custom := components.Location("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `LocationBody`  | body            |
| `LocationQuery` | query           |