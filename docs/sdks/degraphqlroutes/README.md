# DegraphqlRoutes

## Overview

### Available Operations

* [ListDegraphqlRoute](#listdegraphqlroute) - List all Degraphql_routes in a workspace
* [GetDegraphqlRoute](#getdegraphqlroute) - Get a Degraphql_route in a workspace
* [ListDegraphqlRouteWithService](#listdegraphqlroutewithservice) - List all Degraphql_routes associated with a Service in a workspace
* [CreateDegraphqlRouteWithService](#createdegraphqlroutewithservice) - Create a new Degraphql_route associated with a Service in a workspace
* [DeleteDegraphqlRouteWithService](#deletedegraphqlroutewithservice) - Delete a a Degraphql_route associated with a Service in a workspace
* [GetDegraphqlRouteWithService](#getdegraphqlroutewithservice) - Get a Degraphql_route associated with a Service in a workspace
* [UpsertDegraphqlRouteWithService](#upsertdegraphqlroutewithservice) - Upsert a Degraphql_route associated with a Service in a workspace

## ListDegraphqlRoute

List all Degraphql_routes in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-degraphql_route" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/degraphql_routes" -->
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

    res, err := s.DegraphqlRoutes.ListDegraphqlRoute(ctx, operations.ListDegraphqlRouteRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListDegraphqlRouteRequest](../../models/operations/listdegraphqlrouterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListDegraphqlRouteResponse](../../models/operations/listdegraphqlrouteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetDegraphqlRoute

Get a Degraphql_route using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-degraphql_route" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/degraphql_routes/{Degraphql_routeId}" -->
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

    res, err := s.DegraphqlRoutes.GetDegraphqlRoute(ctx, operations.GetDegraphqlRouteRequest{
        DegraphqlRouteID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DegraphqlRoute != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetDegraphqlRouteRequest](../../models/operations/getdegraphqlrouterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetDegraphqlRouteResponse](../../models/operations/getdegraphqlrouteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListDegraphqlRouteWithService

List all Degraphql_routes associated with a Service in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-degraphql_route-with-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/degraphql/routes" -->
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

    res, err := s.DegraphqlRoutes.ListDegraphqlRouteWithService(ctx, operations.ListDegraphqlRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListDegraphqlRouteWithServiceRequest](../../models/operations/listdegraphqlroutewithservicerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListDegraphqlRouteWithServiceResponse](../../models/operations/listdegraphqlroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDegraphqlRouteWithService

Create a new Degraphql_route associated with a Service in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-degraphql_route-with-service" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/degraphql/routes" -->
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

    res, err := s.DegraphqlRoutes.CreateDegraphqlRouteWithService(ctx, operations.CreateDegraphqlRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        Workspace: "team-payments",
        DegraphqlRouteWithoutParents: components.DegraphqlRouteWithoutParents{
            ID: sdkkonnectgo.Pointer("56c4566c-14cc-4132-9011-4139fcbbe50a"),
            Query: "query{ user { email } }",
            URI: "/users",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DegraphqlRoute != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CreateDegraphqlRouteWithServiceRequest](../../models/operations/createdegraphqlroutewithservicerequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CreateDegraphqlRouteWithServiceResponse](../../models/operations/createdegraphqlroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteDegraphqlRouteWithService

Delete a a Degraphql_route associated with a Service using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-degraphql_route-with-service" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/degraphql/routes/{Degraphql_routeId}" -->
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

    res, err := s.DegraphqlRoutes.DeleteDegraphqlRouteWithService(ctx, operations.DeleteDegraphqlRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        DegraphqlRouteID: "",
        Workspace: "team-payments",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.DeleteDegraphqlRouteWithServiceRequest](../../models/operations/deletedegraphqlroutewithservicerequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.DeleteDegraphqlRouteWithServiceResponse](../../models/operations/deletedegraphqlroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDegraphqlRouteWithService

Get a Degraphql_route associated with a Service using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-degraphql_route-with-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/degraphql/routes/{Degraphql_routeId}" -->
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

    res, err := s.DegraphqlRoutes.GetDegraphqlRouteWithService(ctx, operations.GetDegraphqlRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        DegraphqlRouteID: "",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DegraphqlRoute != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetDegraphqlRouteWithServiceRequest](../../models/operations/getdegraphqlroutewithservicerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetDegraphqlRouteWithServiceResponse](../../models/operations/getdegraphqlroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertDegraphqlRouteWithService

Create or Update a Degraphql_route associated with a Service using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-degraphql_route-with-service" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/degraphql/routes/{Degraphql_routeId}" -->
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

    res, err := s.DegraphqlRoutes.UpsertDegraphqlRouteWithService(ctx, operations.UpsertDegraphqlRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        DegraphqlRouteID: "",
        Workspace: "team-payments",
        DegraphqlRouteWithoutParents: components.DegraphqlRouteWithoutParents{
            ID: sdkkonnectgo.Pointer("56c4566c-14cc-4132-9011-4139fcbbe50a"),
            Query: "query{ user { email } }",
            URI: "/users",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DegraphqlRoute != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpsertDegraphqlRouteWithServiceRequest](../../models/operations/upsertdegraphqlroutewithservicerequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.UpsertDegraphqlRouteWithServiceResponse](../../models/operations/upsertdegraphqlroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |