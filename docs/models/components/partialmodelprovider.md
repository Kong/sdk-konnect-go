# PartialModelProvider

AI provider request format - Kong translates requests to and from the specified backend compatible formats.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialModelProviderAnthropic

// Open enum: custom values can be created with a direct type cast
custom := components.PartialModelProvider("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PartialModelProviderAnthropic`   | anthropic                         |
| `PartialModelProviderAzure`       | azure                             |
| `PartialModelProviderBedrock`     | bedrock                           |
| `PartialModelProviderCerebras`    | cerebras                          |
| `PartialModelProviderCohere`      | cohere                            |
| `PartialModelProviderDashscope`   | dashscope                         |
| `PartialModelProviderDatabricks`  | databricks                        |
| `PartialModelProviderDeepseek`    | deepseek                          |
| `PartialModelProviderGemini`      | gemini                            |
| `PartialModelProviderHuggingface` | huggingface                       |
| `PartialModelProviderLlama2`      | llama2                            |
| `PartialModelProviderMistral`     | mistral                           |
| `PartialModelProviderOllama`      | ollama                            |
| `PartialModelProviderOpenai`      | openai                            |
| `PartialModelProviderVllm`        | vllm                              |
| `PartialModelProviderXai`         | xai                               |