# DeleteAuthServerForce

If true, delete the specified auth server and all its associated resources. If false, only allow deletion if no clients, scopes or claims are associated with the auth server.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.DeleteAuthServerForceTrue
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `DeleteAuthServerForceTrue`  | true                         |
| `DeleteAuthServerForceFalse` | false                        |