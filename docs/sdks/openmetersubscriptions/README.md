# OpenMeterSubscriptions

## Overview

Subscriptions are used to track usage of your product or service. Subscriptions can be individuals or organizations that can subscribe to plans and have access to features.

### Available Operations

* [CreateSubscription](#createsubscription) - Create subscription
* [ListSubscriptions](#listsubscriptions) - List subscriptions
* [GetSubscription](#getsubscription) - Get subscription
* [CreateSubscriptionAddon](#createsubscriptionaddon) - Create a new subscription add-on
* [ListSubscriptionAddons](#listsubscriptionaddons) - List subscription addons
* [GetSubscriptionAddon](#getsubscriptionaddon) - Get add-on association for subscription
* [UpdateSubscriptionAddon](#updatesubscriptionaddon) - Update subscription addon
* [CancelSubscription](#cancelsubscription) - Cancel subscription
* [ChangeSubscription](#changesubscription) - Change subscription
* [EditSubscription](#editsubscription) - Edit subscription
* [MigrateSubscription](#migratesubscription) - Migrate subscription
* [RestoreSubscription](#restoresubscription) - Restore subscription
* [UnscheduleSubscription](#unschedulesubscription) - Unschedule subscription
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
        Customer: components.BillingSubscriptionCreateCustomer{
            ID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
        },
        Plan: &components.Plan{
            ID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
            Key: sdkkonnectgo.Pointer("resource_key"),
        },
        CustomPlan: &components.CustomPlan{
            Name: "<value>",
            Labels: map[string]string{
                "env": "test",
            },
            Currency: "USD",
            BillingCadence: "P1Y",
            Phases: []components.BillingPlanPhase{
                components.BillingPlanPhase{
                    Name: "<value>",
                    Labels: map[string]string{
                        "env": "test",
                    },
                    Key: "resource_key",
                    Duration: sdkkonnectgo.Pointer("P1Y"),
                    RateCards: []components.BillingRateCard{
                        components.BillingRateCard{
                            Name: "<value>",
                            Labels: map[string]string{
                                "env": "test",
                            },
                            Key: "resource_key",
                            Feature: &components.FeatureReference{
                                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                            },
                            Currency: sdkkonnectgo.Pointer("USD"),
                            BillingCadence: sdkkonnectgo.Pointer("P1Y"),
                            Price: components.CreatePriceFlat(
                                components.BillingPriceFlat{
                                    Type: components.BillingPriceFlatTypeFlat,
                                    Amount: "112.57",
                                },
                            ),
                            TaxConfig: &components.TaxConfig{
                                Code: &components.TaxCode{
                                    ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                                },
                            },
                            Entitlement: sdkkonnectgo.Pointer(components.CreateEntitlementTemplateBoolean(
                                components.BillingRateCardBooleanEntitlement{
                                    Type: components.BillingRateCardBooleanEntitlementTypeBoolean,
                                },
                            )),
                        },
                    },
                },
            },
        },
        Timing: sdkkonnectgo.Pointer(components.CreateBillingSubscriptionCreateTimingBillingSubscriptionCreateTimingEnum(
            components.BillingSubscriptionCreateTimingEnumImmediate,
        )),
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

    res, err := s.OpenMeterSubscriptions.ListSubscriptions(ctx, operations.ListSubscriptionsRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListSubscriptionsParamsFilter{
            ID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            CustomerID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            PlanID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListSubscriptionsRequest](../../models/operations/listsubscriptionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

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

## CreateSubscriptionAddon

Add add-on to a subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-subscription-addon" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/addons" -->
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

    res, err := s.OpenMeterSubscriptions.CreateSubscriptionAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateSubscriptionAddonRequest{
        Labels: map[string]string{
            "env": "test",
        },
        Addon: components.CreateSubscriptionAddonRequestAddOn{
            ID: "01G65Z755AFWAKHE12NY0CQ9FH",
        },
        Quantity: 207252,
        Timing: components.CreateCreateSubscriptionAddonRequestTimingBillingSubscriptionEditTimingEnum(
            components.BillingSubscriptionEditTimingEnumImmediate,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionAddon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `subscriptionID`                                                                                       | `string`                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    | 01G65Z755AFWAKHE12NY0CQ9FH                                                                             |
| `createSubscriptionAddonRequest`                                                                       | [components.CreateSubscriptionAddonRequest](../../models/components/createsubscriptionaddonrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.CreateSubscriptionAddonResponse](../../models/operations/createsubscriptionaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListSubscriptionAddons

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List the add-ons of a subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-subscription-addons" method="get" path="/v3/openmeter/subscriptions/{subscriptionId}/addons" -->
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

    res, err := s.OpenMeterSubscriptions.ListSubscriptionAddons(ctx, operations.ListSubscriptionAddonsRequest{
        SubscriptionID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionAddonPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListSubscriptionAddonsRequest](../../models/operations/listsubscriptionaddonsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListSubscriptionAddonsResponse](../../models/operations/listsubscriptionaddonsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetSubscriptionAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get an add-on association for a subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-subscription-addon" method="get" path="/v3/openmeter/subscriptions/{subscriptionId}/addons/{subscriptionAddonId}" -->
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

    res, err := s.OpenMeterSubscriptions.GetSubscriptionAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionAddon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `subscriptionAddonID`                                    | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSubscriptionAddonResponse](../../models/operations/getsubscriptionaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateSubscriptionAddon

Update a subscription add-on. Only the quantity is mutable; the timing controls
when the new quantity takes effect. A new entry is appended to the add-on's
timeline.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-subscription-addon" method="patch" path="/v3/openmeter/subscriptions/{subscriptionId}/addons/{subscriptionAddonId}" -->
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

    res, err := s.OpenMeterSubscriptions.UpdateSubscriptionAddon(ctx, operations.UpdateSubscriptionAddonRequest{
        SubscriptionID: "01G65Z755AFWAKHE12NY0CQ9FH",
        SubscriptionAddonID: "01G65Z755AFWAKHE12NY0CQ9FH",
        BillingSubscriptionAddonUpdate: components.BillingSubscriptionAddonUpdate{
            Quantity: 757008,
            Timing: components.CreateBillingSubscriptionAddonUpdateTimingBillingSubscriptionEditTimingEnum(
                components.BillingSubscriptionEditTimingEnumImmediate,
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionAddon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateSubscriptionAddonRequest](../../models/operations/updatesubscriptionaddonrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateSubscriptionAddonResponse](../../models/operations/updatesubscriptionaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CancelSubscription

Cancels the subscription. Will result in a scheduling conflict if there are
other subscriptions scheduled to start after the cancelation time.

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

Closes a running subscription and starts a new one according to the
specification. Can be used for upgrades, downgrades, and plan changes.

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
        Plan: &components.BillingSubscriptionChangePlan{
            ID: sdkkonnectgo.Pointer("01G65Z755AFWAKHE12NY0CQ9FH"),
            Key: sdkkonnectgo.Pointer("resource_key"),
        },
        CustomPlan: &components.BillingSubscriptionChangeCustomPlan{
            Name: "<value>",
            Labels: map[string]string{
                "env": "test",
            },
            Currency: "USD",
            BillingCadence: "P1Y",
            Phases: []components.BillingPlanPhase{
                components.BillingPlanPhase{
                    Name: "<value>",
                    Labels: map[string]string{
                        "env": "test",
                    },
                    Key: "resource_key",
                    Duration: sdkkonnectgo.Pointer("P1Y"),
                    RateCards: []components.BillingRateCard{},
                },
            },
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

## EditSubscription

Edits a running subscription by applying an ordered batch of customizations
(adding or removing items, adding, removing, or stretching phases, or
unscheduling a pending edit). The changes may take effect immediately or at the
next billing cycle. Subscriptions that have add-ons cannot be edited.

### Example Usage

<!-- UsageSnippet language="go" operationID="edit-subscription" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/edit" -->
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

    res, err := s.OpenMeterSubscriptions.EditSubscription(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.BillingSubscriptionEdit{
        Customizations: []components.BillingSubscriptionEditOperation{
            components.CreateBillingSubscriptionEditOperationAddItem(
                components.BillingSubscriptionEditAddItem{
                    Type: components.BillingSubscriptionEditAddItemTypeAddItem,
                    PhaseKey: "<value>",
                    RateCard: components.BillingSubscriptionEditAddItemRateCard{
                        Name: "<value>",
                        Labels: map[string]string{
                            "env": "test",
                        },
                        Key: "resource_key",
                        Feature: &components.BillingSubscriptionEditAddItemFeatureReference{
                            ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                        },
                        Currency: sdkkonnectgo.Pointer("USD"),
                        BillingCadence: sdkkonnectgo.Pointer("P1Y"),
                        Price: components.CreateBillingSubscriptionEditAddItemPriceFree(
                            components.BillingPriceFree{
                                Type: components.BillingPriceFreeTypeFree,
                            },
                        ),
                        TaxConfig: &components.BillingSubscriptionEditAddItemTaxConfig{
                            Code: &components.BillingSubscriptionEditAddItemTaxCode{
                                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                            },
                        },
                        Entitlement: sdkkonnectgo.Pointer(components.CreateBillingSubscriptionEditAddItemEntitlementTemplateStatic(
                            components.BillingRateCardStaticEntitlement{
                                Type: components.BillingRateCardStaticEntitlementTypeStatic,
                                Config: "<value>",
                            },
                        )),
                    },
                },
            ),
        },
        Timing: sdkkonnectgo.Pointer(components.CreateBillingSubscriptionEditTimingBillingSubscriptionEditTimingEnum(
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `subscriptionID`                                                                         | `string`                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                                                               |
| `billingSubscriptionEdit`                                                                | [components.BillingSubscriptionEdit](../../models/components/billingsubscriptionedit.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.EditSubscriptionResponse](../../models/operations/editsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MigrateSubscription

Migrates to a later version of the current plan. With starting_phase omitted and
billing_anchor omitted or unchanged, migration amends the subscription in place:
unchanged items retain their service periods and both response entries have the
same ID. Existing addons must remain compatible with the target plan.
Incompatible phase timelines or billing settings return an error. Providing
starting_phase or a different billing_anchor explicitly requests replacement,
which resets the phase timeline, may produce billing adjustments, and does not
transfer addons. Custom subscriptions cannot be migrated.

### Example Usage

<!-- UsageSnippet language="go" operationID="migrate-subscription" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/migrate" -->
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

    res, err := s.OpenMeterSubscriptions.MigrateSubscription(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.BillingSubscriptionMigrate{
        Timing: sdkkonnectgo.Pointer(components.CreateBillingSubscriptionMigrateTimingBillingSubscriptionEditTimingEnum(
            components.BillingSubscriptionEditTimingEnumImmediate,
        )),
        BillingAnchor: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingSubscriptionMigrateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `subscriptionID`                                                                               | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            | 01G65Z755AFWAKHE12NY0CQ9FH                                                                     |
| `billingSubscriptionMigrate`                                                                   | [components.BillingSubscriptionMigrate](../../models/components/billingsubscriptionmigrate.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.MigrateSubscriptionResponse](../../models/operations/migratesubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RestoreSubscription

Restores the subscription by deleting any later-scheduled successor
subscriptions and continuing this one indefinitely. This is the inverse of a
future-dated change, which schedules a successor. Restore is not available when
multi-subscription is enabled.

### Example Usage

<!-- UsageSnippet language="go" operationID="restore-subscription" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/restore" -->
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

    res, err := s.OpenMeterSubscriptions.RestoreSubscription(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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

**[*operations.RestoreSubscriptionResponse](../../models/operations/restoresubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UnscheduleSubscription

Deletes a scheduled subscription that has not yet become active, removing it and
resolving any scheduling conflict it was holding. This is distinct from
canceling: cancel ends a running subscription, whereas unscheduling removes a
not-yet-active one. Only scheduled subscriptions can be unscheduled;
unscheduling an active or already-started subscription is rejected.

### Example Usage

<!-- UsageSnippet language="go" operationID="unschedule-subscription" method="post" path="/v3/openmeter/subscriptions/{subscriptionId}/unschedule" -->
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

    res, err := s.OpenMeterSubscriptions.UnscheduleSubscription(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.UnscheduleSubscriptionResponse](../../models/operations/unschedulesubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
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