# APIPackageOperations

## Overview

### Available Operations

* [ListAPIPackagesOperations](#listapipackagesoperations) - List API Packages Operations
* [UpdateAPIPackageOperations](#updateapipackageoperations) - Update API Package Operation
* [GetAPIPackagesOperation](#getapipackagesoperation) - Get API Packages Operation
* [RemoveAPIPackageOperation](#removeapipackageoperation) - Remove API Package Operation

## ListAPIPackagesOperations

List API packages operations

### Example Usage

<!-- UsageSnippet language="go" operationID="list-api-packages-operations" method="get" path="/v3/api-packages/{packageId}/operations" -->
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

    res, err := s.APIPackageOperations.ListAPIPackagesOperations(ctx, operations.ListAPIPackagesOperationsRequest{
        PackageID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.APIPackageOperationsFilterParameters{
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
    if res.ListAPIPackagesOperationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListAPIPackagesOperationsRequest](../../models/operations/listapipackagesoperationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListAPIPackagesOperationsResponse](../../models/operations/listapipackagesoperationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPIPackageOperations

Update API Package Operations

### Example Usage

<!-- UsageSnippet language="go" operationID="update-api-package-operations" method="patch" path="/v3/api-packages/{packageId}/operations" -->
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

    res, err := s.APIPackageOperations.UpdateAPIPackageOperations(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", []components.APIPackageJSONPatchItem{
        components.APIPackageJSONPatchItem{
            Op: components.OpRemove,
            Path: "/var/spool",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `packageID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The UUID API Package identifier                                                            | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                                       |
| `requestBody`                                                                              | [][components.APIPackageJSONPatchItem](../../models/components/apipackagejsonpatchitem.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.UpdateAPIPackageOperationsResponse](../../models/operations/updateapipackageoperationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.APISlugConflict   | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAPIPackagesOperation

Get API packages operation

### Example Usage

<!-- UsageSnippet language="go" operationID="get-api-packages-operation" method="get" path="/v3/api-packages/{packageId}/operations/{operationId}" -->
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

    res, err := s.APIPackageOperations.GetAPIPackagesOperation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "9f5061ce-78f6-4452-9108-ad7c02821fd5")
    if err != nil {
        log.Fatal(err)
    }
    if res.APIPackageOperationResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `packageID`                                              | *string*                                                 | :heavy_check_mark:                                       | The UUID API Package identifier                          | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `operationID`                                            | *string*                                                 | :heavy_check_mark:                                       | The UUID API Operation identifier                        | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAPIPackagesOperationResponse](../../models/operations/getapipackagesoperationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RemoveAPIPackageOperation

Remove an operation from an API package.

### Example Usage

<!-- UsageSnippet language="go" operationID="remove-api-package-operation" method="delete" path="/v3/api-packages/{packageId}/operations/{operationId}" -->
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

    res, err := s.APIPackageOperations.RemoveAPIPackageOperation(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "9f5061ce-78f6-4452-9108-ad7c02821fd5")
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
| `operationID`                                            | *string*                                                 | :heavy_check_mark:                                       | The UUID API Operation identifier                        | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RemoveAPIPackageOperationResponse](../../models/operations/removeapipackageoperationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |