# OpenMeterBilling

## Overview

Billing manages the billing profiles, currencies, cost bases, and invoices for customers.

### Available Operations

* [ListBillingProfiles](#listbillingprofiles) - List billing profiles
* [CreateBillingProfile](#createbillingprofile) - Create a new billing profile
* [GetBillingProfile](#getbillingprofile) - Get a billing profile
* [UpdateBillingProfile](#updatebillingprofile) - Update a billing profile
* [DeleteBillingProfile](#deletebillingprofile) - Delete a billing profile

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

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `page`                                                                        | [*metering.PagePaginationQuery](../../models/metering/pagepaginationquery.md) | :heavy_minus_sign:                                                            | Determines which page of the collection to retrieve.                          |
| `opts`                                                                        | [][operations.Option](../../models/operations/option.md)                      | :heavy_minus_sign:                                                            | The options for this request.                                                 |

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
	"github.com/Kong/sdk-konnect-go/models/metering"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.OpenMeterBilling.CreateBillingProfile(ctx, metering.CreateBillingProfileRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Supplier: metering.CreateBillingProfileRequestSupplier{
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
            Addresses: &metering.CreateBillingProfileRequestAddresses{
                BillingAddress: metering.CreateBillingProfileRequestBillingAddress{
                    Country: sdkkonnectgo.Pointer("US"),
                },
            },
        },
        Workflow: metering.CreateBillingProfileRequestWorkflow{
            Collection: &metering.CreateBillingProfileRequestWorkflowCollectionSettings{
                Alignment: sdkkonnectgo.Pointer(metering.CreateCreateBillingProfileRequestAlignmentSubscription(
                    metering.AlignmentBillingWorkflowCollectionAlignmentSubscription{
                        Type: metering.BillingWorkflowCollectionAlignmentSubscriptionAlignmentTypeSubscription,
                    },
                )),
                Interval: sdkkonnectgo.Pointer("P1D"),
            },
            Invoicing: &metering.CreateBillingProfileRequestWorkflowInvoiceSettings{
                DraftPeriod: sdkkonnectgo.Pointer("P1D"),
            },
            Payment: sdkkonnectgo.Pointer(metering.CreateCreateBillingProfileRequestPaymentChargeAutomatically(
                metering.BillingWorkflowPaymentChargeAutomaticallySettings{
                    CollectionMethod: metering.CollectionMethodChargeAutomatically,
                },
            )),
            Tax: &metering.CreateBillingProfileRequestWorkflowTaxSettings{
                DefaultTaxConfig: &metering.CreateBillingProfileRequestDefaultTaxConfig{
                    TaxCode: &metering.CreateBillingProfileRequestTaxCode{
                        ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                    },
                },
            },
        },
        Apps: metering.CreateBillingProfileRequestApps{
            Tax: metering.CreateBillingProfileRequestTax{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
            Invoicing: metering.CreateBillingProfileRequestInvoicing{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
            Payment: metering.CreateBillingProfileRequestMeteringPayment{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [metering.CreateBillingProfileRequest](../../models/metering/createbillingprofilerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

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
	"github.com/Kong/sdk-konnect-go/models/metering"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.OpenMeterBilling.UpdateBillingProfile(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", metering.UpsertBillingProfileRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Supplier: metering.UpsertBillingProfileRequestSupplier{
            Key: sdkkonnectgo.Pointer("019ae40f-4258-7f15-9491-842f42a7d6ac"),
            Addresses: &metering.UpsertBillingProfileRequestAddresses{
                BillingAddress: metering.UpsertBillingProfileRequestBillingAddress{
                    Country: sdkkonnectgo.Pointer("US"),
                },
            },
        },
        Workflow: metering.UpsertBillingProfileRequestWorkflow{
            Collection: &metering.UpsertBillingProfileRequestWorkflowCollectionSettings{
                Alignment: sdkkonnectgo.Pointer(metering.CreateUpsertBillingProfileRequestAlignmentSubscription(
                    metering.UpsertBillingProfileRequestAlignmentBillingWorkflowCollectionAlignmentSubscription{
                        Type: metering.BillingWorkflowCollectionAlignmentSubscriptionAlignmentUpsertBillingProfileRequestTypeSubscription,
                    },
                )),
                Interval: sdkkonnectgo.Pointer("P1D"),
            },
            Invoicing: &metering.UpsertBillingProfileRequestWorkflowInvoiceSettings{
                DraftPeriod: sdkkonnectgo.Pointer("P1D"),
            },
            Payment: sdkkonnectgo.Pointer(metering.CreateUpsertBillingProfileRequestPaymentChargeAutomatically(
                metering.BillingWorkflowPaymentChargeAutomaticallySettings{
                    CollectionMethod: metering.CollectionMethodChargeAutomatically,
                },
            )),
            Tax: &metering.UpsertBillingProfileRequestWorkflowTaxSettings{
                DefaultTaxConfig: &metering.UpsertBillingProfileRequestDefaultTaxConfig{
                    TaxCode: &metering.UpsertBillingProfileRequestTaxCode{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `id`                                                                                         | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          | 01G65Z755AFWAKHE12NY0CQ9FH                                                                   |
| `upsertBillingProfileRequest`                                                                | [metering.UpsertBillingProfileRequest](../../models/metering/upsertbillingprofilerequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

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