# Consumers
(*Consumers*)

## Overview

The consumer object represents a consumer - or a user - of a service.
You can either rely on Kong Gateway as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong Gateway and your existing primary datastore.


### Available Operations

* [ListConsumer](#listconsumer) - List all Consumers
* [CreateConsumer](#createconsumer) - Create a new Consumer
* [DeleteConsumer](#deleteconsumer) - Delete a Consumer
* [GetConsumer](#getconsumer) - Get a Consumer
* [UpsertConsumer](#upsertconsumer) - Upsert a Consumer
* [RemoveConsumerFromAllConsumerGroups](#removeconsumerfromallconsumergroups) - Remove consumer from all consumer groups
* [ListConsumerGroupsForConsumer](#listconsumergroupsforconsumer) - List all Consumer Groups a Consumer belongs to
* [AddConsumerToSpecificConsumerGroup](#addconsumertospecificconsumergroup) - Add consumer to a specific consumer group
* [RemoveConsumerFromConsumerGroup](#removeconsumerfromconsumergroup) - Remove consumer from consumer group

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

### Example Usage

<!-- UsageSnippet language="go" operationID="create-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers" -->
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
| `controlPlaneID`                                                                                                                | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | The UUID of your control plane. This variable is available in the Konnect manager.                                              | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                            |
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
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `consumerID`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer to lookup                                                       | c1059869-6fa7-4329-a5f5-5946d14ca2c5                                               |
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
| `consumerID`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer to lookup                                                       | c1059869-6fa7-4329-a5f5-5946d14ca2c5                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
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

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/consumers/{ConsumerId}" -->
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
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `consumerID`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Consumer to lookup                                                       | c1059869-6fa7-4329-a5f5-5946d14ca2c5                                               |
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