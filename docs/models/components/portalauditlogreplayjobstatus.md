# PortalAuditLogReplayJobStatus

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PortalAuditLogReplayJobStatusUnconfigured

// Open enum: custom values can be created with a direct type cast
custom := components.PortalAuditLogReplayJobStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `PortalAuditLogReplayJobStatusUnconfigured` | unconfigured                                |
| `PortalAuditLogReplayJobStatusAccepted`     | accepted                                    |
| `PortalAuditLogReplayJobStatusPending`      | pending                                     |
| `PortalAuditLogReplayJobStatusRunning`      | running                                     |
| `PortalAuditLogReplayJobStatusCompleted`    | completed                                   |
| `PortalAuditLogReplayJobStatusFailed`       | failed                                      |