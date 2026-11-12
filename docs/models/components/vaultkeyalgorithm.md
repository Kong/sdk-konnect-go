# VaultKeyAlgorithm

The asymmetric algorithm of the vault key.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VaultKeyAlgorithmRsa

// Open enum: custom values can be created with a direct type cast
custom := components.VaultKeyAlgorithm("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `VaultKeyAlgorithmRsa` | RSA                    |
| `VaultKeyAlgorithmEc`  | EC                     |