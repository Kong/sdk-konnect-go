# Upstreams
(*Upstreams*)

## Overview

The upstream object represents a virtual hostname and can be used to load balance incoming requests over multiple services (targets). 
<br><br>
An upstream also includes a [health checker](https://developer.konghq.com/gateway/traffic-control/health-checks-circuit-breakers/), which can enable and disable targets based on their ability or inability to serve requests. 
The configuration for the health checker is stored in the upstream object, and applies to all of its targets.

### Available Operations

* [ListUpstream](#listupstream) - List all Upstreams
* [CreateUpstream](#createupstream) - Create a new Upstream
* [DeleteUpstream](#deleteupstream) - Delete an Upstream
* [GetUpstream](#getupstream) - Fetch an Upstream
* [UpsertUpstream](#upsertupstream) - Upsert a Upstream

## ListUpstream

List all Upstreams

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Upstreams.ListUpstream(ctx, operations.ListUpstreamRequest{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListUpstreamRequest](../../models/operations/listupstreamrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListUpstreamResponse](../../models/operations/listupstreamresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateUpstream

Create a new Upstream

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Upstreams.CreateUpstream(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Upstream{
        Algorithm: components.UpstreamAlgorithmRoundRobin.ToPointer(),
        HashFallback: components.HashFallbackNone.ToPointer(),
        HashOn: components.HashOnNone.ToPointer(),
        HashOnCookiePath: sdkkonnectgo.String("/"),
        Healthchecks: &components.Healthchecks{
            Active: &components.Active{
                Concurrency: sdkkonnectgo.Int64(10),
                Healthy: &components.Healthy{
                    HTTPStatuses: []int64{
                        200,
                        302,
                    },
                    Interval: sdkkonnectgo.Float64(0),
                    Successes: sdkkonnectgo.Int64(0),
                },
                HTTPPath: sdkkonnectgo.String("/"),
                HTTPSVerifyCertificate: sdkkonnectgo.Bool(true),
                Timeout: sdkkonnectgo.Float64(1),
                Type: components.UpstreamTypeHTTP.ToPointer(),
                Unhealthy: &components.Unhealthy{
                    HTTPFailures: sdkkonnectgo.Int64(0),
                    HTTPStatuses: []int64{
                        429,
                        404,
                        500,
                        501,
                        502,
                        503,
                        504,
                        505,
                    },
                    Interval: sdkkonnectgo.Float64(0),
                    TCPFailures: sdkkonnectgo.Int64(0),
                    Timeouts: sdkkonnectgo.Int64(0),
                },
            },
            Passive: &components.Passive{
                Healthy: &components.UpstreamHealthy{
                    HTTPStatuses: []int64{
                        200,
                        201,
                        202,
                        203,
                        204,
                        205,
                        206,
                        207,
                        208,
                        226,
                        300,
                        301,
                        302,
                        303,
                        304,
                        305,
                        306,
                        307,
                        308,
                    },
                    Successes: sdkkonnectgo.Int64(0),
                },
                Type: components.UpstreamHealthchecksTypeHTTP.ToPointer(),
                Unhealthy: &components.UpstreamUnhealthy{
                    HTTPFailures: sdkkonnectgo.Int64(0),
                    HTTPStatuses: []int64{
                        429,
                        500,
                        503,
                    },
                    TCPFailures: sdkkonnectgo.Int64(0),
                    Timeouts: sdkkonnectgo.Int64(0),
                },
            },
            Threshold: sdkkonnectgo.Float64(0),
        },
        ID: sdkkonnectgo.String("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
        Name: "api.example.internal",
        Slots: sdkkonnectgo.Int64(10000),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Upstream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `controlPlaneID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `upstream`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [components.Upstream](../../models/components/upstream.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Description of the new Upstream for creation                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | {<br/>"algorithm": "round-robin",<br/>"hash_fallback": "none",<br/>"hash_on": "none",<br/>"hash_on_cookie_path": "/",<br/>"healthchecks": {<br/>"active": {<br/>"concurrency": 10,<br/>"healthy": {<br/>"http_statuses": [<br/>200,<br/>302<br/>],<br/>"interval": 0,<br/>"successes": 0<br/>},<br/>"http_path": "/",<br/>"https_verify_certificate": true,<br/>"timeout": 1,<br/>"type": "http",<br/>"unhealthy": {<br/>"http_failures": 0,<br/>"http_statuses": [<br/>429,<br/>404,<br/>500,<br/>501,<br/>502,<br/>503,<br/>504,<br/>505<br/>],<br/>"interval": 0,<br/>"tcp_failures": 0,<br/>"timeouts": 0<br/>}<br/>},<br/>"passive": {<br/>"healthy": {<br/>"http_statuses": [<br/>200,<br/>201,<br/>202,<br/>203,<br/>204,<br/>205,<br/>206,<br/>207,<br/>208,<br/>226,<br/>300,<br/>301,<br/>302,<br/>303,<br/>304,<br/>305,<br/>306,<br/>307,<br/>308<br/>],<br/>"successes": 0<br/>},<br/>"type": "http",<br/>"unhealthy": {<br/>"http_failures": 0,<br/>"http_statuses": [<br/>429,<br/>500,<br/>503<br/>],<br/>"tcp_failures": 0,<br/>"timeouts": 0<br/>}<br/>},<br/>"threshold": 0<br/>},<br/>"id": "6eed5e9c-5398-4026-9a4c-d48f18a2431e",<br/>"name": "api.example.internal",<br/>"slots": 10000<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |

### Response

**[*operations.CreateUpstreamResponse](../../models/operations/createupstreamresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteUpstream

Delete an Upstream

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Upstreams.DeleteUpstream(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "426d620c-7058-4ae6-aacc-f85a3204a2c5")
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
| `upstreamID`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Upstream to lookup                                                       | 426d620c-7058-4ae6-aacc-f85a3204a2c5                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteUpstreamResponse](../../models/operations/deleteupstreamresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetUpstream

Get an Upstream using ID or name.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Upstreams.GetUpstream(ctx, "426d620c-7058-4ae6-aacc-f85a3204a2c5", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Upstream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `upstreamID`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Upstream to lookup                                                       | 426d620c-7058-4ae6-aacc-f85a3204a2c5                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetUpstreamResponse](../../models/operations/getupstreamresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertUpstream

Create or Update Upstream using ID or name.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Upstreams.UpsertUpstream(ctx, operations.UpsertUpstreamRequest{
        UpstreamID: "426d620c-7058-4ae6-aacc-f85a3204a2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Upstream: components.Upstream{
            Algorithm: components.UpstreamAlgorithmRoundRobin.ToPointer(),
            HashFallback: components.HashFallbackNone.ToPointer(),
            HashOn: components.HashOnNone.ToPointer(),
            HashOnCookiePath: sdkkonnectgo.String("/"),
            Healthchecks: &components.Healthchecks{
                Active: &components.Active{
                    Concurrency: sdkkonnectgo.Int64(10),
                    Healthy: &components.Healthy{
                        HTTPStatuses: []int64{
                            200,
                            302,
                        },
                        Interval: sdkkonnectgo.Float64(0),
                        Successes: sdkkonnectgo.Int64(0),
                    },
                    HTTPPath: sdkkonnectgo.String("/"),
                    HTTPSVerifyCertificate: sdkkonnectgo.Bool(true),
                    Timeout: sdkkonnectgo.Float64(1),
                    Type: components.UpstreamTypeHTTP.ToPointer(),
                    Unhealthy: &components.Unhealthy{
                        HTTPFailures: sdkkonnectgo.Int64(0),
                        HTTPStatuses: []int64{
                            429,
                            404,
                            500,
                            501,
                            502,
                            503,
                            504,
                            505,
                        },
                        Interval: sdkkonnectgo.Float64(0),
                        TCPFailures: sdkkonnectgo.Int64(0),
                        Timeouts: sdkkonnectgo.Int64(0),
                    },
                },
                Passive: &components.Passive{
                    Healthy: &components.UpstreamHealthy{
                        HTTPStatuses: []int64{
                            200,
                            201,
                            202,
                            203,
                            204,
                            205,
                            206,
                            207,
                            208,
                            226,
                            300,
                            301,
                            302,
                            303,
                            304,
                            305,
                            306,
                            307,
                            308,
                        },
                        Successes: sdkkonnectgo.Int64(0),
                    },
                    Type: components.UpstreamHealthchecksTypeHTTP.ToPointer(),
                    Unhealthy: &components.UpstreamUnhealthy{
                        HTTPFailures: sdkkonnectgo.Int64(0),
                        HTTPStatuses: []int64{
                            429,
                            500,
                            503,
                        },
                        TCPFailures: sdkkonnectgo.Int64(0),
                        Timeouts: sdkkonnectgo.Int64(0),
                    },
                },
                Threshold: sdkkonnectgo.Float64(0),
            },
            ID: sdkkonnectgo.String("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
            Name: "api.example.internal",
            Slots: sdkkonnectgo.Int64(10000),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Upstream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpsertUpstreamRequest](../../models/operations/upsertupstreamrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpsertUpstreamResponse](../../models/operations/upsertupstreamresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |