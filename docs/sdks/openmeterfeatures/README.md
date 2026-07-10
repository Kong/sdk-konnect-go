# OpenMeterFeatures

## Overview

Features represent product capabilities backed by meters, with optional per-unit cost configuration.

### Available Operations

* [ListFeatures](#listfeatures) - List features
* [CreateFeature](#createfeature) - Create feature
* [GetFeature](#getfeature) - Get feature
* [UpdateFeature](#updatefeature) - Update feature
* [DeleteFeature](#deletefeature) - Delete feature
* [QueryFeatureCost](#queryfeaturecost) - Query feature cost

## ListFeatures

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List all features.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-features" method="get" path="/v3/openmeter/features" -->
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

    res, err := s.OpenMeterFeatures.ListFeatures(ctx, operations.ListFeaturesRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListFeatureParamsFilter{
            MeterID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FeaturePagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListFeaturesRequest](../../models/operations/listfeaturesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListFeaturesResponse](../../models/operations/listfeaturesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateFeature

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a feature.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-feature" method="post" path="/v3/openmeter/features" -->
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

    res, err := s.OpenMeterFeatures.CreateFeature(ctx, components.CreateFeatureRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Key: "resource_key",
        Meter: &components.CreateFeatureRequestMeterReference{
            ID: "01G65Z755AFWAKHE12NY0CQ9FH",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Feature != nil {
        switch res.Feature.UnitCost.Type {
            case components.UnitCostUnionTypeManual:
                // res.Feature.UnitCost.BillingFeatureManualUnitCost is populated
            case components.UnitCostUnionTypeLlm:
                // res.Feature.UnitCost.BillingFeatureLLMUnitCost is populated
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.CreateFeatureRequest](../../models/components/createfeaturerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateFeatureResponse](../../models/operations/createfeatureresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetFeature

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get a feature by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-feature" method="get" path="/v3/openmeter/features/{featureId}" -->
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

    res, err := s.OpenMeterFeatures.GetFeature(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.Feature != nil {
        switch res.Feature.UnitCost.Type {
            case components.UnitCostUnionTypeManual:
                // res.Feature.UnitCost.BillingFeatureManualUnitCost is populated
            case components.UnitCostUnionTypeLlm:
                // res.Feature.UnitCost.BillingFeatureLLMUnitCost is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `featureID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetFeatureResponse](../../models/operations/getfeatureresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateFeature

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update a feature by id. Currently only the unit_cost field can be updated.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-feature" method="patch" path="/v3/openmeter/features/{featureId}" -->
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

    res, err := s.OpenMeterFeatures.UpdateFeature(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpdateFeatureRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Feature != nil {
        switch res.Feature.UnitCost.Type {
            case components.UnitCostUnionTypeManual:
                // res.Feature.UnitCost.BillingFeatureManualUnitCost is populated
            case components.UnitCostUnionTypeLlm:
                // res.Feature.UnitCost.BillingFeatureLLMUnitCost is populated
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `featureID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | 01G65Z755AFWAKHE12NY0CQ9FH                                                         |
| `updateFeatureRequest`                                                             | [components.UpdateFeatureRequest](../../models/components/updatefeaturerequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.UpdateFeatureResponse](../../models/operations/updatefeatureresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteFeature

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Delete a feature by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-feature" method="delete" path="/v3/openmeter/features/{featureId}" -->
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

    res, err := s.OpenMeterFeatures.DeleteFeature(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `featureID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteFeatureResponse](../../models/operations/deletefeatureresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## QueryFeatureCost

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Query the cost of a feature.

### Example Usage

<!-- UsageSnippet language="go" operationID="query-feature-cost" method="post" path="/v3/openmeter/features/{featureId}/cost/query" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.OpenMeterFeatures.QueryFeatureCost(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", &components.MeterQueryRequest{
        From: types.MustNewTimeFromString("2023-01-01T00:00:00Z"),
        To: types.MustNewTimeFromString("2023-01-02T00:00:00Z"),
        Granularity: components.MeterQueryRequestGranularityP1D.ToPointer(),
        TimeZone: sdkkonnectgo.Pointer("UTC"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FeatureCostQueryResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `featureID`                                                                                                | `string`                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                 |
| `meterQueryRequest`                                                                                        | [*components.MeterQueryRequest](../../models/components/meterqueryrequest.md)                              | :heavy_minus_sign:                                                                                         | N/A                                                                                                        | {<br/>"from": "2023-01-01T00:00:00Z",<br/>"to": "2023-01-02T00:00:00Z",<br/>"granularity": "P1D",<br/>"time_zone": "UTC"<br/>} |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.QueryFeatureCostResponse](../../models/operations/queryfeaturecostresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |