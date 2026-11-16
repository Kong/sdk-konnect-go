# CatalogMCPImplementations

## Overview

Link an MCP to an AI Gateway MCP.

### Available Operations

* [ListMcpImplementations](#listmcpimplementations) - List MCP implementations
* [CreateMcpImplementation](#createmcpimplementation) - Create MCP implementation
* [GetMcpImplementation](#getmcpimplementation) - Get MCP implementation
* [DeleteMcpImplementation](#deletemcpimplementation) - Delete MCP implementation

## ListMcpImplementations

Returns a paginated list of the MCP's implementation records.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-mcp-implementations" method="get" path="/v1/mcp-servers/{mcpId}/implementations" -->
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

    res, err := s.CatalogMCPImplementations.ListMcpImplementations(ctx, operations.ListMcpImplementationsRequest{
        McpID: "a0119846-f179-4d9f-a168-d701facce7fb",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCatalogMCPImplementationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListMcpImplementationsRequest](../../models/operations/listmcpimplementationsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListMcpImplementationsResponse](../../models/operations/listmcpimplementationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateMcpImplementation

Creates an implementation record for the MCP. Returns `409` if the MCP has an existing implementation.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-implementation" method="post" path="/v1/mcp-servers/{mcpId}/implementations" -->
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

    res, err := s.CatalogMCPImplementations.CreateMcpImplementation(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", components.CreateCreateCatalogMCPImplementationCreateCatalogMCPGatewayImplementation(
        components.CreateCatalogMCPGatewayImplementation{
            Implementation: components.CatalogMCPGatewayImplementationBlock{
                Config: components.CatalogMCPGatewayImplementationConfig{
                    GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
                    GatewayMcpServerID: "79087145-0159-4811-9042-6a9365f234bc",
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPImplementation != nil {
        switch res.CatalogMCPImplementation.Type {
            case components.CatalogMCPImplementationTypeCatalogMCPGatewayImplementation:
                // res.CatalogMCPImplementation.CatalogMCPGatewayImplementation is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `mcpID`                                                                                                | `string`                                                                                               | :heavy_check_mark:                                                                                     | The unique identifier of the MCP.                                                                      | a0119846-f179-4d9f-a168-d701facce7fb                                                                   |
| `createCatalogMCPImplementation`                                                                       | [components.CreateCatalogMCPImplementation](../../models/components/createcatalogmcpimplementation.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.CreateMcpImplementationResponse](../../models/operations/createmcpimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpImplementation

Returns a single MCP implementation record.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-implementation" method="get" path="/v1/mcp-servers/{mcpId}/implementations/{implementationId}" -->
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

    res, err := s.CatalogMCPImplementations.GetMcpImplementation(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", "7703e2f2-f9c0-47e3-9146-5fa50486b871")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPImplementation != nil {
        switch res.CatalogMCPImplementation.Type {
            case components.CatalogMCPImplementationTypeCatalogMCPGatewayImplementation:
                // res.CatalogMCPImplementation.CatalogMCPGatewayImplementation is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `mcpID`                                                  | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP.                        | a0119846-f179-4d9f-a168-d701facce7fb                     |
| `implementationID`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP implementation.         | 7703e2f2-f9c0-47e3-9146-5fa50486b871                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMcpImplementationResponse](../../models/operations/getmcpimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMcpImplementation

Removes a single MCP implementation (link) record by identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mcp-implementation" method="delete" path="/v1/mcp-servers/{mcpId}/implementations/{implementationId}" -->
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

    res, err := s.CatalogMCPImplementations.DeleteMcpImplementation(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", "7703e2f2-f9c0-47e3-9146-5fa50486b871")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `mcpID`                                                  | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP.                        | a0119846-f179-4d9f-a168-d701facce7fb                     |
| `implementationID`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP implementation.         | 7703e2f2-f9c0-47e3-9146-5fa50486b871                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteMcpImplementationResponse](../../models/operations/deletemcpimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |