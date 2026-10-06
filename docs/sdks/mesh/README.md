# Mesh

## Overview

### Available Operations

* [ListMeshControlPlanes](#listmeshcontrolplanes) - List Control Planes
* [CreateCp](#createcp) - Create a new control plane
* [GetMeshControlPlane](#getmeshcontrolplane) - Get the control plane
* [DeleteMeshControlPlane](#deletemeshcontrolplane) - Delete the control plane
* [UpdateCp](#updatecp) - Update a whole control plane
* [PatchMeshControlPlane](#patchmeshcontrolplane) - Patch parts of control plane

## ListMeshControlPlanes

Returns an array of control planes objects.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mesh-control-planes" method="get" path="/v1/mesh/control-planes" -->
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

    res, err := s.Mesh.ListMeshControlPlanes(ctx, operations.ListMeshControlPlanesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMeshControlPlanesResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListMeshControlPlanesRequest](../../models/operations/listmeshcontrolplanesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListMeshControlPlanesResponse](../../models/operations/listmeshcontrolplanesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCp

Create a new control plane

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cp" method="post" path="/v1/mesh/control-planes" -->
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

    res, err := s.Mesh.CreateCp(ctx, components.CreateMeshControlPlaneRequest{
        Name: "Test control plane",
        Description: sdkkonnectgo.Pointer("A control plane to handle traffic on development environment."),
        Labels: map[string]*string{
            "test": sdkkonnectgo.Pointer("true"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeshControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.CreateMeshControlPlaneRequest](../../models/components/createmeshcontrolplanerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateCpResponse](../../models/operations/createcpresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMeshControlPlane

Get the control plane

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mesh-control-plane" method="get" path="/v1/mesh/control-planes/{cpId}" -->
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

    res, err := s.Mesh.GetMeshControlPlane(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b")
    if err != nil {
        log.Fatal(err)
    }
    if res.MeshControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `cpID`                                                   | `string`                                                 | :heavy_check_mark:                                       | Id of the Konnect resource                               | bf138ba2-c9b1-4229-b268-04d9d8a6410b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMeshControlPlaneResponse](../../models/operations/getmeshcontrolplaneresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMeshControlPlane

Delete the control plane

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mesh-control-plane" method="delete" path="/v1/mesh/control-planes/{cpId}" -->
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

    res, err := s.Mesh.DeleteMeshControlPlane(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b")
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
| `cpID`                                                   | `string`                                                 | :heavy_check_mark:                                       | Id of the Konnect resource                               | bf138ba2-c9b1-4229-b268-04d9d8a6410b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteMeshControlPlaneResponse](../../models/operations/deletemeshcontrolplaneresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCp

Update a whole control plane

### Example Usage

<!-- UsageSnippet language="go" operationID="update-cp" method="put" path="/v1/mesh/control-planes/{cpId}" -->
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

    res, err := s.Mesh.UpdateCp(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.PutMeshControlPlaneRequest{
        Name: "Test control plane",
        Description: sdkkonnectgo.Pointer("A control plane to handle traffic on development environment."),
        Labels: map[string]string{
            "test": "true",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeshControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `cpID`                                                                                         | `string`                                                                                       | :heavy_check_mark:                                                                             | Id of the Konnect resource                                                                     | bf138ba2-c9b1-4229-b268-04d9d8a6410b                                                           |
| `putMeshControlPlaneRequest`                                                                   | [components.PutMeshControlPlaneRequest](../../models/components/putmeshcontrolplanerequest.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.UpdateCpResponse](../../models/operations/updatecpresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchMeshControlPlane

Update an individual control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="patch-mesh-control-plane" method="patch" path="/v1/mesh/control-planes/{cpId}" -->
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

    res, err := s.Mesh.PatchMeshControlPlane(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.UpdateMeshControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test control plane"),
        Description: sdkkonnectgo.Pointer("A control plane to handle traffic on development environment."),
        Labels: map[string]*string{
            "test": sdkkonnectgo.Pointer("true"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeshControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `cpID`                                                                                               | `string`                                                                                             | :heavy_check_mark:                                                                                   | Id of the Konnect resource                                                                           | bf138ba2-c9b1-4229-b268-04d9d8a6410b                                                                 |
| `updateMeshControlPlaneRequest`                                                                      | [components.UpdateMeshControlPlaneRequest](../../models/components/updatemeshcontrolplanerequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.PatchMeshControlPlaneResponse](../../models/operations/patchmeshcontrolplaneresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |