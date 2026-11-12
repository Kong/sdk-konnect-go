# CatalogAIModelVersions

## Overview

Manage an AI Model's version.

### Available Operations

* [ListAiModelVersions](#listaimodelversions) - List AI Model versions
* [CreateAiModelVersion](#createaimodelversion) - Create AI Model version
* [GetLatestAiModelVersion](#getlatestaimodelversion) - Get latest AI Model version
* [UpdateLatestAiModelVersion](#updatelatestaimodelversion) - Update latest AI Model version
* [UpsertLatestAiModelVersion](#upsertlatestaimodelversion) - Replace latest AI Model version
* [DeleteLatestAiModelVersion](#deletelatestaimodelversion) - Delete latest AI Model version

## ListAiModelVersions

Returns the AI Model's versions.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-model-versions" method="get" path="/v1/ai-models/{aiModelId}/versions" -->
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

    res, err := s.CatalogAIModelVersions.ListAiModelVersions(ctx, operations.ListAiModelVersionsRequest{
        AiModelID: "123e4567-e89b-12d3-a456-426614174000",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersionList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAiModelVersionsRequest](../../models/operations/listaimodelversionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAiModelVersionsResponse](../../models/operations/listaimodelversionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateAiModelVersion

Creates the `latest` version with its `target_models`. Currently, a model only has a single version. Returns `409` if a version already exists.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-model-version" method="post" path="/v1/ai-models/{aiModelId}/versions" -->
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

    res, err := s.CatalogAIModelVersions.CreateAiModelVersion(ctx, "123e4567-e89b-12d3-a456-426614174000", components.AiModelVersionCreate{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        TargetModels: []components.TargetModel{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `aiModelID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | The unique identifier of the AI Model.                                             | 123e4567-e89b-12d3-a456-426614174000                                               |
| `aiModelVersionCreate`                                                             | [components.AiModelVersionCreate](../../models/components/aimodelversioncreate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreateAiModelVersionResponse](../../models/operations/createaimodelversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetLatestAiModelVersion

Returns the AI Model's `latest` version, including its `target_models`.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-latest-ai-model-version" method="get" path="/v1/ai-models/{aiModelId}/versions/latest" -->
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

    res, err := s.CatalogAIModelVersions.GetLatestAiModelVersion(ctx, "123e4567-e89b-12d3-a456-426614174000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersion != nil {
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

**[*operations.GetLatestAiModelVersionResponse](../../models/operations/getlatestaimodelversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateLatestAiModelVersion

Partially updates the `latest` version. `target_models`, when supplied, is a full-array replacement.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-latest-ai-model-version" method="patch" path="/v1/ai-models/{aiModelId}/versions/latest" -->
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

    res, err := s.CatalogAIModelVersions.UpdateLatestAiModelVersion(ctx, "123e4567-e89b-12d3-a456-426614174000", components.AiModelVersionUpdate{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        TargetModels: []components.TargetModel{
            components.TargetModel{
                Provider: "openai",
                Name: "gpt-4o",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `aiModelID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | The unique identifier of the AI Model.                                             | 123e4567-e89b-12d3-a456-426614174000                                               |
| `aiModelVersionUpdate`                                                             | [components.AiModelVersionUpdate](../../models/components/aimodelversionupdate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.UpdateLatestAiModelVersionResponse](../../models/operations/updatelatestaimodelversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertLatestAiModelVersion

Replaces the `latest` version.


### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-latest-ai-model-version" method="put" path="/v1/ai-models/{aiModelId}/versions/latest" -->
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

    res, err := s.CatalogAIModelVersions.UpsertLatestAiModelVersion(ctx, "123e4567-e89b-12d3-a456-426614174000", components.AiModelVersionCreate{
        Version: sdkkonnectgo.Pointer("1.0.0"),
        TargetModels: []components.TargetModel{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `aiModelID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | The unique identifier of the AI Model.                                             | 123e4567-e89b-12d3-a456-426614174000                                               |
| `aiModelVersionCreate`                                                             | [components.AiModelVersionCreate](../../models/components/aimodelversioncreate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.UpsertLatestAiModelVersionResponse](../../models/operations/upsertlatestaimodelversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteLatestAiModelVersion

Deletes the AI Model's `latest` version. Returns `404` if no version exists yet.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-latest-ai-model-version" method="delete" path="/v1/ai-models/{aiModelId}/versions/latest" -->
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

    res, err := s.CatalogAIModelVersions.DeleteLatestAiModelVersion(ctx, "123e4567-e89b-12d3-a456-426614174000")
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

**[*operations.DeleteLatestAiModelVersionResponse](../../models/operations/deletelatestaimodelversionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |