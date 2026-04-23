# OpenMeterSubscriptions

## Overview

Subscriptions are used to track usage of your product or service. Subscriptions can be individuals or organizations that can subscribe to plans and have access to features.

### Available Operations

* [CreateSubscription](#createsubscription) - Create subscription
* [ListSubscriptions](#listsubscriptions) - List subscriptions
* [GetSubscription](#getsubscription) - Get subscription
* [CancelSubscription](#cancelsubscription) - Cancel subscription
* [ChangeSubscription](#changesubscription) - Change subscription
* [UnscheduleCancelation](#unschedulecancelation) - Unschedule subscription cancelation

## CreateSubscription

Create subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="create-subscription" method="post" path="/v3/openmeter/subscriptions" -->
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

    res, err := s.OpenMeterSubscriptions.CreateSubscription(ctx, components.BillingSubscriptionCreate{
        Labels: map[string]string{
            "env": "test",
        },
        Customer: components.Customer{
            ID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
        },
        Plan: components.Plan{
            ID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
            Key: sdkkonnectgo.Pointer("resource_key"),
        },
        BillingAnchor: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.BillingSubscriptionCreate](../../models/components/billingsubscriptioncreate.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateSubscriptionResponse](../../models/operations/createsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListSubscriptions

List subscriptions

### Example Usage

<!-- UsageSnippet language="go" operationID="list-subscriptions" method="get" path="/v3/openmeter/subscriptions" -->
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

    res, err := s.OpenMeterSubscriptions.ListSubscriptions(ctx, nil, &operations.ListSubscriptionsQueryParamFilter{
        CustomerID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |
| `page`                                                                                                        | [*components.PagePaginationQuery](../../models/components/pagepaginationquery.md)                             | :heavy_minus_sign:                                                                                            | Determines which page of the collection to retrieve.                                                          |
| `filter`                                                                                                      | [*operations.ListSubscriptionsQueryParamFilter](../../models/operations/listsubscriptionsqueryparamfilter.md) | :heavy_minus_sign:                                                                                            | Filter subscriptions.                                                                                         |
| `opts`                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                      | :heavy_minus_sign:                                                                                            | The options for this request.                                                                                 |

### Response

**[*operations.ListSubscriptionsResponse](../../models/operations/listsubscriptionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetSubscription

Get subscription

### Example Usage

<!-- UsageSnippet language="go" operationID="get-subscription" method="get" path="/v3/openmeter/subscriptions/{subscriptionId}" -->
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

    res, err := s.OpenMeterSubscriptions.GetSubscription(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSubscriptionResponse](../../models/operations/getsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CancelSubscription

Cancels the subscription.
Will result in a scheduling conflict if there are other subscriptions scheduled to start after the cancelation time.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancel-subscription" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/cancel" -->
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

    res, err := s.OpenMeterSubscriptions.CancelSubscription(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.BillingSubscriptionCancel{
        Timing: sdkkonnectgo.Pointer(components.CreateTimingBillingSubscriptionEditTimingEnum(
            components.BillingSubscriptionEditTimingEnumImmediate,
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `subscriptionID`                                                                             | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          | 01G65Z755AFWAKHE12NY0CQ9FH                                                                   |
| `billingSubscriptionCancel`                                                                  | [components.BillingSubscriptionCancel](../../models/components/billingsubscriptioncancel.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.CancelSubscriptionResponse](../../models/operations/cancelsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ChangeSubscription

Closes a running subscription and starts a new one according to the specification.
Can be used for upgrades, downgrades, and plan changes.

### Example Usage

<!-- UsageSnippet language="go" operationID="change-subscription" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/change" -->
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

    res, err := s.OpenMeterSubscriptions.ChangeSubscription(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.BillingSubscriptionChange{
        Labels: map[string]string{
            "env": "test",
        },
        Customer: components.BillingSubscriptionChangeCustomer{
            ID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
        },
        Plan: components.BillingSubscriptionChangePlan{
            ID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
            Key: sdkkonnectgo.Pointer("resource_key"),
        },
        BillingAnchor: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
        Timing: components.CreateBillingSubscriptionChangeTimingBillingSubscriptionEditTimingEnum(
            components.BillingSubscriptionEditTimingEnumImmediate,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingSubscriptionChangeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `subscriptionID`                                                                             | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          | 01G65Z755AFWAKHE12NY0CQ9FH                                                                   |
| `billingSubscriptionChange`                                                                  | [components.BillingSubscriptionChange](../../models/components/billingsubscriptionchange.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.ChangeSubscriptionResponse](../../models/operations/changesubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UnscheduleCancelation

Unschedules the subscription cancelation.

### Example Usage

<!-- UsageSnippet language="go" operationID="unschedule-cancelation" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/unschedule-cancelation" -->
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

    res, err := s.OpenMeterSubscriptions.UnscheduleCancelation(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.UnscheduleCancelationResponse](../../models/operations/unschedulecancelationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |