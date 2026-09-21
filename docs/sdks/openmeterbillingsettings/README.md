# OpenMeterBillingSettings

## Overview

Billing settings manages the billing profiles and invoices for customers.

### Available Operations

* [ListInvoices](#listinvoices) - List billing invoices
* [GetInvoice](#getinvoice) - Get a billing invoice
* [UpdateInvoice](#updateinvoice) - Update a billing invoice
* [DeleteInvoice](#deleteinvoice) - Delete a billing invoice
* [AdvanceInvoice](#advanceinvoice) - Advance billing invoice's next status
* [ApproveInvoice](#approveinvoice) - Send the invoice to the customer
* [RetryInvoice](#retryinvoice) - Retry advancing the invoice after a failed attempt
* [SnapshotQuantitiesInvoice](#snapshotquantitiesinvoice) - Snapshot quantities for usage based line items
* [ListBillingProfiles](#listbillingprofiles) - List billing profiles
* [CreateBillingProfile](#createbillingprofile) - Create a new billing profile
* [GetBillingProfile](#getbillingprofile) - Get a billing profile
* [UpdateBillingProfile](#updatebillingprofile) - Update a billing profile
* [DeleteBillingProfile](#deletebillingprofile) - Delete a billing profile

## ListInvoices

List billing invoices.

Returns a page of invoices. Gathering invoices are never included. Use `filter`
to narrow by status, customer, dates, or service period start. Use `sort` to
control ordering.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-invoices" method="get" path="/v3/openmeter/billing/invoices" -->
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

    res, err := s.OpenMeterBillingSettings.ListInvoices(ctx, operations.ListInvoicesRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListInvoicesParamsFilter{
            CustomerID: sdkkonnectgo.Pointer(components.CreateListInvoicesParamsFilterULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            IssuedAt: sdkkonnectgo.Pointer(components.CreateListInvoicesParamsFilterDateTimeFieldFilterListInvoicesParamsFilterDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.ListInvoicesParamsFilterDateTimeFieldFilterDateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            ServicePeriodStart: sdkkonnectgo.Pointer(components.CreateListInvoicesParamsFilterServicePeriodStartDateTimeFieldFilterListInvoicesParamsFilterDateTimeFieldFilterServicePeriodStartDateTimeFieldLTFilter(
                components.ListInvoicesParamsFilterDateTimeFieldFilterServicePeriodStartDateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            CreatedAt: sdkkonnectgo.Pointer(components.CreateListInvoicesParamsFilterCreatedAtDateTimeFieldFilterListInvoicesParamsFilterDateTimeFieldFilterCreatedAtDateTimeFieldGTFilter(
                components.ListInvoicesParamsFilterDateTimeFieldFilterCreatedAtDateTimeFieldGTFilter{
                    Gt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InvoicePagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListInvoicesRequest](../../models/operations/listinvoicesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListInvoicesResponse](../../models/operations/listinvoicesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetInvoice

Get a billing invoice by ID.

Returns the full invoice resource including line items, status details, totals,
and workflow configuration snapshot.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-invoice" method="get" path="/v3/openmeter/billing/invoices/{invoiceId}" -->
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

    res, err := s.OpenMeterBillingSettings.GetInvoice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingInvoice != nil {
        switch res.BillingInvoice.Type {
            case components.BillingInvoiceTypeStandard:
                // res.BillingInvoice.BillingInvoiceStandard is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `invoiceID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetInvoiceResponse](../../models/operations/getinvoiceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateInvoice

Update a billing invoice.

Only the mutable fields of the invoice can be edited: description, labels,
supplier, customer, workflow settings, and top-level lines. Top-level lines are
matched by `id`; lines without an `id` are created, and existing lines omitted
from `lines` are deleted. Detailed (child) lines are always computed and cannot
be edited directly. Only invoices in draft status can be updated.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-invoice" method="put" path="/v3/openmeter/billing/invoices/{invoiceId}" -->
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

    res, err := s.OpenMeterBillingSettings.UpdateInvoice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateUpdateInvoiceRequestStandard(
        components.UpdateInvoiceStandardRequest{
            Supplier: components.UpdateInvoiceStandardRequestSupplier{
                Addresses: &components.UpdateInvoiceStandardRequestAddresses{
                    BillingAddress: components.UpdateInvoiceStandardRequestBillingAddress{
                        Country: sdkkonnectgo.Pointer("US"),
                    },
                },
            },
            Customer: components.UpdateInvoiceStandardRequestCustomer{
                Name: "<value>",
                BillingAddress: &components.UpdateInvoiceStandardRequestCustomerBillingAddress{
                    Country: sdkkonnectgo.Pointer("US"),
                },
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
            },
            Type: components.UpdateInvoiceStandardRequestTypeStandard,
            Workflow: components.UpdateInvoiceStandardRequestWorkflow{
                Workflow: components.UpdateInvoiceStandardRequestWorkflowConfig{
                    Invoicing: &components.UpdateInvoiceStandardRequestInvoicingSettings{
                        DraftPeriod: sdkkonnectgo.Pointer("P1D"),
                    },
                    Payment: sdkkonnectgo.Pointer(components.CreateUpdateInvoiceStandardRequestPaymentSettingsChargeAutomatically(
                        components.UpdateBillingWorkflowPaymentChargeAutomaticallySettings{
                            CollectionMethod: components.UpdateBillingWorkflowPaymentChargeAutomaticallySettingsCollectionMethodChargeAutomatically,
                        },
                    )),
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingInvoice != nil {
        switch res.BillingInvoice.Type {
            case components.BillingInvoiceTypeStandard:
                // res.BillingInvoice.BillingInvoiceStandard is populated
        }

    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `invoiceID`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | 01G65Z755AFWAKHE12NY0CQ9FH                                                         |
| `updateInvoiceRequest`                                                             | [components.UpdateInvoiceRequest](../../models/components/updateinvoicerequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.UpdateInvoiceResponse](../../models/operations/updateinvoiceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteInvoice

Delete a billing invoice.

Only standard invoices in draft status can be deleted. Deleting an invoice will
also delete all associated line items and workflow configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-invoice" method="delete" path="/v3/openmeter/billing/invoices/{invoiceId}" -->
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

    res, err := s.OpenMeterBillingSettings.DeleteInvoice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `invoiceID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteInvoiceResponse](../../models/operations/deleteinvoiceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## AdvanceInvoice

Advance a billing invoice.

Advances the invoice to the next workflow state. The next state is determined by
the invoice's current status and workflow configuration. Only invoices in draft
or issued status can be advanced.

### Example Usage

<!-- UsageSnippet language="go" operationID="advance-invoice" method="post" path="/v3/openmeter/billing/invoices/{invoiceId}/advance" -->
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

    res, err := s.OpenMeterBillingSettings.AdvanceInvoice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingInvoice != nil {
        switch res.BillingInvoice.Type {
            case components.BillingInvoiceTypeStandard:
                // res.BillingInvoice.BillingInvoiceStandard is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `invoiceID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AdvanceInvoiceResponse](../../models/operations/advanceinvoiceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ApproveInvoice

Approve a billing invoice.

This call instantly sends the invoice to the customer using the configured
billing profile app.

This call is valid in two invoice statuses:

- draft: the invoice will be sent to the customer, the invoice state becomes
issued
- manual_approval_needed: the invoice will be sent to the customer, the invoice
state becomes issued

### Example Usage

<!-- UsageSnippet language="go" operationID="approve-invoice" method="post" path="/v3/openmeter/billing/invoices/{invoiceId}/approve" -->
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

    res, err := s.OpenMeterBillingSettings.ApproveInvoice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingInvoice != nil {
        switch res.BillingInvoice.Type {
            case components.BillingInvoiceTypeStandard:
                // res.BillingInvoice.BillingInvoiceStandard is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `invoiceID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ApproveInvoiceResponse](../../models/operations/approveinvoiceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RetryInvoice

Retry sending a billing invoice.

Retry advancing the invoice after a failed attempt.

The action can be called when the invoice's statusDetails' actions field contain
the "retry" action.

### Example Usage

<!-- UsageSnippet language="go" operationID="retry-invoice" method="post" path="/v3/openmeter/billing/invoices/{invoiceId}/retry" -->
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

    res, err := s.OpenMeterBillingSettings.RetryInvoice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingInvoice != nil {
        switch res.BillingInvoice.Type {
            case components.BillingInvoiceTypeStandard:
                // res.BillingInvoice.BillingInvoiceStandard is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `invoiceID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetryInvoiceResponse](../../models/operations/retryinvoiceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## SnapshotQuantitiesInvoice

Snapshot quantities for usage-based line items.

This call will snapshot the quantities for all usage based line items in the
invoice.

This call is only valid in draft.waiting_for_collection status, where the
collection period can be skipped using this action.

### Example Usage

<!-- UsageSnippet language="go" operationID="snapshot-quantities-invoice" method="post" path="/v3/openmeter/billing/invoices/{invoiceId}/snapshot-quantities" -->
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

    res, err := s.OpenMeterBillingSettings.SnapshotQuantitiesInvoice(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingInvoice != nil {
        switch res.BillingInvoice.Type {
            case components.BillingInvoiceTypeStandard:
                // res.BillingInvoice.BillingInvoiceStandard is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `invoiceID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.SnapshotQuantitiesInvoiceResponse](../../models/operations/snapshotquantitiesinvoiceresponse.md), error**

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

    res, err := s.OpenMeterBillingSettings.ListBillingProfiles(ctx, nil)
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

    res, err := s.OpenMeterBillingSettings.CreateBillingProfile(ctx, components.CreateBillingProfileRequest{
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

    res, err := s.OpenMeterBillingSettings.GetBillingProfile(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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

    res, err := s.OpenMeterBillingSettings.UpdateBillingProfile(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpsertBillingProfileRequest{
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

    res, err := s.OpenMeterBillingSettings.DeleteBillingProfile(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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