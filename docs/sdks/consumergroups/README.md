# ConsumerGroups
(*ConsumerGroups*)

## Overview

Consumer groups enable the organization and categorization of consumers (users or applications) within an API ecosystem.
By grouping consumers together, you eliminate the need to manage them individually, providing a scalable, efficient approach to managing configurations.

### Available Operations

* [ListConsumerGroup](#listconsumergroup) - List all Consumer Groups
* [CreateConsumerGroup](#createconsumergroup) - Create a new Consumer Group
* [DeleteConsumerGroup](#deleteconsumergroup) - Delete a Consumer Group
* [GetConsumerGroup](#getconsumergroup) - Fetch a Consumer Group
* [UpsertConsumerGroup](#upsertconsumergroup) - Upsert a Consumer Group
* [RemoveAllConsumersFromConsumerGroup](#removeallconsumersfromconsumergroup) - Remove consumers from consumer group
* [ListConsumersForConsumerGroup](#listconsumersforconsumergroup) - List all Consumers in a Consumer Group
* [AddConsumerToGroup](#addconsumertogroup) - Add consumer to consumer group
* [RemoveConsumerFromGroup](#removeconsumerfromgroup) - Remove consumer from consumer group

## ListConsumerGroup

List all Consumer Groups

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumer_group" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups" -->
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

    res, err := s.ConsumerGroups.ListConsumerGroup(ctx, operations.ListConsumerGroupRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListConsumerGroupRequest](../../models/operations/listconsumergrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListConsumerGroupResponse](../../models/operations/listconsumergroupresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateConsumerGroup

Create a new Consumer Group

### Example Usage

<!-- UsageSnippet language="go" operationID="create-consumer_group" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups" -->
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

    res, err := s.ConsumerGroups.CreateConsumerGroup(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.ConsumerGroup{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConsumerGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `consumerGroup`                                                                    | [components.ConsumerGroup](../../models/components/consumergroup.md)               | :heavy_check_mark:                                                                 | Description of the new Consumer Group for creation                                 |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreateConsumerGroupResponse](../../models/operations/createconsumergroupresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteConsumerGroup

Delete a Consumer Group

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-consumer_group" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}" -->
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

    res, err := s.ConsumerGroups.DeleteConsumerGroup(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "")
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
| `consumerGroupID`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer Group to lookup                                                 |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteConsumerGroupResponse](../../models/operations/deleteconsumergroupresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetConsumerGroup

Get a Consumer Group using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-consumer_group" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}" -->
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

    res, err := s.ConsumerGroups.GetConsumerGroup(ctx, "", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConsumerGroupInsideWrapper != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `consumerGroupID`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer Group to lookup                                                 |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetConsumerGroupResponse](../../models/operations/getconsumergroupresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertConsumerGroup

Create or Update Consumer Group using ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-consumer_group" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}" -->
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

    res, err := s.ConsumerGroups.UpsertConsumerGroup(ctx, operations.UpsertConsumerGroupRequest{
        ConsumerGroupID: "",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerGroup: components.ConsumerGroup{
            Name: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConsumerGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpsertConsumerGroupRequest](../../models/operations/upsertconsumergrouprequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpsertConsumerGroupResponse](../../models/operations/upsertconsumergroupresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## RemoveAllConsumersFromConsumerGroup

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes all consumers from a Consumer Groups. This operation does not delete the consumer group.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-all-consumers-from-consumer-group" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/consumers" -->
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

    res, err := s.ConsumerGroups.RemoveAllConsumersFromConsumerGroup(ctx, "cf4c7e60-11db-49dd-b300-7c7e5f0f7e6b", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
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
| `consumerGroupID`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID or name of the consumer group                                             | cf4c7e60-11db-49dd-b300-7c7e5f0f7e6b                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.RemoveAllConsumersFromConsumerGroupResponse](../../models/operations/removeallconsumersfromconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConsumersForConsumerGroup

List all consumers in a consumer group

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumers-for-consumer-group" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/consumers" -->
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

    res, err := s.ConsumerGroups.ListConsumersForConsumerGroup(ctx, operations.ListConsumersForConsumerGroupRequest{
        ConsumerGroupID: "<id>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListConsumersForConsumerGroupRequest](../../models/operations/listconsumersforconsumergrouprequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListConsumersForConsumerGroupResponse](../../models/operations/listconsumersforconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddConsumerToGroup

Add a consumer to a consumer group

### Example Usage

<!-- UsageSnippet language="go" operationID="add-consumer-to-group" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/consumers" -->
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

    res, err := s.ConsumerGroups.AddConsumerToGroup(ctx, operations.AddConsumerToGroupRequest{
        ConsumerGroupID: "cf4c7e60-11db-49dd-b300-7c7e5f0f7e6b",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        RequestBody: &operations.AddConsumerToGroupRequestBody{
            ConsumerID: sdkkonnectgo.String("cf4c7e60-11db-49dd-b300-7c7e5f0f7e6b"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AddConsumerToGroupRequest](../../models/operations/addconsumertogrouprequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AddConsumerToGroupResponse](../../models/operations/addconsumertogroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveConsumerFromGroup

Remove a consumer from a consumer group

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-consumer-from-group" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/consumer_groups/{ConsumerGroupId}/consumers/{ConsumerId}" -->
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

    res, err := s.ConsumerGroups.RemoveConsumerFromGroup(ctx, operations.RemoveConsumerFromGroupRequest{
        ConsumerGroupID: "cf4c7e60-11db-49dd-b300-7c7e5f0f7e6b",
        ConsumerID: "<id>",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveConsumerFromGroupRequest](../../models/operations/removeconsumerfromgrouprequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveConsumerFromGroupResponse](../../models/operations/removeconsumerfromgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |