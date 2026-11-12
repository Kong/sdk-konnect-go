# BasicAuthCredentials

## Overview

### Available Operations

* [ListBasicAuth](#listbasicauth) - List all Basic-auth credentials in a workspace
* [GetBasicAuth](#getbasicauth) - Get a Basic-auth credential in a workspace
* [ListBasicAuthWithConsumer](#listbasicauthwithconsumer) - List all Basic-auth credentials associated with a Consumer in a workspace
* [CreateBasicAuthWithConsumer](#createbasicauthwithconsumer) - Create a new Basic-auth credential associated with a Consumer in a workspace
* [DeleteBasicAuthWithConsumer](#deletebasicauthwithconsumer) - Delete a a Basic-auth credential associated with a Consumer in a workspace
* [GetBasicAuthWithConsumer](#getbasicauthwithconsumer) - Get a Basic-auth credential associated with a Consumer in a workspace
* [UpsertBasicAuthWithConsumer](#upsertbasicauthwithconsumer) - Upsert a Basic-auth credential associated with a Consumer in a workspace

## ListBasicAuth

List all Basic-auth credentials in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-basic-auth" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/basic-auths" -->
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

    res, err := s.BasicAuthCredentials.ListBasicAuth(ctx, operations.ListBasicAuthRequest{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListBasicAuthRequest](../../models/operations/listbasicauthrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListBasicAuthResponse](../../models/operations/listbasicauthresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetBasicAuth

Get a Basic-auth credential using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-basic-auth" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/basic-auths/{BasicAuthId}" -->
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

    res, err := s.BasicAuthCredentials.GetBasicAuth(ctx, operations.GetBasicAuthRequest{
        BasicAuthID: "80db1b58-ca7c-4d21-b92a-64eb07725872",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetBasicAuthRequest](../../models/operations/getbasicauthrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetBasicAuthResponse](../../models/operations/getbasicauthresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListBasicAuthWithConsumer

List all Basic-auth credentials associated with a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-basic-auth-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/basic-auth" -->
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

    res, err := s.BasicAuthCredentials.ListBasicAuthWithConsumer(ctx, operations.ListBasicAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListBasicAuthWithConsumerRequest](../../models/operations/listbasicauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListBasicAuthWithConsumerResponse](../../models/operations/listbasicauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateBasicAuthWithConsumer

Create a new Basic-auth credential associated with a Consumer in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-basic-auth-with-consumer" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/basic-auth" -->
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

    res, err := s.BasicAuthCredentials.CreateBasicAuthWithConsumer(ctx, operations.CreateBasicAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        Workspace: "team-payments",
        BasicAuthWithoutParents: components.BasicAuthWithoutParents{
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
            Password: "hashedsoopersecretvalue",
            Username: "darius",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateBasicAuthWithConsumerRequest](../../models/operations/createbasicauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateBasicAuthWithConsumerResponse](../../models/operations/createbasicauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteBasicAuthWithConsumer

Delete a a Basic-auth credential associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-basic-auth-with-consumer" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/basic-auth/{BasicAuthId}" -->
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

    res, err := s.BasicAuthCredentials.DeleteBasicAuthWithConsumer(ctx, operations.DeleteBasicAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        BasicAuthID: "80db1b58-ca7c-4d21-b92a-64eb07725872",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteBasicAuthWithConsumerRequest](../../models/operations/deletebasicauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.DeleteBasicAuthWithConsumerResponse](../../models/operations/deletebasicauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetBasicAuthWithConsumer

Get a Basic-auth credential associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-basic-auth-with-consumer" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/basic-auth/{BasicAuthId}" -->
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

    res, err := s.BasicAuthCredentials.GetBasicAuthWithConsumer(ctx, operations.GetBasicAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        BasicAuthID: "80db1b58-ca7c-4d21-b92a-64eb07725872",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetBasicAuthWithConsumerRequest](../../models/operations/getbasicauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetBasicAuthWithConsumerResponse](../../models/operations/getbasicauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertBasicAuthWithConsumer

Create or Update a Basic-auth credential associated with a Consumer using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-basic-auth-with-consumer" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/consumers/{ConsumerIdForNestedEntities}/basic-auth/{BasicAuthId}" -->
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

    res, err := s.BasicAuthCredentials.UpsertBasicAuthWithConsumer(ctx, operations.UpsertBasicAuthWithConsumerRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConsumerIDForNestedEntities: "f28acbfa-c866-4587-b688-0208ac24df21",
        BasicAuthID: "80db1b58-ca7c-4d21-b92a-64eb07725872",
        Workspace: "team-payments",
        BasicAuthWithoutParents: components.BasicAuthWithoutParents{
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
            Password: "hashedsoopersecretvalue",
            Username: "darius",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BasicAuth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpsertBasicAuthWithConsumerRequest](../../models/operations/upsertbasicauthwithconsumerrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpsertBasicAuthWithConsumerResponse](../../models/operations/upsertbasicauthwithconsumerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |