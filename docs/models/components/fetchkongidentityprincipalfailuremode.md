# FetchKongIdentityPrincipalFailureMode

Behavior when the Kong Identity principal lookup fails.
* `error` - fail the authentication if the principal lookup fails.
* `ignore` - proceed without principal metadata if the lookup fails.

**Requires a minimum runtime version of `1.2`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.FetchKongIdentityPrincipalFailureModeError

// Open enum: custom values can be created with a direct type cast
custom := components.FetchKongIdentityPrincipalFailureMode("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `FetchKongIdentityPrincipalFailureModeError`  | error                                         |
| `FetchKongIdentityPrincipalFailureModeIgnore` | ignore                                        |