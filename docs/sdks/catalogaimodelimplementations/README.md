# CatalogAIModelImplementations

## Overview

Link an AI Model to an AI Gateway model.

### Available Operations

* [ListAiModelImplementations](#listaimodelimplementations) - List AI Model implementations
* [CreateAiModelImplementation](#createaimodelimplementation) - Create AI Model implementation
* [GetAiModelImplementation](#getaimodelimplementation) - Get AI Model implementation
* [DeleteAiModelImplementation](#deleteaimodelimplementation) - Delete AI Model implementation

## ListAiModelImplementations

Returns a paginated list of the AI Model's implementation records.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-model-implementations" method="get" path="/v1/ai-models/{aiModelId}/implementations" -->
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

    res, err := s.CatalogAIModelImplementations.ListAiModelImplementations(ctx, operations.ListAiModelImplementationsRequest{
        AiModelID: "123e4567-e89b-12d3-a456-426614174000",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelImplementationList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListAiModelImplementationsRequest](../../models/operations/listaimodelimplementationsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListAiModelImplementationsResponse](../../models/operations/listaimodelimplementationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateAiModelImplementation

Creates an implementation record for the AI Model. Returns `409` if the model has an existing implementation.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-model-implementation" method="post" path="/v1/ai-models/{aiModelId}/implementations" -->
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

    res, err := s.CatalogAIModelImplementations.CreateAiModelImplementation(ctx, "123e4567-e89b-12d3-a456-426614174000", components.AiModelImplementationCreate{
        GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
        GatewayModelID: "gw-model-abc",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelImplementation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `aiModelID`                                                                                      | `string`                                                                                         | :heavy_check_mark:                                                                               | The unique identifier of the AI Model.                                                           | 123e4567-e89b-12d3-a456-426614174000                                                             |
| `aiModelImplementationCreate`                                                                    | [components.AiModelImplementationCreate](../../models/components/aimodelimplementationcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateAiModelImplementationResponse](../../models/operations/createaimodelimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAiModelImplementation

Returns a single AI Model implementation record.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-model-implementation" method="get" path="/v1/ai-models/{aiModelId}/implementations/{implementationId}" -->
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

    res, err := s.CatalogAIModelImplementations.GetAiModelImplementation(ctx, "123e4567-e89b-12d3-a456-426614174000", "d2e1f0a9-8b7c-6d5e-4f3a-2b1c0d9e8f7a")
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelImplementation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `aiModelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model.                   | 123e4567-e89b-12d3-a456-426614174000                     |
| `implementationID`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model implementation.    | d2e1f0a9-8b7c-6d5e-4f3a-2b1c0d9e8f7a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiModelImplementationResponse](../../models/operations/getaimodelimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAiModelImplementation

Removes a single AI Model implementation (link) record by identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-model-implementation" method="delete" path="/v1/ai-models/{aiModelId}/implementations/{implementationId}" -->
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

    res, err := s.CatalogAIModelImplementations.DeleteAiModelImplementation(ctx, "123e4567-e89b-12d3-a456-426614174000", "d2e1f0a9-8b7c-6d5e-4f3a-2b1c0d9e8f7a")
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
| `implementationID`                                       | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model implementation.    | d2e1f0a9-8b7c-6d5e-4f3a-2b1c0d9e8f7a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiModelImplementationResponse](../../models/operations/deleteaimodelimplementationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |