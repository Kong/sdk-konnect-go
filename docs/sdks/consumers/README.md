# Consumers

## Overview

The consumer object represents a consumer - or a user - of a service.
You can either rely on Kong Gateway as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong Gateway and your existing primary datastore.


### Available Operations

* [SearchConsumersInWorkspace](#searchconsumersinworkspace) - Search Consumers in a workspace
* [ListConsumerInWorkspace](#listconsumerinworkspace) - List all Consumers in a workspace
* [CreateConsumerInWorkspace](#createconsumerinworkspace) - Create a new Consumer in a workspace
* [DeleteConsumerInWorkspace](#deleteconsumerinworkspace) - Delete a Consumer in a workspace
* [GetConsumerInWorkspace](#getconsumerinworkspace) - Get a Consumer in a workspace
* [UpsertConsumerInWorkspace](#upsertconsumerinworkspace) - Upsert a Consumer in a workspace
* [RemoveConsumerFromAllConsumerGroupsInWorkspace](#removeconsumerfromallconsumergroupsinworkspace) - Remove consumer from all consumer groups in a workspace
* [ListConsumerGroupsForConsumerInWorkspace](#listconsumergroupsforconsumerinworkspace) - List all Consumer Groups a Consumer belongs to in a workspace
* [AddConsumerToSpecificConsumerGroupInWorkspace](#addconsumertospecificconsumergroupinworkspace) - Add consumer to a specific consumer group in a workspace
* [RemoveConsumerFromConsumerGroupInWorkspace](#removeconsumerfromconsumergroupinworkspace) - Remove consumer from consumer group in a workspace
* [ListConsumer](#listconsumer) - List all Consumers
* [CreateConsumer](#createconsumer) - Create a new Consumer
* [DeleteConsumer](#deleteconsumer) - Delete a Consumer
* [GetConsumer](#getconsumer) - Get a Consumer
* [UpsertConsumer](#upsertconsumer) - Upsert a Consumer
* [RemoveConsumerFromAllConsumerGroups](#removeconsumerfromallconsumergroups) - Remove consumer from all consumer groups
* [ListConsumerGroupsForConsumer](#listconsumergroupsforconsumer) - List all Consumer Groups a Consumer belongs to
* [AddConsumerToSpecificConsumerGroup](#addconsumertospecificconsumergroup) - Add consumer to a specific consumer group
* [RemoveConsumerFromConsumerGroup](#removeconsumerfromconsumergroup) - Remove consumer from consumer group

## SearchConsumersInWorkspace

Search Consumers in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="search-consumers-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/search" -->
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

    res, err := s.Consumers.SearchConsumersInWorkspace(ctx, operations.SearchConsumersInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.SearchConsumersInWorkspaceRequest](../../models/operations/searchconsumersinworkspacerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.SearchConsumersInWorkspaceResponse](../../models/operations/searchconsumersinworkspaceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListConsumerInWorkspace

List all Consumers in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumer-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers" -->
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

    res, err := s.Consumers.ListConsumerInWorkspace(ctx, operations.ListConsumerInWorkspaceRequest{
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
| `request`                                                                                              | [operations.ListConsumerInWorkspaceRequest](../../models/operations/listconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListConsumerInWorkspaceResponse](../../models/operations/listconsumerinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateConsumerInWorkspace

Create a new Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-consumer-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers" -->
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

    res, err := s.Consumers.CreateConsumerInWorkspace(ctx, operations.CreateConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Consumer: components.Consumer{
            CustomID: sdkkonnectgo.Pointer("4200"),
            ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
            Tags: []string{
                "silver-tier",
            },
            Username: sdkkonnectgo.Pointer("bob-the-builder"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateConsumerInWorkspaceRequest](../../models/operations/createconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateConsumerInWorkspaceResponse](../../models/operations/createconsumerinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteConsumerInWorkspace

Delete a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-consumer-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" -->
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

    res, err := s.Consumers.DeleteConsumerInWorkspace(ctx, operations.DeleteConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
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
| `request`                                                                                                  | [operations.DeleteConsumerInWorkspaceRequest](../../models/operations/deleteconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.DeleteConsumerInWorkspaceResponse](../../models/operations/deleteconsumerinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetConsumerInWorkspace

Get a Consumer using ID or username in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-consumer-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" -->
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

    res, err := s.Consumers.GetConsumerInWorkspace(ctx, operations.GetConsumerInWorkspaceRequest{
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetConsumerInWorkspaceRequest](../../models/operations/getconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetConsumerInWorkspaceResponse](../../models/operations/getconsumerinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertConsumerInWorkspace

Create or Update Consumer using ID or username in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-consumer-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" -->
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

    res, err := s.Consumers.UpsertConsumerInWorkspace(ctx, operations.UpsertConsumerInWorkspaceRequest{
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Consumer: components.Consumer{
            CustomID: sdkkonnectgo.Pointer("4200"),
            ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
            Tags: []string{
                "silver-tier",
            },
            Username: sdkkonnectgo.Pointer("bob-the-builder"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpsertConsumerInWorkspaceRequest](../../models/operations/upsertconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpsertConsumerInWorkspaceResponse](../../models/operations/upsertconsumerinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## RemoveConsumerFromAllConsumerGroupsInWorkspace

Removes a consumer from all Consumer Groups. This operation does not delete the consumer group in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-consumer-from-all-consumer-groups-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups" -->
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

    res, err := s.Consumers.RemoveConsumerFromAllConsumerGroupsInWorkspace(ctx, operations.RemoveConsumerFromAllConsumerGroupsInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.RemoveConsumerFromAllConsumerGroupsInWorkspaceRequest](../../models/operations/removeconsumerfromallconsumergroupsinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.RemoveConsumerFromAllConsumerGroupsInWorkspaceResponse](../../models/operations/removeconsumerfromallconsumergroupsinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConsumerGroupsForConsumerInWorkspace

List all Consumer Groups a Consumer belongs to in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumer-groups-for-consumer-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups" -->
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

    res, err := s.Consumers.ListConsumerGroupsForConsumerInWorkspace(ctx, operations.ListConsumerGroupsForConsumerInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.ListConsumerGroupsForConsumerInWorkspaceRequest](../../models/operations/listconsumergroupsforconsumerinworkspacerequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.ListConsumerGroupsForConsumerInWorkspaceResponse](../../models/operations/listconsumergroupsforconsumerinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddConsumerToSpecificConsumerGroupInWorkspace

Add a consumer to a consumer group in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="add-consumer-to-specific-consumer-group-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups" -->
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

    res, err := s.Consumers.AddConsumerToSpecificConsumerGroupInWorkspace(ctx, operations.AddConsumerToSpecificConsumerGroupInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        Workspace: "team-payments",
        RequestBody: &operations.AddConsumerToSpecificConsumerGroupInWorkspaceRequestBody{
            Group: sdkkonnectgo.Pointer("fedee695-2ae2-4e45-877a-776d9b2fc793"),
        },
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.AddConsumerToSpecificConsumerGroupInWorkspaceRequest](../../models/operations/addconsumertospecificconsumergroupinworkspacerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.AddConsumerToSpecificConsumerGroupInWorkspaceResponse](../../models/operations/addconsumertospecificconsumergroupinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveConsumerFromConsumerGroupInWorkspace

Removes a consumer from a Consumer Group. This operation does not delete the consumer group in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-consumer-from-consumer-group-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups/{ConsumerGroupId}" -->
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

    res, err := s.Consumers.RemoveConsumerFromConsumerGroupInWorkspace(ctx, operations.RemoveConsumerFromConsumerGroupInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ConsumerGroupID: "",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.RemoveConsumerFromConsumerGroupInWorkspaceRequest](../../models/operations/removeconsumerfromconsumergroupinworkspacerequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.RemoveConsumerFromConsumerGroupInWorkspaceResponse](../../models/operations/removeconsumerfromconsumergroupinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConsumer

List all Consumers

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers" -->
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

    res, err := s.Consumers.ListConsumer(ctx, operations.ListConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        CustomID: sdkkonnectgo.Pointer("my-custom-id"),
        FilterNameContains: sdkkonnectgo.Pointer("john"),
        FilterNameEq: sdkkonnectgo.Pointer("john"),
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
| `request`                                                                        | [operations.ListConsumerRequest](../../models/operations/listconsumerrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListConsumerResponse](../../models/operations/listconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateConsumer

Create a new Consumer

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers" example="DuplicateApiKey" -->
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

    res, err := s.Consumers.CreateConsumer(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Consumer{
        CustomID: sdkkonnectgo.Pointer("4200"),
        ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
        Tags: []string{
            "silver-tier",
        },
        Username: sdkkonnectgo.Pointer("bob-the-builder"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers" example="InvalidAuthCred" -->
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

    res, err := s.Consumers.CreateConsumer(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Consumer{
        CustomID: sdkkonnectgo.Pointer("4200"),
        ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
        Tags: []string{
            "silver-tier",
        },
        Username: sdkkonnectgo.Pointer("bob-the-builder"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers" example="NoAPIKey" -->
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

    res, err := s.Consumers.CreateConsumer(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Consumer{
        CustomID: sdkkonnectgo.Pointer("4200"),
        ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
        Tags: []string{
            "silver-tier",
        },
        Username: sdkkonnectgo.Pointer("bob-the-builder"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     | Example                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |                                                                                                                                 |
| `controlPlaneID`                                                                                                                | `string`                                                                                                                        | :heavy_check_mark:                                                                                                              | The UUID of your control plane. This variable is available in the Konnect manager.                                              | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                            |
| `consumer`                                                                                                                      | [components.Consumer](../../models/components/consumer.md)                                                                      | :heavy_check_mark:                                                                                                              | Description of the new Consumer for creation                                                                                    | {<br/>"custom_id": "4200",<br/>"id": "8a388226-80e8-4027-a486-25e4f7db5d21",<br/>"tags": [<br/>"silver-tier"<br/>],<br/>"username": "bob-the-builder"<br/>} |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |                                                                                                                                 |

### Response

**[*operations.CreateConsumerResponse](../../models/operations/createconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteConsumer

Delete a Consumer

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-consumer" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}" -->
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

    res, err := s.Consumers.DeleteConsumer(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "c1059869-6fa7-4329-a5f5-5946d14ca2c5")
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
| `consumerID`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer to lookup                                                       | c1059869-6fa7-4329-a5f5-5946d14ca2c5                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteConsumerResponse](../../models/operations/deleteconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetConsumer

Get a Consumer using ID or username.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}" -->
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

    res, err := s.Consumers.GetConsumer(ctx, "c1059869-6fa7-4329-a5f5-5946d14ca2c5", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `consumerID`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer to lookup                                                       | c1059869-6fa7-4329-a5f5-5946d14ca2c5                                               |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetConsumerResponse](../../models/operations/getconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertConsumer

Create or Update Consumer using ID or username.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}" example="DuplicateApiKey" -->
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

    res, err := s.Consumers.UpsertConsumer(ctx, operations.UpsertConsumerRequest{
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Consumer: components.Consumer{
            CustomID: sdkkonnectgo.Pointer("4200"),
            ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
            Tags: []string{
                "silver-tier",
            },
            Username: sdkkonnectgo.Pointer("bob-the-builder"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}" example="InvalidAuthCred" -->
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

    res, err := s.Consumers.UpsertConsumer(ctx, operations.UpsertConsumerRequest{
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Consumer: components.Consumer{
            CustomID: sdkkonnectgo.Pointer("4200"),
            ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
            Tags: []string{
                "silver-tier",
            },
            Username: sdkkonnectgo.Pointer("bob-the-builder"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}" example="NoAPIKey" -->
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

    res, err := s.Consumers.UpsertConsumer(ctx, operations.UpsertConsumerRequest{
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Consumer: components.Consumer{
            CustomID: sdkkonnectgo.Pointer("4200"),
            ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
            Tags: []string{
                "silver-tier",
            },
            Username: sdkkonnectgo.Pointer("bob-the-builder"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Consumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpsertConsumerRequest](../../models/operations/upsertconsumerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpsertConsumerResponse](../../models/operations/upsertconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## RemoveConsumerFromAllConsumerGroups

Removes a consumer from all Consumer Groups. This operation does not delete the consumer group.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-consumer-from-all-consumer-groups" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}/consumer_groups" -->
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

    res, err := s.Consumers.RemoveConsumerFromAllConsumerGroups(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "c1059869-6fa7-4329-a5f5-5946d14ca2c5")
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
| `consumerID`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer to lookup                                                       | c1059869-6fa7-4329-a5f5-5946d14ca2c5                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.RemoveConsumerFromAllConsumerGroupsResponse](../../models/operations/removeconsumerfromallconsumergroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConsumerGroupsForConsumer

List all Consumer Groups a Consumer belongs to

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumer-groups-for-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}/consumer_groups" -->
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

    res, err := s.Consumers.ListConsumerGroupsForConsumer(ctx, operations.ListConsumerGroupsForConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListConsumerGroupsForConsumerRequest](../../models/operations/listconsumergroupsforconsumerrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListConsumerGroupsForConsumerResponse](../../models/operations/listconsumergroupsforconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddConsumerToSpecificConsumerGroup

Add a consumer to a consumer group

### Example Usage

<!-- UsageSnippet language="go" operationID="add-consumer-to-specific-consumer-group" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}/consumer_groups" -->
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

    res, err := s.Consumers.AddConsumerToSpecificConsumerGroup(ctx, operations.AddConsumerToSpecificConsumerGroupRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        RequestBody: &operations.AddConsumerToSpecificConsumerGroupRequestBody{
            Group: sdkkonnectgo.Pointer("fedee695-2ae2-4e45-877a-776d9b2fc793"),
        },
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.AddConsumerToSpecificConsumerGroupRequest](../../models/operations/addconsumertospecificconsumergrouprequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.AddConsumerToSpecificConsumerGroupResponse](../../models/operations/addconsumertospecificconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveConsumerFromConsumerGroup

Removes a consumer from a Consumer Group. This operation does not delete the consumer group.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-consumer-from-consumer-group" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}/consumer_groups/{ConsumerGroupId}" -->
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

    res, err := s.Consumers.RemoveConsumerFromConsumerGroup(ctx, operations.RemoveConsumerFromConsumerGroupRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ConsumerGroupID: "",
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
| `request`                                                                                                              | [operations.RemoveConsumerFromConsumerGroupRequest](../../models/operations/removeconsumerfromconsumergrouprequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.RemoveConsumerFromConsumerGroupResponse](../../models/operations/removeconsumerfromconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |