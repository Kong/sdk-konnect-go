# OpenMeterDefaults

## Overview

Organization-level default configuration.

### Available Operations

* [GetOrganizationDefaultTaxCodes](#getorganizationdefaulttaxcodes) - Get organization default tax codes
* [UpdateOrganizationDefaultTaxCodes](#updateorganizationdefaulttaxcodes) - Update organization default tax codes

## GetOrganizationDefaultTaxCodes

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-organization-default-tax-codes" method="get" path="/v3/openmeter/defaults/tax-codes" -->
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

    res, err := s.OpenMeterDefaults.GetOrganizationDefaultTaxCodes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationDefaultTaxCodes != nil {
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

**[*operations.GetOrganizationDefaultTaxCodesResponse](../../models/operations/getorganizationdefaulttaxcodesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateOrganizationDefaultTaxCodes

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-organization-default-tax-codes" method="put" path="/v3/openmeter/defaults/tax-codes" -->
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

    res, err := s.OpenMeterDefaults.UpdateOrganizationDefaultTaxCodes(ctx, components.UpdateOrganizationDefaultTaxCodesRequest{
        InvoicingTaxCode: &components.UpdateOrganizationDefaultTaxCodesRequestInvoicingTaxCode{
            ID: "01G65Z755AFWAKHE12NY0CQ9FH",
        },
        CreditGrantTaxCode: &components.UpdateOrganizationDefaultTaxCodesRequestCreditGrantTaxCode{
            ID: "01G65Z755AFWAKHE12NY0CQ9FH",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationDefaultTaxCodes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [components.UpdateOrganizationDefaultTaxCodesRequest](../../models/components/updateorganizationdefaulttaxcodesrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpdateOrganizationDefaultTaxCodesResponse](../../models/operations/updateorganizationdefaulttaxcodesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |