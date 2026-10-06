# MCPNamedArgumentFormat

Specifies the input format.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPNamedArgumentFormatString

// Open enum: custom values can be created with a direct type cast
custom := components.MCPNamedArgumentFormat("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `MCPNamedArgumentFormatString`   | string                           |
| `MCPNamedArgumentFormatNumber`   | number                           |
| `MCPNamedArgumentFormatBoolean`  | boolean                          |
| `MCPNamedArgumentFormatFilepath` | filepath                         |