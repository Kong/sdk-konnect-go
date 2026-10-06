# CatalogAIModelSpecs

## Overview

Manage an AI Model version's oas specification.

### Available Operations

* [GetAiModelSpec](#getaimodelspec) - Get AI Model spec
* [GetAiModelVersionSpec](#getaimodelversionspec) - Get AI Model version spec
* [UpsertAiModelVersionSpec](#upsertaimodelversionspec) - Upsert AI Model version spec
* [DeleteAiModelVersionSpec](#deleteaimodelversionspec) - Delete AI Model version spec

## GetAiModelSpec

Convenience shortcut to the `latest` version's oas specification. Returns `404`  if either the version or the specification is missing, distinguished by error code in the response body (`ai_model_version_not_found` / `ai_model_spec_not_found`).


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-model-spec" method="get" path="/v1/ai-models/{aiModelId}/spec" -->
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

    res, err := s.CatalogAIModelSpecs.GetAiModelSpec(ctx, "123e4567-e89b-12d3-a456-426614174000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersionSpec != nil {
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

**[*operations.GetAiModelSpecResponse](../../models/operations/getaimodelspecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAiModelVersionSpec

Returns the `latest` version's oas specification contents along with metadata.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-model-version-spec" method="get" path="/v1/ai-models/{aiModelId}/versions/latest/spec" -->
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

    res, err := s.CatalogAIModelSpecs.GetAiModelVersionSpec(ctx, "123e4567-e89b-12d3-a456-426614174000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersionSpec != nil {
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

**[*operations.GetAiModelVersionSpecResponse](../../models/operations/getaimodelversionspecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertAiModelVersionSpec

Creates or replaces the `latest` version's oas specification (upsert).


### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-ai-model-version-spec" method="put" path="/v1/ai-models/{aiModelId}/versions/latest/spec" -->
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

    res, err := s.CatalogAIModelSpecs.UpsertAiModelVersionSpec(ctx, "123e4567-e89b-12d3-a456-426614174000", components.AiModelVersionSpecWrite{
        SpecContent: "{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"My AI Model\",\"version\":\"1.0.0\"},\"paths\":{}}",
        SpecProvider: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersionSpec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `aiModelID`                                                                              | `string`                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the AI Model.                                                   | 123e4567-e89b-12d3-a456-426614174000                                                     |
| `aiModelVersionSpecWrite`                                                                | [components.AiModelVersionSpecWrite](../../models/components/aimodelversionspecwrite.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.UpsertAiModelVersionSpecResponse](../../models/operations/upsertaimodelversionspecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAiModelVersionSpec

Deletes the `latest` version's oas specification.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-model-version-spec" method="delete" path="/v1/ai-models/{aiModelId}/versions/latest/spec" -->
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

    res, err := s.CatalogAIModelSpecs.DeleteAiModelVersionSpec(ctx, "123e4567-e89b-12d3-a456-426614174000")
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

**[*operations.DeleteAiModelVersionSpecResponse](../../models/operations/deleteaimodelversionspecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |