# AIGWOpenIDConnectGeneratedConfigSource

Where to obtain the actor token. `header` reads it from a request header. `config` uses a fixed value. `none` disables actor token support.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGWOpenIDConnectGeneratedConfigSourceConfig

// Open enum: custom values can be created with a direct type cast
custom := components.AIGWOpenIDConnectGeneratedConfigSource("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `AIGWOpenIDConnectGeneratedConfigSourceConfig` | config                                         |
| `AIGWOpenIDConnectGeneratedConfigSourceHeader` | header                                         |
| `AIGWOpenIDConnectGeneratedConfigSourceNone`   | none                                           |