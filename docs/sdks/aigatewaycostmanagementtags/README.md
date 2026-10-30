# AIGatewayCostManagementTags

## Overview

Configure which request headers are captured as cost tags on AI usage events. The configuration applies to the organization in the region the request is sent to.

### Available Operations

* [GetCostManagementTags](#getcostmanagementtags) - Get cost management tags
* [UpdateCostManagementTags](#updatecostmanagementtags) - Update cost management tags

## GetCostManagementTags

Returns the cost management tags configuration of the organization in this region.
If no configuration has been set, the defaults are returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-management-tags" method="get" path="/v1/ai-cost-management/tag-capture" -->
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

    res, err := s.AIGatewayCostManagementTags.GetCostManagementTags(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CostManagementTags != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostManagementTagsResponse](../../models/operations/getcostmanagementtagsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateCostManagementTags

Replaces the cost management tags configuration of the organization in this region.
Header names and prefixes are compared case-insensitively.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-cost-management-tags" method="put" path="/v1/ai-cost-management/tag-capture" -->
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

    res, err := s.AIGatewayCostManagementTags.UpdateCostManagementTags(ctx, components.UpdateCostManagementTags{
        Headers: []components.CostTagHeader{},
        Prefixes: []components.CostTagPrefix{
            components.CostTagPrefix{
                Name: "my-co",
            },
        },
        Labels: map[string]string{
            "category": "finance",
        },
        ManagedBy: map[string]string{
            "owner": "terraform",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CostManagementTags != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.UpdateCostManagementTags](../../models/components/updatecostmanagementtags.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateCostManagementTagsResponse](../../models/operations/updatecostmanagementtagsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |