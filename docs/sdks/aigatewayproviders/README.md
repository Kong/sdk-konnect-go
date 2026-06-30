# AIGatewayProviders

## Overview

Providers that define the backend AI service connections for the AI Gateway.

### Available Operations

* [ListAiGatewayProviders](#listaigatewayproviders) - List AI Gateway Providers
* [CreateAiGatewayProvider](#createaigatewayprovider) - Create an AI Gateway Provider
* [GetAiGatewayProvider](#getaigatewayprovider) - Get an AI Gateway Provider
* [UpdateAiGatewayProvider](#updateaigatewayprovider) - Update an AI Gateway Provider
* [DeleteAiGatewayProvider](#deleteaigatewayprovider) - Delete an AI Gateway Provider

## ListAiGatewayProviders

Returns a list of providers configured for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-providers" method="get" path="/v1/ai-gateways/{gatewayId}/providers" -->
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

    res, err := s.AIGatewayProviders.ListAiGatewayProviders(ctx, operations.ListAiGatewayProvidersRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayProvidersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAiGatewayProvidersRequest](../../models/operations/listaigatewayprovidersrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAiGatewayProvidersResponse](../../models/operations/listaigatewayprovidersresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayProvider

Registers a new provider for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-provider" method="post" path="/v1/ai-gateways/{gatewayId}/providers" -->
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

    res, err := s.AIGatewayProviders.CreateAiGatewayProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayProviderRequestDatabricks(
        components.AIGatewayProviderDatabricks{
            Type: components.AIGatewayProviderDatabricksTypeDatabricks,
            DisplayName: "Azure AI SE",
            Name: "azure-ai-se",
            Config: components.AIGatewayProviderDatabricksConfig{
                Auth: components.AIGatewayProviderConfigAuthBasic{
                    Type: components.AIGatewayProviderConfigAuthBasicTypeBasic,
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayProvider != nil {
        switch res.AIGatewayProvider.Type {
            case components.AIGatewayProviderTypeAnthropic:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderAnthropic is populated
            case components.AIGatewayProviderTypeAzure:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderAzure is populated
            case components.AIGatewayProviderTypeBedrock:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderBedrock is populated
            case components.AIGatewayProviderTypeCerebras:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderCerebras is populated
            case components.AIGatewayProviderTypeCohere:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderCohere is populated
            case components.AIGatewayProviderTypeDashscope:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDashscope is populated
            case components.AIGatewayProviderTypeDatabricks:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDatabricks is populated
            case components.AIGatewayProviderTypeDeepseek:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDeepseek is populated
            case components.AIGatewayProviderTypeGemini:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderGemini is populated
            case components.AIGatewayProviderTypeHuggingface:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderHuggingface is populated
            case components.AIGatewayProviderTypeKimi:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderKimi is populated
            case components.AIGatewayProviderTypeLlama2:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderLlama2 is populated
            case components.AIGatewayProviderTypeMistral:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderMistral is populated
            case components.AIGatewayProviderTypeOllama:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderOllama is populated
            case components.AIGatewayProviderTypeOpenai:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderOpenai is populated
            case components.AIGatewayProviderTypeVercel:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVercel is populated
            case components.AIGatewayProviderTypeVllm:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVllm is populated
            case components.AIGatewayProviderTypeXai:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderXai is populated
            case components.AIGatewayProviderTypeVertex:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVertex is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `gatewayID`                                                                                            | `string`                                                                                               | :heavy_check_mark:                                                                                     | The unique ID of the AI Gateway.                                                                       | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                   |
| `createAIGatewayProviderRequest`                                                                       | [components.CreateAIGatewayProviderRequest](../../models/components/createaigatewayproviderrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.CreateAiGatewayProviderResponse](../../models/operations/createaigatewayproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayProvider

Returns the details of a specific AI Gateway provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-provider" method="get" path="/v1/ai-gateways/{gatewayId}/providers/{providerIdOrName}" -->
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

    res, err := s.AIGatewayProviders.GetAiGatewayProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayProvider != nil {
        switch res.AIGatewayProvider.Type {
            case components.AIGatewayProviderTypeAnthropic:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderAnthropic is populated
            case components.AIGatewayProviderTypeAzure:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderAzure is populated
            case components.AIGatewayProviderTypeBedrock:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderBedrock is populated
            case components.AIGatewayProviderTypeCerebras:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderCerebras is populated
            case components.AIGatewayProviderTypeCohere:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderCohere is populated
            case components.AIGatewayProviderTypeDashscope:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDashscope is populated
            case components.AIGatewayProviderTypeDatabricks:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDatabricks is populated
            case components.AIGatewayProviderTypeDeepseek:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDeepseek is populated
            case components.AIGatewayProviderTypeGemini:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderGemini is populated
            case components.AIGatewayProviderTypeHuggingface:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderHuggingface is populated
            case components.AIGatewayProviderTypeKimi:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderKimi is populated
            case components.AIGatewayProviderTypeLlama2:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderLlama2 is populated
            case components.AIGatewayProviderTypeMistral:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderMistral is populated
            case components.AIGatewayProviderTypeOllama:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderOllama is populated
            case components.AIGatewayProviderTypeOpenai:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderOpenai is populated
            case components.AIGatewayProviderTypeVercel:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVercel is populated
            case components.AIGatewayProviderTypeVllm:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVllm is populated
            case components.AIGatewayProviderTypeXai:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderXai is populated
            case components.AIGatewayProviderTypeVertex:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVertex is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `providerIDOrName`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway provider.        | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayProviderResponse](../../models/operations/getaigatewayproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayProvider

Updates the configuration of an existing AI Gateway provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-provider" method="put" path="/v1/ai-gateways/{gatewayId}/providers/{providerIdOrName}" -->
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

    res, err := s.AIGatewayProviders.UpdateAiGatewayProvider(ctx, operations.UpdateAiGatewayProviderRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ProviderIDOrName: "my-entity-name",
        UpdateAIGatewayProviderRequest: components.CreateUpdateAIGatewayProviderRequestDeepseek(
            components.AIGatewayProviderDeepseek{
                Type: components.AIGatewayProviderDeepseekTypeDeepseek,
                DisplayName: "Azure AI SE",
                Name: "azure-ai-se",
                Config: components.AIGatewayProviderDeepseekConfig{
                    Auth: components.AIGatewayProviderConfigAuthBasic{
                        Type: components.AIGatewayProviderConfigAuthBasicTypeBasic,
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayProvider != nil {
        switch res.AIGatewayProvider.Type {
            case components.AIGatewayProviderTypeAnthropic:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderAnthropic is populated
            case components.AIGatewayProviderTypeAzure:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderAzure is populated
            case components.AIGatewayProviderTypeBedrock:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderBedrock is populated
            case components.AIGatewayProviderTypeCerebras:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderCerebras is populated
            case components.AIGatewayProviderTypeCohere:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderCohere is populated
            case components.AIGatewayProviderTypeDashscope:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDashscope is populated
            case components.AIGatewayProviderTypeDatabricks:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDatabricks is populated
            case components.AIGatewayProviderTypeDeepseek:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderDeepseek is populated
            case components.AIGatewayProviderTypeGemini:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderGemini is populated
            case components.AIGatewayProviderTypeHuggingface:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderHuggingface is populated
            case components.AIGatewayProviderTypeKimi:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderKimi is populated
            case components.AIGatewayProviderTypeLlama2:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderLlama2 is populated
            case components.AIGatewayProviderTypeMistral:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderMistral is populated
            case components.AIGatewayProviderTypeOllama:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderOllama is populated
            case components.AIGatewayProviderTypeOpenai:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderOpenai is populated
            case components.AIGatewayProviderTypeVercel:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVercel is populated
            case components.AIGatewayProviderTypeVllm:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVllm is populated
            case components.AIGatewayProviderTypeXai:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderXai is populated
            case components.AIGatewayProviderTypeVertex:
                // res.AIGatewayProvider.AIGatewayProviderAIGatewayProviderVertex is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAiGatewayProviderRequest](../../models/operations/updateaigatewayproviderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAiGatewayProviderResponse](../../models/operations/updateaigatewayproviderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayProvider

Removes a specific AI Gateway provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-provider" method="delete" path="/v1/ai-gateways/{gatewayId}/providers/{providerIdOrName}" -->
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

    res, err := s.AIGatewayProviders.DeleteAiGatewayProvider(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `providerIDOrName`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway provider.        | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayProviderResponse](../../models/operations/deleteaigatewayproviderresponse.md), error**

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