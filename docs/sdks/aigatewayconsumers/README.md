# AIGatewayConsumers

## Overview

Individual consumers with credentials and group memberships for AI Gateway access control.

### Available Operations

* [ListAiGatewayConsumers](#listaigatewayconsumers) - List AI Gateway Consumers
* [CreateAiGatewayConsumer](#createaigatewayconsumer) - Create an AI Gateway Consumer
* [ListAiGatewayConsumerCredentials](#listaigatewayconsumercredentials) - List AI Gateway Consumer Credentials
* [CreateAiGatewayConsumerCredential](#createaigatewayconsumercredential) - Create an AI Gateway Consumer Credential
* [GetAiGatewayConsumerCredential](#getaigatewayconsumercredential) - Get an AI Gateway Consumer Credential
* [DeleteAiGatewayConsumerCredential](#deleteaigatewayconsumercredential) - Delete an AI Gateway Consumer Credential
* [GetAiGatewayConsumer](#getaigatewayconsumer) - Get an AI Gateway Consumer
* [UpdateAiGatewayConsumer](#updateaigatewayconsumer) - Update an AI Gateway Consumer
* [DeleteAiGatewayConsumer](#deleteaigatewayconsumer) - Delete an AI Gateway Consumer
* [ListAiGatewayConsumerGroupsForConsumer](#listaigatewayconsumergroupsforconsumer) - List Consumer Groups a Consumer belongs to
* [UpdateAiGatewayConsumerGroupsForConsumer](#updateaigatewayconsumergroupsforconsumer) - Updates Consumer Groups a Consumer belongs to

## ListAiGatewayConsumers

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of all consumers for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-consumers" method="get" path="/v1/ai-gateways/{gatewayId}/consumers" -->
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

    res, err := s.AIGatewayConsumers.ListAiGatewayConsumers(ctx, operations.ListAiGatewayConsumersRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayConsumersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAiGatewayConsumersRequest](../../models/operations/listaigatewayconsumersrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAiGatewayConsumersResponse](../../models/operations/listaigatewayconsumersresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayConsumer

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Creates a new consumer for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-consumer" method="post" path="/v1/ai-gateways/{gatewayId}/consumers" -->
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

    res, err := s.AIGatewayConsumers.CreateAiGatewayConsumer(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateAIGatewayConsumerRequest{
        DisplayName: "Greg's Dev Consumer",
        Name: "gregs-dev-consumer",
        Type: components.CreateAIGatewayConsumerRequestTypeOauth,
        CustomID: sdkkonnectgo.Pointer("dev-users"),
        Labels: map[string]string{
            "category": "finance",
        },
        ManagedBy: map[string]string{
            "owner": "terraform",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConsumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `gatewayID`                                                                                            | `string`                                                                                               | :heavy_check_mark:                                                                                     | The unique ID of the AI Gateway.                                                                       | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                   |
| `createAIGatewayConsumerRequest`                                                                       | [components.CreateAIGatewayConsumerRequest](../../models/components/createaigatewayconsumerrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.CreateAiGatewayConsumerResponse](../../models/operations/createaigatewayconsumerresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## ListAiGatewayConsumerCredentials

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of all credentials for an AI Gateway consumer.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-consumer-credentials" method="get" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerId}/credentials" -->
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

    res, err := s.AIGatewayConsumers.ListAiGatewayConsumerCredentials(ctx, operations.ListAiGatewayConsumerCredentialsRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConsumerID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayConsumerCredentialsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListAiGatewayConsumerCredentialsRequest](../../models/operations/listaigatewayconsumercredentialsrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ListAiGatewayConsumerCredentialsResponse](../../models/operations/listaigatewayconsumercredentialsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayConsumerCredential

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Creates a new credential for an AI Gateway consumer.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-consumer-credential" method="post" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerId}/credentials" -->
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

    res, err := s.AIGatewayConsumers.CreateAiGatewayConsumerCredential(ctx, operations.CreateAiGatewayConsumerCredentialRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConsumerID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CreateAIGatewayConsumerCredentialRequest: components.CreateAIGatewayConsumerCredentialRequest{
            DisplayName: "Greg's Dev Key",
            Name: "gregs-dev-key",
            Labels: map[string]string{
                "category": "finance",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
            Type: components.CreateAIGatewayConsumerCredentialRequestTypeAPIKey,
            TTL: sdkkonnectgo.Pointer[int64](86400),
            APIKey: sdkkonnectgo.Pointer("sk-387788hd3xnej"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConsumerCredentialWithKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.CreateAiGatewayConsumerCredentialRequest](../../models/operations/createaigatewayconsumercredentialrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.CreateAiGatewayConsumerCredentialResponse](../../models/operations/createaigatewayconsumercredentialresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayConsumerCredential

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific credential for an AI Gateway consumer.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-consumer-credential" method="get" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerId}/credentials/{credentialId}" -->
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

    res, err := s.AIGatewayConsumers.GetAiGatewayConsumerCredential(ctx, operations.GetAiGatewayConsumerCredentialRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConsumerID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CredentialID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConsumerCredential != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAiGatewayConsumerCredentialRequest](../../models/operations/getaigatewayconsumercredentialrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.GetAiGatewayConsumerCredentialResponse](../../models/operations/getaigatewayconsumercredentialresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayConsumerCredential

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific credential for an AI Gateway consumer.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-consumer-credential" method="delete" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerId}/credentials/{credentialId}" -->
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

    res, err := s.AIGatewayConsumers.DeleteAiGatewayConsumerCredential(ctx, operations.DeleteAiGatewayConsumerCredentialRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConsumerID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        CredentialID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteAiGatewayConsumerCredentialRequest](../../models/operations/deleteaigatewayconsumercredentialrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.DeleteAiGatewayConsumerCredentialResponse](../../models/operations/deleteaigatewayconsumercredentialresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayConsumer

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific AI Gateway consumer.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-consumer" method="get" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerIdOrName}" -->
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

    res, err := s.AIGatewayConsumers.GetAiGatewayConsumer(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConsumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `consumerIDOrName`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway consumer.        | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayConsumerResponse](../../models/operations/getaigatewayconsumerresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayConsumer

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the configuration of an existing AI Gateway consumer.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-consumer" method="put" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerIdOrName}" -->
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

    res, err := s.AIGatewayConsumers.UpdateAiGatewayConsumer(ctx, operations.UpdateAiGatewayConsumerRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConsumerIDOrName: "my-entity-name",
        UpdateAIGatewayConsumerRequest: components.UpdateAIGatewayConsumerRequest{
            DisplayName: "Greg's Dev Consumer",
            Name: "gregs-dev-consumer",
            Type: components.UpdateAIGatewayConsumerRequestTypeAPIKey,
            CustomID: sdkkonnectgo.Pointer("dev-users"),
            Labels: map[string]string{
                "category": "finance",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConsumer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAiGatewayConsumerRequest](../../models/operations/updateaigatewayconsumerrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAiGatewayConsumerResponse](../../models/operations/updateaigatewayconsumerresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayConsumer

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific AI Gateway consumer.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-consumer" method="delete" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerIdOrName}" -->
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

    res, err := s.AIGatewayConsumers.DeleteAiGatewayConsumer(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `consumerIDOrName`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway consumer.        | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayConsumerResponse](../../models/operations/deleteaigatewayconsumerresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## ListAiGatewayConsumerGroupsForConsumer

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List AI Gateway Consumer Groups an Consumer belongs to

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-consumer-groups-for-consumer" method="get" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerIdOrName}/consumer-groups" -->
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

    res, err := s.AIGatewayConsumers.ListAiGatewayConsumerGroupsForConsumer(ctx, operations.ListAiGatewayConsumerGroupsForConsumerRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConsumerIDOrName: "my-entity-name",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayConsumerGroupsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.ListAiGatewayConsumerGroupsForConsumerRequest](../../models/operations/listaigatewayconsumergroupsforconsumerrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ListAiGatewayConsumerGroupsForConsumerResponse](../../models/operations/listaigatewayconsumergroupsforconsumerresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayConsumerGroupsForConsumer

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates AI Gateway Consumer Groups a Consumer belongs to

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-consumer-groups-for-consumer" method="put" path="/v1/ai-gateways/{gatewayId}/consumers/{consumerIdOrName}/consumer-groups" -->
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

    res, err := s.AIGatewayConsumers.UpdateAiGatewayConsumerGroupsForConsumer(ctx, operations.UpdateAiGatewayConsumerGroupsForConsumerRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConsumerIDOrName: "my-entity-name",
        RequestBody: operations.UpdateAiGatewayConsumerGroupsForConsumerRequestBody{},
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
| `request`                                                                                                                                | [operations.UpdateAiGatewayConsumerGroupsForConsumerRequest](../../models/operations/updateaigatewayconsumergroupsforconsumerrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.UpdateAiGatewayConsumerGroupsForConsumerResponse](../../models/operations/updateaigatewayconsumergroupsforconsumerresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |