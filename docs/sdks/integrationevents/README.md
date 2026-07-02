# IntegrationEvents

## Overview

Integration events represent the activity within an integration instance.


### Available Operations

* [ListIntegrationEvents](#listintegrationevents) - List Integration Events
* [CreateIntegrationEvents](#createintegrationevents) - Bulk Create Integration Events
* [ListCatalogServiceIntegrationEvents](#listcatalogserviceintegrationevents) - List a catalog service integration Events

## ListIntegrationEvents

Returns a paginated collection of integration events.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-integration-events" method="get" path="/v1/integration-events" -->
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

    res, err := s.IntegrationEvents.ListIntegrationEvents(ctx, operations.ListIntegrationEventsRequest{
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Filter: &components.IntegrationEventsFilterParameters{
            Ts: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTFilter(
                components.DateTimeFieldGTFilter{
                    Gt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListIntegrationEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListIntegrationEventsRequest](../../models/operations/listintegrationeventsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListIntegrationEventsResponse](../../models/operations/listintegrationeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateIntegrationEvents

Bulk-creates integration events for an integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-integration-events" method="post" path="/v1/integration-instances/{integrationInstance}/integration-events" -->
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

    res, err := s.IntegrationEvents.CreateIntegrationEvents(ctx, "gateway-manager", components.CreateIntegrationEventsBody{
        Data: []components.CreateIntegrationEvent{
            components.CreateIntegrationEvent{
                Type: "incident.triggered",
                Ts: types.MustTimeFromString("2024-01-01T00:00:00Z"),
                DeduplicationID: "1d34191ebfe305a56bf4f3303d1a8dcf",
                Actor: &components.IntegrationEventActor{
                    Type: sdkkonnectgo.Pointer("user"),
                    ID: "PLH1HKV",
                    Name: sdkkonnectgo.Pointer("John Doe"),
                },
                Attributes: map[string]any{
                    "service_id": "PLK33TY",
                    "incident_id": "PGR0VU2",
                    "escalation_policy_id": "P005NFS",
                    "priority_id": "P9BO35Z",
                    "urgency": "low",
                    "title": "Example API Service Incident",
                    "html_url": "https://acme.pagerduty.com/incidents/PGR0VU2",
                },
                Data: map[string]any{
                    "pagerduty_service_id": "PLK33TY",
                },
                Binding: "pagerduty_service",
                Detail: sdkkonnectgo.Pointer("Example API Service Incident"),
                LinkHTML: sdkkonnectgo.Pointer("https://acme.pagerduty.com/incidents/PGR0VU2"),
            },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `integrationInstance`                                                                            | `string`                                                                                         | :heavy_check_mark:                                                                               | Machine name of the integration.                                                                 | gateway-manager                                                                                  |
| `createIntegrationEventsBody`                                                                    | [components.CreateIntegrationEventsBody](../../models/components/createintegrationeventsbody.md) | :heavy_check_mark:                                                                               | Request body schema for bulk-creating integration events.                                        |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateIntegrationEventsResponse](../../models/operations/createintegrationeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCatalogServiceIntegrationEvents

Returns a paginated collection of integration events belonging to the given catalog service.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-catalog-service-integration-events" method="get" path="/v1/catalog-services/{serviceId}/integration-events" -->
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

    res, err := s.IntegrationEvents.ListCatalogServiceIntegrationEvents(ctx, operations.ListCatalogServiceIntegrationEventsRequest{
        ServiceID: "7f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Filter: &components.CatalogServiceIntegrationEventsFilterParameters{
            Ts: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListIntegrationEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListCatalogServiceIntegrationEventsRequest](../../models/operations/listcatalogserviceintegrationeventsrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListCatalogServiceIntegrationEventsResponse](../../models/operations/listcatalogserviceintegrationeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |