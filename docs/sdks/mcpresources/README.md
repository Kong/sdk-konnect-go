# MCPResources

## Overview

### Available Operations

* [ListMcpResources](#listmcpresources) - List MCP Resources
* [CreateMcpResource](#createmcpresource) - Create an MCP Resource
* [GetMcpResource](#getmcpresource) - Get an MCP Resource
* [DeleteMcpResource](#deletemcpresource) - Delete an MCP Resource
* [UpdateMcpResource](#updatemcpresource) - Update an MCP Resource

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

    res, err := s.MCPResources.ListMcpResources(ctx, operations.ListMcpResourcesRequest{
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

    res, err := s.MCPResources.CreateMcpResource(ctx, operations.CreateCreateMcpResourceRequestBodyAPIResourcePayload(
        components.APIResourcePayload{
            Name: "my-api-resource",
            DisplayName: "Richard_Bogisich-Haley",
            Labels: map[string]string{
                "env": "test",
            },
            Source: components.CreateAPIResourceSourcePayloadAPICatalog(
                components.MCPResourceSourceAPICatalog{
                    Type: components.MCPResourceSourceAPICatalogTypeAPICatalog,
                    Config: components.MCPResourceSourceAPICatalogConfig{
                        APIID: "158a0fc7-479a-4928-8a64-7656c1824f7f",
                        APIVersionID: "bc0b9df1-67ba-475f-9670-37de3cc2e4f6",
                    },
                },
            ),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Type {
            case components.MCPResourceInfoTypeRawAPIResource:
                // res.MCPResourceInfo.RawAPIResource is populated
            case components.MCPResourceInfoTypeCatalogAPIResource:
                // res.MCPResourceInfo.CatalogAPIResource is populated
            case components.MCPResourceInfoTypeRemoteMcpServerResource:
                // res.MCPResourceInfo.RemoteMcpServerResource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateMcpResourceRequestBody](../../models/operations/createmcpresourcerequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

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

    res, err := s.MCPResources.GetMcpResource(ctx, "1c4cc520-45e3-4087-8b08-6465521d7fc2")
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Type {
            case components.MCPResourceInfoTypeRawAPIResource:
                // res.MCPResourceInfo.RawAPIResource is populated
            case components.MCPResourceInfoTypeCatalogAPIResource:
                // res.MCPResourceInfo.CatalogAPIResource is populated
            case components.MCPResourceInfoTypeRemoteMcpServerResource:
                // res.MCPResourceInfo.RemoteMcpServerResource is populated
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

    res, err := s.MCPResources.DeleteMcpResource(ctx, "3d073fde-7988-464a-81f2-9949151b4314")
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

    res, err := s.MCPResources.UpdateMcpResource(ctx, "fe6ae9ac-feda-46df-81e4-112aab1e4e51", operations.CreateUpdateMcpResourceRequestBodyAPIResourcePayload(
        components.APIResourcePayload{
            Name: "my-api-resource",
            DisplayName: "Daron_Graham",
            Labels: map[string]string{
                "env": "test",
            },
            Source: components.CreateAPIResourceSourcePayloadAPICatalog(
                components.MCPResourceSourceAPICatalog{
                    Type: components.MCPResourceSourceAPICatalogTypeAPICatalog,
                    Config: components.MCPResourceSourceAPICatalogConfig{
                        APIID: "158a0fc7-479a-4928-8a64-7656c1824f7f",
                        APIVersionID: "bc0b9df1-67ba-475f-9670-37de3cc2e4f6",
                    },
                },
            ),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.MCPResourceInfo != nil {
        switch res.MCPResourceInfo.Type {
            case components.MCPResourceInfoTypeRawAPIResource:
                // res.MCPResourceInfo.RawAPIResource is populated
            case components.MCPResourceInfoTypeCatalogAPIResource:
                // res.MCPResourceInfo.CatalogAPIResource is populated
            case components.MCPResourceInfoTypeRemoteMcpServerResource:
                // res.MCPResourceInfo.RemoteMcpServerResource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `resourceID`                                                                                       | `string`                                                                                           | :heavy_check_mark:                                                                                 | The ID of the MCP resource.                                                                        |
| `requestBody`                                                                                      | [operations.UpdateMcpResourceRequestBody](../../models/operations/updatemcpresourcerequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

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