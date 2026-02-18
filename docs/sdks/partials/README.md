# Partials

## Overview

Some entities in Kong Gateway share common configuration settings that often need to be repeated. For example, multiple plugins that connect to Redis may require the same connection settings. Without Partials, you would need to replicate this configuration across all plugins. If the settings change, you would need to update each plugin individually.

### Available Operations

* [ListPartial](#listpartial) - List all Partials
* [CreatePartial](#createpartial) - Create a new Partial
* [DeletePartial](#deletepartial) - Delete a Partial
* [GetPartial](#getpartial) - Get a Partial
* [UpsertPartial](#upsertpartial) - Upsert a Partial

## ListPartial

List all Partials

### Example Usage

<!-- UsageSnippet language="go" operationID="list-partial" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/partials" -->
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

    res, err := s.Partials.ListPartial(ctx, operations.ListPartialRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListPartialRequest](../../models/operations/listpartialrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListPartialResponse](../../models/operations/listpartialresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreatePartial

Create a new Partial

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-partial" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/partials" example="DuplicateApiKey" -->
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

    res, err := s.Partials.CreatePartial(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreatePartialRedisCe(
        components.PartialRedisCe{
            Config: components.PartialRedisCeConfig{
                Host: sdkkonnectgo.Pointer("localhost"),
                Password: sdkkonnectgo.Pointer("password"),
                ServerName: sdkkonnectgo.Pointer("redis"),
                Username: sdkkonnectgo.Pointer("username"),
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Partial != nil {
        switch res.Partial.Type {
            case components.PartialTypeRedisCe:
                // res.Partial.PartialRedisCe is populated
            case components.PartialTypeRedisEe:
                // res.Partial.PartialRedisEe is populated
        }

    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-partial" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/partials" example="InvalidAuthCred" -->
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

    res, err := s.Partials.CreatePartial(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreatePartialRedisEe(
        components.PartialRedisEe{
            Config: components.PartialRedisEeConfig{
                ClusterNodes: []components.ClusterNodes{
                    components.ClusterNodes{
                        IP: sdkkonnectgo.Pointer("192.168.1.10"),
                        Port: sdkkonnectgo.Pointer[int64](6380),
                    },
                },
                Host: sdkkonnectgo.Pointer("localhost"),
                Password: sdkkonnectgo.Pointer("password"),
                Port: sdkkonnectgo.Pointer(components.CreatePortInteger(
                    6379,
                )),
                ReadTimeout: sdkkonnectgo.Pointer[int64](1000),
                SendTimeout: sdkkonnectgo.Pointer[int64](1000),
                SentinelNodes: []components.SentinelNodes{
                    components.SentinelNodes{
                        Host: sdkkonnectgo.Pointer("sentinel1.redis.server"),
                        Port: sdkkonnectgo.Pointer[int64](26379),
                    },
                },
                ServerName: sdkkonnectgo.Pointer("redis-ee"),
                Username: sdkkonnectgo.Pointer("username"),
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Partial != nil {
        switch res.Partial.Type {
            case components.PartialTypeRedisCe:
                // res.Partial.PartialRedisCe is populated
            case components.PartialTypeRedisEe:
                // res.Partial.PartialRedisEe is populated
        }

    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-partial" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/partials" example="NoAPIKey" -->
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

    res, err := s.Partials.CreatePartial(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreatePartialRedisCe(
        components.PartialRedisCe{
            Config: components.PartialRedisCeConfig{
                Host: sdkkonnectgo.Pointer("localhost"),
                Password: sdkkonnectgo.Pointer("password"),
                ServerName: sdkkonnectgo.Pointer("redis"),
                Username: sdkkonnectgo.Pointer("username"),
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Partial != nil {
        switch res.Partial.Type {
            case components.PartialTypeRedisCe:
                // res.Partial.PartialRedisCe is populated
            case components.PartialTypeRedisEe:
                // res.Partial.PartialRedisEe is populated
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `partial`                                                                          | [components.Partial](../../models/components/partial.md)                           | :heavy_check_mark:                                                                 | Description of the new Partial for creation                                        |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreatePartialResponse](../../models/operations/createpartialresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeletePartial

Delete a Partial

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-partial" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/partials/{PartialId}" -->
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

    res, err := s.Partials.DeletePartial(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "")
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
| `partialID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Partial to lookup                                                        |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeletePartialResponse](../../models/operations/deletepartialresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetPartial

Get a Partial using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-partial" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/partials/{PartialId}" -->
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

    res, err := s.Partials.GetPartial(ctx, "", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Partial != nil {
        switch res.Partial.Type {
            case components.PartialTypeRedisCe:
                // res.Partial.PartialRedisCe is populated
            case components.PartialTypeRedisEe:
                // res.Partial.PartialRedisEe is populated
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `partialID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Partial to lookup                                                        |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetPartialResponse](../../models/operations/getpartialresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertPartial

Create or Update Partial using ID.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-partial" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/partials/{PartialId}" example="DuplicateApiKey" -->
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

    res, err := s.Partials.UpsertPartial(ctx, operations.UpsertPartialRequest{
        PartialID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Partial: components.CreatePartialRedisEe(
            components.PartialRedisEe{
                Config: components.PartialRedisEeConfig{
                    ClusterNodes: []components.ClusterNodes{
                        components.ClusterNodes{
                            IP: sdkkonnectgo.Pointer("192.168.1.10"),
                            Port: sdkkonnectgo.Pointer[int64](6380),
                        },
                    },
                    Host: sdkkonnectgo.Pointer("localhost"),
                    Password: sdkkonnectgo.Pointer("password"),
                    Port: sdkkonnectgo.Pointer(components.CreatePortInteger(
                        6379,
                    )),
                    ReadTimeout: sdkkonnectgo.Pointer[int64](1000),
                    SendTimeout: sdkkonnectgo.Pointer[int64](1000),
                    SentinelNodes: []components.SentinelNodes{
                        components.SentinelNodes{
                            Host: sdkkonnectgo.Pointer("sentinel1.redis.server"),
                            Port: sdkkonnectgo.Pointer[int64](26379),
                        },
                    },
                    ServerName: sdkkonnectgo.Pointer("redis-ee"),
                    Username: sdkkonnectgo.Pointer("username"),
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Partial != nil {
        switch res.Partial.Type {
            case components.PartialTypeRedisCe:
                // res.Partial.PartialRedisCe is populated
            case components.PartialTypeRedisEe:
                // res.Partial.PartialRedisEe is populated
        }

    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-partial" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/partials/{PartialId}" example="InvalidAuthCred" -->
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

    res, err := s.Partials.UpsertPartial(ctx, operations.UpsertPartialRequest{
        PartialID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Partial: components.CreatePartialRedisCe(
            components.PartialRedisCe{
                Config: components.PartialRedisCeConfig{
                    Host: sdkkonnectgo.Pointer("localhost"),
                    Password: sdkkonnectgo.Pointer("password"),
                    ServerName: sdkkonnectgo.Pointer("redis"),
                    Username: sdkkonnectgo.Pointer("username"),
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Partial != nil {
        switch res.Partial.Type {
            case components.PartialTypeRedisCe:
                // res.Partial.PartialRedisCe is populated
            case components.PartialTypeRedisEe:
                // res.Partial.PartialRedisEe is populated
        }

    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-partial" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/partials/{PartialId}" example="NoAPIKey" -->
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

    res, err := s.Partials.UpsertPartial(ctx, operations.UpsertPartialRequest{
        PartialID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Partial: components.CreatePartialRedisCe(
            components.PartialRedisCe{
                Config: components.PartialRedisCeConfig{
                    Host: sdkkonnectgo.Pointer("localhost"),
                    Password: sdkkonnectgo.Pointer("password"),
                    ServerName: sdkkonnectgo.Pointer("redis"),
                    Username: sdkkonnectgo.Pointer("username"),
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Partial != nil {
        switch res.Partial.Type {
            case components.PartialTypeRedisCe:
                // res.Partial.PartialRedisCe is populated
            case components.PartialTypeRedisEe:
                // res.Partial.PartialRedisEe is populated
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpsertPartialRequest](../../models/operations/upsertpartialrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpsertPartialResponse](../../models/operations/upsertpartialresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |