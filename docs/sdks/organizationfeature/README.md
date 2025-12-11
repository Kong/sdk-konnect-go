# OrganizationFeature

## Overview

### Available Operations

* [GetOrganizationFeature](#getorganizationfeature) - Get Feature Configuration
* [UpsertOrganizationFeature](#upsertorganizationfeature) - Upsert Feature Configuration

## GetOrganizationFeature

Returns the feature configuration of an organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-organization-feature" method="get" path="/v2/control-planes/features/{featureName}" -->
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

    res, err := s.OrganizationFeature.GetOrganizationFeature(ctx, components.FeatureNameEnterpriseAiPlugins)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationFeature != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `featureName`                                                    | [components.FeatureName](../../models/components/featurename.md) | :heavy_check_mark:                                               | The feature name for the configuration.                          |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.GetOrganizationFeatureResponse](../../models/operations/getorganizationfeatureresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertOrganizationFeature

Upsert Feature Configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-organization-feature" method="put" path="/v2/control-planes/features/{featureName}" -->
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

    res, err := s.OrganizationFeature.UpsertOrganizationFeature(ctx, components.FeatureNameEnterpriseAiPlugins, &components.UpsertOrganizationFeatureRequest{
        Enabled: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationFeature != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `featureName`                                                                                               | [components.FeatureName](../../models/components/featurename.md)                                            | :heavy_check_mark:                                                                                          | The feature name for the configuration.                                                                     |
| `upsertOrganizationFeatureRequest`                                                                          | [*components.UpsertOrganizationFeatureRequest](../../models/components/upsertorganizationfeaturerequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.UpsertOrganizationFeatureResponse](../../models/operations/upsertorganizationfeatureresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.BaseError         | 500, 503                    | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |