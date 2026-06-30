# SuggestedResourceActions

## Overview

Provides suggested actions to take on resources. Suggested actions are generated from suggestion rules.


### Available Operations

* [ListSuggestedResourceAction](#listsuggestedresourceaction) - List Suggested Resource Actions
* [GetSuggestedResourceAction](#getsuggestedresourceaction) - Get a Suggested Resource Action
* [UpdateSuggestedResourceAction](#updatesuggestedresourceaction) - Update Suggested Resource Action

## ListSuggestedResourceAction

Returns a paginated collection of a suggested resource actions.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-suggested-resource-action" method="get" path="/v1/suggested-resource-actions" -->
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

    res, err := s.SuggestedResourceActions.ListSuggestedResourceAction(ctx, operations.ListSuggestedResourceActionRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.SuggestedResourceActionFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTFilter(
                components.DateTimeFieldGTFilter{
                    Gt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Gte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSuggestedResourceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListSuggestedResourceActionRequest](../../models/operations/listsuggestedresourceactionrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListSuggestedResourceActionResponse](../../models/operations/listsuggestedresourceactionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetSuggestedResourceAction

Returns information about a suggested resource action.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-suggested-resource-action" method="get" path="/v1/suggested-resource-actions/{suggestedResourceActionId}" -->
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

    res, err := s.SuggestedResourceActions.GetSuggestedResourceAction(ctx, "76ca0689-f08e-4c04-b2a6-4c992d89d554")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuggestedResourceActionObject != nil {
        switch res.SuggestedResourceActionObject.Action.Type {
            case components.SuggestedResourceActionTypeArchiveActionPayload:
                // res.SuggestedResourceActionObject.Action.ArchiveActionPayload is populated
            case components.SuggestedResourceActionTypeIgnoreActionPayload:
                // res.SuggestedResourceActionObject.Action.IgnoreActionPayload is populated
            case components.SuggestedResourceActionTypeMapServiceAction:
                // res.SuggestedResourceActionObject.Action.MapServiceAction is populated
            case components.SuggestedResourceActionTypeCreateAndMapServiceActionPayload:
                // res.SuggestedResourceActionObject.Action.CreateAndMapServiceActionPayload is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `suggestedResourceActionID`                              | `string`                                                 | :heavy_check_mark:                                       | ID of the suggested resource action.                     | 76ca0689-f08e-4c04-b2a6-4c992d89d554                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSuggestedResourceActionResponse](../../models/operations/getsuggestedresourceactionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateSuggestedResourceAction

Updates the given suggested resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-suggested-resource-action" method="patch" path="/v1/suggested-resource-actions/{suggestedResourceActionId}" -->
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

    res, err := s.SuggestedResourceActions.UpdateSuggestedResourceAction(ctx, "76ca0689-f08e-4c04-b2a6-4c992d89d554", components.UpdateSuggestedResourceAction{
        Ignored: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SuggestedResourceActionObject != nil {
        switch res.SuggestedResourceActionObject.Action.Type {
            case components.SuggestedResourceActionTypeArchiveActionPayload:
                // res.SuggestedResourceActionObject.Action.ArchiveActionPayload is populated
            case components.SuggestedResourceActionTypeIgnoreActionPayload:
                // res.SuggestedResourceActionObject.Action.IgnoreActionPayload is populated
            case components.SuggestedResourceActionTypeMapServiceAction:
                // res.SuggestedResourceActionObject.Action.MapServiceAction is populated
            case components.SuggestedResourceActionTypeCreateAndMapServiceActionPayload:
                // res.SuggestedResourceActionObject.Action.CreateAndMapServiceActionPayload is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `suggestedResourceActionID`                                                                          | `string`                                                                                             | :heavy_check_mark:                                                                                   | ID of the suggested resource action.                                                                 | 76ca0689-f08e-4c04-b2a6-4c992d89d554                                                                 |
| `updateSuggestedResourceAction`                                                                      | [components.UpdateSuggestedResourceAction](../../models/components/updatesuggestedresourceaction.md) | :heavy_check_mark:                                                                                   | Request body schema for updating a suggested resource action.                                        |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.UpdateSuggestedResourceActionResponse](../../models/operations/updatesuggestedresourceactionresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |