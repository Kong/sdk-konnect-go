# Exports

## Overview

Exports copy the organization's cost and usage data into files to download. These operations manage one-time exports and export schedules, and download the files.

### Available Operations

* [ListCostExportSchedules](#listcostexportschedules) - List export schedules
* [CreateCostExportSchedule](#createcostexportschedule) - Create an export schedule
* [GetCostExportSchedule](#getcostexportschedule) - Get an export schedule
* [UpsertCostExportSchedule](#upsertcostexportschedule) - Replace an export schedule
* [DeleteCostExportSchedule](#deletecostexportschedule) - Delete an export schedule
* [ListCostExports](#listcostexports) - List exports
* [CreateCostExport](#createcostexport) - Create an export
* [GetCostExport](#getcostexport) - Get an export
* [DeleteCostExport](#deletecostexport) - Delete an export
* [CreateCostExportDownload](#createcostexportdownload) - Download an export file

## ListCostExportSchedules

Returns a page of the organization's export schedules, sorted by `name` unless the request sets `sort`. Filters on `name`, `status`, and `cadence` narrow the list.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-export-schedules" method="get" path="/v1/ai-cost-management/export-schedules" -->
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

    res, err := s.Exports.ListCostExportSchedules(ctx, operations.ListCostExportSchedulesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCostExportSchedulesRequest](../../models/operations/listcostexportschedulesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListCostExportSchedulesResponse](../../models/operations/listcostexportschedulesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostExportSchedule

Creates an export schedule and returns it. The server assigns the identifier.

An `active` schedule creates one export of each of its periods as the period ends, starting with the period in progress.

If the body is invalid, the request fails with `400`. If the organization already has 20 export schedules, the request fails with `409`.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-export-schedule" method="post" path="/v1/ai-cost-management/export-schedules" -->
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

    res, err := s.Exports.CreateCostExportSchedule(ctx, components.CreateAICostExportScheduleInputAICostExportScheduleInput1(
        components.AICostExportScheduleInput1{
            Format: components.AICostExportScheduleInputFormatCsv,
            Compression: components.AICostExportScheduleInputCompressionGzip,
            Name: "<value>",
            Cadence: components.AICostExportScheduleInputCadenceDaily,
            Dataset: components.AICostExportScheduleInputDatasetUsageEvents,
            Labels: map[string]string{
                "env": "test",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostExportSchedule != nil {
        switch res.AICostExportSchedule.Type {
            case components.AICostExportScheduleTypeAICostExportSchedule1:
                // res.AICostExportSchedule.AICostExportSchedule1 is populated
            case components.AICostExportScheduleTypeTwo:
                // res.AICostExportSchedule.Two is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.AICostExportScheduleInput](../../models/components/aicostexportscheduleinput.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateCostExportScheduleResponse](../../models/operations/createcostexportscheduleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCostExportSchedule

Returns one of the organization's export schedules. If the organization has no export schedule with this identifier, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-export-schedule" method="get" path="/v1/ai-cost-management/export-schedules/{scheduleId}" -->
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

    res, err := s.Exports.GetCostExportSchedule(ctx, "62a510af-81e8-4838-94bb-dd1ae73d1591")
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostExportSchedule != nil {
        switch res.AICostExportSchedule.Type {
            case components.AICostExportScheduleTypeAICostExportSchedule1:
                // res.AICostExportSchedule.AICostExportSchedule1 is populated
            case components.AICostExportScheduleTypeTwo:
                // res.AICostExportSchedule.Two is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `scheduleID`                                             | `string`                                                 | :heavy_check_mark:                                       | The `id` of the export schedule.                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostExportScheduleResponse](../../models/operations/getcostexportscheduleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertCostExportSchedule

Replaces an export schedule's `name`, `status`, `dataset`, `focus`, `format`, `compression`, and `filters`, and returns the result. The schedule keeps its identifier.

- The request carries the full schedule, so an omitted `filters` clears the filters, an omitted `status` resumes a paused schedule, and an omitted `managed_by` resets to `{}`.
- Setting `status` to `paused` pauses the schedule, and setting it to `active` resumes it.
- The replacement applies to the periods that end after it. Exports that the schedule already created don't change. If runs are late, the replacement first creates the missed exports with the earlier settings, within the schedule's limit of 31 periods.
- `cadence` and `time_zone` can't change, so the body repeats their stored values. If either differs, the request fails with `400` naming the field. Redefining a schedule's periods takes a delete and a create.
- An omitted `time_zone` counts as `UTC`, so a replacement of a schedule in another time zone that omits it fails with `400`.
- If the organization has no export schedule with this identifier, or the identifier isn't a valid UUID, the request fails with `404`. The request never creates a schedule.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-cost-export-schedule" method="put" path="/v1/ai-cost-management/export-schedules/{scheduleId}" -->
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

    res, err := s.Exports.UpsertCostExportSchedule(ctx, "5be7a065-61dc-4a7a-9daf-2d15b2125748", components.CreateAICostExportScheduleInputAICostExportScheduleInput1(
        components.AICostExportScheduleInput1{
            Format: components.AICostExportScheduleInputFormatCsv,
            Compression: components.AICostExportScheduleInputCompressionGzip,
            Name: "<value>",
            Cadence: components.AICostExportScheduleInputCadenceMonthly,
            Dataset: components.AICostExportScheduleInputDatasetUsageEvents,
            Labels: map[string]string{
                "env": "test",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostExportSchedule != nil {
        switch res.AICostExportSchedule.Type {
            case components.AICostExportScheduleTypeAICostExportSchedule1:
                // res.AICostExportSchedule.AICostExportSchedule1 is populated
            case components.AICostExportScheduleTypeTwo:
                // res.AICostExportSchedule.Two is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `scheduleID`                                                                                 | `string`                                                                                     | :heavy_check_mark:                                                                           | The `id` of the export schedule.                                                             |
| `aiCostExportScheduleInput`                                                                  | [components.AICostExportScheduleInput](../../models/components/aicostexportscheduleinput.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpsertCostExportScheduleResponse](../../models/operations/upsertcostexportscheduleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCostExportSchedule

Deletes an export schedule. The exports it created remain, with `schedule` set to `null`, and a `pending` one still runs.

If the organization has no export schedule with this identifier, or the identifier isn't a valid UUID, the request fails with `404`. This includes deleting the same schedule again.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-cost-export-schedule" method="delete" path="/v1/ai-cost-management/export-schedules/{scheduleId}" -->
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

    res, err := s.Exports.DeleteCostExportSchedule(ctx, "35245638-87c3-43fc-8095-92ab7bf78587")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `scheduleID`                                             | `string`                                                 | :heavy_check_mark:                                       | The `id` of the export schedule.                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCostExportScheduleResponse](../../models/operations/deletecostexportscheduleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCostExports

Returns a page of the organization's exports, newest first unless the request sets `sort`. Filters on `status`, `dataset`, `schedule.id`, and `created_at` narrow the list.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-exports" method="get" path="/v1/ai-cost-management/exports" -->
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

    res, err := s.Exports.ListCostExports(ctx, operations.ListCostExportsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListCostExportsRequest](../../models/operations/listcostexportsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListCostExportsResponse](../../models/operations/listcostexportsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostExport

Queues a one-time export and returns it with `status` set to `pending`. The server assigns the identifier.

The export runs asynchronously. A client reads it again with `GET /v1/ai-cost-management/exports/{exportId}` to follow its progress. Once it succeeds, the client requests a download URL for one file at a time with `POST /v1/ai-cost-management/exports/{exportId}/downloads`.

If the body is invalid, the request fails with `400`, including when `ends_at` is later than the end of the day in progress in `time_zone`. That check on `ends_at` runs last: a body with other errors gets those errors first, and the corrected body then gets the error on `ends_at`.

If the organization already has 20 `pending` one-time exports, the request fails with `409`. Exports that a schedule creates don't count toward that limit.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-export" method="post" path="/v1/ai-cost-management/exports" -->
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

    res, err := s.Exports.CreateCostExport(ctx, components.CreateAICostExportInputAICostExportInput1(
        components.AICostExportInput1{
            Format: components.AICostExportInputFormatCsv,
            Compression: components.AICostExportInputCompressionGzip,
            Dataset: components.AICostExportInputDatasetUsageEvents,
            StartsAt: types.MustTimeFromString("2025-08-11T22:20:42.721Z"),
            EndsAt: types.MustTimeFromString("2024-08-25T04:17:22.848Z"),
            Labels: map[string]string{
                "env": "test",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostExport != nil {
        switch res.AICostExport.Type {
            case components.AICostExportTypeAICostExport1:
                // res.AICostExport.AICostExport1 is populated
            case components.AICostExportTypeAICostExport2:
                // res.AICostExport.AICostExport2 is populated
        }

    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.AICostExportInput](../../models/components/aicostexportinput.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateCostExportResponse](../../models/operations/createcostexportresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCostExport

Returns one of the organization's exports. If the organization has no export with this identifier, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-export" method="get" path="/v1/ai-cost-management/exports/{exportId}" -->
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

    res, err := s.Exports.GetCostExport(ctx, "c72d63de-9b87-4b05-af70-067f5f7c18e1")
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostExport != nil {
        switch res.AICostExport.Type {
            case components.AICostExportTypeAICostExport1:
                // res.AICostExport.AICostExport1 is populated
            case components.AICostExportTypeAICostExport2:
                // res.AICostExport.AICostExport2 is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `exportID`                                               | `string`                                                 | :heavy_check_mark:                                       | The `id` of the export.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostExportResponse](../../models/operations/getcostexportresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCostExport

Deletes an export and its files before it responds, so the files can't be downloaded afterwards. Deleting a `pending` or `running` export cancels it.

If the files can't be removed, the request fails with `500` and the export remains, so a retry deletes both. If the organization has no export with this identifier, or the identifier isn't a valid UUID, the request fails with `404`. This includes deleting the same export again.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-cost-export" method="delete" path="/v1/ai-cost-management/exports/{exportId}" -->
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

    res, err := s.Exports.DeleteCostExport(ctx, "aaa27a02-f18c-473f-a6fa-063ad2a0b0a3")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `exportID`                                               | `string`                                                 | :heavy_check_mark:                                       | The `id` of the export.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCostExportResponse](../../models/operations/deletecostexportresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostExportDownload

Returns a URL that serves one file of a `succeeded` export, either a data file or `manifest.json`, to a plain `GET` for 5 minutes. The service doesn't store the URL, and no operation returns it again.

- If `file` names no file of the export, the request fails with `400`.
- If the organization has no export with this identifier, or the identifier isn't a valid UUID, the request fails with `404`.
- If the export hasn't succeeded, or it has expired, the request fails with `409`.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-export-download" method="post" path="/v1/ai-cost-management/exports/{exportId}/downloads" -->
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

    res, err := s.Exports.CreateCostExportDownload(ctx, "aee926f2-30b5-4f9f-b1af-1a4691581e4e", operations.CreateCostExportDownloadRequestBody{
        File: "<value>",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `exportID`                                                                                                       | `string`                                                                                                         | :heavy_check_mark:                                                                                               | The `id` of the export.                                                                                          |
| `requestBody`                                                                                                    | [operations.CreateCostExportDownloadRequestBody](../../models/operations/createcostexportdownloadrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateCostExportDownloadResponse](../../models/operations/createcostexportdownloadresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |