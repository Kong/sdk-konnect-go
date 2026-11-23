# UsageEvents

## Overview

Usage events record each call to an AI model. These operations report and list the events, total them, forecast the month's spend, and list the values seen in each field.

### Available Operations

* [ListCostUsageEvents](#listcostusageevents) - List usage events
* [CreateCostUsageEvents](#createcostusageevents) - Report usage events
* [ListCostUsageFieldValues](#listcostusagefieldvalues) - List observed usage event field values
* [ForecastCostUsageEvents](#forecastcostusageevents) - Get the monthly spend forecast
* [QueryCostUsageEvents](#querycostusageevents) - Query usage events

## ListCostUsageEvents

Returns a page of the organization's usage events, newest first unless the request sets `sort`. The list includes usage from the gateway and usage reported with `POST /v1/ai-cost-management/events`, each in the `ai.usage.v1` form that producers report. For a caller whose access covers only particular principals or groups, the operation reads only the events of those principals and groups.

The service may return fewer events than requested when the page reaches its byte budget. Follow the page links to continue.

Listed events differ from what producers send:

- `arrival` isn't included, though `filter[arrival]=late` selects late events.
- A quantity or cost the producer didn't report shows as `"0"`.
- `datacontenttype`, `dataschema`, and provenance keys other than the ones this schema lists are left out.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-usage-events" method="get" path="/v1/ai-cost-management/events" -->
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

    res, err := s.UsageEvents.ListCostUsageEvents(ctx, operations.ListCostUsageEventsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCostUsageEventsRequest](../../models/operations/listcostusageeventsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListCostUsageEventsResponse](../../models/operations/listcostusageeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostUsageEvents

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Reports usage from AI model calls that didn't pass through the gateway, such as the model calls of an agent tool. Each event is an `ai.usage.v1` event without `source`. The service assigns `source` as `//principals/{uuid}` for the reporting principal. If an event includes `source`, the request fails with `400`.

- The service validates the whole request before it accepts anything. If any event is invalid, the request fails with `400` and no event is accepted.
- The service responds only after every event is queued for storage. It then stores the events asynchronously. They appear in the event list, where `filter[source]=//principals/{uuid}` selects them, and they count toward budgets like gateway usage.
- If the service can't confirm every event in time, the request fails with a server error, and some events may already be accepted. Retrying the same request is safe, because a repeated event collapses into the stored one. The reporting principal's `source`, the event's `id`, and its `time` identify a stored event. A retry that changes `id` or `time` stores a second event, and a retry that changes only other data replaces the stored event with whichever copy is stored last.
- An event that reaches the service more than 1 hour after its `time` is accepted, and `filter[arrival]=late` selects it.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-usage-events" method="post" path="/v1/ai-cost-management/events" -->
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

    res, err := s.UsageEvents.CreateCostUsageEvents(ctx, operations.CreateCostUsageEventsRequestBody{
        Data: []operations.Data{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateCostUsageEventsRequestBody](../../models/operations/createcostusageeventsrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateCostUsageEventsResponse](../../models/operations/createcostusageeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCostUsageFieldValues

Returns a page of the distinct values seen so far in one field of the organization's usage events, in ascending `value` order unless `sort` is `value desc`. A client can offer them as choices for usage filters.

Subjects come back as subject references, such as `principal/{uuid}`, which the usage filters accept as they are. For `field=tags`, a request without `key` lists the tag keys, and a request with `key` lists the values of that tag.

For a caller with access to the whole organization, a value stays in the list after the events that carried it leave retention or are deleted. For a caller whose access covers only particular principals or groups, the `subject` and `cost_center` lists hold only the subjects and cost centers of the events of those principals and groups, so a value leaves the list with its events. Such a caller's `groups` and `teams` lists hold only the groups the caller's access names. Every other field lists the organization's values, the same as for a caller with access to the whole organization.

The service may return fewer values than requested when the page reaches its byte budget. Follow the page links to continue.

The list doesn't take a time range or other usage filters, and it doesn't count usage. A cursor belongs to its `field` and `key`. With a cursor from another selection, the request fails with `400`.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-usage-field-values" method="get" path="/v1/ai-cost-management/events/field-values" -->
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

    res, err := s.UsageEvents.ListCostUsageFieldValues(ctx, operations.ListCostUsageFieldValuesRequest{
        Field: operations.FieldModelID,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListCostUsageFieldValuesRequest](../../models/operations/listcostusagefieldvaluesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListCostUsageFieldValuesResponse](../../models/operations/listcostusagefieldvaluesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ForecastCostUsageEvents

Returns the organization's recorded and predicted spend, in US dollars, for each day of the current calendar month in `time_zone`. For a caller whose access covers only particular principals or groups, the operation reads only the events of those principals and groups. The prediction trains on up to 28 completed days, and the day in progress never counts. Each series covers at most one calendar month, so the response isn't paginated.

Predictions are statistical estimates. They don't check or enforce any budget.

### Example Usage

<!-- UsageSnippet language="go" operationID="forecast-cost-usage-events" method="get" path="/v1/ai-cost-management/events/forecast" -->
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

    res, err := s.UsageEvents.ForecastCostUsageEvents(ctx, operations.MethodAuto.ToPointer(), sdkkonnectgo.Pointer("UTC"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `method`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [*operations.Method](../../models/operations/method.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The prediction method:<br/><br/>- `auto`: `weekly_seasonal` once 14 completed training days exist, and `exponential_smoothing` before that.<br/>- `weekly_seasonal`: the mean daily spend for each weekday. Needs 14 completed training days.<br/>- `linear_regression`: a straight-line fit of daily spend, extended to the end of the month. Needs 2 completed training days.<br/>- `exponential_smoothing`: a weighted average of recent daily spend, carried forward. Needs 2 completed training days.<br/><br/>A method named in the request never falls back to another when the history is too short. The response's `status` is then `insufficient_history`. Defaults to `auto`. |
| `timeZone`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The IANA time zone whose calendar sets the month and its days.<br/><br/>The name must be spelled as the time zone database spells it, such as `Europe/Budapest`, not `europe/budapest`. Offsets such as `+02:00` aren't accepted. Defaults to `UTC`.                                                                                                                                                                                                                                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Response

**[*operations.ForecastCostUsageEventsResponse](../../models/operations/forecastcostusageeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## QueryCostUsageEvents

Returns totals over the organization's usage events in a time range, grouped by up to 2 dimensions and optionally split into calendar buckets. Groups come largest `reported_cost` first unless the request sets `sort`. For a caller whose access covers only particular principals or groups, the operation reads only the events of those principals and groups.

The operation doesn't change anything; the query goes in the body because it can be large. Each page holds `page.size` groups with their complete `series`. To read the next or previous page, a client sends the response's `meta.page.next` or `meta.page.previous` cursor back as `page.after` or `page.before` in an otherwise identical body.

The totals come from the same stored events as the usage event list and budget balances, so a query and a budget balance over the same scope and range report the same numbers. A redelivered copy of an event counts in the totals until it is merged into the stored event, so for a short time the totals can exceed the usage event list.

### Example Usage

<!-- UsageSnippet language="go" operationID="query-cost-usage-events" method="post" path="/v1/ai-cost-management/events/query" -->
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

    res, err := s.UsageEvents.QueryCostUsageEvents(ctx, operations.QueryCostUsageEventsRequestBody{
        StartsAt: types.MustTimeFromString("2026-08-19T09:13:20.040Z"),
        EndsAt: types.MustTimeFromString("2026-10-29T16:57:13.664Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.QueryCostUsageEventsRequestBody](../../models/operations/querycostusageeventsrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.QueryCostUsageEventsResponse](../../models/operations/querycostusageeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |