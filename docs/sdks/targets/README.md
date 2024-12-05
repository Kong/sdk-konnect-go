# Targets
(*Targets*)

## Overview

### Available Operations

* [ListTargetWithUpstream](#listtargetwithupstream) - List all Targets associated with an Upstream
* [CreateTargetWithUpstream](#createtargetwithupstream) - Create a new Target associated with an Upstream
* [DeleteTargetWithUpstream](#deletetargetwithupstream) - Delete a a Target associated with an Upstream
* [GetTargetWithUpstream](#gettargetwithupstream) - Fetch a Target associated with an Upstream
* [UpsertTargetWithUpstream](#upserttargetwithupstream) - Upsert a Target associated with an Upstream

## ListTargetWithUpstream

List all Targets associated with an Upstream

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
    res, err := s.Targets.ListTargetWithUpstream(ctx, operations.ListTargetWithUpstreamRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        UpstreamIDForTarget: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListTargetWithUpstreamRequest](../../models/operations/listtargetwithupstreamrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListTargetWithUpstreamResponse](../../models/operations/listtargetwithupstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTargetWithUpstream

Create a new Target associated with an Upstream

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
    res, err := s.Targets.CreateTargetWithUpstream(ctx, operations.CreateTargetWithUpstreamRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        UpstreamIDForTarget: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
        TargetWithoutParents: components.TargetWithoutParents{
            ID: sdkkonnectgo.String("089292a7-ba3d-4d88-acf0-97b4b2e2621a"),
            Target: sdkkonnectgo.String("203.0.113.42"),
            Weight: sdkkonnectgo.Int64(100),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Target != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateTargetWithUpstreamRequest](../../models/operations/createtargetwithupstreamrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateTargetWithUpstreamResponse](../../models/operations/createtargetwithupstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTargetWithUpstream

Delete a a Target associated with an Upstream using ID or target.

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
    res, err := s.Targets.DeleteTargetWithUpstream(ctx, operations.DeleteTargetWithUpstreamRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        UpstreamIDForTarget: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
        TargetID: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeleteTargetWithUpstreamRequest](../../models/operations/deletetargetwithupstreamrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DeleteTargetWithUpstreamResponse](../../models/operations/deletetargetwithupstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTargetWithUpstream

Get a Target associated with an Upstream using ID or target.

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
    res, err := s.Targets.GetTargetWithUpstream(ctx, operations.GetTargetWithUpstreamRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        UpstreamIDForTarget: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
        TargetID: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Target != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetTargetWithUpstreamRequest](../../models/operations/gettargetwithupstreamrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetTargetWithUpstreamResponse](../../models/operations/gettargetwithupstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertTargetWithUpstream

Create or Update a Target associated with an Upstream using ID or target.

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
    res, err := s.Targets.UpsertTargetWithUpstream(ctx, operations.UpsertTargetWithUpstreamRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        UpstreamIDForTarget: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
        TargetID: "5a078780-5d4c-4aae-984a-bdc6f52113d8",
        TargetWithoutParents: components.TargetWithoutParents{
            ID: sdkkonnectgo.String("089292a7-ba3d-4d88-acf0-97b4b2e2621a"),
            Target: sdkkonnectgo.String("203.0.113.42"),
            Weight: sdkkonnectgo.Int64(100),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Target != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpsertTargetWithUpstreamRequest](../../models/operations/upserttargetwithupstreamrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpsertTargetWithUpstreamResponse](../../models/operations/upserttargetwithupstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |