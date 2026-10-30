# AIGatewayModelModelCapabilities

**`decisions` requires a minimum runtime version of `2.2`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelModelCapabilitiesWildcard

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelModelCapabilities("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AIGatewayModelModelCapabilitiesWildcard`           | *                                                   |
| `AIGatewayModelModelCapabilitiesGenerate`           | generate                                            |
| `AIGatewayModelModelCapabilitiesAgentic`            | agentic                                             |
| `AIGatewayModelModelCapabilitiesRealtime`           | realtime                                            |
| `AIGatewayModelModelCapabilitiesEmbeddings`         | embeddings                                          |
| `AIGatewayModelModelCapabilitiesImage`              | image                                               |
| `AIGatewayModelModelCapabilitiesAudioSpeech`        | audio/speech                                        |
| `AIGatewayModelModelCapabilitiesAudioTranscription` | audio/transcription                                 |
| `AIGatewayModelModelCapabilitiesAudioTranslation`   | audio/translation                                   |
| `AIGatewayModelModelCapabilitiesVideo`              | video                                               |
| `AIGatewayModelModelCapabilitiesRerank`             | rerank                                              |
| `AIGatewayModelModelCapabilitiesDecisions`          | decisions                                           |