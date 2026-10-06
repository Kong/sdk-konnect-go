# MCPPositionalArgumentFormat

Specifies the input format.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPPositionalArgumentFormatString

// Open enum: custom values can be created with a direct type cast
custom := components.MCPPositionalArgumentFormat("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `MCPPositionalArgumentFormatString`   | string                                |
| `MCPPositionalArgumentFormatNumber`   | number                                |
| `MCPPositionalArgumentFormatBoolean`  | boolean                               |
| `MCPPositionalArgumentFormatFilepath` | filepath                              |