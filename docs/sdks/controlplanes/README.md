# ControlPlanes
(*ControlPlanes*)

## Overview

### Available Operations

* [List](#list) - List Control Planes
* [Create](#create) - Create Control Plane
* [Get](#get) - Fetch Control Plane
* [Update](#update) - Update Control Plane
* [Delete](#delete) - Delete Control Plane

## List

Returns an array of control plane objects containing information about the Konnect Control Planes.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"github.com/Kong/sdk-konnect-go-internal/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlanes.List(ctx, operations.ListControlPlanesRequest{
        FilterNameEq: sdkkonnectgointernal.String("test"),
        FilterName: sdkkonnectgointernal.String("test"),
        FilterNameContains: sdkkonnectgointernal.String("test"),
        FilterNameNeq: sdkkonnectgointernal.String("test"),
        FilterIDEq: sdkkonnectgointernal.String("7f9fd312-a987-4628-b4c5-bb4f4fddd5f7"),
        FilterID: sdkkonnectgointernal.String("7f9fd312-a987-4628-b4c5-bb4f4fddd5f7"),
        FilterIDOeq: sdkkonnectgointernal.String("some-value,some-other-value"),
        FilterClusterTypeEq: sdkkonnectgointernal.String("CLUSTER_TYPE_CONTROL_PLANE"),
        FilterClusterType: sdkkonnectgointernal.String("CLUSTER_TYPE_CONTROL_PLANE"),
        FilterClusterTypeNeq: sdkkonnectgointernal.String("test"),
        Labels: sdkkonnectgointernal.String("key:value,existCheck"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListControlPlanesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListControlPlanesRequest](../../models/operations/listcontrolplanesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListControlPlanesResponse](../../models/operations/listcontrolplanesresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## Create

Create a control plane in the Konnect Organization.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlanes.Create(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgointernal.String("A test control plane for exploration."),
        ClusterType: components.ClusterTypeClusterTypeK8SIngressController.ToPointer(),
        CloudGateway: sdkkonnectgointernal.Bool(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateControlPlaneRequest](../../models/components/createcontrolplanerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateControlPlaneResponse](../../models/operations/createcontrolplaneresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.ConflictError       | 409                           | application/problem+json      |
| sdkerrors.InternalServerError | 500                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## Get

Returns information about a team from a given team ID.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlanes.Get(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The control plane ID                                     | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetControlPlaneResponse](../../models/operations/getcontrolplaneresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## Update

Update an individual control plane.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlanes.Update(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgointernal.String("Test Control Plane"),
        Description: sdkkonnectgointernal.String("A test control plane for exploration."),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "development",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | The control plane ID                                                                         | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                         |
| `updateControlPlaneRequest`                                                                  | [components.UpdateControlPlaneRequest](../../models/components/updatecontrolplanerequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.UpdateControlPlaneResponse](../../models/operations/updatecontrolplaneresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.InternalServerError | 500                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete an individual control plane.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ControlPlanes.Delete(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The control plane ID                                     | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteControlPlaneResponse](../../models/operations/deletecontrolplaneresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.InternalServerError | 500                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |