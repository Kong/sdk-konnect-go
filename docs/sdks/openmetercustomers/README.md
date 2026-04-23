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
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
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

Create a [Stripe Checkout Session](https://docs.stripe.com/payments/checkout) for the customer.

Creates a Checkout Session for collecting payment method information from customers.
The session operates in "setup" mode, which collects payment details without charging
the customer immediately. The collected payment method can be used for future
subscription billing.

For hosted checkout sessions, redirect customers to the returned URL. For embedded
sessions, use the client_secret to initialize Stripe.js in your application.

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

Useful to redirect the customer to the Stripe Customer Portal to manage their payment methods,
change their billing address and access their invoice history.
Only returns URL if the customer billing profile is linked to a stripe app and customer.

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