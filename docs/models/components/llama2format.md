# Llama2Format

If using llama2 provider, select the upstream message format.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.Llama2FormatOllama

// Open enum: custom values can be created with a direct type cast
custom := components.Llama2Format("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `Llama2FormatOllama` | ollama               |
| `Llama2FormatOpenai` | openai               |
| `Llama2FormatRaw`    | raw                  |