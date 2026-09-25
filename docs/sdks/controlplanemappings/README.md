# ControlPlaneMappings

## Overview

### Available Operations

* [ListControlPlaneMappings](#listcontrolplanemappings) - Get a list of CPs mapped to Context Interface
* [CreateContextInterfaceControlPlaneMapping](#createcontextinterfacecontrolplanemapping) - Create Context Interface CP Mapping
* [GetControlPlaneMapping](#getcontrolplanemapping) - Get Context Interface Control Plane Mapping
* [PatchContextInterfaceControlPlaneMapping](#patchcontextinterfacecontrolplanemapping) - Patch Context Interface Control Plane Mapping
* [DeleteContextInterfaceControlPlaneMapping](#deletecontextinterfacecontrolplanemapping) - Delete Context Interface CP Mapping

## ListControlPlaneMappings

Get a list of all Control Planes that have been mapped to the given Context Interface.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-control-plane-mappings" method="get" path="/v1/context-interfaces/{interfaceId}/control-plane-mappings" -->
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
        InterfaceID: "2f594555-a2ac-438e-b375-6e7bbbb984de",
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

## CreateContextInterfaceControlPlaneMapping

Create a Context Interface - Control Plane mapping

### Example Usage

<!-- UsageSnippet language="go" operationID="create-context-interface-control-plane-mapping" method="post" path="/v1/context-interfaces/{interfaceId}/control-plane-mappings" -->
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

    res, err := s.ControlPlaneMappings.CreateContextInterfaceControlPlaneMapping(ctx, "5745c1d6-c749-4470-b361-29a6c2b8ac78", components.CreateMCPServerControlPlaneMappingRequest{
        ControlPlaneID: "2bed8e06-002c-4137-82a8-e68c19286226",
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
| `interfaceID`                                                                                                                | `string`                                                                                                                     | :heavy_check_mark:                                                                                                           | The ID of the Context Interface.                                                                                             |
| `createMCPServerControlPlaneMappingRequest`                                                                                  | [components.CreateMCPServerControlPlaneMappingRequest](../../models/components/createmcpservercontrolplanemappingrequest.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CreateContextInterfaceControlPlaneMappingResponse](../../models/operations/createcontextinterfacecontrolplanemappingresponse.md), error**

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

<!-- UsageSnippet language="go" operationID="get-control-plane-mapping" method="get" path="/v1/context-interfaces/{interfaceId}/control-plane-mappings/{mappingId}" -->
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
| `interfaceID`                                            | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface.                         |
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface - Control Plane mapping. |
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

## PatchContextInterfaceControlPlaneMapping

Update a mapping between a Context Interface and a Control Plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-context-interface-control-plane-mapping" method="patch" path="/v1/context-interfaces/{interfaceId}/control-plane-mappings/{mappingId}" -->
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

    res, err := s.ControlPlaneMappings.PatchContextInterfaceControlPlaneMapping(ctx, operations.PatchContextInterfaceControlPlaneMappingRequest{
        InterfaceID: "495ce8f7-8775-4767-829f-166d855dbdc6",
        MappingID: "fa5f4042-d070-450b-beee-8f42edde987a",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.PatchContextInterfaceControlPlaneMappingRequest](../../models/operations/patchcontextinterfacecontrolplanemappingrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.PatchContextInterfaceControlPlaneMappingResponse](../../models/operations/patchcontextinterfacecontrolplanemappingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteContextInterfaceControlPlaneMapping

Delete Context Interface Control Plane Mapping.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-context-interface-control-plane-mapping" method="delete" path="/v1/context-interfaces/{interfaceId}/control-plane-mappings/{mappingId}" -->
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

    res, err := s.ControlPlaneMappings.DeleteContextInterfaceControlPlaneMapping(ctx, "db226339-9181-462f-b47c-29a6b6637076", "34b91af4-ce67-44f2-8b3d-04802c193a4e")
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
| `mappingID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the Context Interface - Control Plane mapping. |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteContextInterfaceControlPlaneMappingResponse](../../models/operations/deletecontextinterfacecontrolplanemappingresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |