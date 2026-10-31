# CatalogAIModels

## Overview

Manage AI Models - a catalogable proxy for AI Gateway models

### Available Operations

* [ListAiModels](#listaimodels) - List AI Models
* [CreateAiModel](#createaimodel) - Create AI Model
* [GetAiModel](#getaimodel) - Get AI Model
* [UpdateAiModel](#updateaimodel) - Update AI Model
* [UpsertAiModel](#upsertaimodel) - Replace AI Model
* [DeleteAiModel](#deleteaimodel) - Delete AI Model

## ListAiModels

Returns a paginated list of AI Models for the current organization. Each row includes a `managed_by` object derived from whether the model is linked to an AI Gateway model.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-models" method="get" path="/v1/ai-models" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.CatalogAIModels.ListAiModels(ctx, operations.ListAiModelsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.AiModelFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListAiModelsRequest](../../models/operations/listaimodelsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListAiModelsResponse](../../models/operations/listaimodelsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateAiModel

Creates a bare AI Model. No `target_models` are accepted here — create the version separately with `POST /ai-models/{aiModelId}/versions`.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-model" method="post" path="/v1/ai-models" -->
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

    res, err := s.CatalogAIModels.CreateAiModel(ctx, components.AiModelCreate{
        Name: "my-ai-model",
        DisplayName: "My AI Model",
        Description: sdkkonnectgo.Pointer("Customer support triage model"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.AiModelCreate](../../models/components/aimodelcreate.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreateAiModelResponse](../../models/operations/createaimodelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAiModel

Returns a single AI Model.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-model" method="get" path="/v1/ai-models/{aiModelId}" -->
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

    res, err := s.CatalogAIModels.GetAiModel(ctx, "123e4567-e89b-12d3-a456-426614174000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `aiModelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model.                   | 123e4567-e89b-12d3-a456-426614174000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiModelResponse](../../models/operations/getaimodelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAiModel

Partially updates an AI Model.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-model" method="patch" path="/v1/ai-models/{aiModelId}" -->
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

    res, err := s.CatalogAIModels.UpdateAiModel(ctx, "123e4567-e89b-12d3-a456-426614174000", components.AiModelUpdate{
        Name: sdkkonnectgo.Pointer("my-ai-model"),
        DisplayName: sdkkonnectgo.Pointer("My AI Model"),
        Description: sdkkonnectgo.Pointer("Customer support triage model"),
        Labels: map[string]*string{
            "env": sdkkonnectgo.Pointer("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `aiModelID`                                                          | `string`                                                             | :heavy_check_mark:                                                   | The unique identifier of the AI Model.                               | 123e4567-e89b-12d3-a456-426614174000                                 |
| `aiModelUpdate`                                                      | [components.AiModelUpdate](../../models/components/aimodelupdate.md) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |

### Response

**[*operations.UpdateAiModelResponse](../../models/operations/updateaimodelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertAiModel

Replaces an AI Model.


### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-ai-model" method="put" path="/v1/ai-models/{aiModelId}" -->
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

    res, err := s.CatalogAIModels.UpsertAiModel(ctx, "123e4567-e89b-12d3-a456-426614174000", components.AiModelReplace{
        Name: "my-ai-model",
        DisplayName: "My AI Model",
        Description: sdkkonnectgo.Pointer("Customer support triage model"),
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `aiModelID`                                                            | `string`                                                               | :heavy_check_mark:                                                     | The unique identifier of the AI Model.                                 | 123e4567-e89b-12d3-a456-426614174000                                   |
| `aiModelReplace`                                                       | [components.AiModelReplace](../../models/components/aimodelreplace.md) | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |                                                                        |

### Response

**[*operations.UpsertAiModelResponse](../../models/operations/upsertaimodelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAiModel

Deletes an AI Model.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-model" method="delete" path="/v1/ai-models/{aiModelId}" -->
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

    res, err := s.CatalogAIModels.DeleteAiModel(ctx, "123e4567-e89b-12d3-a456-426614174000")
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
| `aiModelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model.                   | 123e4567-e89b-12d3-a456-426614174000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiModelResponse](../../models/operations/deleteaimodelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |