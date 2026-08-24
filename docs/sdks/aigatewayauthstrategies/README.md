# AIGatewayAuthStrategies

## Overview

Auth strategies for authenticating clients accessing AI Gateway resources.

### Available Operations

* [ListAiGatewayAuthStrategies](#listaigatewayauthstrategies) - List AI Gateway Auth Strategies
* [CreateAiGatewayAuthStrategy](#createaigatewayauthstrategy) - Create an AI Gateway Auth Strategy
* [GetAiGatewayAuthStrategy](#getaigatewayauthstrategy) - Get an AI Gateway Auth Strategy
* [UpdateAiGatewayAuthStrategy](#updateaigatewayauthstrategy) - Update an AI Gateway Auth Strategy
* [DeleteAiGatewayAuthStrategy](#deleteaigatewayauthstrategy) - Delete an AI Gateway Auth Strategy

## ListAiGatewayAuthStrategies

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of auth strategies configured for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-auth-strategies" method="get" path="/v1/ai-gateways/{gatewayId}/auth-strategies" -->
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

    res, err := s.AIGatewayAuthStrategies.ListAiGatewayAuthStrategies(ctx, operations.ListAiGatewayAuthStrategiesRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayAuthStrategiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListAiGatewayAuthStrategiesRequest](../../models/operations/listaigatewayauthstrategiesrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListAiGatewayAuthStrategiesResponse](../../models/operations/listaigatewayauthstrategiesresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayAuthStrategy

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Registers a new auth strategy for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-auth-strategy" method="post" path="/v1/ai-gateways/{gatewayId}/auth-strategies" -->
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

    res, err := s.AIGatewayAuthStrategies.CreateAiGatewayAuthStrategy(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayAuthStrategyRequestOpenidConnect(
        components.AIGatewayAuthStrategyOpenIDConnect{
            DisplayName: "Okta AI SE",
            Name: "okta-ai-se",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayAuthStrategy != nil {
        switch res.AIGatewayAuthStrategy.Type {
            case components.AIGatewayAuthStrategyTypeKeyAuth:
                // res.AIGatewayAuthStrategy.AIGatewayAuthStrategyKeyAuthResponse is populated
            case components.AIGatewayAuthStrategyTypeOpenidConnect:
                // res.AIGatewayAuthStrategy.AIGatewayAuthStrategyOpenIDConnectResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `gatewayID`                                                                                                    | `string`                                                                                                       | :heavy_check_mark:                                                                                             | The unique ID of the AI Gateway.                                                                               | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                           |
| `createAIGatewayAuthStrategyRequest`                                                                           | [components.CreateAIGatewayAuthStrategyRequest](../../models/components/createaigatewayauthstrategyrequest.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.CreateAiGatewayAuthStrategyResponse](../../models/operations/createaigatewayauthstrategyresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayAuthStrategy

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific AI Gateway auth strategy.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-auth-strategy" method="get" path="/v1/ai-gateways/{gatewayId}/auth-strategies/{authStrategyIdOrName}" -->
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

    res, err := s.AIGatewayAuthStrategies.GetAiGatewayAuthStrategy(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayAuthStrategy != nil {
        switch res.AIGatewayAuthStrategy.Type {
            case components.AIGatewayAuthStrategyTypeKeyAuth:
                // res.AIGatewayAuthStrategy.AIGatewayAuthStrategyKeyAuthResponse is populated
            case components.AIGatewayAuthStrategyTypeOpenidConnect:
                // res.AIGatewayAuthStrategy.AIGatewayAuthStrategyOpenIDConnectResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `authStrategyIDOrName`                                   | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway auth strategy.   | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayAuthStrategyResponse](../../models/operations/getaigatewayauthstrategyresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayAuthStrategy

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the configuration of an existing AI Gateway auth strategy.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-auth-strategy" method="put" path="/v1/ai-gateways/{gatewayId}/auth-strategies/{authStrategyIdOrName}" -->
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

    res, err := s.AIGatewayAuthStrategies.UpdateAiGatewayAuthStrategy(ctx, operations.UpdateAiGatewayAuthStrategyRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        AuthStrategyIDOrName: "my-entity-name",
        UpdateAIGatewayAuthStrategyRequest: components.CreateUpdateAIGatewayAuthStrategyRequestOpenidConnect(
            components.AIGatewayAuthStrategyOpenIDConnect{
                DisplayName: "Okta AI SE",
                Name: "okta-ai-se",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayAuthStrategy != nil {
        switch res.AIGatewayAuthStrategy.Type {
            case components.AIGatewayAuthStrategyTypeKeyAuth:
                // res.AIGatewayAuthStrategy.AIGatewayAuthStrategyKeyAuthResponse is populated
            case components.AIGatewayAuthStrategyTypeOpenidConnect:
                // res.AIGatewayAuthStrategy.AIGatewayAuthStrategyOpenIDConnectResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateAiGatewayAuthStrategyRequest](../../models/operations/updateaigatewayauthstrategyrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateAiGatewayAuthStrategyResponse](../../models/operations/updateaigatewayauthstrategyresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayAuthStrategy

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific AI Gateway auth strategy.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-auth-strategy" method="delete" path="/v1/ai-gateways/{gatewayId}/auth-strategies/{authStrategyIdOrName}" -->
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

    res, err := s.AIGatewayAuthStrategies.DeleteAiGatewayAuthStrategy(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `authStrategyIDOrName`                                   | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway auth strategy.   | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayAuthStrategyResponse](../../models/operations/deleteaigatewayauthstrategyresponse.md), error**

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