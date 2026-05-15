# ProduceFailureMode

Describes how to handle a failure in a policy applied to produced records.
* `reject` - rejects the record batch.
* `passthrough` - passes the record silently to the backend cluster even though policy execution failed.
* `mark` - passes the record to the backend cluster but marks it with a `kong/policy-failure-<id>` header whose value is the reason for the policy failure (truncated to 512 characters).


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProduceFailureModeReject

// Open enum: custom values can be created with a direct type cast
custom := components.ProduceFailureMode("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ProduceFailureModeReject`      | reject                          |
| `ProduceFailureModePassthrough` | passthrough                     |
| `ProduceFailureModeMark`        | mark                            |