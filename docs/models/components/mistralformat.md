# MistralFormat

If using mistral provider, select the upstream message format.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MistralFormatOllama

// Open enum: custom values can be created with a direct type cast
custom := components.MistralFormat("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `MistralFormatOllama` | ollama                |
| `MistralFormatOpenai` | openai                |