# MCPServerRuntime

## Overview

### Available Operations

* [GetMcpServerRuntimeConfig](#getmcpserverruntimeconfig) - Get MCP Server Runtime Configuration

## GetMcpServerRuntimeConfig

Returns the runtime configuration for a deployed MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-runtime-config" method="get" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers/{mcpServerId}/runtime-config" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServerRuntime.GetMcpServerRuntimeConfig(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "e6753948-dd59-41a8-a93d-f27fad29f648")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerRuntime != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `mcpServerID`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | The ID of the MCP server.                                                          |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetMcpServerRuntimeConfigResponse](../../models/operations/getmcpserverruntimeconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |