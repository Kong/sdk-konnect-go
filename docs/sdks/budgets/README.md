# Budgets

## Overview

Budgets set spending limits, or exemptions from them, for parts of the organization's usage. These operations manage budgets and their notification rules, and read balances, firings, and notification events.

### Available Operations

* [ListCostBudgetBalances](#listcostbudgetbalances) - List budget balances
* [ListCostBudgets](#listcostbudgets) - List budgets
* [CreateCostBudget](#createcostbudget) - Create a budget
* [GetCostBudget](#getcostbudget) - Get a budget
* [UpsertCostBudget](#upsertcostbudget) - Replace a budget's updatable fields
* [DeleteCostBudget](#deletecostbudget) - Delete a budget
* [GetCostBudgetBalance](#getcostbudgetbalance) - Get a budget's balance
* [GetCostNotificationRule](#getcostnotificationrule) - Get a budget's notification rule
* [UpsertCostNotificationRule](#upsertcostnotificationrule) - Create or replace a notification rule
* [DeleteCostNotificationRule](#deletecostnotificationrule) - Delete a budget's notification rule
* [ListCostNotificationRuleFirings](#listcostnotificationrulefirings) - List a notification rule's firings
* [ListCostNotificationEvents](#listcostnotificationevents) - List notification events
* [GetCostNotificationEvent](#getcostnotificationevent) - Get a notification event

## ListCostBudgetBalances

Returns a page of the balances of the organization's limit budgets, highest consumed share first unless the request sets `sort`. Exemptions have no balance.

- A limit budget without `each` has one balance for its whole scope, with `value` set to `null`.
- A limit budget with `each` has one balance for every value of the `each` field with usage in the budget window.
- A value that an `eq` or `in` predicate on the `each` field selects has a balance even without usage.

`filter[budget.id]` narrows the list to some budgets, and `filter[scope.subject]` and `filter[scope.group]` to the balances of one entity. `budget` is a reference unless the request sends `expand=budget`, which renders the full budget. For a caller whose access covers only particular principals or groups, the list holds the balances of the budgets that the budget list shows that caller. For a budget that the caller sees only through their own derived budget, it holds that one balance, even without usage.

`evaluated_at` fixes every budget window for every page, unless a budget's window changes, but not the usage: usage that arrives while a client reads the pages can change the amounts. A zero limit has no consumed share, and its balances sort as the highest share.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-budget-balances" method="get" path="/v1/ai-cost-management/budget-balances" -->
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

    res, err := s.Budgets.ListCostBudgetBalances(ctx, operations.ListCostBudgetBalancesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListCostBudgetBalancesRequest](../../models/operations/listcostbudgetbalancesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListCostBudgetBalancesResponse](../../models/operations/listcostbudgetbalancesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCostBudgets

Returns a page of the organization's budgets, newest first unless the request sets `sort`. Filters on `name`, `effect`, `priority`, `posture`, `labels.<key>`, and `managed_by.<key>` narrow the list.

For a caller whose access covers only particular principals or groups, the list holds only the budgets whose scope lies within those principals and groups. It also holds each budget with `each` set to `subject` or `subject_id` whose scope covers one of the caller's principals.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-budgets" method="get" path="/v1/ai-cost-management/budgets" -->
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

    res, err := s.Budgets.ListCostBudgets(ctx, operations.ListCostBudgetsRequest{})
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
| `request`                                                                              | [operations.ListCostBudgetsRequest](../../models/operations/listcostbudgetsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListCostBudgetsResponse](../../models/operations/listcostbudgetsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostBudget

Creates a budget and returns it. The server assigns the identifier and records the caller as `created_by`.

`effect` selects the kind of budget: a limit budget takes `limit`, `window`, and `posture`, and an exemption takes `validity`. A limit budget can split its limit with `scope.each`.

- If the body is invalid, the request fails with `400`.
- For a caller whose access covers only particular principals or groups, a `scope` that doesn't lie within those principals and groups fails with `400`, and `invalid_parameters` names `scope`.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-budget" method="post" path="/v1/ai-cost-management/budgets" -->
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

    res, err := s.Budgets.CreateCostBudget(ctx, operations.CreateCreateCostBudgetRequestBodyLimit(
        components.AICostLimitBudgetInput{
            Name: "<value>",
            Scope: components.AICostLimitBudgetInputScope{
                Filters: "<value>",
            },
            Limit: components.Limit{
                Value: "<value>",
                Currency: components.AICostCurrencyUsd,
            },
            Window: components.AICostLimitBudgetInputWindow{
                Interval: "<value>",
                AnchoredAt: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostBudget != nil {
        switch res.AICostBudget.Type {
            case components.AICostBudgetTypeLimit:
                // res.AICostBudget.AICostLimitBudget is populated
            case components.AICostBudgetTypeExempt:
                // res.AICostBudget.AICostExemptBudget is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateCostBudgetRequestBody](../../models/operations/createcostbudgetrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateCostBudgetResponse](../../models/operations/createcostbudgetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCostBudget

Returns one of the organization's budgets. If the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-budget" method="get" path="/v1/ai-cost-management/budgets/{budgetId}" -->
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

    res, err := s.Budgets.GetCostBudget(ctx, "bcf6e5b6-0072-497e-b048-c7ef21426f0c")
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostBudget != nil {
        switch res.AICostBudget.Type {
            case components.AICostBudgetTypeLimit:
                // res.AICostBudget.AICostLimitBudget is populated
            case components.AICostBudgetTypeExempt:
                // res.AICostBudget.AICostExemptBudget is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `budgetID`                                               | `string`                                                 | :heavy_check_mark:                                       | The `id` of the budget.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostBudgetResponse](../../models/operations/getcostbudgetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertCostBudget

Replaces a budget's `name`, `labels`, `priority`, and `scope`, with `limit`, `window`, and `posture` on a limit budget or `validity` on an exemption, and returns the result. The request carries the full budget. The budget keeps its identifier, notification rule, and recorded firings.

- An omitted `labels`, `managed_by`, `priority`, or `posture` resets to its default.
- `effect` can't change. If the request changes it, it fails with `400` naming `effect`.
- A changed `limit`, `scope`, or `window` applies to the whole current budget window.
- Removing `scope.each` while the notification rule's message uses `{{entity.field}}` or `{{entity.value}}` fails with `400` naming `scope.each`. Change the message first.
- Every replacement evaluates the budget's notification rule against the stored budget. A threshold that spend in the current budget window has reached fires at once, unless it already fired for that window and limit. A changed limit, or a changed window that moves the current window's start, therefore lets the thresholds fire again. A changed scope doesn't, but a changed `each` can fire for derived values new to the window. The evaluation runs after the budget is stored, so if it fails, retrying the same request completes it.
- A caller whose access covers only particular principals or groups replaces a budget that lies within its groups alone, or one within its principals and groups that it created. Replacing any other budget it can see fails with `403`. The replacement's `scope` must keep the budget replaceable by the caller, or the request fails with `400` naming `scope`.
- The server assigns identifiers, so if the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404` and doesn't create one.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-cost-budget" method="put" path="/v1/ai-cost-management/budgets/{budgetId}" -->
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

    res, err := s.Budgets.UpsertCostBudget(ctx, "0109b403-2dad-489a-86ce-cf99101a53c5", operations.CreateUpsertCostBudgetRequestBodyLimit(
        components.AICostLimitBudgetInput{
            Name: "<value>",
            Scope: components.AICostLimitBudgetInputScope{
                Filters: "<value>",
            },
            Limit: components.Limit{
                Value: "<value>",
                Currency: components.AICostCurrencyUsd,
            },
            Window: components.AICostLimitBudgetInputWindow{
                Interval: "<value>",
                AnchoredAt: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostBudget != nil {
        switch res.AICostBudget.Type {
            case components.AICostBudgetTypeLimit:
                // res.AICostBudget.AICostLimitBudget is populated
            case components.AICostBudgetTypeExempt:
                // res.AICostBudget.AICostExemptBudget is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `budgetID`                                                                                       | `string`                                                                                         | :heavy_check_mark:                                                                               | The `id` of the budget.                                                                          |
| `requestBody`                                                                                    | [operations.UpsertCostBudgetRequestBody](../../models/operations/upsertcostbudgetrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpsertCostBudgetResponse](../../models/operations/upsertcostbudgetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCostBudget

Deletes a budget, with its notification rule and the rule's firings and deliveries. Deleting the same budget again fails with `404`.

A caller whose access covers only particular principals or groups deletes a budget that lies within its groups alone, or one within its principals and groups that it created. Deleting any other budget it can see fails with `403`. If the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-cost-budget" method="delete" path="/v1/ai-cost-management/budgets/{budgetId}" -->
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

    res, err := s.Budgets.DeleteCostBudget(ctx, "1102ce41-b0b3-47e8-80b0-9d42763834cf")
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
| `budgetID`                                               | `string`                                                 | :heavy_check_mark:                                       | The `id` of the budget.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCostBudgetResponse](../../models/operations/deletecostbudgetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCostBudgetBalance

Returns a limit budget's spend, remaining amount, and projected spend in the budget window that contains `evaluated_at`. `derived` adds the balances of the derived budgets with the highest spend.

A caller who sees only their own derived budget reads that one. Its amounts count that caller's usage alone, and `derived` holds the caller's value, or is empty without usage in the budget window.

The amounts count the usage stored when the request runs. `latest_record_at` shows when the newest counted usage reached the service. Only a limit budget has a balance, so if the budget is an exemption, the request fails with `400`. If the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-budget-balance" method="get" path="/v1/ai-cost-management/budgets/{budgetId}/balance" -->
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

    res, err := s.Budgets.GetCostBudgetBalance(ctx, operations.GetCostBudgetBalanceRequest{
        BudgetID: "439bea69-b34c-4ad4-8ab8-2d7f48646fa4",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetCostBudgetBalanceRequest](../../models/operations/getcostbudgetbalancerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetCostBudgetBalanceResponse](../../models/operations/getcostbudgetbalanceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCostNotificationRule

Returns a limit budget's notification rule: the thresholds at which it fires, the channels that receive the alerts, and the wording of its messages. A budget has at most one rule.

If the budget has no rule, the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404`. An exemption has no rule, so if the budget is an exemption, the request fails with `400`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-notification-rule" method="get" path="/v1/ai-cost-management/budgets/{budgetId}/notification-rule" -->
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

    res, err := s.Budgets.GetCostNotificationRule(ctx, "96073acc-35ba-4b16-b806-5d6bbefdb11d")
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostNotificationRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `budgetID`                                               | `string`                                                 | :heavy_check_mark:                                       | The `id` of the budget.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostNotificationRuleResponse](../../models/operations/getcostnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertCostNotificationRule

Creates a limit budget's notification rule, or replaces its `thresholds`, `channels`, and `message`, and returns the rule. The response status is `201` when the request creates the rule and `200` when it replaces one.

- The service evaluates the rule at once, so a threshold that spend in the current budget window has already reached fires right away. A threshold already recorded for this budget window, derived budget, and limit doesn't fire again. The evaluation runs after the rule is stored, so if it fails, retrying the same request completes it.
- Omitting `message` or sending `null` restores the built-in wording. An omitted `managed_by` resets to `{}`.
- Removing a threshold or channel keeps the firings and deliveries already recorded, and pending deliveries are still sent.
- If a listed channel doesn't exist in the organization, the request fails with `400`, and `invalid_parameters` names `channels[i].id` with the rule `missing_reference`.
- If `message` has an unknown placeholder or a `{{` without a closing `}}`, the request fails with `400`. It fails the same way when `message` uses `{{entity.field}}` or `{{entity.value}}` on a budget without `each`.
- If the budget is an exemption, the request fails with `400`. If the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-cost-notification-rule" method="put" path="/v1/ai-cost-management/budgets/{budgetId}/notification-rule" -->
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

    res, err := s.Budgets.UpsertCostNotificationRule(ctx, "b6423463-cf85-4ad1-a987-ef1ddbb33771", operations.UpsertCostNotificationRuleRequestBody{
        Thresholds: []components.AICostNotificationThreshold{
            components.AICostNotificationThreshold{
                Percent: 690517,
            },
        },
        ManagedBy: map[string]string{
            "owner": "terraform",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostNotificationRule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `budgetID`                                                                                                           | `string`                                                                                                             | :heavy_check_mark:                                                                                                   | The `id` of the budget.                                                                                              |
| `requestBody`                                                                                                        | [operations.UpsertCostNotificationRuleRequestBody](../../models/operations/upsertcostnotificationrulerequestbody.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.UpsertCostNotificationRuleResponse](../../models/operations/upsertcostnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCostNotificationRule

Deletes a budget's notification rule, with its firings and their deliveries. The budget stays. Deleting the rule again fails with `404`.

If the budget has no rule, the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404`. If the budget is an exemption, the request fails with `400`.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-cost-notification-rule" method="delete" path="/v1/ai-cost-management/budgets/{budgetId}/notification-rule" -->
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

    res, err := s.Budgets.DeleteCostNotificationRule(ctx, "8210798d-ff4a-4ae5-baa4-bd19c4b35460")
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
| `budgetID`                                               | `string`                                                 | :heavy_check_mark:                                       | The `id` of the budget.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCostNotificationRuleResponse](../../models/operations/deletecostnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCostNotificationRuleFirings

Returns a page of the firings of a budget's notification rule, newest first unless the request sets `sort`. A firing is a threshold crossing that the rule recorded. It has the fields of a threshold event, except that `budget` holds only the budget's `id` and there's no `kind`. Firings are kept for as long as the rule exists.

A caller who sees only their own derived budget reads only that derived budget's firings, and none on a budget with `each` set to `subject_id`.

If the budget has no rule, the organization has no budget with this identifier that the caller can see, or the identifier isn't a valid UUID, the request fails with `404`. If the budget is an exemption, the request fails with `400`.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-notification-rule-firings" method="get" path="/v1/ai-cost-management/budgets/{budgetId}/notification-rule/firings" -->
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

    res, err := s.Budgets.ListCostNotificationRuleFirings(ctx, operations.ListCostNotificationRuleFiringsRequest{
        BudgetID: "1df9f9a4-c0f5-4585-a450-290de66b4ca1",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListCostNotificationRuleFiringsRequest](../../models/operations/listcostnotificationrulefiringsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListCostNotificationRuleFiringsResponse](../../models/operations/listcostnotificationrulefiringsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCostNotificationEvents

Returns a page of the organization's threshold events, newest first unless the request sets `sort`. Each event is a recorded threshold crossing of a budget's notification rule. `budget` is a reference unless the request sends `expand=budget`, which renders the full budget, so a list can show each budget's name and scope without reading the budget. Channel tests are notification events too, but this list never includes them.

Events with the same `fired_at` are ordered by `budget.id`, then by `threshold.percent`, then by `value`. Filters on `deliveries.status` and `deliveries.channel.id` select events with at least one matching delivery. When both are present, they must match the same delivery.

For a caller whose access covers only particular principals or groups, the list holds only the threshold events of the budgets that caller can read. A caller who sees only their own derived budget reads only that derived budget's events, and none on a budget with `each` set to `subject_id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-notification-events" method="get" path="/v1/ai-cost-management/notification-events" -->
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

    res, err := s.Budgets.ListCostNotificationEvents(ctx, operations.ListCostNotificationEventsRequest{
        Filter: &components.AICostNotificationEventFilterParameters{
            FiredAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCostNotificationEventsRequest](../../models/operations/listcostnotificationeventsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListCostNotificationEventsResponse](../../models/operations/listcostnotificationeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCostNotificationEvent

Returns one notification event of either kind: a threshold event, with `kind` set to `threshold`, or a channel test, with `kind` set to `test`. A threshold event's `budget` is a reference unless the request sends `expand=budget`, which renders the full budget. Reading a test event shows the outcome of a test sent with `POST /v1/ai-cost-management/notification-channels/{channelId}/tests`.

A caller whose access covers only particular principals or groups reads only the threshold events that the event list shows them. For such a caller, any other event, including every channel test, fails with `404`. If the organization has no notification event with this identifier, or the identifier isn't a valid UUID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-notification-event" method="get" path="/v1/ai-cost-management/notification-events/{eventId}" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
	"github.com/Kong/sdk-konnect-go/models/operations"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Budgets.GetCostNotificationEvent(ctx, "d9231635-1342-4b3f-9d8d-3954c292bc2e", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        switch res.OneOf.Type {
            case operations.GetCostNotificationEventResponseBodyTypeThreshold:
                // res.OneOf.AICostNotificationThresholdEvent is populated
            case operations.GetCostNotificationEventResponseBodyTypeTest:
                // res.OneOf.AICostNotificationTestEvent is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `eventID`                                                                                                                                                                | `string`                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                       | The `id` of the notification event.                                                                                                                                      |
| `expand`                                                                                                                                                                 | `*string`                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                       | The references to render as full resources, as a comma-separated list of paths. Accepts `budget`. Without it, a reference renders as an object with the resource's `id`. |
| `opts`                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.GetCostNotificationEventResponse](../../models/operations/getcostnotificationeventresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |