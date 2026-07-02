# MCPServers

## Overview

### Available Operations

* [ListMcpServerConfigs](#listmcpserverconfigs) - List all MCP Servers
* [CreateMcpServerConfig](#createmcpserverconfig) - Create an MCP Server
* [GetMcpServerConfig](#getmcpserverconfig) - Get MCP Server
* [UpdateMcpServerConfig](#updatemcpserverconfig) - Update MCP Server
* [PatchMcpServerConfig](#patchmcpserverconfig) - Partially Update MCP Server
* [DeleteMcpServerConfig](#deletemcpserverconfig) - Delete MCP Server
* [GetMcpServerStatus](#getmcpserverstatus) - Get MCP Server deployment status
* [GetMcpServerGeneratedCode](#getmcpservergeneratedcode) - Get generated Python code for an MCP Server
* [ListMcpResources](#listmcpresources) - List MCP Resources
* [CreateMcpResource](#createmcpresource) - Create an MCP Resource
* [GetMcpResource](#getmcpresource) - Get an MCP Resource
* [DeleteMcpResource](#deletemcpresource) - Delete an MCP Resource
* [UpdateMcpResource](#updatemcpresource) - Update an MCP Resource
* [GetMcpServerSignals](#getmcpserversignals) - Get MCP Server Signals
* [ListMcpServersByControlPlane](#listmcpserversbycontrolplane) - List MCP Servers by Control Plane
* [GetMcpServerByControlPlane](#getmcpserverbycontrolplane) - Get MCP Server by Control Plane
* [GetMcpServerKongEntities](#getmcpserverkongentities) - Get Kong entities for the MCP Server Gateway
* [PostMcpServerStatus](#postmcpserverstatus) - Report MCP Server deployment status
* [GetMcpServerCode](#getmcpservercode) - Get generated Python code for the MCP Server

## ListMcpServerConfigs

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
        DisplayName: "Kellie85",
        Description: "scuffle lost zowie eek boohoo",
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServer != nil {
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
    if res.MCPServer != nil {
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

    res, err := s.MCPServers.UpdateMcpServerConfig(ctx, "fd3b7872-2dae-4834-a227-54b3bea26ffa", components.CreateMCPServerRequest{
        Name: "<value>",
        DisplayName: "Weston_Dicki",
        Description: "astride reference wolf delightfully deafening inasmuch",
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `mcpServerID`                                                                          | `string`                                                                               | :heavy_check_mark:                                                                     | The ID of the MCP server.                                                              |
| `createMCPServerRequest`                                                               | [components.CreateMCPServerRequest](../../models/components/createmcpserverrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
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

    res, err := s.MCPServers.PatchMcpServerConfig(ctx, "2caa5873-34f4-44c1-bcfe-6f21646006c8", components.PatchMCPServerRequest{
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServer != nil {
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

## GetMcpServerGeneratedCode

Generates MCP server Python code from the OpenAPI specifications associated with the MCP server.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-server-generated-code" method="get" path="/v1/mcp-servers/{mcpServerId}/code" -->
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

    res, err := s.MCPServers.GetMcpServerGeneratedCode(ctx, "f732f541-e15d-45b8-ba94-dbf4f86dff3a")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPServerCodeResponse != nil {
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

**[*operations.GetMcpServerGeneratedCodeResponse](../../models/operations/getmcpservergeneratedcoderesponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListMcpResources

Returns a list of MCP resource objects.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-resources" method="get" path="/v1/mcp-resources" -->
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

    res, err := s.MCPServers.ListMcpResources(ctx, operations.ListMcpResourcesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMCPResourcesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListMcpResourcesRequest](../../models/operations/listmcpresourcesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListMcpResourcesResponse](../../models/operations/listmcpresourcesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateMcpResource

Create an MCP resource in the Konnect Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-resource" method="post" path="/v1/mcp-resources" -->
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

    res, err := s.MCPServers.CreateMcpResource(ctx, components.MCPResourceInput{
        Name: "my-api-resource",
        DisplayName: "Richard_Bogisich-Haley",
        Labels: map[string]string{
            "env": "test",
        },
        Source: components.CreateMCPResourceSourceDetailInputAPICatalog(
            components.MCPResourceSourceAPICatalog{
                Type: components.MCPResourceSourceAPICatalogTypeAPICatalog,
                Config: components.MCPResourceSourceAPICatalogConfig{
                    APIID: "158a0fc7-479a-4928-8a64-7656c1824f7f",
                    APIVersionID: "bc0b9df1-67ba-475f-9670-37de3cc2e4f6",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Source.Type {
            case components.MCPResourceSourceDetailTypeAPICatalog:
                // res.MCPResourceInfo.Source.MCPResourceSourceAPICatalog is populated
            case components.MCPResourceSourceDetailTypeRaw:
                // res.MCPResourceInfo.Source.MCPResourceSourceRaw is populated
        }

    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.MCPResourceInput](../../models/components/mcpresourceinput.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateMcpResourceResponse](../../models/operations/createmcpresourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpResource

Retrieve an MCP resource by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-resource" method="get" path="/v1/mcp-resources/{resourceId}" -->
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

    res, err := s.MCPServers.GetMcpResource(ctx, "1c4cc520-45e3-4087-8b08-6465521d7fc2")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Source.Type {
            case components.MCPResourceSourceDetailTypeAPICatalog:
                // res.MCPResourceInfo.Source.MCPResourceSourceAPICatalog is populated
            case components.MCPResourceSourceDetailTypeRaw:
                // res.MCPResourceInfo.Source.MCPResourceSourceRaw is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `resourceID`                                             | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP resource.                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMcpResourceResponse](../../models/operations/getmcpresourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMcpResource

Delete an MCP resource by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mcp-resource" method="delete" path="/v1/mcp-resources/{resourceId}" -->
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

    res, err := s.MCPServers.DeleteMcpResource(ctx, "3d073fde-7988-464a-81f2-9949151b4314")
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
| `resourceID`                                             | `string`                                                 | :heavy_check_mark:                                       | The ID of the MCP resource.                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteMcpResourceResponse](../../models/operations/deletemcpresourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateMcpResource

Update an MCP resource by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-mcp-resource" method="put" path="/v1/mcp-resources/{resourceId}" -->
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

    res, err := s.MCPServers.UpdateMcpResource(ctx, "fe6ae9ac-feda-46df-81e4-112aab1e4e51", components.MCPResourceInput{
        Name: "my-api-resource",
        DisplayName: "Daron_Graham",
        Labels: map[string]string{
            "env": "test",
        },
        Source: components.CreateMCPResourceSourceDetailInputAPICatalog(
            components.MCPResourceSourceAPICatalog{
                Type: components.MCPResourceSourceAPICatalogTypeAPICatalog,
                Config: components.MCPResourceSourceAPICatalogConfig{
                    APIID: "158a0fc7-479a-4928-8a64-7656c1824f7f",
                    APIVersionID: "bc0b9df1-67ba-475f-9670-37de3cc2e4f6",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Source.Type {
            case components.MCPResourceSourceDetailTypeAPICatalog:
                // res.MCPResourceInfo.Source.MCPResourceSourceAPICatalog is populated
            case components.MCPResourceSourceDetailTypeRaw:
                // res.MCPResourceInfo.Source.MCPResourceSourceRaw is populated
        }

    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `resourceID`                                                               | `string`                                                                   | :heavy_check_mark:                                                         | The ID of the MCP resource.                                                |
| `mcpResourceInput`                                                         | [components.MCPResourceInput](../../models/components/mcpresourceinput.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.UpdateMcpResourceResponse](../../models/operations/updatemcpresourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpServerSignals

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

## ListMcpServersByControlPlane

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

## GetMcpServerKongEntities

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

## GetMcpServerCode

Generates MCP server Python code from the OpenAPI specifications associated with the MCP server.


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