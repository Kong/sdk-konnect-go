# APIPackageImage

## Overview

### Available Operations

* [UpsertAPIPackageImage](#upsertapipackageimage) - Create or Replace an API Package Image
* [FetchAPIPackageImage](#fetchapipackageimage) - Get API Package Image
* [DeleteAPIPackageImage](#deleteapipackageimage) - Delete an API Package Image
* [FetchAPIPackageRawImage](#fetchapipackagerawimage) - Get an API Package Raw Image

## UpsertAPIPackageImage

Create or Replace an API Package Image. The replacement image is not applied to the API Package until asynchronous validation completes successfully.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-api-package-image" method="put" path="/v3/api-packages/{packageId}/images/{imageType}" -->
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

    res, err := s.APIPackageImage.UpsertAPIPackageImage(ctx, operations.UpsertAPIPackageImageRequest{
        ImageType: components.ImageTypeSchemaIcon,
        PackageID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        ReplaceImageRequestSchema: components.ReplaceImageRequestSchema{
            Data: "data:image/png,YW5faW1hZ2VfZmlsZQ==",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APIPackageImage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpsertAPIPackageImageRequest](../../models/operations/upsertapipackageimagerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpsertAPIPackageImageResponse](../../models/operations/upsertapipackageimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIPackageImage

Retrieves the status and metadata associated with an API package Image. Returns the latest image if it has not yet been applied (status: uploading, validating, or invalid), otherwise returns the details of currently applied image (status: valid).

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-package-image" method="get" path="/v3/api-packages/{packageId}/images/{imageType}" -->
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

    res, err := s.APIPackageImage.FetchAPIPackageImage(ctx, components.ImageTypeSchemaIcon, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIPackageImage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `imageType`                                                              | [components.ImageTypeSchema](../../models/components/imagetypeschema.md) | :heavy_check_mark:                                                       | The Supported image type.                                                | icon                                                                     |
| `packageID`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | The UUID API Package identifier                                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.FetchAPIPackageImageResponse](../../models/operations/fetchapipackageimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAPIPackageImage

Delete an API Package Image.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-package-image" method="delete" path="/v3/api-packages/{packageId}/images/{imageType}" -->
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

    res, err := s.APIPackageImage.DeleteAPIPackageImage(ctx, components.ImageTypeSchemaIcon, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
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
| `imageType`                                                              | [components.ImageTypeSchema](../../models/components/imagetypeschema.md) | :heavy_check_mark:                                                       | The Supported image type.                                                | icon                                                                     |
| `packageID`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | The UUID API Package identifier                                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.DeleteAPIPackageImageResponse](../../models/operations/deleteapipackageimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPIPackageRawImage

Retrieves the raw image of an API Package. Only the currently applied image (status: valid) can be retrieved.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-api-package-raw-image" method="get" path="/v3/api-packages/{packageId}/images/{imageType}/raw" -->
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

    res, err := s.APIPackageImage.FetchAPIPackageRawImage(ctx, components.ImageTypeSchemaIcon, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
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
| `imageType`                                                              | [components.ImageTypeSchema](../../models/components/imagetypeschema.md) | :heavy_check_mark:                                                       | The Supported image type.                                                | icon                                                                     |
| `packageID`                                                              | `string`                                                                 | :heavy_check_mark:                                                       | The UUID API Package identifier                                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.FetchAPIPackageRawImageResponse](../../models/operations/fetchapipackagerawimageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |