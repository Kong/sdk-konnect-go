# AIGatewayIdentityProviders

## Overview

Identity providers for authenticating users and accessing AI Gateway resources.

### Available Operations

* [ListAiGatewayIdentityProviders](#listaigatewayidentityproviders) - List AI Gateway Identity Providers
* [CreateAiGatewayIdentityProvider](#createaigatewayidentityprovider) - Create an AI Gateway Identity Provider
* [GetAiGatewayIdentityProvider](#getaigatewayidentityprovider) - Get an AI Gateway Identity Provider
* [UpdateAiGatewayIdentityProvider](#updateaigatewayidentityprovider) - Update an AI Gateway Identity Provider
* [DeleteAiGatewayIdentityProvider](#deleteaigatewayidentityprovider) - Delete an AI Gateway Identity Provider

## ListAiGatewayIdentityProviders

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of identity providers configured for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-identity-providers" method="get" path="/v1/ai-gateways/{gatewayId}/identity" -->
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

    res, err := s.AIGatewayIdentityProviders.ListAiGatewayIdentityProviders(ctx, operations.ListAiGatewayIdentityProvidersRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayIdentityProvidersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListAiGatewayIdentityProvidersRequest](../../models/operations/listaigatewayidentityprovidersrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListAiGatewayIdentityProvidersResponse](../../models/operations/listaigatewayidentityprovidersresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Registers a new identity provider for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-identity-provider" method="post" path="/v1/ai-gateways/{gatewayId}/identity" -->
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

    res, err := s.AIGatewayIdentityProviders.CreateAiGatewayIdentityProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayIdentityProviderRequestKeyAuth(
        components.AIGatewayIdentityProviderKeyAuth{
            DisplayName: "Okta AI SE",
            Name: "okta-ai-se",
            Type: components.AIGatewayIdentityProviderKeyAuthTypeKeyAuth,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayIdentityProvider != nil {
        switch res.AIGatewayIdentityProvider.Type {
            case components.AIGatewayIdentityProviderTypeKeyAuth:
                // res.AIGatewayIdentityProvider.AIGatewayIdentityProviderKeyAuthResponse is populated
            case components.AIGatewayIdentityProviderTypeOpenidConnect:
                // res.AIGatewayIdentityProvider.AIGatewayIdentityProviderOpenIDConnectResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `gatewayID`                                                                                                            | `string`                                                                                                               | :heavy_check_mark:                                                                                                     | The unique ID of the AI Gateway.                                                                                       | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                                   |
| `createAIGatewayIdentityProviderRequest`                                                                               | [components.CreateAIGatewayIdentityProviderRequest](../../models/components/createaigatewayidentityproviderrequest.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.CreateAiGatewayIdentityProviderResponse](../../models/operations/createaigatewayidentityproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific AI Gateway identity provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-identity-provider" method="get" path="/v1/ai-gateways/{gatewayId}/identity/{identityProviderIdOrName}" -->
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

    res, err := s.AIGatewayIdentityProviders.GetAiGatewayIdentityProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayIdentityProvider != nil {
        switch res.AIGatewayIdentityProvider.Type {
            case components.AIGatewayIdentityProviderTypeKeyAuth:
                // res.AIGatewayIdentityProvider.AIGatewayIdentityProviderKeyAuthResponse is populated
            case components.AIGatewayIdentityProviderTypeOpenidConnect:
                // res.AIGatewayIdentityProvider.AIGatewayIdentityProviderOpenIDConnectResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `gatewayID`                                                | `string`                                                   | :heavy_check_mark:                                         | The unique ID of the AI Gateway.                           | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                       |
| `identityProviderIDOrName`                                 | `string`                                                   | :heavy_check_mark:                                         | The unique ID or name of the AI Gateway Identity provider. | my-entity-name                                             |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.GetAiGatewayIdentityProviderResponse](../../models/operations/getaigatewayidentityproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the configuration of an existing AI Gateway Identity provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-identity-provider" method="put" path="/v1/ai-gateways/{gatewayId}/identity/{identityProviderIdOrName}" -->
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

    res, err := s.AIGatewayIdentityProviders.UpdateAiGatewayIdentityProvider(ctx, operations.UpdateAiGatewayIdentityProviderRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        IdentityProviderIDOrName: "my-entity-name",
        UpdateAIGatewayIdentityProviderRequest: components.CreateUpdateAIGatewayIdentityProviderRequestKeyAuth(
            components.AIGatewayIdentityProviderKeyAuth{
                DisplayName: "Okta AI SE",
                Name: "okta-ai-se",
                Type: components.AIGatewayIdentityProviderKeyAuthTypeKeyAuth,
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayIdentityProvider != nil {
        switch res.AIGatewayIdentityProvider.Type {
            case components.AIGatewayIdentityProviderTypeKeyAuth:
                // res.AIGatewayIdentityProvider.AIGatewayIdentityProviderKeyAuthResponse is populated
            case components.AIGatewayIdentityProviderTypeOpenidConnect:
                // res.AIGatewayIdentityProvider.AIGatewayIdentityProviderOpenIDConnectResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.UpdateAiGatewayIdentityProviderRequest](../../models/operations/updateaigatewayidentityproviderrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.UpdateAiGatewayIdentityProviderResponse](../../models/operations/updateaigatewayidentityproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayIdentityProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific AI Gateway Identity provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-identity-provider" method="delete" path="/v1/ai-gateways/{gatewayId}/identity/{identityProviderIdOrName}" -->
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

    res, err := s.AIGatewayIdentityProviders.DeleteAiGatewayIdentityProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `gatewayID`                                                | `string`                                                   | :heavy_check_mark:                                         | The unique ID of the AI Gateway.                           | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                       |
| `identityProviderIDOrName`                                 | `string`                                                   | :heavy_check_mark:                                         | The unique ID or name of the AI Gateway Identity provider. | my-entity-name                                             |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.DeleteAiGatewayIdentityProviderResponse](../../models/operations/deleteaigatewayidentityproviderresponse.md), error**

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