# ConflictErrorStatus

The HTTP status code of the error. Useful when passing the response
body to child properties in a frontend UI. Must be returned as an integer.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
)

value := sdkerrors.ConflictErrorStatusFourHundredAndNine
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ConflictErrorStatusFourHundredAndNine` | 409                                     |