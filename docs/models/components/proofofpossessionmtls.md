# ProofOfPossessionMtls

Enable mtls proof of possession. If set to strict, all tokens (from supported auth_methods: bearer, introspection, and session granted with bearer or introspection) are verified, if set to optional, only tokens that contain the certificate hash claim are verified. If the verification fails, the request will be rejected with 401.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProofOfPossessionMtlsOff

// Open enum: custom values can be created with a direct type cast
custom := components.ProofOfPossessionMtls("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ProofOfPossessionMtlsOff`      | off                             |
| `ProofOfPossessionMtlsOptional` | optional                        |
| `ProofOfPossessionMtlsStrict`   | strict                          |