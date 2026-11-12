# ClientAlg

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ClientAlgEs256

// Open enum: custom values can be created with a direct type cast
custom := components.ClientAlg("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `ClientAlgEs256` | ES256            |
| `ClientAlgEs384` | ES384            |
| `ClientAlgEs512` | ES512            |
| `ClientAlgEdDsa` | EdDSA            |
| `ClientAlgHs256` | HS256            |
| `ClientAlgHs384` | HS384            |
| `ClientAlgHs512` | HS512            |
| `ClientAlgPs256` | PS256            |
| `ClientAlgPs384` | PS384            |
| `ClientAlgPs512` | PS512            |
| `ClientAlgRs256` | RS256            |
| `ClientAlgRs384` | RS384            |
| `ClientAlgRs512` | RS512            |