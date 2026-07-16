# OpenMeterCustomers

## Overview

Customers are used to track usage of your product or service. Customers can be individuals or organizations that can subscribe to plans and have access to features.

### Available Operations

* [CreateCustomer](#createcustomer) - Create customer
* [ListCustomers](#listcustomers) - List customers
* [GetCustomer](#getcustomer) - Get customer
* [UpsertCustomer](#upsertcustomer) - Upsert customer
* [DeleteCustomer](#deletecustomer) - Delete customer
* [GetCustomerBilling](#getcustomerbilling) - Get customer billing data
* [UpdateCustomerBilling](#updatecustomerbilling) - Update customer billing data
* [UpdateCustomerBillingAppData](#updatecustomerbillingappdata) - Update customer billing app data
* [CreateCustomerStripeCheckoutSession](#createcustomerstripecheckoutsession) - Create Stripe Checkout Session
* [CreateCustomerStripePortalSession](#createcustomerstripeportalsession) - Create Stripe customer portal session
* [ListCustomerCharges](#listcustomercharges) - List customer charges
* [CreateCustomerCharges](#createcustomercharges) - Create customer charge
* [CreateCreditAdjustment](#createcreditadjustment) - Create a credit adjustment
* [GetCustomerCreditBalance](#getcustomercreditbalance) - Get a customer's credit balance
* [CreateCreditGrant](#createcreditgrant) - Create a new credit grant
* [ListCreditGrants](#listcreditgrants) - List credit grants
* [GetCreditGrant](#getcreditgrant) - Get a credit grant
* [UpdateCreditGrantExternalSettlement](#updatecreditgrantexternalsettlement) - Update credit grant external settlement status
* [ListCreditTransactions](#listcredittransactions) - List credit transactions

## CreateCustomer

Create customer

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer" method="post" path="/v3/openmeter/customers" -->
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

    res, err := s.OpenMeterCustomers.CreateCustomer(ctx, components.CreateCustomerRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Key: "019ae40f-4258-7f15-9491-842f42a7d6ac",
        Currency: sdkkonnectgo.Pointer("USD"),
        BillingAddress: &components.BillingAddress{
            Country: sdkkonnectgo.Pointer("US"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreateCustomerRequest](../../models/components/createcustomerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateCustomerResponse](../../models/operations/createcustomerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCustomers

List customers

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customers" method="get" path="/v3/openmeter/customers" -->
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

    res, err := s.OpenMeterCustomers.ListCustomers(ctx, operations.ListCustomersRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListCustomersParamsFilter{
            BillingProfileID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListCustomersRequest](../../models/operations/listcustomersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListCustomersResponse](../../models/operations/listcustomersresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomer

Get customer

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer" method="get" path="/v3/openmeter/customers/{customerId}" -->
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

    res, err := s.OpenMeterCustomers.GetCustomer(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCustomerResponse](../../models/operations/getcustomerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertCustomer

Upsert customer

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-customer" method="put" path="/v3/openmeter/customers/{customerId}" -->
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

    res, err := s.OpenMeterCustomers.UpsertCustomer(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpsertCustomerRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Currency: sdkkonnectgo.Pointer("USD"),
        BillingAddress: &components.UpsertCustomerRequestBillingAddress{
            Country: sdkkonnectgo.Pointer("US"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `customerID`                                                                         | `string`                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  | 01G65Z755AFWAKHE12NY0CQ9FH                                                           |
| `upsertCustomerRequest`                                                              | [components.UpsertCustomerRequest](../../models/components/upsertcustomerrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.UpsertCustomerResponse](../../models/operations/upsertcustomerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCustomer

Delete customer

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-customer" method="delete" path="/v3/openmeter/customers/{customerId}" -->
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

    res, err := s.OpenMeterCustomers.DeleteCustomer(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCustomerResponse](../../models/operations/deletecustomerresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomerBilling

Get customer billing data

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer-billing" method="get" path="/v3/openmeter/customers/{customerId}/billing" -->
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

    res, err := s.OpenMeterCustomers.GetCustomerBilling(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCustomerData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCustomerBillingResponse](../../models/operations/getcustomerbillingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCustomerBilling

Update customer billing data

### Example Usage

<!-- UsageSnippet language="go" operationID="update-customer-billing" method="put" path="/v3/openmeter/customers/{customerId}/billing" -->
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

    res, err := s.OpenMeterCustomers.UpdateCustomerBilling(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpsertCustomerBillingDataRequest{
        BillingProfile: &components.UpsertCustomerBillingDataRequestBillingProfile{
            ID: "01G65Z755AFWAKHE12NY0CQ9FH",
        },
        AppData: &components.UpsertCustomerBillingDataRequestAppCustomerData{
            Stripe: &components.UpsertCustomerBillingDataRequestStripe{
                CustomerID: sdkkonnectgo.Pointer("cus_1234567890"),
                DefaultPaymentMethodID: sdkkonnectgo.Pointer("pm_1234567890"),
                Labels: map[string]string{
                    "env": "test",
                },
            },
            ExternalInvoicing: &components.UpsertCustomerBillingDataRequestExternalInvoicing{
                Labels: map[string]string{
                    "env": "test",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCustomerData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `customerID`                                                                                               | `string`                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                 |
| `upsertCustomerBillingDataRequest`                                                                         | [components.UpsertCustomerBillingDataRequest](../../models/components/upsertcustomerbillingdatarequest.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.UpdateCustomerBillingResponse](../../models/operations/updatecustomerbillingresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCustomerBillingAppData

Update customer billing app data

### Example Usage

<!-- UsageSnippet language="go" operationID="update-customer-billing-app-data" method="put" path="/v3/openmeter/customers/{customerId}/billing/app-data" -->
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

    res, err := s.OpenMeterCustomers.UpdateCustomerBillingAppData(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpsertAppCustomerDataRequest{
        Stripe: &components.Stripe{
            CustomerID: sdkkonnectgo.Pointer("cus_1234567890"),
            DefaultPaymentMethodID: sdkkonnectgo.Pointer("pm_1234567890"),
            Labels: map[string]string{
                "env": "test",
            },
        },
        ExternalInvoicing: &components.ExternalInvoicing{
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingAppCustomerData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `customerID`                                                                                       | `string`                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                | 01G65Z755AFWAKHE12NY0CQ9FH                                                                         |
| `upsertAppCustomerDataRequest`                                                                     | [components.UpsertAppCustomerDataRequest](../../models/components/upsertappcustomerdatarequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.UpdateCustomerBillingAppDataResponse](../../models/operations/updatecustomerbillingappdataresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCustomerStripeCheckoutSession

Create a [Stripe Checkout Session](https://docs.stripe.com/payments/checkout)
for the customer.

Creates a Checkout Session for collecting payment method information from
customers. The session operates in "setup" mode, which collects payment details
without charging the customer immediately. The collected payment method can be
used for future subscription billing.

For hosted checkout sessions, redirect customers to the returned URL. For
embedded sessions, use the client_secret to initialize Stripe.js in your
application.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer-stripe-checkout-session" method="post" path="/v3/openmeter/customers/{customerId}/billing/stripe/checkout-sessions" -->
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

    res, err := s.OpenMeterCustomers.CreateCustomerStripeCheckoutSession(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.BillingCustomerStripeCreateCheckoutSessionRequest{
        StripeOptions: components.StripeOptions{
            Currency: sdkkonnectgo.Pointer("USD"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingAppStripeCreateCheckoutSessionResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  | Example                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |                                                                                                                                              |
| `customerID`                                                                                                                                 | `string`                                                                                                                                     | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                                                   |
| `billingCustomerStripeCreateCheckoutSessionRequest`                                                                                          | [components.BillingCustomerStripeCreateCheckoutSessionRequest](../../models/components/billingcustomerstripecreatecheckoutsessionrequest.md) | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |                                                                                                                                              |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |                                                                                                                                              |

### Response

**[*operations.CreateCustomerStripeCheckoutSessionResponse](../../models/operations/createcustomerstripecheckoutsessionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCustomerStripePortalSession

Create Stripe Customer Portal Session.

Useful to redirect the customer to the Stripe Customer Portal to manage their
payment methods, change their billing address and access their invoice history.
Only returns URL if the customer billing profile is linked to a stripe app and
customer.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer-stripe-portal-session" method="post" path="/v3/openmeter/customers/{customerId}/billing/stripe/portal-sessions" -->
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

    res, err := s.OpenMeterCustomers.CreateCustomerStripePortalSession(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.BillingCustomerStripeCreateCustomerPortalSessionRequest{
        StripeOptions: components.BillingCustomerStripeCreateCustomerPortalSessionRequestStripeOptions{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingAppStripeCreateCustomerPortalSessionResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              | Example                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |                                                                                                                                                          |
| `customerID`                                                                                                                                             | `string`                                                                                                                                                 | :heavy_check_mark:                                                                                                                                       | N/A                                                                                                                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                                                               |
| `billingCustomerStripeCreateCustomerPortalSessionRequest`                                                                                                | [components.BillingCustomerStripeCreateCustomerPortalSessionRequest](../../models/components/billingcustomerstripecreatecustomerportalsessionrequest.md) | :heavy_check_mark:                                                                                                                                       | N/A                                                                                                                                                      |                                                                                                                                                          |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |                                                                                                                                                          |

### Response

**[*operations.CreateCustomerStripePortalSessionResponse](../../models/operations/createcustomerstripeportalsessionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCustomerCharges

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

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

    res, err := s.OpenMeterCustomers.ListCustomerCharges(ctx, operations.ListCustomerChargesRequest{
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

## CreateCustomerCharges

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create customer charge.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer-charges" method="post" path="/v3/openmeter/customers/{customerId}/charges" -->
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

    res, err := s.OpenMeterCustomers.CreateCustomerCharges(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateCreateChargeRequestUsageBased(
        components.CreateChargeUsageBasedRequest{
            Name: "<value>",
            Type: components.CreateChargeUsageBasedRequestTypeUsageBased,
            Currency: "USD",
            InvoiceAt: types.MustTimeFromString("2023-01-01T01:01:01.001Z"),
            ServicePeriod: components.CreateChargeUsageBasedRequestServicePeriod{
                From: types.MustTimeFromString("2023-01-01T01:01:01.001Z"),
                To: types.MustTimeFromString("2023-01-01T01:01:01.001Z"),
            },
            SettlementMode: components.CreateChargeUsageBasedRequestSettlementModeCreditOnly,
            FeatureKey: "<value>",
            Price: components.CreateCreateChargeUsageBasedRequestPriceFlat(
                components.BillingPriceFlat{
                    Type: components.BillingPriceFlatTypeFlat,
                    Amount: "478.68",
                },
            ),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCharge != nil {
        switch res.BillingCharge.Type {
            case components.BillingChargeTypeFlatFee:
                // res.BillingCharge.BillingChargeFlatFee is populated
            case components.BillingChargeTypeUsageBased:
                // res.BillingCharge.BillingChargeUsageBased is populated
        }

    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `customerID`                                                                     | `string`                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | 01G65Z755AFWAKHE12NY0CQ9FH                                                       |
| `createChargeRequest`                                                            | [components.CreateChargeRequest](../../models/components/createchargerequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateCustomerChargesResponse](../../models/operations/createcustomerchargesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCreditAdjustment

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

A credit adjustment can be used to make manual adjustments to a customer's
credit balance.

Supported use-cases:

- Usage correction

### Example Usage

<!-- UsageSnippet language="go" operationID="create-credit-adjustment" method="post" path="/v3/openmeter/customers/{customerId}/credits/adjustments" -->
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

    res, err := s.OpenMeterCustomers.CreateCreditAdjustment(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateCreditAdjustmentRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Currency: "USD",
        Amount: "259.30",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCreditAdjustment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `customerID`                                                                                         | `string`                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  | 01G65Z755AFWAKHE12NY0CQ9FH                                                                           |
| `createCreditAdjustmentRequest`                                                                      | [components.CreateCreditAdjustmentRequest](../../models/components/createcreditadjustmentrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CreateCreditAdjustmentResponse](../../models/operations/createcreditadjustmentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomerCreditBalance

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get a credit balance.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer-credit-balance" method="get" path="/v3/openmeter/customers/{customerId}/credits/balance" -->
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

    res, err := s.OpenMeterCustomers.GetCustomerCreditBalance(ctx, operations.GetCustomerCreditBalanceRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Timestamp: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCreditBalances != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCustomerCreditBalanceRequest](../../models/operations/getcustomercreditbalancerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetCustomerCreditBalanceResponse](../../models/operations/getcustomercreditbalanceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCreditGrant

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a new credit grant. A credit grant represents an allocation of prepaid
credits to a customer.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-credit-grant" method="post" path="/v3/openmeter/customers/{customerId}/credits/grants" -->
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

    res, err := s.OpenMeterCustomers.CreateCreditGrant(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateCreditGrantRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        FundingMethod: components.CreditFundingMethodInvoice,
        Currency: "USD",
        Amount: "517.19",
        Purchase: &components.Purchase{
            Currency: "USD",
        },
        TaxConfig: &components.TaxConfigurationForACreditGrant{
            TaxCode: &components.TaxCode{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
        },
        Filters: &components.CreateCreditGrantFilters{
            Features: []string{
                "input_tokens",
                "output_tokens",
            },
        },
        EffectiveAt: types.MustNewTimeFromString("2023-01-01T01:01:01.001Z"),
        ExpiresAfter: sdkkonnectgo.Pointer("P1Y"),
        Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCreditGrant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `customerID`                                                                               | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        | 01G65Z755AFWAKHE12NY0CQ9FH                                                                 |
| `createCreditGrantRequest`                                                                 | [components.CreateCreditGrantRequest](../../models/components/createcreditgrantrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.CreateCreditGrantResponse](../../models/operations/createcreditgrantresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCreditGrants

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List credit grants.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-credit-grants" method="get" path="/v3/openmeter/customers/{customerId}/credits/grants" -->
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

    res, err := s.OpenMeterCustomers.ListCreditGrants(ctx, operations.ListCreditGrantsRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Filter: &components.ListCreditGrantsParamsFilter{
            Currency: sdkkonnectgo.Pointer("USD"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditGrantPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListCreditGrantsRequest](../../models/operations/listcreditgrantsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListCreditGrantsResponse](../../models/operations/listcreditgrantsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCreditGrant

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get a credit grant.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-credit-grant" method="get" path="/v3/openmeter/customers/{customerId}/credits/grants/{creditGrantId}" -->
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

    res, err := s.OpenMeterCustomers.GetCreditGrant(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCreditGrant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customerID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `creditGrantID`                                          | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCreditGrantResponse](../../models/operations/getcreditgrantresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCreditGrantExternalSettlement

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update the payment settlement status of an externally funded credit grant.

Use this endpoint to synchronize the payment state of an external payment with
the system so that revenue recognition and credit availability work as expected.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-credit-grant-external-settlement" method="post" path="/v3/openmeter/customers/{customerId}/credits/grants/{creditGrantId}/settlement/external" -->
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

    res, err := s.OpenMeterCustomers.UpdateCreditGrantExternalSettlement(ctx, operations.UpdateCreditGrantExternalSettlementRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        CreditGrantID: "01G65Z755AFWAKHE12NY0CQ9FH",
        UpdateCreditGrantExternalSettlementRequest: components.UpdateCreditGrantExternalSettlementRequest{
            Status: components.CreditPurchasePaymentSettlementStatusSettled,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingCreditGrant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.UpdateCreditGrantExternalSettlementRequest](../../models/operations/updatecreditgrantexternalsettlementrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.UpdateCreditGrantExternalSettlementResponse](../../models/operations/updatecreditgrantexternalsettlementresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCreditTransactions

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List credit transactions for a customer.

Returns an immutable, chronological record of credit movements: funded credits
and consumed credits. Transactions are returned in reverse chronological order
by default.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-credit-transactions" method="get" path="/v3/openmeter/customers/{customerId}/credits/transactions" -->
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

    res, err := s.OpenMeterCustomers.ListCreditTransactions(ctx, operations.ListCreditTransactionsRequest{
        CustomerID: "01G65Z755AFWAKHE12NY0CQ9FH",
        Filter: &components.ListCreditTransactionsParamsFilter{
            Currency: sdkkonnectgo.Pointer("USD"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditTransactionPaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListCreditTransactionsRequest](../../models/operations/listcredittransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListCreditTransactionsResponse](../../models/operations/listcredittransactionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |