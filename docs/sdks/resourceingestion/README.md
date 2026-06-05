# ResourceIngestion

## Overview

Get details and trigger resource ingestion from authorized integrations.
Resource ingestion allows for manually syncing resources from authorized integrations.
It provides details on when the resources were last synced from an authorized integration.


### Available Operations

* [ScheduleResourceIngestion](#scheduleresourceingestion) - Schedule Resource Ingestion
* [FetchResourceIngestion](#fetchresourceingestion) - Get a Resource Ingestion

## ScheduleResourceIngestion

Schedules a resource ingestion job for the given integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="schedule-resource-ingestion" method="put" path="/v1/integration-instances/{integrationInstanceId}/resource-ingestion" -->
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

    res, err := s.ResourceIngestion.ScheduleResourceIngestion(ctx, "004d6d17-d21f-4518-89a2-b9177cb5485c", components.ScheduleResourceIngestion{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResourceIngestionStatus != nil {
        switch res.ResourceIngestionStatus.SchedulerState.Type {
            case components.SchedulerStateTypeResourceSchedulerStateOk:
                // res.ResourceIngestionStatus.SchedulerState.ResourceSchedulerStateOk is populated
            case components.SchedulerStateTypeResourceSchedulerStateNotOk:
                // res.ResourceIngestionStatus.SchedulerState.ResourceSchedulerStateNotOk is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `integrationInstanceID`                                                                      | `string`                                                                                     | :heavy_check_mark:                                                                           | The `id` of the integration instance.                                                        | 004d6d17-d21f-4518-89a2-b9177cb5485c                                                         |
| `scheduleResourceIngestion`                                                                  | [components.ScheduleResourceIngestion](../../models/components/scheduleresourceingestion.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.ScheduleResourceIngestionResponse](../../models/operations/scheduleresourceingestionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchResourceIngestion

Fetches a resource ingestion job for the given integration instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-resource-ingestion" method="get" path="/v1/integration-instances/{integrationInstanceId}/resource-ingestion" -->
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

    res, err := s.ResourceIngestion.FetchResourceIngestion(ctx, "004d6d17-d21f-4518-89a2-b9177cb5485c")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResourceIngestionStatus != nil {
        switch res.ResourceIngestionStatus.SchedulerState.Type {
            case components.SchedulerStateTypeResourceSchedulerStateOk:
                // res.ResourceIngestionStatus.SchedulerState.ResourceSchedulerStateOk is populated
            case components.SchedulerStateTypeResourceSchedulerStateNotOk:
                // res.ResourceIngestionStatus.SchedulerState.ResourceSchedulerStateNotOk is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `integrationInstanceID`                                  | `string`                                                 | :heavy_check_mark:                                       | The `id` of the integration instance.                    | 004d6d17-d21f-4518-89a2-b9177cb5485c                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchResourceIngestionResponse](../../models/operations/fetchresourceingestionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |