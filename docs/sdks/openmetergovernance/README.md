# OpenMeterGovernance

## Overview

Governance evaluation of customers to check their feature access.

### Available Operations

* [QueryGovernanceAccess](#querygovernanceaccess) - Query governance access

## QueryGovernanceAccess

Query feature access for a list of customers.

The endpoint resolves each provided identifier to a customer and returns the
access status for the requested features, plus optional credit balance
availability.

_Designed to be called on a fixed refresh interval and the query response is
intended to be cached._

### Example Usage

<!-- UsageSnippet language="go" operationID="query-governance-access" method="post" path="/v3/openmeter/governance/query" -->
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

    res, err := s.OpenMeterGovernance.QueryGovernanceAccess(ctx, components.GovernanceQueryRequest{
        Customer: components.GovernanceQueryRequestCustomer{
            Keys: []string{
                "<value 1>",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GovernanceQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `governanceQueryRequest`                                                                      | [components.GovernanceQueryRequest](../../models/components/governancequeryrequest.md)        | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `page`                                                                                        | [*components.CursorPaginationQueryPage](../../models/components/cursorpaginationquerypage.md) | :heavy_minus_sign:                                                                            | Determines which page of the collection to retrieve.                                          |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*operations.QueryGovernanceAccessResponse](../../models/operations/querygovernanceaccessresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |