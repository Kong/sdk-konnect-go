# ConsumeFailureMode

Describes how to handle a failure in a policy applied to consumed records.
* `error` - the batch is not delivered to the client. Use sparingly: erroring on a batch causes clients to get stuck on the problematic offset and requires manual intervention to skip it.
A single record that always fails the policy, for example one written by a producer with access to the topic, pins every consumer group on that offset.
The error the client receives contains the policy failure reason, which can include runtime-specific configuration details.
* `skip` - the record is not delivered to the client.
* `passthrough` - passes the record to the client even though policy execution failed.
* `mark` - passes the record to the client but marks it with a `kong/policy-failure-<id>` header whose value is the reason for the policy failure (truncated to 512 characters).
The failure reason can contain runtime-specific configuration details, such as schema IDs, registry URLs, field paths, KMS ARNs and key IDs, KMS error text, JWT issuer and audience data, or expression snippets.
If you don't want to share this information with clients, remove the header with a Modify Headers consume policy.

**Requires a minimum runtime version of `1.2`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ConsumeFailureModeError

// Open enum: custom values can be created with a direct type cast
custom := components.ConsumeFailureMode("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ConsumeFailureModeError`       | error                           |
| `ConsumeFailureModeSkip`        | skip                            |
| `ConsumeFailureModePassthrough` | passthrough                     |
| `ConsumeFailureModeMark`        | mark                            |