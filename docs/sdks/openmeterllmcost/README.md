# OpenMeterLLMCost

## Overview

LLM cost database providing normalized pricing data for language models across providers. Prices are synced from multiple external sources and auto-confirmed when they agree.

### Available Operations

* [ListLlmCostOverrides](#listllmcostoverrides) - List LLM cost overrides
* [CreateLlmCostOverride](#createllmcostoverride) - Create LLM cost override
* [DeleteLlmCostOverride](#deletellmcostoverride) - Delete LLM cost override
* [ListLlmCostPrices](#listllmcostprices) - List LLM cost prices
* [GetLlmCostPrice](#getllmcostprice) - Get LLM cost price

## ListLlmCostOverrides

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List per-namespace price overrides.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-llm-cost-overrides" method="get" path="/v3/openmeter/llm-cost/overrides" -->
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

    res, err := s.OpenMeterLLMCost.ListLlmCostOverrides(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PricePagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `filter`                                                                                          | [*metering.ListLLMCostPricesParamsFilter](../../models/metering/listllmcostpricesparamsfilter.md) | :heavy_minus_sign:                                                                                | Filter options for listing LLM cost prices.                                                       |
| `page`                                                                                            | [*metering.PagePaginationQuery](../../models/metering/pagepaginationquery.md)                     | :heavy_minus_sign:                                                                                | Determines which page of the collection to retrieve.                                              |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.ListLlmCostOverridesResponse](../../models/operations/listllmcostoverridesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateLlmCostOverride

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a per-namespace price override.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-llm-cost-override" method="post" path="/v3/openmeter/llm-cost/overrides" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/metering"
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

    res, err := s.OpenMeterLLMCost.CreateLlmCostOverride(ctx, metering.LLMCostOverrideCreate{
        Provider: "<value>",
        ModelID: "<id>",
        Pricing: metering.LLMCostOverrideCreatePricing{
            InputPerToken: "<value>",
            OutputPerToken: "<value>",
        },
        Currency: "USD",
        EffectiveFrom: types.MustTimeFromString("2023-01-01T01:01:01.001Z"),
        EffectiveTo: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LLMCostPrice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [metering.LLMCostOverrideCreate](../../models/metering/llmcostoverridecreate.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateLlmCostOverrideResponse](../../models/operations/createllmcostoverrideresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteLlmCostOverride

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Delete a per-namespace price override.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-llm-cost-override" method="delete" path="/v3/openmeter/llm-cost/overrides/{priceId}" -->
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

    res, err := s.OpenMeterLLMCost.DeleteLlmCostOverride(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `priceID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteLlmCostOverrideResponse](../../models/operations/deletellmcostoverrideresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListLlmCostPrices

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List global LLM cost prices. Returns prices with overrides applied if any.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-llm-cost-prices" method="get" path="/v3/openmeter/llm-cost/prices" -->
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

    res, err := s.OpenMeterLLMCost.ListLlmCostPrices(ctx, operations.ListLlmCostPricesRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PricePagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLlmCostPricesRequest](../../models/operations/listllmcostpricesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListLlmCostPricesResponse](../../models/operations/listllmcostpricesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetLlmCostPrice

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get a specific LLM cost price by ID. Returns the price with overrides applied if
any.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-llm-cost-price" method="get" path="/v3/openmeter/llm-cost/prices/{priceId}" -->
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

    res, err := s.OpenMeterLLMCost.GetLlmCostPrice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.LLMCostPrice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `priceID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetLlmCostPriceResponse](../../models/operations/getllmcostpriceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |