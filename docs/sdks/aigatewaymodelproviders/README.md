# AIGatewayModelProviders

## Overview

Model providers that define the backend AI service connections for the AI Gateway.

### Available Operations

* [ListAiGatewayModelProviders](#listaigatewaymodelproviders) - List AI Gateway Model Providers
* [CreateAiGatewayModelProvider](#createaigatewaymodelprovider) - Create an AI Gateway Model Provider
* [GetAiGatewayModelProvider](#getaigatewaymodelprovider) - Get an AI Gateway Model Provider
* [UpdateAiGatewayModelProvider](#updateaigatewaymodelprovider) - Update an AI Gateway Model Provider
* [DeleteAiGatewayModelProvider](#deleteaigatewaymodelprovider) - Delete an AI Gateway Model Provider

## ListAiGatewayModelProviders

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns a list of model providers configured for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-model-providers" method="get" path="/v1/ai-gateways/{gatewayId}/model-providers" -->
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

    res, err := s.AIGatewayModelProviders.ListAiGatewayModelProviders(ctx, operations.ListAiGatewayModelProvidersRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayModelProvidersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListAiGatewayModelProvidersRequest](../../models/operations/listaigatewaymodelprovidersrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListAiGatewayModelProvidersResponse](../../models/operations/listaigatewaymodelprovidersresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayModelProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Registers a new model provider for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-model-provider" method="post" path="/v1/ai-gateways/{gatewayId}/model-providers" -->
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

    res, err := s.AIGatewayModelProviders.CreateAiGatewayModelProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayModelProviderRequestVertex(
        components.AIGatewayModelProviderVertex{
            Type: components.AIGatewayModelProviderVertexTypeVertex,
            DisplayName: "Azure AI SE",
            Name: "azure-ai-se",
            Config: components.AIGatewayModelProviderVertexConfig{
                Auth: components.CreateAIGatewayModelProviderVertexAuthGcp(
                    components.AIGatewayModelProviderConfigAuthGCP{
                        Type: components.AIGatewayModelProviderConfigAuthGCPTypeGcp,
                    },
                ),
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayModelProvider != nil {
        switch res.AIGatewayModelProvider.Type {
            case components.AIGatewayModelProviderTypeAnthropic:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAnthropic is populated
            case components.AIGatewayModelProviderTypeAzure:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAzure is populated
            case components.AIGatewayModelProviderTypeBedrock:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderBedrock is populated
            case components.AIGatewayModelProviderTypeCerebras:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCerebras is populated
            case components.AIGatewayModelProviderTypeCohere:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCohere is populated
            case components.AIGatewayModelProviderTypeDashscope:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDashscope is populated
            case components.AIGatewayModelProviderTypeDatabricks:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDatabricks is populated
            case components.AIGatewayModelProviderTypeDeepseek:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDeepseek is populated
            case components.AIGatewayModelProviderTypeGemini:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderGemini is populated
            case components.AIGatewayModelProviderTypeHuggingface:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderHuggingface is populated
            case components.AIGatewayModelProviderTypeKimi:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderKimi is populated
            case components.AIGatewayModelProviderTypeLlama2:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderLlama2 is populated
            case components.AIGatewayModelProviderTypeMistral:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderMistral is populated
            case components.AIGatewayModelProviderTypeOllama:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOllama is populated
            case components.AIGatewayModelProviderTypeOpenai:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOpenai is populated
            case components.AIGatewayModelProviderTypeVercel:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVercel is populated
            case components.AIGatewayModelProviderTypeVllm:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVllm is populated
            case components.AIGatewayModelProviderTypeXai:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderXai is populated
            case components.AIGatewayModelProviderTypeVertex:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVertex is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `gatewayID`                                                                                                      | `string`                                                                                                         | :heavy_check_mark:                                                                                               | The unique ID of the AI Gateway.                                                                                 | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                             |
| `createAIGatewayModelProviderRequest`                                                                            | [components.CreateAIGatewayModelProviderRequest](../../models/components/createaigatewaymodelproviderrequest.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.CreateAiGatewayModelProviderResponse](../../models/operations/createaigatewaymodelproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayModelProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the details of a specific AI Gateway model provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-model-provider" method="get" path="/v1/ai-gateways/{gatewayId}/model-providers/{modelProviderIdOrName}" -->
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

    res, err := s.AIGatewayModelProviders.GetAiGatewayModelProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayModelProvider != nil {
        switch res.AIGatewayModelProvider.Type {
            case components.AIGatewayModelProviderTypeAnthropic:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAnthropic is populated
            case components.AIGatewayModelProviderTypeAzure:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAzure is populated
            case components.AIGatewayModelProviderTypeBedrock:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderBedrock is populated
            case components.AIGatewayModelProviderTypeCerebras:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCerebras is populated
            case components.AIGatewayModelProviderTypeCohere:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCohere is populated
            case components.AIGatewayModelProviderTypeDashscope:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDashscope is populated
            case components.AIGatewayModelProviderTypeDatabricks:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDatabricks is populated
            case components.AIGatewayModelProviderTypeDeepseek:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDeepseek is populated
            case components.AIGatewayModelProviderTypeGemini:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderGemini is populated
            case components.AIGatewayModelProviderTypeHuggingface:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderHuggingface is populated
            case components.AIGatewayModelProviderTypeKimi:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderKimi is populated
            case components.AIGatewayModelProviderTypeLlama2:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderLlama2 is populated
            case components.AIGatewayModelProviderTypeMistral:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderMistral is populated
            case components.AIGatewayModelProviderTypeOllama:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOllama is populated
            case components.AIGatewayModelProviderTypeOpenai:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOpenai is populated
            case components.AIGatewayModelProviderTypeVercel:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVercel is populated
            case components.AIGatewayModelProviderTypeVllm:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVllm is populated
            case components.AIGatewayModelProviderTypeXai:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderXai is populated
            case components.AIGatewayModelProviderTypeVertex:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVertex is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `modelProviderIDOrName`                                  | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway model provider.  | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayModelProviderResponse](../../models/operations/getaigatewaymodelproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayModelProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the configuration of an existing AI Gateway model provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-model-provider" method="put" path="/v1/ai-gateways/{gatewayId}/model-providers/{modelProviderIdOrName}" -->
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

    res, err := s.AIGatewayModelProviders.UpdateAiGatewayModelProvider(ctx, operations.UpdateAiGatewayModelProviderRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ModelProviderIDOrName: "my-entity-name",
        UpdateAIGatewayModelProviderRequest: components.CreateUpdateAIGatewayModelProviderRequestDeepseek(
            components.AIGatewayModelProviderDeepseek{
                Type: components.AIGatewayModelProviderDeepseekTypeDeepseek,
                DisplayName: "Azure AI SE",
                Name: "azure-ai-se",
                Config: components.AIGatewayModelProviderDeepseekConfig{
                    Auth: components.AIGatewayModelProviderConfigAuthBasic{
                        Type: components.AIGatewayModelProviderConfigAuthBasicTypeBasic,
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayModelProvider != nil {
        switch res.AIGatewayModelProvider.Type {
            case components.AIGatewayModelProviderTypeAnthropic:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAnthropic is populated
            case components.AIGatewayModelProviderTypeAzure:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAzure is populated
            case components.AIGatewayModelProviderTypeBedrock:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderBedrock is populated
            case components.AIGatewayModelProviderTypeCerebras:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCerebras is populated
            case components.AIGatewayModelProviderTypeCohere:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCohere is populated
            case components.AIGatewayModelProviderTypeDashscope:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDashscope is populated
            case components.AIGatewayModelProviderTypeDatabricks:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDatabricks is populated
            case components.AIGatewayModelProviderTypeDeepseek:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDeepseek is populated
            case components.AIGatewayModelProviderTypeGemini:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderGemini is populated
            case components.AIGatewayModelProviderTypeHuggingface:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderHuggingface is populated
            case components.AIGatewayModelProviderTypeKimi:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderKimi is populated
            case components.AIGatewayModelProviderTypeLlama2:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderLlama2 is populated
            case components.AIGatewayModelProviderTypeMistral:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderMistral is populated
            case components.AIGatewayModelProviderTypeOllama:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOllama is populated
            case components.AIGatewayModelProviderTypeOpenai:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOpenai is populated
            case components.AIGatewayModelProviderTypeVercel:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVercel is populated
            case components.AIGatewayModelProviderTypeVllm:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVllm is populated
            case components.AIGatewayModelProviderTypeXai:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderXai is populated
            case components.AIGatewayModelProviderTypeVertex:
                // res.AIGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVertex is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateAiGatewayModelProviderRequest](../../models/operations/updateaigatewaymodelproviderrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateAiGatewayModelProviderResponse](../../models/operations/updateaigatewaymodelproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayModelProvider

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Removes a specific AI Gateway model provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-model-provider" method="delete" path="/v1/ai-gateways/{gatewayId}/model-providers/{modelProviderIdOrName}" -->
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

    res, err := s.AIGatewayModelProviders.DeleteAiGatewayModelProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `modelProviderIDOrName`                                  | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway model provider.  | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayModelProviderResponse](../../models/operations/deleteaigatewaymodelproviderresponse.md), error**

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