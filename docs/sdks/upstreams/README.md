# Upstreams

## Overview

The upstream object represents a virtual hostname and can be used to load balance incoming requests over multiple services (targets).
<br><br>
An upstream also includes a [health checker](https://developer.konghq.com/gateway/traffic-control/health-checks-circuit-breakers/), which can enable and disable targets based on their ability or inability to serve requests.
The configuration for the health checker is stored in the upstream object, and applies to all of its targets.

### Available Operations

* [ListUpstreamInWorkspace](#listupstreaminworkspace) - List all Upstreams in a workspace
* [CreateUpstreamInWorkspace](#createupstreaminworkspace) - Create a new Upstream in a workspace
* [DeleteUpstreamInWorkspace](#deleteupstreaminworkspace) - Delete an Upstream in a workspace
* [GetUpstreamInWorkspace](#getupstreaminworkspace) - Get an Upstream in a workspace
* [UpsertUpstreamInWorkspace](#upsertupstreaminworkspace) - Upsert a Upstream in a workspace
* [ListUpstream](#listupstream) - List all Upstreams
* [CreateUpstream](#createupstream) - Create a new Upstream
* [DeleteUpstream](#deleteupstream) - Delete an Upstream
* [GetUpstream](#getupstream) - Get an Upstream
* [UpsertUpstream](#upsertupstream) - Upsert a Upstream

## ListUpstreamInWorkspace

List all Upstreams in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-upstream-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/upstreams" -->
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

    res, err := s.Upstreams.ListUpstreamInWorkspace(ctx, operations.ListUpstreamInWorkspaceRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListUpstreamInWorkspaceRequest](../../models/operations/listupstreaminworkspacerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListUpstreamInWorkspaceResponse](../../models/operations/listupstreaminworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateUpstreamInWorkspace

Create a new Upstream in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-upstream-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/upstreams" -->
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

    res, err := s.Upstreams.CreateUpstreamInWorkspace(ctx, operations.CreateUpstreamInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Upstream: components.Upstream{
            Healthchecks: &components.Healthchecks{
                Active: &components.Active{
                    Healthy: &components.Healthy{
                        HTTPStatuses: []int64{
                            200,
                            302,
                        },
                    },
                    Unhealthy: &components.Unhealthy{
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
                    },
                    Unhealthy: &components.UpstreamUnhealthy{
                        HTTPStatuses: []int64{
                            429,
                            500,
                            503,
                        },
                    },
                },
            },
            ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
            Name: "api.example.internal",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateUpstreamInWorkspaceRequest](../../models/operations/createupstreaminworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateUpstreamInWorkspaceResponse](../../models/operations/createupstreaminworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteUpstreamInWorkspace

Delete an Upstream in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-upstream-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/upstreams/{UpstreamId}" -->
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

    res, err := s.Upstreams.DeleteUpstreamInWorkspace(ctx, operations.DeleteUpstreamInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        UpstreamID: "426d620c-7058-4ae6-aacc-f85a3204a2c5",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeleteUpstreamInWorkspaceRequest](../../models/operations/deleteupstreaminworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.DeleteUpstreamInWorkspaceResponse](../../models/operations/deleteupstreaminworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetUpstreamInWorkspace

Get an Upstream using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-upstream-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/upstreams/{UpstreamId}" -->
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

    res, err := s.Upstreams.GetUpstreamInWorkspace(ctx, operations.GetUpstreamInWorkspaceRequest{
        UpstreamID: "426d620c-7058-4ae6-aacc-f85a3204a2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUpstreamInWorkspaceRequest](../../models/operations/getupstreaminworkspacerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetUpstreamInWorkspaceResponse](../../models/operations/getupstreaminworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertUpstreamInWorkspace

Create or Update Upstream using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-upstream-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/upstreams/{UpstreamId}" -->
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

    res, err := s.Upstreams.UpsertUpstreamInWorkspace(ctx, operations.UpsertUpstreamInWorkspaceRequest{
        UpstreamID: "426d620c-7058-4ae6-aacc-f85a3204a2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Upstream: components.Upstream{
            Healthchecks: &components.Healthchecks{
                Active: &components.Active{
                    Healthy: &components.Healthy{
                        HTTPStatuses: []int64{
                            200,
                            302,
                        },
                    },
                    Unhealthy: &components.Unhealthy{
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
                    },
                    Unhealthy: &components.UpstreamUnhealthy{
                        HTTPStatuses: []int64{
                            429,
                            500,
                            503,
                        },
                    },
                },
            },
            ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
            Name: "api.example.internal",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpsertUpstreamInWorkspaceRequest](../../models/operations/upsertupstreaminworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpsertUpstreamInWorkspaceResponse](../../models/operations/upsertupstreaminworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListUpstream

List all Upstreams

### Example Usage

<!-- UsageSnippet language="go" operationID="list-upstream" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams" -->
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

    res, err := s.Upstreams.ListUpstream(ctx, operations.ListUpstreamRequest{
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

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-upstream" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams" example="DuplicateApiKey" -->
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

    res, err := s.Upstreams.CreateUpstream(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Upstream{
        Healthchecks: &components.Healthchecks{
            Active: &components.Active{
                Healthy: &components.Healthy{
                    HTTPStatuses: []int64{
                        200,
                        302,
                    },
                },
                Unhealthy: &components.Unhealthy{
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
                },
                Unhealthy: &components.UpstreamUnhealthy{
                    HTTPStatuses: []int64{
                        429,
                        500,
                        503,
                    },
                },
            },
        },
        ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
        Name: "api.example.internal",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Upstream != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-upstream" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams" example="InvalidAuthCred" -->
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

    res, err := s.Upstreams.CreateUpstream(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Upstream{
        Healthchecks: &components.Healthchecks{
            Active: &components.Active{
                Healthy: &components.Healthy{
                    HTTPStatuses: []int64{
                        200,
                        302,
                    },
                },
                Unhealthy: &components.Unhealthy{
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
                },
                Unhealthy: &components.UpstreamUnhealthy{
                    HTTPStatuses: []int64{
                        429,
                        500,
                        503,
                    },
                },
            },
        },
        ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
        Name: "api.example.internal",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Upstream != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-upstream" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams" example="NoAPIKey" -->
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

    res, err := s.Upstreams.CreateUpstream(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Upstream{
        Healthchecks: &components.Healthchecks{
            Active: &components.Active{
                Healthy: &components.Healthy{
                    HTTPStatuses: []int64{
                        200,
                        302,
                    },
                },
                Unhealthy: &components.Unhealthy{
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
                },
                Unhealthy: &components.UpstreamUnhealthy{
                    HTTPStatuses: []int64{
                        429,
                        500,
                        503,
                    },
                },
            },
        },
        ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
        Name: "api.example.internal",
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
| `controlPlaneID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | `string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
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

<!-- UsageSnippet language="go" operationID="delete-upstream" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams/{UpstreamId}" -->
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
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `upstreamID`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the Upstream to lookup                                                       | 426d620c-7058-4ae6-aacc-f85a3204a2c5                                               |
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

<!-- UsageSnippet language="go" operationID="get-upstream" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams/{UpstreamId}" -->
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
| `upstreamID`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the Upstream to lookup                                                       | 426d620c-7058-4ae6-aacc-f85a3204a2c5                                               |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
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

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-upstream" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams/{UpstreamId}" example="DuplicateApiKey" -->
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

    res, err := s.Upstreams.UpsertUpstream(ctx, operations.UpsertUpstreamRequest{
        UpstreamID: "426d620c-7058-4ae6-aacc-f85a3204a2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Upstream: components.Upstream{
            Healthchecks: &components.Healthchecks{
                Active: &components.Active{
                    Healthy: &components.Healthy{
                        HTTPStatuses: []int64{
                            200,
                            302,
                        },
                    },
                    Unhealthy: &components.Unhealthy{
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
                    },
                    Unhealthy: &components.UpstreamUnhealthy{
                        HTTPStatuses: []int64{
                            429,
                            500,
                            503,
                        },
                    },
                },
            },
            ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
            Name: "api.example.internal",
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
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-upstream" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams/{UpstreamId}" example="InvalidAuthCred" -->
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

    res, err := s.Upstreams.UpsertUpstream(ctx, operations.UpsertUpstreamRequest{
        UpstreamID: "426d620c-7058-4ae6-aacc-f85a3204a2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Upstream: components.Upstream{
            Healthchecks: &components.Healthchecks{
                Active: &components.Active{
                    Healthy: &components.Healthy{
                        HTTPStatuses: []int64{
                            200,
                            302,
                        },
                    },
                    Unhealthy: &components.Unhealthy{
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
                    },
                    Unhealthy: &components.UpstreamUnhealthy{
                        HTTPStatuses: []int64{
                            429,
                            500,
                            503,
                        },
                    },
                },
            },
            ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
            Name: "api.example.internal",
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
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-upstream" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/upstreams/{UpstreamId}" example="NoAPIKey" -->
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

    res, err := s.Upstreams.UpsertUpstream(ctx, operations.UpsertUpstreamRequest{
        UpstreamID: "426d620c-7058-4ae6-aacc-f85a3204a2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Upstream: components.Upstream{
            Healthchecks: &components.Healthchecks{
                Active: &components.Active{
                    Healthy: &components.Healthy{
                        HTTPStatuses: []int64{
                            200,
                            302,
                        },
                    },
                    Unhealthy: &components.Unhealthy{
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
                    },
                    Unhealthy: &components.UpstreamUnhealthy{
                        HTTPStatuses: []int64{
                            429,
                            500,
                            503,
                        },
                    },
                },
            },
            ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
            Name: "api.example.internal",
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