# MCPInputFormat

Specifies the input format.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPInputFormatString

// Open enum: custom values can be created with a direct type cast
custom := components.MCPInputFormat("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `MCPInputFormatString`   | string                   |
| `MCPInputFormatNumber`   | number                   |
| `MCPInputFormatBoolean`  | boolean                  |
| `MCPInputFormatFilepath` | filepath                 |