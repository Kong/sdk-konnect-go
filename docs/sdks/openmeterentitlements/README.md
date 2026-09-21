# OpenMeterEntitlements

## Overview

Entitlements are used to control access to features for customers.

### Available Operations

* [ListCustomerEntitlementAccess](#listcustomerentitlementaccess) - List customer entitlement access
* [QueryEntitlementAccess](#queryentitlementaccess) - Query entitlement access

## ListCustomerEntitlementAccess

List customer entitlement access

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customer-entitlement-access" method="get" path="/v3/openmeter/customers/{customerId}/entitlement-access" -->
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

    res, err := s.OpenMeterEntitlements.ListCustomerEntitlementAccess(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCustomerEntitlementAccessResponseData != nil {
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

**[*operations.ListCustomerEntitlementAccessResponse](../../models/operations/listcustomerentitlementaccessresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## QueryEntitlementAccess

Query feature access for a list of customers.

The endpoint resolves each provided identifier to a customer and returns the
access status for the requested features, plus optional credit balance
availability.

_Designed to be called on a fixed refresh interval and the query response is
intended to be cached._

### Example Usage

<!-- UsageSnippet language="go" operationID="query-entitlement-access" method="post" path="/v3/openmeter/entitlement-access/query" -->
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

    res, err := s.OpenMeterEntitlements.QueryEntitlementAccess(ctx, components.EntitlementAccessQueryRequest{
        Customer: components.EntitlementAccessQueryRequestCustomer{
            Keys: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementAccessQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `entitlementAccessQueryRequest`                                                                      | [components.EntitlementAccessQueryRequest](../../models/components/entitlementaccessqueryrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `page`                                                                                               | [*components.CursorPaginationQueryPage](../../models/components/cursorpaginationquerypage.md)        | :heavy_minus_sign:                                                                                   | Determines which page of the collection to retrieve.                                                 |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.QueryEntitlementAccessResponse](../../models/operations/queryentitlementaccessresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |