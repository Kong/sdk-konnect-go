# APIPackageSpecification

## Overview

### Available Operations

* [GetAPIPackageComputedSpecification](#getapipackagecomputedspecification) - Get the API package computed specification
* [UpdateAPIPackageCurrentSpecification](#updateapipackagecurrentspecification) - Update API package current specification
* [GetAPIPackageCurrentSpecification](#getapipackagecurrentspecification) - Get the API package current specification
* [DeleteAPIPackageCurrentSpecification](#deleteapipackagecurrentspecification) - Delete API Package current Specification

## GetAPIPackageComputedSpecification

Fetches the computed specification of an API Package.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-api-package-computed-specification" method="get" path="/v3/api-packages/{packageId}/computed-specification" -->
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

    res, err := s.APIPackageSpecification.GetAPIPackageComputedSpecification(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `packageID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID API Package identifier                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAPIPackageComputedSpecificationResponse](../../models/operations/getapipackagecomputedspecificationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPIPackageCurrentSpecification

Updates the current specification of an API package.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-api-package-current-specification" method="put" path="/v3/api-packages/{packageId}/current-specification" -->
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

    res, err := s.APIPackageSpecification.UpdateAPIPackageCurrentSpecification(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.APISpec{
        Content: sdkkonnectgo.Pointer("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
        Type: components.APISpecAPISpecTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `packageID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID API Package identifier                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `apiSpec`                                                | [components.APISpec](../../models/components/apispec.md) | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.UpdateAPIPackageCurrentSpecificationResponse](../../models/operations/updateapipackagecurrentspecificationresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.APISlugConflict           | 409                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## GetAPIPackageCurrentSpecification

Fetches the current specification of an API Package.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-api-package-current-specification" method="get" path="/v3/api-packages/{packageId}/current-specification" -->
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

    res, err := s.APIPackageSpecification.GetAPIPackageCurrentSpecification(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `packageID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID API Package identifier                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAPIPackageCurrentSpecificationResponse](../../models/operations/getapipackagecurrentspecificationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAPIPackageCurrentSpecification

Deletes the current specification of an API Package.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-package-current-specification" method="delete" path="/v3/api-packages/{packageId}/current-specification" -->
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

    res, err := s.APIPackageSpecification.DeleteAPIPackageCurrentSpecification(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5")
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
| `packageID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID API Package identifier                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPIPackageCurrentSpecificationResponse](../../models/operations/deleteapipackagecurrentspecificationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |