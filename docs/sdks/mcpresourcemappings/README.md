# MCPResourceMappings

## Overview

### Available Operations

* [CreateMcpServerResourceMapping](#createmcpserverresourcemapping) - Create MCP Server MCP Resource Mapping
* [ListMcpServerResourceMappings](#listmcpserverresourcemappings) - List MCP Server MCP Resource Mappings
* [DeleteMcpServerResourceMapping](#deletemcpserverresourcemapping) - Delete MCP Server MCP Resource Mapping
* [GetMcpServerResourceMapping](#getmcpserverresourcemapping) - Get MCP Server MCP Resource Mapping

## CreateMcpServerResourceMapping

Create a mapping between an MCP server and an MCP resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-server-resource-mapping" method="post" path="/v1/mcp-servers/{mcpServerId}/mcp-resource-mappings" -->
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

    res, err := s.MCPResourceMappings.CreateMcpServerResourceMapping(ctx, "e2514dfc-c6c2-4005-b56d-bfa15c3a74c3", components.CreateMCPServerMCPResourceMappingRequest{
        McpResourceID: "a60eceeb-dd7c-4884-8b26-806489c2f704",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerMCPResourceMapping != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `mcpServerID`                                                                                                              | `string`                                                                                                                   | :heavy_check_mark:                                                                                                         | The ID of the MCP server.                                                                                                  |
| `createMCPServerMCPResourceMappingRequest`                                                                                 | [components.CreateMCPServerMCPResourceMappingRequest](../../models/components/createmcpservermcpresourcemappingrequest.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.CreateMcpServerResourceMappingResponse](../../models/operations/createmcpserverresourcemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListMcpServerResourceMappings

Returns a list of MCP resource mappings for the given MCP server.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-server-resource-mappings" method="get" path="/v1/mcp-servers/{mcpServerId}/mcp-resource-mappings" -->
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

    res, err := s.MCPResourceMappings.ListMcpServerResourceMappings(ctx, operations.ListMcpServerResourceMappingsRequest{
        McpServerID: "cb866058-5f07-4cbf-82c6-0cb189b9ed57",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMCPServerMCPResourceMappingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListMcpServerResourceMappingsRequest](../../models/operations/listmcpserverresourcemappingsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListMcpServerResourceMappingsResponse](../../models/operations/listmcpserverresourcemappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMcpServerResourceMapping

Delete a mapping between an MCP server and an MCP resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mcp-server-resource-mapping" method="delete" path="/v1/mcp-servers/{mcpServerId}/mcp-resource-mappings/{mappingId}" -->
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

    res, err := s.MCPResourceMappings.DeleteMcpServerResourceMapping(ctx, "7775aac7-ea1b-40f3-a6d3-1fb6937f197d", "b8156c3e-5e1b-4d68-82dd-9e1cf8505813")
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
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server MCP resource mapping.           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteMcpServerResourceMappingResponse](../../models/operations/deletemcpserverresourcemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerResourceMapping

Retrieve a specific MCP resource mapping by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-resource-mapping" method="get" path="/v1/mcp-servers/{mcpServerId}/mcp-resource-mappings/{mappingId}" -->
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

    res, err := s.MCPResourceMappings.GetMcpServerResourceMapping(ctx, "b30d8ac1-7a4e-4161-a394-13537b6b4e01", "b4036286-f265-4a6e-906d-36922afc368e")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerMCPResourceMapping != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mcpServerID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server.                                |
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP server MCP resource mapping.           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMcpServerResourceMappingResponse](../../models/operations/getmcpserverresourcemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |