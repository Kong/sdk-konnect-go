# AIGatewayConfigStoreSecrets

## Overview

### Available Operations

* [ListAiGatewayConfigStoreSecrets](#listaigatewayconfigstoresecrets) - List AI Gateway Config Store Secrets
* [CreateAiGatewayConfigStoreSecret](#createaigatewayconfigstoresecret) - Create an AI Gateway Config Store Secret
* [GetAiGatewayConfigStoreSecret](#getaigatewayconfigstoresecret) - Get an AI Gateway Config Store Secret
* [UpdateAiGatewayConfigStoreSecret](#updateaigatewayconfigstoresecret) - Update an AI Gateway Config Store Secret
* [DeleteAiGatewayConfigStoreSecret](#deleteaigatewayconfigstoresecret) - Delete an AI Gateway Config Store Secret

## ListAiGatewayConfigStoreSecrets

Returns a collection of all secrets for an AI Gateway Config Store.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-config-store-secrets" method="get" path="/v1/ai-gateways/{gatewayId}/config-stores/{configStoreIdOrName}/secrets" -->
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

    res, err := s.AIGatewayConfigStoreSecrets.ListAiGatewayConfigStoreSecrets(ctx, operations.ListAiGatewayConfigStoreSecretsRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConfigStoreIDOrName: "my-entity-name",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayConfigStoreSecretsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListAiGatewayConfigStoreSecretsRequest](../../models/operations/listaigatewayconfigstoresecretsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListAiGatewayConfigStoreSecretsResponse](../../models/operations/listaigatewayconfigstoresecretsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayConfigStoreSecret

Creates a secret for an AI Gateway Config Store.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-config-store-secret" method="post" path="/v1/ai-gateways/{gatewayId}/config-stores/{configStoreIdOrName}/secrets" -->
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

    res, err := s.AIGatewayConfigStoreSecrets.CreateAiGatewayConfigStoreSecret(ctx, operations.CreateAiGatewayConfigStoreSecretRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConfigStoreIDOrName: "my-entity-name",
        CreateAIGatewayConfigStoreSecretRequest: components.CreateAIGatewayConfigStoreSecretRequest{
            Key: "my-secret-key",
            Value: "my-secret-value",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConfigStoreSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.CreateAiGatewayConfigStoreSecretRequest](../../models/operations/createaigatewayconfigstoresecretrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreateAiGatewayConfigStoreSecretResponse](../../models/operations/createaigatewayconfigstoresecretresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayConfigStoreSecret

Returns the secret entity for the Config Store. Secret values once stored cannot be retrieved.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-config-store-secret" method="get" path="/v1/ai-gateways/{gatewayId}/config-stores/{configStoreIdOrName}/secrets/{key}" -->
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

    res, err := s.AIGatewayConfigStoreSecrets.GetAiGatewayConfigStoreSecret(ctx, operations.GetAiGatewayConfigStoreSecretRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConfigStoreIDOrName: "my-entity-name",
        Key: "my-secret-key",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConfigStoreSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAiGatewayConfigStoreSecretRequest](../../models/operations/getaigatewayconfigstoresecretrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetAiGatewayConfigStoreSecretResponse](../../models/operations/getaigatewayconfigstoresecretresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayConfigStoreSecret

Updates a secret for an AI Gateway Config Store.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-config-store-secret" method="put" path="/v1/ai-gateways/{gatewayId}/config-stores/{configStoreIdOrName}/secrets/{key}" -->
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

    res, err := s.AIGatewayConfigStoreSecrets.UpdateAiGatewayConfigStoreSecret(ctx, operations.UpdateAiGatewayConfigStoreSecretRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConfigStoreIDOrName: "my-entity-name",
        Key: "my-secret-key",
        UpdateAIGatewayConfigStoreSecretRequest: components.UpdateAIGatewayConfigStoreSecretRequest{
            Value: "my-secret-value",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayConfigStoreSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateAiGatewayConfigStoreSecretRequest](../../models/operations/updateaigatewayconfigstoresecretrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateAiGatewayConfigStoreSecretResponse](../../models/operations/updateaigatewayconfigstoresecretresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayConfigStoreSecret

Removes a secret from an AI Gateway Config Store.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-config-store-secret" method="delete" path="/v1/ai-gateways/{gatewayId}/config-stores/{configStoreIdOrName}/secrets/{key}" -->
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

    res, err := s.AIGatewayConfigStoreSecrets.DeleteAiGatewayConfigStoreSecret(ctx, operations.DeleteAiGatewayConfigStoreSecretRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ConfigStoreIDOrName: "my-entity-name",
        Key: "my-secret-key",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteAiGatewayConfigStoreSecretRequest](../../models/operations/deleteaigatewayconfigstoresecretrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.DeleteAiGatewayConfigStoreSecretResponse](../../models/operations/deleteaigatewayconfigstoresecretresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |