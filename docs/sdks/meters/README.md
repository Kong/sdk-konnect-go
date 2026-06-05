# Meters

## Overview

Meters specify how to aggregate events for billing and analytics purposes. Meters can be configured with multiple aggregation methods and groupings. Multiple meters can be created for the same event type, enabling flexible metering scenarios.

### Available Operations

* [CreateMeter](#createmeter) - Create meter
* [ListMeters](#listmeters) - List meters
* [GetMeter](#getmeter) - Get meter
* [UpdateMeter](#updatemeter) - Update meter
* [DeleteMeter](#deletemeter) - Delete meter
* [QueryMeter](#querymeter) - Query meter

## CreateMeter

Create a meter.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-meter" method="post" path="/v3/openmeter/meters" -->
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

    res, err := s.Meters.CreateMeter(ctx, components.CreateMeterRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Key: "resource_key",
        Aggregation: components.AggregationSum,
        EventType: "prompt",
        EventsFrom: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
        ValueProperty: sdkkonnectgo.Pointer("$.tokens"),
        Dimensions: map[string]string{
            "type": "$.type",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Meter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.CreateMeterRequest](../../models/components/createmeterrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateMeterResponse](../../models/operations/createmeterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListMeters

List meters.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-meters" method="get" path="/v3/openmeter/meters" -->
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

    res, err := s.Meters.ListMeters(ctx, operations.ListMetersRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeterPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListMetersRequest](../../models/operations/listmetersrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListMetersResponse](../../models/operations/listmetersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMeter

Get a meter by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-meter" method="get" path="/v3/openmeter/meters/{meterId}" -->
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

    res, err := s.Meters.GetMeter(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.Meter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `meterID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMeterResponse](../../models/operations/getmeterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateMeter

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update a meter.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-meter" method="put" path="/v3/openmeter/meters/{meterId}" -->
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

    res, err := s.Meters.UpdateMeter(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpdateMeterRequest{
        Labels: map[string]string{
            "env": "test",
        },
        Dimensions: map[string]string{
            "type": "$.type",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Meter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `meterID`                                                                      | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | 01G65Z755AFWAKHE12NY0CQ9FH                                                     |
| `updateMeterRequest`                                                           | [components.UpdateMeterRequest](../../models/components/updatemeterrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.UpdateMeterResponse](../../models/operations/updatemeterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMeter

Delete a meter.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-meter" method="delete" path="/v3/openmeter/meters/{meterId}" -->
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

    res, err := s.Meters.DeleteMeter(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `meterID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteMeterResponse](../../models/operations/deletemeterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## QueryMeter

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Query a meter for usage.

Set `Accept: application/json` (the default) to get a structured JSON response.
Set `Accept: text/csv` to download the same data as a CSV file suitable for
spreadsheets. The CSV columns, in order, are:

`from, to, [subject,] [customer_id, customer_key, customer_name,] <dimensions...>, value`

The `subject` column is emitted only when `subject` is in the query's
`group_by_dimensions`. The three `customer_*` columns are emitted together only
when `customer_id` is in the query's `group_by_dimensions`.

### Example Usage

<!-- UsageSnippet language="go" operationID="query-meter" method="post" path="/v3/openmeter/meters/{meterId}/query" -->
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

    res, err := s.Meters.QueryMeter(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.MeterQueryRequest{
        From: types.MustNewTimeFromString("2023-01-01T00:00:00Z"),
        To: types.MustNewTimeFromString("2023-01-02T00:00:00Z"),
        Granularity: components.MeterQueryRequestGranularityP1D.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeterQueryResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `meterID`                                                                                                  | `string`                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                 |
| `meterQueryRequest`                                                                                        | [components.MeterQueryRequest](../../models/components/meterqueryrequest.md)                               | :heavy_check_mark:                                                                                         | N/A                                                                                                        | {<br/>"from": "2023-01-01T00:00:00Z",<br/>"to": "2023-01-02T00:00:00Z",<br/>"granularity": "P1D",<br/>"time_zone": "UTC"<br/>} |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.QueryMeterResponse](../../models/operations/querymeterresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |