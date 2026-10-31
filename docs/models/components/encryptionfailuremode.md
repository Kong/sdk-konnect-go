# EncryptionFailureMode

Describes how to handle failing encryption or decryption.
Use `error` if the record should be rejected if encryption or decryption fails.
Use `passthrough` to ignore encryption or decryption failure and continue proxying the record.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EncryptionFailureModeError

// Open enum: custom values can be created with a direct type cast
custom := components.EncryptionFailureMode("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `EncryptionFailureModeError`       | error                              |
| `EncryptionFailureModePassthrough` | passthrough                        |