# Routes
(*Routes*)

## Overview

Route entities define rules to match client requests. Each route is associated with a service, and a service may have multiple routes associated to it. Every request matching a given route will be proxied to the associated service. You need at least one matching rule that applies to the protocol being matched by the route.
<br><br>
The combination of routes and services, and the separation of concerns between them, offers a powerful routing mechanism with which it is possible to define fine-grained entrypoints in Kong Gateway leading to different upstream services of your infrastructure.
<br><br>
Depending on the protocol, one of the following attributes must be set:
<br>

- `http`: At least one of `methods`, `hosts`, `headers`, or `paths`
- `https`: At least one of `methods`, `hosts`, `headers`, `paths`, or `snis`
- `tcp`: At least one of `sources` or `destinations`
- `tls`: at least one of `sources`, `destinations`, or `snis`
- `tls_passthrough`: set `snis`
- `grpc`: At least one of `hosts`, `headers`, or `paths`
- `grpcs`: At least one of `hosts`, `headers`, `paths`, or `snis`
- `ws`: At least one of `hosts`, `headers`, or `paths`
- `wss`: At least one of `hosts`, `headers`, `paths`, or `snis`
  <br>
  A route can't have both `tls` and `tls_passthrough` protocols at same time.
  <br><br>
  Learn more about the router:
- [Configure routes using expressions](https://docs.konghq.com/gateway/latest/key-concepts/routes/expressions)
- [Router Expressions language reference](https://docs.konghq.com/gateway/latest/reference/router-expressions-language/)


### Available Operations

* [ListRoute](#listroute) - List all Routes
* [CreateRoute](#createroute) - Create a new Route
* [DeleteRoute](#deleteroute) - Delete a Route
* [GetRoute](#getroute) - Fetch a Route
* [UpsertRoute](#upsertroute) - Upsert a Route
* [ListRouteWithService](#listroutewithservice) - List all Routes associated with a Service
* [CreateRouteWithService](#createroutewithservice) - Create a new Route associated with a Service
* [DeleteRouteWithService](#deleteroutewithservice) - Delete a a Route associated with a Service
* [GetRouteWithService](#getroutewithservice) - Fetch a Route associated with a Service
* [UpsertRouteWithService](#upsertroutewithservice) - Upsert a Route associated with a Service

## ListRoute

List all Routes

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.ListRoute(ctx, operations.ListRouteRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.String("tag1,tag2"),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListRouteRequest](../../models/operations/listrouterequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.ListRouteResponse](../../models/operations/listrouteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateRoute

Create a new Route

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.CreateRoute(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.RouteInput{
        Hosts: []string{
            "foo.example.com",
            "foo.example.us",
        },
        ID: sdkkonnectgo.String("56c4566c-14cc-4132-9011-4139fcbbe50a"),
        Name: sdkkonnectgo.String("example-route"),
        Paths: []string{
            "/v1",
            "/v2",
        },
        Service: &components.RouteService{
            ID: sdkkonnectgo.String("bd380f99-659d-415e-b0e7-72ea05df3218"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Route != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                           | Type                                                                                                                                                                                                                | Required                                                                                                                                                                                                            | Description                                                                                                                                                                                                         | Example                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                 |                                                                                                                                                                                                                     |
| `controlPlaneID`                                                                                                                                                                                                    | *string*                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                  | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                                  | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                                |
| `route`                                                                                                                                                                                                             | [components.RouteInput](../../models/components/routeinput.md)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                  | Description of the new Route for creation                                                                                                                                                                           | {<br/>"hosts": [<br/>"foo.example.com",<br/>"foo.example.us"<br/>],<br/>"id": "56c4566c-14cc-4132-9011-4139fcbbe50a",<br/>"name": "example-route",<br/>"paths": [<br/>"/v1",<br/>"/v2"<br/>],<br/>"service": {<br/>"id": "bd380f99-659d-415e-b0e7-72ea05df3218"<br/>}<br/>} |
| `opts`                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                       |                                                                                                                                                                                                                     |

### Response

**[*operations.CreateRouteResponse](../../models/operations/createrouteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteRoute

Delete a Route

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.DeleteRoute(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `routeID`                                                                          | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Route to lookup                                                          | a4326a41-aa12-44e3-93e4-6b6e58bfb9d7                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteRouteResponse](../../models/operations/deleterouteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetRoute

Get a Route using ID or name.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.GetRoute(ctx, "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Route != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `routeID`                                                                          | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Route to lookup                                                          | a4326a41-aa12-44e3-93e4-6b6e58bfb9d7                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetRouteResponse](../../models/operations/getrouteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertRoute

Create or Update Route using ID or name.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.UpsertRoute(ctx, operations.UpsertRouteRequest{
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Route: components.RouteInput{
            Hosts: []string{
                "foo.example.com",
                "foo.example.us",
            },
            ID: sdkkonnectgo.String("56c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: sdkkonnectgo.String("example-route"),
            Paths: []string{
                "/v1",
                "/v2",
            },
            Service: &components.RouteService{
                ID: sdkkonnectgo.String("bd380f99-659d-415e-b0e7-72ea05df3218"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Route != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpsertRouteRequest](../../models/operations/upsertrouterequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpsertRouteResponse](../../models/operations/upsertrouteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListRouteWithService

List all Routes associated with a Service

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.ListRouteWithService(ctx, operations.ListRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        Tags: sdkkonnectgo.String("tag1,tag2"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListRouteWithServiceRequest](../../models/operations/listroutewithservicerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListRouteWithServiceResponse](../../models/operations/listroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateRouteWithService

Create a new Route associated with a Service

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.CreateRouteWithService(ctx, operations.CreateRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        RouteWithoutParents: components.RouteWithoutParents{
            Hosts: []string{
                "foo.example.com",
                "foo.example.us",
            },
            ID: sdkkonnectgo.String("56c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: sdkkonnectgo.String("example-route"),
            Paths: []string{
                "/v1",
                "/v2",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Route != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateRouteWithServiceRequest](../../models/operations/createroutewithservicerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateRouteWithServiceResponse](../../models/operations/createroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteRouteWithService

Delete a a Route associated with a Service using ID or name.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.DeleteRouteWithService(ctx, operations.DeleteRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteRouteWithServiceRequest](../../models/operations/deleteroutewithservicerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.DeleteRouteWithServiceResponse](../../models/operations/deleteroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRouteWithService

Get a Route associated with a Service using ID or name.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.GetRouteWithService(ctx, operations.GetRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Route != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetRouteWithServiceRequest](../../models/operations/getroutewithservicerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetRouteWithServiceResponse](../../models/operations/getroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertRouteWithService

Create or Update a Route associated with a Service using ID or name.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Routes.UpsertRouteWithService(ctx, operations.UpsertRouteWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        RouteID: "a4326a41-aa12-44e3-93e4-6b6e58bfb9d7",
        RouteWithoutParents: components.RouteWithoutParents{
            Hosts: []string{
                "foo.example.com",
                "foo.example.us",
            },
            ID: sdkkonnectgo.String("56c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: sdkkonnectgo.String("example-route"),
            Paths: []string{
                "/v1",
                "/v2",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Route != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpsertRouteWithServiceRequest](../../models/operations/upsertroutewithservicerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpsertRouteWithServiceResponse](../../models/operations/upsertroutewithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |