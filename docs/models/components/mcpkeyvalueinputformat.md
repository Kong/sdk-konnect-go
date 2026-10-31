# MCPKeyValueInputFormat

Specifies the input format.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPKeyValueInputFormatString

// Open enum: custom values can be created with a direct type cast
custom := components.MCPKeyValueInputFormat("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `MCPKeyValueInputFormatString`   | string                           |
| `MCPKeyValueInputFormatNumber`   | number                           |
| `MCPKeyValueInputFormatBoolean`  | boolean                          |
| `MCPKeyValueInputFormatFilepath` | filepath                         |