# NotificationChannels

## Overview

Notification channels are the email recipient lists and Slack incoming webhooks that budget alerts go to. These operations manage the channels and send test messages.

### Available Operations

* [ListCostNotificationChannels](#listcostnotificationchannels) - List notification channels
* [CreateCostNotificationChannel](#createcostnotificationchannel) - Create a notification channel
* [GetCostNotificationChannel](#getcostnotificationchannel) - Get a notification channel
* [UpsertCostNotificationChannel](#upsertcostnotificationchannel) - Replace a notification channel
* [DeleteCostNotificationChannel](#deletecostnotificationchannel) - Delete a notification channel
* [CreateCostNotificationChannelTest](#createcostnotificationchanneltest) - Send a test notification

## ListCostNotificationChannels

Returns a page of the organization's notification channels, sorted by `name` unless the request sets `sort`. Filters on `name`, `type`, `labels.<key>`, and `managed_by.<key>` narrow the list.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-notification-channels" method="get" path="/v1/ai-cost-management/notification-channels" -->
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

    res, err := s.NotificationChannels.ListCostNotificationChannels(ctx, operations.ListCostNotificationChannelsRequest{})
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
| `request`                                                                                                        | [operations.ListCostNotificationChannelsRequest](../../models/operations/listcostnotificationchannelsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListCostNotificationChannelsResponse](../../models/operations/listcostnotificationchannelsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostNotificationChannel

Creates a notification channel in the organization and returns it. The server assigns the identifier.

`type` selects the kind of channel: an email channel takes its recipients in `email`, and a Slack channel takes an incoming webhook URL in `slack`. The response shows the webhook URL with its secret token replaced by `…`. If the body is invalid, such as a webhook URL that isn't a Slack incoming webhook URL, the request fails with `400`.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-notification-channel" method="post" path="/v1/ai-cost-management/notification-channels" -->
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

    res, err := s.NotificationChannels.CreateCostNotificationChannel(ctx, operations.CreateCreateCostNotificationChannelRequestBodySlack(
        components.AICostSlackNotificationChannelInput{
            Name: "<value>",
            Slack: components.AICostSlackNotificationChannelInputSlack{
                WebhookURL: "https://monumental-lift.com/",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostNotificationChannel != nil {
        switch res.AICostNotificationChannel.Type {
            case components.AICostNotificationChannelTypeEmail:
                // res.AICostNotificationChannel.AICostEmailNotificationChannel is populated
            case components.AICostNotificationChannelTypeSlack:
                // res.AICostNotificationChannel.AICostSlackNotificationChannel is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.CreateCostNotificationChannelRequestBody](../../models/operations/createcostnotificationchannelrequestbody.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.CreateCostNotificationChannelResponse](../../models/operations/createcostnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCostNotificationChannel

Returns one of the organization's notification channels. If the organization has no notification channel with this identifier, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-notification-channel" method="get" path="/v1/ai-cost-management/notification-channels/{channelId}" -->
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

    res, err := s.NotificationChannels.GetCostNotificationChannel(ctx, "22997f0f-2642-41e8-b62e-b74bf190bdaa")
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostNotificationChannel != nil {
        switch res.AICostNotificationChannel.Type {
            case components.AICostNotificationChannelTypeEmail:
                // res.AICostNotificationChannel.AICostEmailNotificationChannel is populated
            case components.AICostNotificationChannelTypeSlack:
                // res.AICostNotificationChannel.AICostSlackNotificationChannel is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `channelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the notification channel.                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostNotificationChannelResponse](../../models/operations/getcostnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertCostNotificationChannel

Replaces a notification channel's `name`, `labels`, and settings, and returns the result. The channel keeps its identifier and type.

- The request carries the full channel, so an omitted `labels` or `managed_by` clears it.
- A Slack channel keeps its stored webhook URL when the request omits `slack.webhook_url`. This is the only exception to full replacement, because reads never return the full URL.
- If `type` differs from the stored type, the request fails with `400`, and `invalid_parameters` names `type`.
- If the organization has no notification channel with this identifier, or the identifier isn't a valid UUID, the request fails with `404`. The request never creates a channel.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-cost-notification-channel" method="put" path="/v1/ai-cost-management/notification-channels/{channelId}" -->
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

    res, err := s.NotificationChannels.UpsertCostNotificationChannel(ctx, "b23d4c19-1d1a-40ae-8e37-3b3ddc6f1b16", operations.CreateUpsertCostNotificationChannelRequestBodySlack(
        components.AICostSlackNotificationChannelReplaceInput{
            Name: "<value>",
            Slack: components.AICostSlackNotificationChannelReplaceInputSlack{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostNotificationChannel != nil {
        switch res.AICostNotificationChannel.Type {
            case components.AICostNotificationChannelTypeEmail:
                // res.AICostNotificationChannel.AICostEmailNotificationChannel is populated
            case components.AICostNotificationChannelTypeSlack:
                // res.AICostNotificationChannel.AICostSlackNotificationChannel is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `channelID`                                                                                                                | `string`                                                                                                                   | :heavy_check_mark:                                                                                                         | The `id` of the notification channel.                                                                                      |
| `requestBody`                                                                                                              | [operations.UpsertCostNotificationChannelRequestBody](../../models/operations/upsertcostnotificationchannelrequestbody.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpsertCostNotificationChannelResponse](../../models/operations/upsertcostnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCostNotificationChannel

Deletes a notification channel, removes it from every notification rule that lists it, and deletes its deliveries and channel tests, including any still pending.

If the organization has no notification channel with this identifier, or the identifier isn't a valid UUID, the request fails with `404`. This includes deleting the same channel again.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-cost-notification-channel" method="delete" path="/v1/ai-cost-management/notification-channels/{channelId}" -->
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

    res, err := s.NotificationChannels.DeleteCostNotificationChannel(ctx, "27e9d2cd-7d44-4561-b5dd-dceb7f73f48a")
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
| `channelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the notification channel.                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCostNotificationChannelResponse](../../models/operations/deletecostnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostNotificationChannelTest

Queues a test message for the channel and returns the notification event that records it, with `kind` set to `test`. The test is the message an alert sends, with test wording and example figures, so what arrives shows what an alert will look like.

The service sends the test asynchronously through the channel's real delivery path, within seconds. A client reads the event again with `GET /v1/ai-cost-management/notification-events/{eventId}` to see the outcome. A test isn't retried, so if its first attempt fails, the event's `error` says why.

While the channel already has a pending test, the request returns that test instead of queueing another. Sending a test needs permission to edit the channel. If the organization has no notification channel with this identifier, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-notification-channel-test" method="post" path="/v1/ai-cost-management/notification-channels/{channelId}/tests" -->
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

    res, err := s.NotificationChannels.CreateCostNotificationChannelTest(ctx, "a522fa06-ef38-4a71-8da9-c224ade858e5")
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostNotificationTestEvent != nil {
        switch res.AICostNotificationTestEvent.Type {
            case components.AICostNotificationTestEventTypePending:
                // res.AICostNotificationTestEvent.AICostNotificationTestEventPending is populated
            case components.AICostNotificationTestEventTypeDelivered:
                // res.AICostNotificationTestEvent.AICostNotificationTestEventDelivered is populated
            case components.AICostNotificationTestEventTypeFailed:
                // res.AICostNotificationTestEvent.AICostNotificationTestEventFailed is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `channelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The `id` of the notification channel.                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateCostNotificationChannelTestResponse](../../models/operations/createcostnotificationchanneltestresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |