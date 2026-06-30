# AIGatewayModels

## Overview

Models that define routing, capabilities, and backend targets for the AI Gateway.

### Available Operations

* [ListAiGatewayModels](#listaigatewaymodels) - List AI Gateway Models
* [CreateAiGatewayModel](#createaigatewaymodel) - Create an AI Gateway Model
* [GetAiGatewayModel](#getaigatewaymodel) - Get an AI Gateway Model
* [UpdateAiGatewayModel](#updateaigatewaymodel) - Update an AI Gateway Model
* [DeleteAiGatewayModel](#deleteaigatewaymodel) - Delete an AI Gateway Model

## ListAiGatewayModels

Returns a list of all models registered in the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-models" method="get" path="/v1/ai-gateways/{gatewayId}/models" -->
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

    res, err := s.AIGatewayModels.ListAiGatewayModels(ctx, operations.ListAiGatewayModelsRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayModelsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAiGatewayModelsRequest](../../models/operations/listaigatewaymodelsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAiGatewayModelsResponse](../../models/operations/listaigatewaymodelsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayModel

Registers a new model with routing, capabilities, and target backends.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-model" method="post" path="/v1/ai-gateways/{gatewayId}/models" -->
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

    res, err := s.AIGatewayModels.CreateAiGatewayModel(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayModelRequestAPI(
        components.AIGatewayModelAPI{
            DisplayName: "My GPT 5 model",
            Name: "my-gpt-5-model",
            Formats: []components.AIGatewayModelFormat{
                components.AIGatewayModelFormat{
                    Type: components.AIGatewayModelFormatTypeOpenai.ToPointer(),
                },
            },
            Targets: []components.AIGatewayTarget{},
            Type: components.AIGatewayModelAPITypeAPI,
            Config: components.AIGatewayModelAPIConfig{
                Route: components.AIGatewayRouteConfig{
                    Headers: map[string]any{
                        "version": []any{
                            "v1",
                            "v2",
                        },
                    },
                    Hosts: []string{
                        "foo.example.com",
                    },
                },
                Balancer: sdkkonnectgo.Pointer(components.CreateAIGatewayModelBalancerConfigLowestUsage(
                    components.AIGatewayModelBalancerLowestUsageConfig{
                        Algorithm: components.AIGatewayModelBalancerLowestUsageConfigAlgorithmLowestUsage,
                    },
                )),
                Model: components.AIGatewayModelAPIConfigModel{},
            },
            Capabilities: []components.Capabilities{
                components.CapabilitiesFiles,
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayModel != nil {
        switch res.AIGatewayModel.Type {
            case components.AIGatewayModelTypeAPI:
                // res.AIGatewayModel.AIGatewayModelAIGatewayModelAPI is populated
            case components.AIGatewayModelTypeModel:
                // res.AIGatewayModel.AIGatewayModelAIGatewayModelModel is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `gatewayID`                                                                                      | `string`                                                                                         | :heavy_check_mark:                                                                               | The unique ID of the AI Gateway.                                                                 | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                             |
| `createAIGatewayModelRequest`                                                                    | [components.CreateAIGatewayModelRequest](../../models/components/createaigatewaymodelrequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateAiGatewayModelResponse](../../models/operations/createaigatewaymodelresponse.md), error**

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

## GetAiGatewayModel

Returns the details of a specific AI Gateway model.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-model" method="get" path="/v1/ai-gateways/{gatewayId}/models/{modelIdOrName}" -->
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

    res, err := s.AIGatewayModels.GetAiGatewayModel(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayModel != nil {
        switch res.AIGatewayModel.Type {
            case components.AIGatewayModelTypeAPI:
                // res.AIGatewayModel.AIGatewayModelAIGatewayModelAPI is populated
            case components.AIGatewayModelTypeModel:
                // res.AIGatewayModel.AIGatewayModelAIGatewayModelModel is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `modelIDOrName`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway model.           | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayModelResponse](../../models/operations/getaigatewaymodelresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayModel

Updates the configuration of an existing AI Gateway model.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-model" method="put" path="/v1/ai-gateways/{gatewayId}/models/{modelIdOrName}" -->
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

    res, err := s.AIGatewayModels.UpdateAiGatewayModel(ctx, operations.UpdateAiGatewayModelRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        ModelIDOrName: "my-entity-name",
        UpdateAIGatewayModelRequest: components.CreateUpdateAIGatewayModelRequestAPI(
            components.AIGatewayModelAPI{
                DisplayName: "My GPT 5 model",
                Name: "my-gpt-5-model",
                Formats: []components.AIGatewayModelFormat{},
                Targets: []components.AIGatewayTarget{
                    components.AIGatewayTarget{
                        Name: "gpt-5-model",
                        Provider: "azure-ai-se",
                        Config: components.CreateAIGatewayTargetConfigOpenai(
                            components.AIGatewayTargetOpenaiConfig{
                                Type: components.AIGatewayTargetOpenaiConfigTypeOpenai,
                            },
                        ),
                    },
                },
                Type: components.AIGatewayModelAPITypeAPI,
                Config: components.AIGatewayModelAPIConfig{
                    Route: components.AIGatewayRouteConfig{
                        Headers: map[string]any{
                            "version": []any{
                                "v1",
                                "v2",
                            },
                        },
                        Hosts: []string{
                            "foo.example.com",
                        },
                    },
                    Balancer: sdkkonnectgo.Pointer(components.CreateAIGatewayModelBalancerConfigLowestUsage(
                        components.AIGatewayModelBalancerLowestUsageConfig{
                            Algorithm: components.AIGatewayModelBalancerLowestUsageConfigAlgorithmLowestUsage,
                        },
                    )),
                    Model: components.AIGatewayModelAPIConfigModel{},
                },
                Capabilities: []components.Capabilities{
                    components.CapabilitiesFiles,
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayModel != nil {
        switch res.AIGatewayModel.Type {
            case components.AIGatewayModelTypeAPI:
                // res.AIGatewayModel.AIGatewayModelAIGatewayModelAPI is populated
            case components.AIGatewayModelTypeModel:
                // res.AIGatewayModel.AIGatewayModelAIGatewayModelModel is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAiGatewayModelRequest](../../models/operations/updateaigatewaymodelrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAiGatewayModelResponse](../../models/operations/updateaigatewaymodelresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayModel

Removes a specific AI Gateway model.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-model" method="delete" path="/v1/ai-gateways/{gatewayId}/models/{modelIdOrName}" -->
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

    res, err := s.AIGatewayModels.DeleteAiGatewayModel(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `modelIDOrName`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway model.           | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayModelResponse](../../models/operations/deleteaigatewaymodelresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |