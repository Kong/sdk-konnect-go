# OrgConfigurations

## Overview

Operations related to org scoped notifications and their org subscriptions. The http channel of an org scoped notification is set up for the whole organization, and no user can switch it on or off. Only an org admin can use these operations. The routes exist only when http destinations are enabled.

### Available Operations

* [ListOrgConfigurations](#listorgconfigurations) - List available org events
* [ListOrgSubscriptions](#listorgsubscriptions) - List org subscriptions
* [CreateOrgSubscription](#createorgsubscription) - Create an org subscription
* [GetOrgSubscription](#getorgsubscription) - Get an org subscription
* [UpdateOrgSubscription](#updateorgsubscription) - Update an org subscription
* [DeleteOrgSubscription](#deleteorgsubscription) - Delete an org subscription

## ListOrgConfigurations

Returns the organization event types that you can subscribe to, with each event's title, description, and namespace.
Only available to organization administrators.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-org-configurations" method="get" path="/v1/notifications/configurations/organization" -->
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

    res, err := s.OrgConfigurations.ListOrgConfigurations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgConfigurationListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrgConfigurationsResponse](../../models/operations/listorgconfigurationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListOrgSubscriptions

List existing organization subscriptions, including the subscription destinations.
Only available to organization administrators.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-org-subscriptions" method="get" path="/v1/notifications/configurations/organization/{eventId}/subscriptions" -->
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

    res, err := s.OrgConfigurations.ListOrgSubscriptions(ctx, operations.ListOrgSubscriptionsRequest{
        EventID: "invoice-ready",
        PageBefore: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscriptionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListOrgSubscriptionsRequest](../../models/operations/listorgsubscriptionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListOrgSubscriptionsResponse](../../models/operations/listorgsubscriptionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateOrgSubscription

Create a subscription for an org scoped event.
The event ID to subscribe is placed in the URL. A destination may be used in multiple subscriptions.
Only available to organization administrators.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-org-subscription" method="post" path="/v1/notifications/configurations/organization/{eventId}/subscriptions" -->
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

    res, err := s.OrgConfigurations.CreateOrgSubscription(ctx, "invoice-ready", components.OrgSubscriptionCreatePayload{
        DestinationIds: []string{},
        Regions: []components.NotificationRegion{
            components.NotificationRegionIn,
        },
        Entities: []string{
            "<value 1>",
            "<value 2>",
        },
        Enabled: false,
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `eventID`                                                                                          | `string`                                                                                           | :heavy_check_mark:                                                                                 | Formatted string ID of the notification event.                                                     | invoice-ready                                                                                      |
| `orgSubscriptionCreatePayload`                                                                     | [components.OrgSubscriptionCreatePayload](../../models/components/orgsubscriptioncreatepayload.md) | :heavy_check_mark:                                                                                 | Request body schema for creating an org subscription.                                              |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.CreateOrgSubscriptionResponse](../../models/operations/createorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetOrgSubscription

Fetch the details of a single org scoped subscription.
Only available to organization administrators.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-org-subscription" method="get" path="/v1/notifications/configurations/organization/{eventId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.OrgConfigurations.GetOrgSubscription(ctx, "invoice-ready", "2efea526-5094-430a-8577-7c0653ac52d0")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `eventID`                                                | `string`                                                 | :heavy_check_mark:                                       | Formatted string ID of the notification event.           | invoice-ready                                            |
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of the org subscription, a UUID v7.                   |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetOrgSubscriptionResponse](../../models/operations/getorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateOrgSubscription

Update the details of an org scoped subscription.
Only available to organization administrators.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-org-subscription" method="patch" path="/v1/notifications/configurations/organization/{eventId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.OrgConfigurations.UpdateOrgSubscription(ctx, operations.UpdateOrgSubscriptionRequest{
        EventID: "invoice-ready",
        SubscriptionID: "9c1c16fb-7821-4b7f-9ce9-8d6216959572",
        OrgSubscriptionUpdatePayload: components.OrgSubscriptionUpdatePayload{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateOrgSubscriptionRequest](../../models/operations/updateorgsubscriptionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateOrgSubscriptionResponse](../../models/operations/updateorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteOrgSubscription

Delete an org scoped subscription.
Only available to organization administrators.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-org-subscription" method="delete" path="/v1/notifications/configurations/organization/{eventId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.OrgConfigurations.DeleteOrgSubscription(ctx, "invoice-ready", "8e7f1c6a-cbea-4be6-a2a4-40f3e593e6d6")
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
| `eventID`                                                | `string`                                                 | :heavy_check_mark:                                       | Formatted string ID of the notification event.           | invoice-ready                                            |
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of the org subscription, a UUID v7.                   |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteOrgSubscriptionResponse](../../models/operations/deleteorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |