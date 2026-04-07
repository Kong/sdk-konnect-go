# MCPServers

## Overview

### Available Operations

* [ListMcpServersByControlPlane](#listmcpserversbycontrolplane) - List MCP Servers by Control Plane
* [GetMcpServerByControlPlane](#getmcpserverbycontrolplane) - Get MCP Server by Control Plane
* [GetMcpServerCode](#getmcpservercode) - Get generated Python code for the MCP Server
* [GetMcpServerKongEntities](#getmcpserverkongentities) - Get Kong entities for the MCP Server Gateway
* [PostMcpServerStatus](#postmcpserverstatus) - Report MCP Server deployment status
* [GetMcpServerSignals](#getmcpserversignals) - Get MCP Server Signals
* [ListMcpServerConfigs](#listmcpserverconfigs) - List all MCP Servers
* [CreateMcpServerConfig](#createmcpserverconfig) - Create an MCP Server
* [GetMcpServerConfig](#getmcpserverconfig) - Get MCP Server
* [UpdateMcpServerConfig](#updatemcpserverconfig) - Update MCP Server
* [PatchMcpServerConfig](#patchmcpserverconfig) - Partially Update MCP Server
* [DeleteMcpServerConfig](#deletemcpserverconfig) - Delete MCP Server
* [GetMcpServerStatus](#getmcpserverstatus) - Get MCP Server deployment status
* [ListMcpServerTools](#listmcpservertools) - List MCP Server Tools
* [CreateMcpServerTool](#createmcpservertool) - Create MCP Server Tool
* [GetMcpServerTool](#getmcpservertool) - Get MCP Server Tool
* [UpdateMcpServerTool](#updatemcpservertool) - Update MCP Server Tool
* [PatchMcpServerTool](#patchmcpservertool) - Partially Update MCP Server Tool
* [DeleteMcpServerTool](#deletemcpservertool) - Delete MCP Server Tool

## ListMcpServersByControlPlane

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns an array of MCP server objects for the specified control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-servers-by-control-plane" method="get" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServers.ListMcpServersByControlPlane(ctx, operations.ListMcpServersByControlPlaneRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMCPServersCPInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListMcpServersByControlPlaneRequest](../../models/operations/listmcpserversbycontrolplanerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListMcpServersByControlPlaneResponse](../../models/operations/listmcpserversbycontrolplaneresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerByControlPlane

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Retrieves an MCP server by its ID within a specific control plane context.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-by-control-plane" method="get" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers/{mcpServerId}" -->
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

    res, err := s.MCPServers.GetMcpServerByControlPlane(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "d50e98e4-a114-4b7b-8375-f4f6aa5b4a5a")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerCPInfo != nil {
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

**[*operations.GetMcpServerByControlPlaneResponse](../../models/operations/getmcpserverbycontrolplaneresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerCode

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Retrieves generated Python code for the MCP Server implementation.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-code" method="get" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers/{mcpServerId}/code" -->
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

    res, err := s.MCPServers.GetMcpServerCode(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "c561f6dd-36f0-404e-882e-b647d9b67032")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerCodeResponse != nil {
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

**[*operations.GetMcpServerCodeResponse](../../models/operations/getmcpservercoderesponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerKongEntities

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Retrieves Kong entities (Routes, Services, Plugins) for the Gateway in front of the MCP Server.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-kong-entities" method="get" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers/{mcpServerId}/kong-entities" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServers.GetMcpServerKongEntities(ctx, operations.GetMcpServerKongEntitiesRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        McpServerID: "d38237b3-7d32-4cf5-bc8b-652fc6e3a2d6",
        RuntimeVersion: sdkkonnectgo.Pointer("3.13.0.0"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongEntitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetMcpServerKongEntitiesRequest](../../models/operations/getmcpserverkongentitiesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetMcpServerKongEntitiesResponse](../../models/operations/getmcpserverkongentitiesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PostMcpServerStatus

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Reports the current deployment status of an MCP server, including replica counts and per-version pod statuses.

### Example Usage

<!-- UsageSnippet language="go" operationID="post-mcp-server-status" method="post" path="/v1/mcp-cp/{controlPlaneId}/mcp-servers/{mcpServerId}/status" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServers.PostMcpServerStatus(ctx, operations.PostMcpServerStatusRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        McpServerID: "b4797f3d-7c96-462e-95e9-0ab14490e2f3",
        RequestBody: []components.MCPServerVersionStatus{
            components.MCPServerVersionStatus{
                Version: "v1",
                DesiredReplicas: 2,
                CreatedReplicas: 2,
                PodsStatus: components.MCPServerPodsStatus{
                    Ready: 2,
                    Starting: 0,
                    Failing: 0,
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PostMcpServerStatusRequest](../../models/operations/postmcpserverstatusrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostMcpServerStatusResponse](../../models/operations/postmcpserverstatusresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerSignals

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Client-initiated signals fetch with long-polling capabilities.
The client supplies `capabilities` as a deepObject query parameter:


  - Keys are capability names (e.g. `mcp`)
  - Values contain the capability request fields (`version`)

Example:


  GET /mcp-cp/{controlPlaneId}/signals?capabilities[mcp][version]=5

The CP responds with a JSON payload containing zero or more signals that the client should process.
If no signals are available, the CP responds with HTTP 304 Not Modified.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-signals" method="get" path="/v1/mcp-cp/{controlPlaneId}/signals" -->
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

    res, err := s.MCPServers.GetMcpServerSignals(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", &components.MCPCapabilitiesMap{
        Mcp: &components.MCPCapabilityRequest{
            Version: "5",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerSignals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                       | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     | Example                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                           | :heavy_check_mark:                                                                                                                              | The context to use for the request.                                                                                                             |                                                                                                                                                 |
| `controlPlaneID`                                                                                                                                | `string`                                                                                                                                        | :heavy_check_mark:                                                                                                                              | The UUID of your control plane. This variable is available in the Konnect manager.                                                              | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                            |
| `capabilities`                                                                                                                                  | [*components.MCPCapabilitiesMap](../../models/components/mcpcapabilitiesmap.md)                                                                 | :heavy_minus_sign:                                                                                                                              | Capability requests as a deepObject map.<br/>- Keys are capability names (e.g. `mcp`).<br/>- Values contain the capability request fields (`version`).<br/> | {<br/>"mcp": {<br/>"version": "5"<br/>}<br/>}                                                                                                   |
| `opts`                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                        | :heavy_minus_sign:                                                                                                                              | The options for this request.                                                                                                                   |                                                                                                                                                 |

### Response

**[*operations.GetMcpServerSignalsResponse](../../models/operations/getmcpserversignalsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListMcpServerConfigs

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns an array of MCP server objects.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-server-configs" method="get" path="/v1/mcp-servers" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServers.ListMcpServerConfigs(ctx, operations.ListMcpServerConfigsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMCPServersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListMcpServerConfigsRequest](../../models/operations/listmcpserverconfigsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListMcpServerConfigsResponse](../../models/operations/listmcpserverconfigsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateMcpServerConfig

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create an MCP Server in the Konnect Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-server-config" method="post" path="/v1/mcp-servers" -->
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

    res, err := s.MCPServers.CreateMcpServerConfig(ctx, components.CreateMCPServerRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateMCPServerRequest](../../models/components/createmcpserverrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateMcpServerConfigResponse](../../models/operations/createmcpserverconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerConfig

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns information about an individual MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-config" method="get" path="/v1/mcp-servers/{mcpServerId}" -->
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

    res, err := s.MCPServers.GetMcpServerConfig(ctx, "1913c3b9-3d2d-47ed-978a-14b63dc661a3")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mcpServerID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMcpServerConfigResponse](../../models/operations/getmcpserverconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateMcpServerConfig

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update an individual MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-mcp-server-config" method="put" path="/v1/mcp-servers/{mcpServerId}" -->
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

    res, err := s.MCPServers.UpdateMcpServerConfig(ctx, "fd3b7872-2dae-4834-a227-54b3bea26ffa", components.UpdateMCPServerRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `mcpServerID`                                                                          | `string`                                                                               | :heavy_check_mark:                                                                     | The ID of the MCP server.                                                              |
| `updateMCPServerRequest`                                                               | [components.UpdateMCPServerRequest](../../models/components/updatemcpserverrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateMcpServerConfigResponse](../../models/operations/updatemcpserverconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchMcpServerConfig

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Partially update an individual MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-mcp-server-config" method="patch" path="/v1/mcp-servers/{mcpServerId}" -->
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

    res, err := s.MCPServers.PatchMcpServerConfig(ctx, "2caa5873-34f4-44c1-bcfe-6f21646006c8", components.PatchMCPServerRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `mcpServerID`                                                                        | `string`                                                                             | :heavy_check_mark:                                                                   | The ID of the MCP server.                                                            |
| `patchMCPServerRequest`                                                              | [components.PatchMCPServerRequest](../../models/components/patchmcpserverrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchMcpServerConfigResponse](../../models/operations/patchmcpserverconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMcpServerConfig

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Delete an MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mcp-server-config" method="delete" path="/v1/mcp-servers/{mcpServerId}" -->
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

    res, err := s.MCPServers.DeleteMcpServerConfig(ctx, "4b4da211-e969-4d21-a344-b82be05a7256")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mcpServerID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteMcpServerConfigResponse](../../models/operations/deletemcpserverconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerStatus

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the current aggregated deployment status of an MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-status" method="get" path="/v1/mcp-servers/{mcpServerId}/status" -->
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

    res, err := s.MCPServers.GetMcpServerStatus(ctx, "e3bb4453-e73c-413b-964a-bd160b2101c7")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mcpServerID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMcpServerStatusResponse](../../models/operations/getmcpserverstatusresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListMcpServerTools

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns an array of MCP server tool objects.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-server-tools" method="get" path="/v1/mcp-servers/{mcpServerId}/tools" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServers.ListMcpServerTools(ctx, operations.ListMcpServerToolsRequest{
        McpServerID: "8143deaf-bb0c-4c1c-ae3b-77a8194a2223",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMCPServerToolsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListMcpServerToolsRequest](../../models/operations/listmcpservertoolsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListMcpServerToolsResponse](../../models/operations/listmcpservertoolsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateMcpServerTool

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create an MCP server tool under the specified MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-server-tool" method="post" path="/v1/mcp-servers/{mcpServerId}/tools" -->
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

    res, err := s.MCPServers.CreateMcpServerTool(ctx, "f4d40e0c-1334-4edb-ad42-f0a1ba080d3d", components.MCPServerToolInput{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerTool != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `mcpServerID`                                                                  | `string`                                                                       | :heavy_check_mark:                                                             | The ID of the MCP server.                                                      |
| `mcpServerToolInput`                                                           | [components.MCPServerToolInput](../../models/components/mcpservertoolinput.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateMcpServerToolResponse](../../models/operations/createmcpservertoolresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerTool

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns information about an individual MCP server tool.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-tool" method="get" path="/v1/mcp-servers/{mcpServerId}/tools/{toolId}" -->
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

    res, err := s.MCPServers.GetMcpServerTool(ctx, "b5b68314-ac99-4b60-8641-559c0109235d", "717a4cd0-b59e-4615-9880-e38dd2e4170c")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerTool != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mcpServerID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server.                                |
| `toolID`                                                 | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server tool.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMcpServerToolResponse](../../models/operations/getmcpservertoolresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateMcpServerTool

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update an individual MCP server tool.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-mcp-server-tool" method="put" path="/v1/mcp-servers/{mcpServerId}/tools/{toolId}" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServers.UpdateMcpServerTool(ctx, operations.UpdateMcpServerToolRequest{
        McpServerID: "6335311e-27cb-4256-9170-f60fbac984ae",
        ToolID: "970ed3e8-4e5d-484e-9cf0-2b68d7483d23",
        MCPServerToolInput: components.MCPServerToolInput{
            Name: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerTool != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateMcpServerToolRequest](../../models/operations/updatemcpservertoolrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateMcpServerToolResponse](../../models/operations/updatemcpservertoolresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchMcpServerTool

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Partially update an individual MCP server tool.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-mcp-server-tool" method="patch" path="/v1/mcp-servers/{mcpServerId}/tools/{toolId}" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.MCPServers.PatchMcpServerTool(ctx, operations.PatchMcpServerToolRequest{
        McpServerID: "2b9cb4a0-b1be-45f7-bb94-25df94948363",
        ToolID: "bd9c8f21-19df-42fa-944f-32320d13c33d",
        PatchMCPServerToolRequest: components.PatchMCPServerToolRequest{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerTool != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchMcpServerToolRequest](../../models/operations/patchmcpservertoolrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchMcpServerToolResponse](../../models/operations/patchmcpservertoolresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMcpServerTool

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Delete an MCP server tool under the specified MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mcp-server-tool" method="delete" path="/v1/mcp-servers/{mcpServerId}/tools/{toolId}" -->
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

    res, err := s.MCPServers.DeleteMcpServerTool(ctx, "2a9438e5-7db7-4586-abf7-9326f2cfeac5", "a037decd-08db-48eb-88b6-243572fe38d6")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mcpServerID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server.                                |
| `toolID`                                                 | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server tool.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteMcpServerToolResponse](../../models/operations/deletemcpservertoolresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |