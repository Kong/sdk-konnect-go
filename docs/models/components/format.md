# Format

The request format to use when communicating with the Llama2 model.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.FormatOllama

// Open enum: custom values can be created with a direct type cast
custom := components.Format("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `FormatOllama` | ollama         |
| `FormatOpenai` | openai         |
| `FormatRaw`    | raw            |