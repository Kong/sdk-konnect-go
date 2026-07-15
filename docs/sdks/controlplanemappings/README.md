# ControlPlaneMappings

## Overview

### Available Operations

* [ListControlPlaneMappings](#listcontrolplanemappings) - Get a list of CPs mapped to MCP Server
* [CreateMcpServerControlPlaneMapping](#createmcpservercontrolplanemapping) - Create MCP Server CP Mapping
* [GetControlPlaneMapping](#getcontrolplanemapping) - Get a mapping between an MCP Server and a CP
* [PatchMcpServerControlPlaneMapping](#patchmcpservercontrolplanemapping) - Patch MCP Server Control Plane Mapping
* [DeleteMcpServerControlPlaneMapping](#deletemcpservercontrolplanemapping) - Delete MCP Server Control Plane Mapping

## ListControlPlaneMappings

Get a list of all Control Planes that have been mapped to the given MCP Server.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-control-plane-mappings" method="get" path="/v1/mcp-servers/{mcpServerId}/control-plane-mappings" -->
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

    res, err := s.ControlPlaneMappings.ListControlPlaneMappings(ctx, operations.ListControlPlaneMappingsRequest{
        McpServerID: "2f594555-a2ac-438e-b375-6e7bbbb984de",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMCPServerControlPlaneMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListControlPlaneMappingsRequest](../../models/operations/listcontrolplanemappingsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListControlPlaneMappingsResponse](../../models/operations/listcontrolplanemappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateMcpServerControlPlaneMapping

Create an MCP Server - Control Plane mapping

### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-server-control-plane-mapping" method="post" path="/v1/mcp-servers/{mcpServerId}/control-plane-mappings" -->
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

    res, err := s.ControlPlaneMappings.CreateMcpServerControlPlaneMapping(ctx, "8929d049-bab0-4b7d-b442-a77a71ff08ee", components.CreateMCPServerControlPlaneMappingRequest{
        ControlPlaneID: "0ae82373-342d-4224-bdb3-28e7167b126e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerControlPlaneMappingInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `mcpServerID`                                                                                                                | `string`                                                                                                                     | :heavy_check_mark:                                                                                                           | The ID of the MCP server.                                                                                                    |
| `createMCPServerControlPlaneMappingRequest`                                                                                  | [components.CreateMCPServerControlPlaneMappingRequest](../../models/components/createmcpservercontrolplanemappingrequest.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CreateMcpServerControlPlaneMappingResponse](../../models/operations/createmcpservercontrolplanemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetControlPlaneMapping

Retrieve a mapping by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-control-plane-mapping" method="get" path="/v1/mcp-servers/{mcpServerId}/control-plane-mappings/{mappingId}" -->
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

    res, err := s.ControlPlaneMappings.GetControlPlaneMapping(ctx, "ade854e2-778b-48d0-998f-9587340f7b84", "8e260d77-f92d-45ec-95c2-dacf1ac8a9ec")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerControlPlaneMappingInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mcpServerID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server.                                |
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP Server - Control Plane mapping.        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetControlPlaneMappingResponse](../../models/operations/getcontrolplanemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchMcpServerControlPlaneMapping

Update a mapping between an MCP Server and a Control Plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-mcp-server-control-plane-mapping" method="patch" path="/v1/mcp-servers/{mcpServerId}/control-plane-mappings/{mappingId}" -->
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

    res, err := s.ControlPlaneMappings.PatchMcpServerControlPlaneMapping(ctx, operations.PatchMcpServerControlPlaneMappingRequest{
        McpServerID: "1fcda8ef-10c5-4856-a107-54895723bc0f",
        MappingID: "1b7d7ac5-1b4e-4d99-9906-a42b9fa8eba6",
        PatchMCPServerControlPlaneMappingRequest: components.PatchMCPServerControlPlaneMappingRequest{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerControlPlaneMappingInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PatchMcpServerControlPlaneMappingRequest](../../models/operations/patchmcpservercontrolplanemappingrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.PatchMcpServerControlPlaneMappingResponse](../../models/operations/patchmcpservercontrolplanemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMcpServerControlPlaneMapping

Delete a mapping between an MCP Server and a Control Plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mcp-server-control-plane-mapping" method="delete" path="/v1/mcp-servers/{mcpServerId}/control-plane-mappings/{mappingId}" -->
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

    res, err := s.ControlPlaneMappings.DeleteMcpServerControlPlaneMapping(ctx, "485d8ed6-b56e-40e9-b241-1218ef25eb73", "08e9b37a-0b11-42f6-a8eb-5f3b6c132825")
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
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP Server - Control Plane mapping.        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteMcpServerControlPlaneMappingResponse](../../models/operations/deletemcpservercontrolplanemappingresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |