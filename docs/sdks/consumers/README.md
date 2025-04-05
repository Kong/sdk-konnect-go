# Consumers
(*Consumers*)

## Overview

The consumer object represents a consumer - or a user - of a service.
You can either rely on Kong Gateway as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong Gateway and your existing primary datastore.


### Available Operations

* [ListConsumer](#listconsumer) - List all Consumers
* [CreateConsumer](#createconsumer) - Create a new Consumer
* [DeleteConsumer](#deleteconsumer) - Delete a Consumer
* [GetConsumer](#getconsumer) - Fetch a Consumer
* [UpsertConsumer](#upsertconsumer) - Upsert a Consumer

## ListConsumer

List all Consumers

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

    res, err := s.Consumers.ListConsumer(ctx, operations.ListConsumerRequest{
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

    res, err := s.Consumers.CreateConsumer(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Consumer{
        CustomID: sdkkonnectgo.String("4200"),
        ID: sdkkonnectgo.String("8a388226-80e8-4027-a486-25e4f7db5d21"),
        Tags: []string{
            "silver-tier",
        },
        Username: sdkkonnectgo.String("bob-the-builder"),
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

    res, err := s.Consumers.UpsertConsumer(ctx, operations.UpsertConsumerRequest{
        ConsumerID: "c1059869-6fa7-4329-a5f5-5946d14ca2c5",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Consumer: components.Consumer{
            CustomID: sdkkonnectgo.String("4200"),
            ID: sdkkonnectgo.String("8a388226-80e8-4027-a486-25e4f7db5d21"),
            Tags: []string{
                "silver-tier",
            },
            Username: sdkkonnectgo.String("bob-the-builder"),
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