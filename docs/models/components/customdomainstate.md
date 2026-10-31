# CustomDomainState

The current state of the custom domain. Possible values:
- `created` — The domain has been registered but TLS provisioning has not yet started.
- `initializing` — Konnect is provisioning the TLS certificate and configuring SNI routing.
- `ready` — The domain is fully provisioned and serving traffic.
- `terminating` — The domain is being deleted and its TLS certificate is being removed.
- `terminated` — The domain has been fully deleted and is no longer available.
- `error` — Provisioning failed; check `state_metadata` for details.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CustomDomainStateCreated

// Open enum: custom values can be created with a direct type cast
custom := components.CustomDomainState("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CustomDomainStateCreated`      | created                         |
| `CustomDomainStateInitializing` | initializing                    |
| `CustomDomainStateReady`        | ready                           |
| `CustomDomainStateTerminating`  | terminating                     |
| `CustomDomainStateTerminated`   | terminated                      |
| `CustomDomainStateError`        | error                           |