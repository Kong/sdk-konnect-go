# RouteType

The model's operation implementation, for this provider. 

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RouteTypeAudioV1AudioSpeech

// Open enum: custom values can be created with a direct type cast
custom := components.RouteType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `RouteTypeAudioV1AudioSpeech`         | audio/v1/audio/speech                 |
| `RouteTypeAudioV1AudioTranscriptions` | audio/v1/audio/transcriptions         |
| `RouteTypeAudioV1AudioTranslations`   | audio/v1/audio/translations           |
| `RouteTypeImageV1ImagesEdits`         | image/v1/images/edits                 |
| `RouteTypeImageV1ImagesGenerations`   | image/v1/images/generations           |
| `RouteTypeLlmV1Assistants`            | llm/v1/assistants                     |
| `RouteTypeLlmV1Batches`               | llm/v1/batches                        |
| `RouteTypeLlmV1Chat`                  | llm/v1/chat                           |
| `RouteTypeLlmV1Completions`           | llm/v1/completions                    |
| `RouteTypeLlmV1Embeddings`            | llm/v1/embeddings                     |
| `RouteTypeLlmV1Files`                 | llm/v1/files                          |
| `RouteTypeLlmV1Responses`             | llm/v1/responses                      |
| `RouteTypePreserve`                   | preserve                              |
| `RouteTypeRealtimeV1Realtime`         | realtime/v1/realtime                  |
| `RouteTypeVideoV1VideosGenerations`   | video/v1/videos/generations           |