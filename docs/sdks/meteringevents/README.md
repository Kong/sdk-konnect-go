# MeteringEvents

## Overview

Metering events are used to track usage of your product or service. Events are processed asynchronously by the meters, so they may not be immediately available for querying.

### Available Operations

* [ListMeteringEvents](#listmeteringevents) - List metering events
* [IngestMeteringEvents](#ingestmeteringevents) - Ingest metering events

## ListMeteringEvents

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List ingested events.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-metering-events" method="get" path="/v3/openmeter/events" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/metering"
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

    res, err := s.MeteringEvents.ListMeteringEvents(ctx, operations.ListMeteringEventsRequest{
        Filter: &metering.ListEventsParamsFilter{
            CustomerID: sdkkonnectgo.Pointer(metering.CreateListEventsParamsFilterULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            Time: sdkkonnectgo.Pointer(metering.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                metering.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            IngestedAt: sdkkonnectgo.Pointer(metering.CreateListEventsParamsFilterDateTimeFieldFilterDateTimeFieldFilterDateTimeFieldLTEFilter(
                metering.DateTimeFieldFilterDateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            StoredAt: sdkkonnectgo.Pointer(metering.CreateListEventsParamsFilterMeteringDateTimeFieldFilterListEventsParamsFilterDateTimeFieldFilterDateTimeFieldEqualsFilter(
                metering.ListEventsParamsFilterDateTimeFieldFilterDateTimeFieldEqualsFilter{
                    Eq: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IngestedEventPaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListMeteringEventsRequest](../../models/operations/listmeteringeventsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListMeteringEventsResponse](../../models/operations/listmeteringeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## IngestMeteringEvents

Ingests an event or batch of events following the CloudEvents specification.

### Example Usage

<!-- UsageSnippet language="go" operationID="ingest-metering-events" method="post" path="/v3/openmeter/events" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
	"github.com/Kong/sdk-konnect-go/models/metering"
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

    res, err := s.MeteringEvents.IngestMeteringEvents(ctx, operations.CreateIngestMeteringEventsRequestBodyMeteringEvent(
        metering.MeteringEvent{
            ID: "5c10fade-1c9e-4d6c-8275-c52c36731d3c",
            Source: "service-name",
            Type: "prompt",
            Subject: "customer-id",
            Time: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
            Data: map[string]any{
                "prompt": "Hello, world!",
                "tokens": 100,
                "model": "gpt-4o",
                "type": "input",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.IngestMeteringEventsRequestBody](../../models/operations/ingestmeteringeventsrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.IngestMeteringEventsResponse](../../models/operations/ingestmeteringeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |