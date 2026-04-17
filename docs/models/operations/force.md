# Force

If true, delete specified config store and all secrets, even if there are secrets linked to the config store If false, do not allow deletion if there are secrets linked to the config store

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.ForceTrue
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `ForceTrue`  | true         |
| `ForceFalse` | false        |