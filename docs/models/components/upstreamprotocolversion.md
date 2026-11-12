# UpstreamProtocolVersion

The MCP protocol revision Kong speaks to the upstream MCP server. Leave unset to
negotiate a handshake revision with an `initialize` exchange, which is the default. Set a
per-request revision to reach an upstream that answers no handshake and mints no session.

**Requires a minimum runtime version of `2.1`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpstreamProtocolVersionTwoThousandAndTwentySixMinus07Minus28

// Open enum: custom values can be created with a direct type cast
custom := components.UpstreamProtocolVersion("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `UpstreamProtocolVersionTwoThousandAndTwentySixMinus07Minus28`  | 2026-07-28                                                      |
| `UpstreamProtocolVersionTwoThousandAndTwentyFiveMinus11Minus25` | 2025-11-25                                                      |
| `UpstreamProtocolVersionTwoThousandAndTwentyFiveMinus06Minus18` | 2025-06-18                                                      |
| `UpstreamProtocolVersionTwoThousandAndTwentyFiveMinus03Minus26` | 2025-03-26                                                      |