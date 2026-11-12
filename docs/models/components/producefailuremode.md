# ProduceFailureMode

Describes how to handle a failure in a policy applied to produced records.
* `reject` - rejects the record batch.
* `passthrough` - passes the record silently to the backend cluster even though policy execution failed.
* `mark` - passes the record to the backend cluster but marks it with a `kong/policy-failure-<id>` header whose value is the reason for the policy failure (truncated to 512 characters).
The header is stored with the record in the backend cluster, so every consumer of the topic sees it, including systems that never pass through the Event Gateway.
The failure reason can contain runtime-specific configuration details, such as schema IDs, registry URLs, field paths, KMS ARNs and key IDs, KMS error text, JWT issuer and audience data, or expression snippets, and a producer with access to the topic can trigger policy failures on purpose to expose it.
If you don't want to share this information, remove the header with a Modify Headers produce policy, or use a different failure mode.

**Requires a minimum runtime version of `1.2`**.

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