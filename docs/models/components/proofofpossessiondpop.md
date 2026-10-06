# ProofOfPossessionDpop

Enable Demonstrating Proof-of-Possession (DPoP). If set to strict, all request are verified despite the presence of the DPoP key claim (cnf.jkt). If set to optional, only tokens bound with DPoP's key are verified with the proof.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProofOfPossessionDpopOff

// Open enum: custom values can be created with a direct type cast
custom := components.ProofOfPossessionDpop("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ProofOfPossessionDpopOff`      | off                             |
| `ProofOfPossessionDpopOptional` | optional                        |
| `ProofOfPossessionDpopStrict`   | strict                          |