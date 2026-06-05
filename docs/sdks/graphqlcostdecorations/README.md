# GraphQLCostDecorations

## Overview

### Available Operations

* [ListGraphqlRateLimitingAdvancedCostInWorkspace](#listgraphqlratelimitingadvancedcostinworkspace) - List all GraphQL Cost Decorations in a workspace
* [GetGraphqlRateLimitingAdvancedCostInWorkspace](#getgraphqlratelimitingadvancedcostinworkspace) - Get a GraphQL Cost Decoration in a workspace
* [ListGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](#listgraphqlratelimitingadvancedcostwithserviceinworkspace) - List all GraphQL Cost Decorations associated with a Service in a workspace
* [CreateGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](#creategraphqlratelimitingadvancedcostwithserviceinworkspace) - Create a new GraphQL Cost Decoration associated with a Service in a workspace
* [DeleteGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](#deletegraphqlratelimitingadvancedcostwithserviceinworkspace) - Delete a a GraphQL Cost Decoration associated with a Service in a workspace
* [GetGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](#getgraphqlratelimitingadvancedcostwithserviceinworkspace) - Get a GraphQL Cost Decoration associated with a Service in a workspace
* [UpsertGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](#upsertgraphqlratelimitingadvancedcostwithserviceinworkspace) - Upsert a GraphQL Cost Decoration associated with a Service in a workspace
* [ListGraphqlRateLimitingAdvancedCost](#listgraphqlratelimitingadvancedcost) - List all GraphQL Cost Decorations
* [GetGraphqlRateLimitingAdvancedCost](#getgraphqlratelimitingadvancedcost) - Get a GraphQL Cost Decoration
* [ListGraphqlRateLimitingAdvancedCostWithService](#listgraphqlratelimitingadvancedcostwithservice) - List all GraphQL Cost Decorations associated with a Service
* [CreateGraphqlRateLimitingAdvancedCostWithService](#creategraphqlratelimitingadvancedcostwithservice) - Create a new GraphQL Cost Decoration associated with a Service
* [DeleteGraphqlRateLimitingAdvancedCostWithService](#deletegraphqlratelimitingadvancedcostwithservice) - Delete a a GraphQL Cost Decoration associated with a Service
* [GetGraphqlRateLimitingAdvancedCostWithService](#getgraphqlratelimitingadvancedcostwithservice) - Get a GraphQL Cost Decoration associated with a Service
* [UpsertGraphqlRateLimitingAdvancedCostWithService](#upsertgraphqlratelimitingadvancedcostwithservice) - Upsert a GraphQL Cost Decoration associated with a Service

## ListGraphqlRateLimitingAdvancedCostInWorkspace

List all GraphQL Cost Decorations in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-graphql-rate-limiting-advanced-cost-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.ListGraphqlRateLimitingAdvancedCostInWorkspace(ctx, operations.ListGraphqlRateLimitingAdvancedCostInWorkspaceRequest{
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.ListGraphqlRateLimitingAdvancedCostInWorkspaceRequest](../../models/operations/listgraphqlratelimitingadvancedcostinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.ListGraphqlRateLimitingAdvancedCostInWorkspaceResponse](../../models/operations/listgraphqlratelimitingadvancedcostinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetGraphqlRateLimitingAdvancedCostInWorkspace

Get a GraphQL Cost Decoration using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-graphql-rate-limiting-advanced-cost-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.GetGraphqlRateLimitingAdvancedCostInWorkspace(ctx, operations.GetGraphqlRateLimitingAdvancedCostInWorkspaceRequest{
        GraphQLCostDecorationID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.GetGraphqlRateLimitingAdvancedCostInWorkspaceRequest](../../models/operations/getgraphqlratelimitingadvancedcostinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.GetGraphqlRateLimitingAdvancedCostInWorkspaceResponse](../../models/operations/getgraphqlratelimitingadvancedcostinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListGraphqlRateLimitingAdvancedCostWithServiceInWorkspace

List all GraphQL Cost Decorations associated with a Service in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-graphql-rate-limiting-advanced-cost-with-service-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.ListGraphqlRateLimitingAdvancedCostWithServiceInWorkspace(ctx, operations.ListGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest{
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

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `request`                                                                                                                                                                  | [operations.ListGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest](../../models/operations/listgraphqlratelimitingadvancedcostwithserviceinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                                         | The request object to use for the request.                                                                                                                                 |
| `opts`                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.ListGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceResponse](../../models/operations/listgraphqlratelimitingadvancedcostwithserviceinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGraphqlRateLimitingAdvancedCostWithServiceInWorkspace

Create a new GraphQL Cost Decoration associated with a Service in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-graphql-rate-limiting-advanced-cost-with-service-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.CreateGraphqlRateLimitingAdvancedCostWithServiceInWorkspace(ctx, operations.CreateGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        Workspace: "team-payments",
        GraphQLCostDecorationWithoutParents: components.GraphQLCostDecorationWithoutParents{
            TypePath: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.CreateGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest](../../models/operations/creategraphqlratelimitingadvancedcostwithserviceinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.CreateGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceResponse](../../models/operations/creategraphqlratelimitingadvancedcostwithserviceinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteGraphqlRateLimitingAdvancedCostWithServiceInWorkspace

Delete a a GraphQL Cost Decoration associated with a Service using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-graphql-rate-limiting-advanced-cost-with-service-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.DeleteGraphqlRateLimitingAdvancedCostWithServiceInWorkspace(ctx, operations.DeleteGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationID: "",
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

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.DeleteGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest](../../models/operations/deletegraphqlratelimitingadvancedcostwithserviceinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.DeleteGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceResponse](../../models/operations/deletegraphqlratelimitingadvancedcostwithserviceinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGraphqlRateLimitingAdvancedCostWithServiceInWorkspace

Get a GraphQL Cost Decoration associated with a Service using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-graphql-rate-limiting-advanced-cost-with-service-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.GetGraphqlRateLimitingAdvancedCostWithServiceInWorkspace(ctx, operations.GetGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationID: "",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.GetGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest](../../models/operations/getgraphqlratelimitingadvancedcostwithserviceinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |
| `opts`                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.GetGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceResponse](../../models/operations/getgraphqlratelimitingadvancedcostwithserviceinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertGraphqlRateLimitingAdvancedCostWithServiceInWorkspace

Create or Update a GraphQL Cost Decoration associated with a Service using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-graphql-rate-limiting-advanced-cost-with-service-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/services/{ServiceId}/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.UpsertGraphqlRateLimitingAdvancedCostWithServiceInWorkspace(ctx, operations.UpsertGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationID: "",
        Workspace: "team-payments",
        GraphQLCostDecorationWithoutParents: components.GraphQLCostDecorationWithoutParents{
            TypePath: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                      | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                          | :heavy_check_mark:                                                                                                                                                             | The context to use for the request.                                                                                                                                            |
| `request`                                                                                                                                                                      | [operations.UpsertGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceRequest](../../models/operations/upsertgraphqlratelimitingadvancedcostwithserviceinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                                             | The request object to use for the request.                                                                                                                                     |
| `opts`                                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                             | The options for this request.                                                                                                                                                  |

### Response

**[*operations.UpsertGraphqlRateLimitingAdvancedCostWithServiceInWorkspaceResponse](../../models/operations/upsertgraphqlratelimitingadvancedcostwithserviceinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGraphqlRateLimitingAdvancedCost

List all GraphQL Cost Decorations

### Example Usage

<!-- UsageSnippet language="go" operationID="list-graphql-rate-limiting-advanced-cost" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.ListGraphqlRateLimitingAdvancedCost(ctx, operations.ListGraphqlRateLimitingAdvancedCostRequest{
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListGraphqlRateLimitingAdvancedCostRequest](../../models/operations/listgraphqlratelimitingadvancedcostrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListGraphqlRateLimitingAdvancedCostResponse](../../models/operations/listgraphqlratelimitingadvancedcostresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetGraphqlRateLimitingAdvancedCost

Get a GraphQL Cost Decoration using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-graphql-rate-limiting-advanced-cost" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.GetGraphqlRateLimitingAdvancedCost(ctx, "", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `graphQLCostDecorationID`                                                          | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the GraphQL Cost Decoration to lookup                                        |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetGraphqlRateLimitingAdvancedCostResponse](../../models/operations/getgraphqlratelimitingadvancedcostresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListGraphqlRateLimitingAdvancedCostWithService

List all GraphQL Cost Decorations associated with a Service

### Example Usage

<!-- UsageSnippet language="go" operationID="list-graphql-rate-limiting-advanced-cost-with-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.ListGraphqlRateLimitingAdvancedCostWithService(ctx, operations.ListGraphqlRateLimitingAdvancedCostWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.ListGraphqlRateLimitingAdvancedCostWithServiceRequest](../../models/operations/listgraphqlratelimitingadvancedcostwithservicerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.ListGraphqlRateLimitingAdvancedCostWithServiceResponse](../../models/operations/listgraphqlratelimitingadvancedcostwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGraphqlRateLimitingAdvancedCostWithService

Create a new GraphQL Cost Decoration associated with a Service

### Example Usage

<!-- UsageSnippet language="go" operationID="create-graphql-rate-limiting-advanced-cost-with-service" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/graphql-rate-limiting-advanced/costs" -->
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

    res, err := s.GraphQLCostDecorations.CreateGraphqlRateLimitingAdvancedCostWithService(ctx, operations.CreateGraphqlRateLimitingAdvancedCostWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationWithoutParents: components.GraphQLCostDecorationWithoutParents{
            TypePath: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.CreateGraphqlRateLimitingAdvancedCostWithServiceRequest](../../models/operations/creategraphqlratelimitingadvancedcostwithservicerequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.CreateGraphqlRateLimitingAdvancedCostWithServiceResponse](../../models/operations/creategraphqlratelimitingadvancedcostwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteGraphqlRateLimitingAdvancedCostWithService

Delete a a GraphQL Cost Decoration associated with a Service using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-graphql-rate-limiting-advanced-cost-with-service" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.DeleteGraphqlRateLimitingAdvancedCostWithService(ctx, operations.DeleteGraphqlRateLimitingAdvancedCostWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationID: "",
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

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.DeleteGraphqlRateLimitingAdvancedCostWithServiceRequest](../../models/operations/deletegraphqlratelimitingadvancedcostwithservicerequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.DeleteGraphqlRateLimitingAdvancedCostWithServiceResponse](../../models/operations/deletegraphqlratelimitingadvancedcostwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGraphqlRateLimitingAdvancedCostWithService

Get a GraphQL Cost Decoration associated with a Service using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-graphql-rate-limiting-advanced-cost-with-service" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.GetGraphqlRateLimitingAdvancedCostWithService(ctx, operations.GetGraphqlRateLimitingAdvancedCostWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationID: "",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.GetGraphqlRateLimitingAdvancedCostWithServiceRequest](../../models/operations/getgraphqlratelimitingadvancedcostwithservicerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.GetGraphqlRateLimitingAdvancedCostWithServiceResponse](../../models/operations/getgraphqlratelimitingadvancedcostwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertGraphqlRateLimitingAdvancedCostWithService

Create or Update a GraphQL Cost Decoration associated with a Service using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-graphql-rate-limiting-advanced-cost-with-service" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/services/{ServiceId}/graphql-rate-limiting-advanced/costs/{GraphQLCostDecorationId}" -->
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

    res, err := s.GraphQLCostDecorations.UpsertGraphqlRateLimitingAdvancedCostWithService(ctx, operations.UpsertGraphqlRateLimitingAdvancedCostWithServiceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ServiceID: "7fca84d6-7d37-4a74-a7b0-93e576089a41",
        GraphQLCostDecorationID: "",
        GraphQLCostDecorationWithoutParents: components.GraphQLCostDecorationWithoutParents{
            TypePath: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GraphQLCostDecoration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.UpsertGraphqlRateLimitingAdvancedCostWithServiceRequest](../../models/operations/upsertgraphqlratelimitingadvancedcostwithservicerequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.UpsertGraphqlRateLimitingAdvancedCostWithServiceResponse](../../models/operations/upsertgraphqlratelimitingadvancedcostwithserviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |