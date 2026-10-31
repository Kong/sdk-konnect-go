# APIImage

## Overview

### Available Operations

* [UpsertAPIImage](#upsertapiimage) - Create or Replace an API Image
* [FetchAPIImage](#fetchapiimage) - Get API Image
* [DeleteAPIImage](#deleteapiimage) - Delete an API Image
* [FetchAPIRawImage](#fetchapirawimage) - Get an API Raw Image

## UpsertAPIImage

Create or Replace an API Image. The replacement image is not applied to the API until asynchronous validation completes successfully.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-api-image" method="put" path="/v3/apis/{apiId}/images/{imageType}" -->
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

    res, err := s.APIImage.UpsertAPIImage(ctx, operations.UpsertAPIImageRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        ImageType: components.ImageTypeSchemaIcon,
        ReplaceImageRequestSchema: components.ReplaceImageRequestSchema{
            Data: "data:image/png,YW5faW1hZ2VfZmlsZQ==",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpsertAPIImageRequest](../../models/operations/upsertapiimagerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpsertAPIImageResponse](../../models/operations/upsertapiimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIImage

Retrieves the status and metadata associated with an API Image. Returns the latest image if it has not yet been applied (status: uploading, validating, or invalid), otherwise returns the details of currently applied image (status: valid).

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-image" method="get" path="/v3/apis/{apiId}/images/{imageType}" -->
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

    res, err := s.APIImage.FetchAPIImage(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.ImageTypeSchemaIcon)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIImage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `apiID`                                                                  | `string`                                                                 | :heavy_check_mark:                                                       | The UUID API identifier                                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                     |
| `imageType`                                                              | [components.ImageTypeSchema](../../models/components/imagetypeschema.md) | :heavy_check_mark:                                                       | The Supported image type.                                                | icon                                                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.FetchAPIImageResponse](../../models/operations/fetchapiimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAPIImage

Delete an API Image.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-image" method="delete" path="/v3/apis/{apiId}/images/{imageType}" -->
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

    res, err := s.APIImage.DeleteAPIImage(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.ImageTypeSchemaIcon)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `apiID`                                                                  | `string`                                                                 | :heavy_check_mark:                                                       | The UUID API identifier                                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                     |
| `imageType`                                                              | [components.ImageTypeSchema](../../models/components/imagetypeschema.md) | :heavy_check_mark:                                                       | The Supported image type.                                                | icon                                                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.DeleteAPIImageResponse](../../models/operations/deleteapiimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIRawImage

Retrieves the raw image of an API. Only the currently applied image (status: valid) can be retrieved.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-raw-image" method="get" path="/v3/apis/{apiId}/images/{imageType}/raw" -->
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

    res, err := s.APIImage.FetchAPIRawImage(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.ImageTypeSchemaIcon)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `apiID`                                                                  | `string`                                                                 | :heavy_check_mark:                                                       | The UUID API identifier                                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                     |
| `imageType`                                                              | [components.ImageTypeSchema](../../models/components/imagetypeschema.md) | :heavy_check_mark:                                                       | The Supported image type.                                                | icon                                                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.FetchAPIRawImageResponse](../../models/operations/fetchapirawimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |