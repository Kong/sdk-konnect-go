# QueryParamForce

If true, allows operations to be removed from the current version when using access control enforcement.
If false, operations removal will be rejected with a 409 error.
Omitting the value means true.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.QueryParamForceTrue
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `QueryParamForceTrue`  | true                   |
| `QueryParamForceFalse` | false                  |