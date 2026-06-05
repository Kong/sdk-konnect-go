# OpenMeterBilling

## Overview

Billing manages the billing profiles, currencies, cost bases, and invoices for customers.

### Available Operations

* [ListCurrencies](#listcurrencies) - List currencies
* [CreateCustomCurrency](#createcustomcurrency) - Create custom currency
* [ListCostBases](#listcostbases) - List cost bases
* [CreateCostBasis](#createcostbasis) - Create cost basis
* [ListCustomerCharges](#listcustomercharges) - List customer charges
* [ListBillingProfiles](#listbillingprofiles) - List billing profiles
* [CreateBillingProfile](#createbillingprofile) - Create a new billing profile
* [GetBillingProfile](#getbillingprofile) - Get a billing profile
* [UpdateBillingProfile](#updatebillingprofile) - Update a billing profile
* [DeleteBillingProfile](#deletebillingprofile) - Delete a billing profile

## ListCurrencies

List currencies supported by the billing system.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-currencies" method="get" path="/v3/openmeter/currencies" -->
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

    res, err := s.OpenMeterBilling.ListCurrencies(ctx, operations.ListCurrenciesRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CurrencyPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListCurrenciesRequest](../../models/operations/listcurrenciesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListCurrenciesResponse](../../models/operations/listcurrenciesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCustomCurrency

Create a custom currency. This operation allows defining your own custom
currency for billing purposes.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-custom-currency" method="post" path="/v3/openmeter/currencies/custom" -->
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

    res, err := s.OpenMeterBilling.CreateCustomCurrency(ctx, components.CreateCurrencyCustomRequest{
        Name: "<value>",
        Code: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCurrencyCustom != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateCurrencyCustomRequest](../../models/components/createcurrencycustomrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateCustomCurrencyResponse](../../models/operations/createcustomcurrencyresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCostBases

List cost bases for a currency. For custom currencies, there can be multiple
cost bases with different `effective_from` dates.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-bases" method="get" path="/v3/openmeter/currencies/custom/{currencyId}/cost-bases" -->
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

    res, err := s.OpenMeterBilling.ListCostBases(ctx, operations.ListCostBasesRequest{
        CurrencyID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Filter: &components.ListCostBasesParamsFilter{
            FiatCode: sdkkonnectgo.Pointer("USD"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CostBasisPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListCostBasesRequest](../../models/operations/listcostbasesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListCostBasesResponse](../../models/operations/listcostbasesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCostBasis

Create a cost basis for a currency.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-cost-basis" method="post" path="/v3/openmeter/currencies/custom/{currencyId}/cost-bases" -->
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

    res, err := s.OpenMeterBilling.CreateCostBasis(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateCostBasisRequest{
        FiatCode: "USD",
        Rate: "<value>",
        EffectiveFrom: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCostBasis != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `currencyID`                                                                           | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | 01G65Z755AFWAKHE12NY0CQ9FH                                                             |
| `createCostBasisRequest`                                                               | [components.CreateCostBasisRequest](../../models/components/createcostbasisrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.CreateCostBasisResponse](../../models/operations/createcostbasisresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCustomerCharges

List customer charges.

Returns the customer's charges that are represented as either flat fee or
usage-based charges.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customer-charges" method="get" path="/v3/openmeter/customers/{customerId}/charges" -->
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

    res, err := s.OpenMeterBilling.ListCustomerCharges(ctx, operations.ListCustomerChargesRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ChargePagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCustomerChargesRequest](../../models/operations/listcustomerchargesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListCustomerChargesResponse](../../models/operations/listcustomerchargesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListBillingProfiles

List billing profiles.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-billing-profiles" method="get" path="/v3/openmeter/profiles" -->
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

    res, err := s.OpenMeterBilling.ListBillingProfiles(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingProfilePagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `page`                                                                            | [*components.PagePaginationQuery](../../models/components/pagepaginationquery.md) | :heavy_minus_sign:                                                                | Determines which page of the collection to retrieve.                              |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.ListBillingProfilesResponse](../../models/operations/listbillingprofilesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateBillingProfile

Create a new billing profile.

Billing profiles contain the settings for billing and controls invoice
generation. An organization can have multiple billing profiles defined. A
billing profile is linked to a specific app. This association is established
during the billing profile's creation and remains immutable.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-billing-profile" method="post" path="/v3/openmeter/profiles" -->
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

    res, err := s.OpenMeterBilling.CreateBillingProfile(ctx, components.CreateBillingProfileRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Supplier: components.CreateBillingProfileRequestSupplier{
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
            Addresses: &components.CreateBillingProfileRequestAddresses{
                BillingAddress: components.CreateBillingProfileRequestBillingAddress{
                    Country: sdkkonnectgo.Pointer("US"),
                },
            },
        },
        Workflow: components.CreateBillingProfileRequestWorkflow{
            Collection: &components.CreateBillingProfileRequestWorkflowCollectionSettings{
                Alignment: sdkkonnectgo.Pointer(components.CreateCreateBillingProfileRequestAlignmentSubscription(
                    components.AlignmentBillingWorkflowCollectionAlignmentSubscription{
                        Type: components.BillingWorkflowCollectionAlignmentSubscriptionAlignmentTypeSubscription,
                    },
                )),
                Interval: sdkkonnectgo.Pointer("P1D"),
            },
            Invoicing: &components.CreateBillingProfileRequestWorkflowInvoiceSettings{
                DraftPeriod: sdkkonnectgo.Pointer("P1D"),
            },
            Payment: sdkkonnectgo.Pointer(components.CreateCreateBillingProfileRequestPaymentChargeAutomatically(
                components.BillingWorkflowPaymentChargeAutomaticallySettings{
                    CollectionMethod: components.CollectionMethodChargeAutomatically,
                },
            )),
            Tax: &components.CreateBillingProfileRequestWorkflowTaxSettings{
                DefaultTaxConfig: &components.CreateBillingProfileRequestDefaultTaxConfig{
                    TaxCode: &components.CreateBillingProfileRequestTaxCode{
                        ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                    },
                },
            },
        },
        Apps: components.CreateBillingProfileRequestApps{
            Tax: components.CreateBillingProfileRequestTax{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
            Invoicing: components.CreateBillingProfileRequestInvoicing{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
            Payment: components.CreateBillingProfileRequestAppsPayment{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
        },
        Default: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingProfile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateBillingProfileRequest](../../models/components/createbillingprofilerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateBillingProfileResponse](../../models/operations/createbillingprofileresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetBillingProfile

Get a billing profile.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-billing-profile" method="get" path="/v3/openmeter/profiles/{id}" -->
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

    res, err := s.OpenMeterBilling.GetBillingProfile(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingProfile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetBillingProfileResponse](../../models/operations/getbillingprofileresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateBillingProfile

Update a billing profile.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-billing-profile" method="put" path="/v3/openmeter/profiles/{id}" -->
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

    res, err := s.OpenMeterBilling.UpdateBillingProfile(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpsertBillingProfileRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Supplier: components.UpsertBillingProfileRequestSupplier{
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
            Addresses: &components.UpsertBillingProfileRequestAddresses{
                BillingAddress: components.UpsertBillingProfileRequestBillingAddress{
                    Country: sdkkonnectgo.Pointer("US"),
                },
            },
        },
        Workflow: components.UpsertBillingProfileRequestWorkflow{
            Collection: &components.UpsertBillingProfileRequestWorkflowCollectionSettings{
                Alignment: sdkkonnectgo.Pointer(components.CreateUpsertBillingProfileRequestAlignmentSubscription(
                    components.UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentSubscription{
                        Type: components.BillingWorkflowCollectionAlignmentSubscriptionAlignmentUpsertBillingProfileRequestTypeSubscription,
                    },
                )),
                Interval: sdkkonnectgo.Pointer("P1D"),
            },
            Invoicing: &components.UpsertBillingProfileRequestWorkflowInvoiceSettings{
                DraftPeriod: sdkkonnectgo.Pointer("P1D"),
            },
            Payment: sdkkonnectgo.Pointer(components.CreateUpsertBillingProfileRequestPaymentChargeAutomatically(
                components.BillingWorkflowPaymentChargeAutomaticallySettings{
                    CollectionMethod: components.CollectionMethodChargeAutomatically,
                },
            )),
            Tax: &components.UpsertBillingProfileRequestWorkflowTaxSettings{
                DefaultTaxConfig: &components.UpsertBillingProfileRequestDefaultTaxConfig{
                    TaxCode: &components.UpsertBillingProfileRequestTaxCode{
                        ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                    },
                },
            },
        },
        Default: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingProfile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `id`                                                                                             | `string`                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              | 01G65Z755AFWAKHE12NY0CQ9FH                                                                       |
| `upsertBillingProfileRequest`                                                                    | [components.UpsertBillingProfileRequest](../../models/components/upsertbillingprofilerequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.UpdateBillingProfileResponse](../../models/operations/updatebillingprofileresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteBillingProfile

Delete a billing profile.

Only such billing profiles can be deleted that are:

- not the default profile
- not pinned to any customer using customer overrides
- only have finalized invoices

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-billing-profile" method="delete" path="/v3/openmeter/profiles/{id}" -->
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

    res, err := s.OpenMeterBilling.DeleteBillingProfile(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteBillingProfileResponse](../../models/operations/deletebillingprofileresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |