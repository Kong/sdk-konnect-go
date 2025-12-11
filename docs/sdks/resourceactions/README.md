# ResourceActions

## Overview

Resource actions provides an overview of actions taken on resources.
It provides information on when one of the following actions are taken on a resource:
- Map
- Unmap
- Archive
- Restore


### Available Operations

* [ListResourceActions](#listresourceactions) - List Resource Actions

## ListResourceActions

Returns a paginated collection of a resource actions.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-resource-actions" method="get" path="/v1/resource-actions" -->
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

    res, err := s.ResourceActions.ListResourceActions(ctx, &components.CursorPageParameters{
        Size: sdkkonnectgo.Pointer[int64](10),
        After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    }, &components.ResourceActionFilterParameters{
        CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
            components.DateTimeFieldLTFilter{
                Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |
| `page`                                                                                                  | [*components.CursorPageParameters](../../models/components/cursorpageparameters.md)                     | :heavy_minus_sign:                                                                                      | Determines which page of the collection to retrieve.                                                    |
| `filter`                                                                                                | [*components.ResourceActionFilterParameters](../../models/components/resourceactionfilterparameters.md) | :heavy_minus_sign:                                                                                      | Filters a collection of resource actions.                                                               |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |

### Response

**[*operations.ListResourceActionsResponse](../../models/operations/listresourceactionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |