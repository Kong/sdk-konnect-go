# PortalAuditLogs
(*PortalAuditLogs*)

## Overview

### Available Operations

* [UpdatePortalAuditLogReplayJob](#updateportalauditlogreplayjob) - Update Portal Audit Log Replay Job
* [GetPortalAuditLogReplayJob](#getportalauditlogreplayjob) - Get Portal Audit Log Replay Job
* [UpdatePortalAuditLogWebhook](#updateportalauditlogwebhook) - Update Portal Audit Log Webhook
* [GetPortalAuditLogWebhook](#getportalauditlogwebhook) - Get Portal Audit Log Webhook
* [GetPortalAuditLogWebhookStatus](#getportalauditlogwebhookstatus) - Get Portal Audit Log Webhook Status

## UpdatePortalAuditLogReplayJob

Updates a job to re-send audit logs to an portal's webhook.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-audit-log-replay-job" method="put" path="/v3/portals/{portalId}/audit-log-replay-job" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuditLogs.UpdatePortalAuditLogReplayJob(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuditLogReplayJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `portalID`                                                                                              | *string*                                                                                                | :heavy_check_mark:                                                                                      | ID of the portal.                                                                                       | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                    |
| `replacePortalAuditLogReplayJob`                                                                        | [*components.ReplacePortalAuditLogReplayJob](../../models/components/replaceportalauditlogreplayjob.md) | :heavy_minus_sign:                                                                                      | The request schema to replace a portal audit log replay job.                                            |                                                                                                         |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.UpdatePortalAuditLogReplayJobResponse](../../models/operations/updateportalauditlogreplayjobresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalAuditLogReplayJob

Returns the audit log replay job's configuration and status.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-audit-log-replay-job" method="get" path="/v3/portals/{portalId}/audit-log-replay-job" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuditLogs.GetPortalAuditLogReplayJob(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuditLogReplayJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalAuditLogReplayJobResponse](../../models/operations/getportalauditlogreplayjobresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalAuditLogWebhook

Updates the configuration for a webhook to receive audit logs.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-audit-log-webhook" method="patch" path="/v3/portals/{portalId}/audit-log-webhook" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuditLogs.UpdatePortalAuditLogWebhook(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.UpdatePortalAuditLogWebhook{
        Enabled: sdkkonnectgo.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuditLogWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `portalID`                                                                                        | *string*                                                                                          | :heavy_check_mark:                                                                                | ID of the portal.                                                                                 | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                              |
| `updatePortalAuditLogWebhook`                                                                     | [*components.UpdatePortalAuditLogWebhook](../../models/components/updateportalauditlogwebhook.md) | :heavy_minus_sign:                                                                                | The request schema to modify an portal audit log webhook.                                         |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.UpdatePortalAuditLogWebhookResponse](../../models/operations/updateportalauditlogwebhookresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalAuditLogWebhook

Returns configuration for the audit log webhook.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-audit-log-webhook" method="get" path="/v3/portals/{portalId}/audit-log-webhook" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuditLogs.GetPortalAuditLogWebhook(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuditLogWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalAuditLogWebhookResponse](../../models/operations/getportalauditlogwebhookresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalAuditLogWebhookStatus

Returns status of the audit log webhook.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-audit-log-webhook-status" method="get" path="/v3/portals/{portalId}/audit-log-webhook/status" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalAuditLogs.GetPortalAuditLogWebhookStatus(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalAuditLogWebhookStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalAuditLogWebhookStatusResponse](../../models/operations/getportalauditlogwebhookstatusresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |