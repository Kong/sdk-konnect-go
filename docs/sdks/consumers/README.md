# Consumers

## Overview

The consumer object represents a consumer - or a user - of a service.
You can either rely on Kong Gateway as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong Gateway and your existing primary datastore.


### Available Operations

* [ListConsumer](#listconsumer) - List all Consumers in a workspace
* [CreateConsumer](#createconsumer) - Create a new Consumer in a workspace
* [DeleteConsumer](#deleteconsumer) - Delete a Consumer in a workspace
* [GetConsumer](#getconsumer) - Get a Consumer in a workspace
* [UpsertConsumer](#upsertconsumer) - Upsert a Consumer in a workspace
* [RemoveConsumerFromAllConsumerGroups](#removeconsumerfromallconsumergroups) - Remove consumer from all consumer groups in a workspace
* [ListConsumerGroupsForConsumer](#listconsumergroupsforconsumer) - List all Consumer Groups a Consumer belongs to in a workspace
* [AddConsumerToSpecificConsumerGroup](#addconsumertospecificconsumergroup) - Add consumer to a specific consumer group in a workspace
* [RemoveConsumerFromConsumerGroup](#removeconsumerfromconsumergroup) - Remove consumer from consumer group in a workspace

## ListConsumer

List all Consumers in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers" -->
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

Create a new Consumer in a workspace

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers" example="DuplicateApiKey" -->
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

    res, err := s.Consumers.CreateConsumer(ctx, operations.CreateConsumerRequest{
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
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers" example="InvalidAuthCred" -->
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

    res, err := s.Consumers.CreateConsumer(ctx, operations.CreateConsumerRequest{
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
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers" example="NoAPIKey" -->
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

    res, err := s.Consumers.CreateConsumer(ctx, operations.CreateConsumerRequest{
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateConsumerRequest](../../models/operations/createconsumerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateConsumerResponse](../../models/operations/createconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteConsumer

Delete a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-consumer" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" -->
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

    res, err := s.Consumers.DeleteConsumer(ctx, operations.DeleteConsumerRequest{
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteConsumerRequest](../../models/operations/deleteconsumerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteConsumerResponse](../../models/operations/deleteconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetConsumer

Get a Consumer using ID or username in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" -->
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

    res, err := s.Consumers.GetConsumer(ctx, operations.GetConsumerRequest{
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetConsumerRequest](../../models/operations/getconsumerrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetConsumerResponse](../../models/operations/getconsumerresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertConsumer

Create or Update Consumer using ID or username in a workspace.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" example="DuplicateApiKey" -->
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
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" example="InvalidAuthCred" -->
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
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}" example="NoAPIKey" -->
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

Removes a consumer from all Consumer Groups. This operation does not delete the consumer group in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-consumer-from-all-consumer-groups" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups" -->
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

    res, err := s.Consumers.RemoveConsumerFromAllConsumerGroups(ctx, operations.RemoveConsumerFromAllConsumerGroupsRequest{
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.RemoveConsumerFromAllConsumerGroupsRequest](../../models/operations/removeconsumerfromallconsumergroupsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.RemoveConsumerFromAllConsumerGroupsResponse](../../models/operations/removeconsumerfromallconsumergroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConsumerGroupsForConsumer

List all Consumer Groups a Consumer belongs to in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-consumer-groups-for-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups" -->
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
| `request`                                                                                                          | [operations.ListConsumerGroupsForConsumerRequest](../../models/operations/listconsumergroupsforconsumerrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListConsumerGroupsForConsumerResponse](../../models/operations/listconsumergroupsforconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddConsumerToSpecificConsumerGroup

Add a consumer to a consumer group in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="add-consumer-to-specific-consumer-group" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups" -->
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
        Workspace: "team-payments",
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

Removes a consumer from a Consumer Group. This operation does not delete the consumer group in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-consumer-from-consumer-group" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerId}/consumer_groups/{ConsumerGroupId}" -->
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
| `request`                                                                                                              | [operations.RemoveConsumerFromConsumerGroupRequest](../../models/operations/removeconsumerfromconsumergrouprequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.RemoveConsumerFromConsumerGroupResponse](../../models/operations/removeconsumerfromconsumergroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |