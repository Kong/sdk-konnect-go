# PartialEmbeddingsProvider

AI provider format to use for embeddings API

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialEmbeddingsProviderAzure

// Open enum: custom values can be created with a direct type cast
custom := components.PartialEmbeddingsProvider("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `PartialEmbeddingsProviderAzure`       | azure                                  |
| `PartialEmbeddingsProviderBedrock`     | bedrock                                |
| `PartialEmbeddingsProviderGemini`      | gemini                                 |
| `PartialEmbeddingsProviderHuggingface` | huggingface                            |
| `PartialEmbeddingsProviderMistral`     | mistral                                |
| `PartialEmbeddingsProviderOllama`      | ollama                                 |
| `PartialEmbeddingsProviderOpenai`      | openai                                 |