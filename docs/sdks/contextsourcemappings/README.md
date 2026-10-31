# ContextSourceMappings

## Overview

### Available Operations

* [CreateContextInterfaceSourceMapping](#createcontextinterfacesourcemapping) - Create Context Interface Source Mapping
* [ListContextInterfaceSourceMappings](#listcontextinterfacesourcemappings) - List Context Interface Source Mappings
* [DeleteContextInterfaceSourceMapping](#deletecontextinterfacesourcemapping) - Delete Context Interface Source Mapping
* [GetContextInterfaceSourceMapping](#getcontextinterfacesourcemapping) - Get Context Interface Source Mapping

## CreateContextInterfaceSourceMapping

Create a mapping between a Context Interface and a Context Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-context-interface-source-mapping" method="post" path="/v1/context-interfaces/{interfaceId}/context-source-mappings" -->
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

    res, err := s.ContextSourceMappings.CreateContextInterfaceSourceMapping(ctx, "cadc8cb9-271a-48ae-8e52-15826ba72f3e", components.CreateMCPServerMCPResourceMappingRequest{
        ContextSourceID: "a60eceeb-dd7c-4884-8b26-806489c2f704",
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
| `interfaceID`                                                                                                              | `string`                                                                                                                   | :heavy_check_mark:                                                                                                         | The ID of the Context Interface.                                                                                           |
| `createMCPServerMCPResourceMappingRequest`                                                                                 | [components.CreateMCPServerMCPResourceMappingRequest](../../models/components/createmcpservermcpresourcemappingrequest.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.CreateContextInterfaceSourceMappingResponse](../../models/operations/createcontextinterfacesourcemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListContextInterfaceSourceMappings

Returns a list of Context Interface Source Mappings for the given Context Interface.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-context-interface-source-mappings" method="get" path="/v1/context-interfaces/{interfaceId}/context-source-mappings" -->
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

    res, err := s.ContextSourceMappings.ListContextInterfaceSourceMappings(ctx, operations.ListContextInterfaceSourceMappingsRequest{
        InterfaceID: "e4cb8eae-2c8c-438c-bbf4-2ec8ee7d3876",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListContextInterfaceSourceMappingsRequest](../../models/operations/listcontextinterfacesourcemappingsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.ListContextInterfaceSourceMappingsResponse](../../models/operations/listcontextinterfacesourcemappingsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteContextInterfaceSourceMapping

Delete a mapping between a Context Interface and a Context Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-context-interface-source-mapping" method="delete" path="/v1/context-interfaces/{interfaceId}/context-source-mappings/{mappingId}" -->
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

    res, err := s.ContextSourceMappings.DeleteContextInterfaceSourceMapping(ctx, "6615204a-8dbc-435a-ac96-5dee444ca291", "6ca7397f-801e-4dd2-ad61-bb05b6451a16")
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
| `interfaceID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface.                         |
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface Source Mapping.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteContextInterfaceSourceMappingResponse](../../models/operations/deletecontextinterfacesourcemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetContextInterfaceSourceMapping

Retrieve a specific Context Interface Source Mapping by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-context-interface-source-mapping" method="get" path="/v1/context-interfaces/{interfaceId}/context-source-mappings/{mappingId}" -->
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

    res, err := s.ContextSourceMappings.GetContextInterfaceSourceMapping(ctx, "0be381db-1a3f-469c-ae75-c785fed8e7b8", "a31e8a2d-5885-437d-a862-86c2a5fdfc97")
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
| `interfaceID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface.                         |
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface Source Mapping.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContextInterfaceSourceMappingResponse](../../models/operations/getcontextinterfacesourcemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |