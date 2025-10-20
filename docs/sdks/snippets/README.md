# Snippets
(*Snippets*)

## Overview

APIs related to Konnect Developer Portal Custom Snippets.

### Available Operations

* [ListPortalSnippets](#listportalsnippets) - List Snippets
* [CreatePortalSnippet](#createportalsnippet) - Create Snippet
* [GetPortalSnippet](#getportalsnippet) - Fetch Snippet
* [UpdatePortalSnippet](#updateportalsnippet) - Update Snippet
* [DeletePortalSnippet](#deleteportalsnippet) - Delete Snippet

## ListPortalSnippets

Returns the paginated list of custom snippets that have been created for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-portal-snippets" method="get" path="/v3/portals/{portalId}/snippets" -->
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

    res, err := s.Snippets.ListPortalSnippets(ctx, operations.ListPortalSnippetsRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.PortalSnippetsFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalSnippetsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListPortalSnippetsRequest](../../models/operations/listportalsnippetsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListPortalSnippetsResponse](../../models/operations/listportalsnippetsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortalSnippet

Creates a new custom snippet for this portal. Custom snippets can be used to display static content, documentation, or other information to developers. Title and Description properties may be provided in the frontmatter section of `content`. If you set values in both the `POST` request _and_ in the frontmatter, the values in the frontmatter will take precedence.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-portal-snippet" method="post" path="/v3/portals/{portalId}/snippets" -->
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

    res, err := s.Snippets.CreatePortalSnippet(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreatePortalSnippetRequest{
        Name: "my-snippet",
        Title: sdkkonnectgo.Pointer("My Snippet"),
        Content: "# Welcome to My Snippet",
        Visibility: components.SnippetVisibilityStatusPublic.ToPointer(),
        Status: components.PublishedStatusPublished.ToPointer(),
        Description: sdkkonnectgo.Pointer("A custom page about developer portals"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalSnippetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `portalID`                                                                                     | *string*                                                                                       | :heavy_check_mark:                                                                             | ID of the portal.                                                                              | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                           |
| `createPortalSnippetRequest`                                                                   | [components.CreatePortalSnippetRequest](../../models/components/createportalsnippetrequest.md) | :heavy_check_mark:                                                                             | Create a snippet in a portal.                                                                  |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.CreatePortalSnippetResponse](../../models/operations/createportalsnippetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalSnippet

Returns the configuration of a single custom snippet for this portal. Custom snippets can be used to display static content, documentation, or other information to developers.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-snippet" method="get" path="/v3/portals/{portalId}/snippets/{snippetId}" -->
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

    res, err := s.Snippets.GetPortalSnippet(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "ebbac5b0-ac89-45c3-9d2e-c4542c657e79")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalSnippetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `snippetID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the snippet.                                       | ebbac5b0-ac89-45c3-9d2e-c4542c657e79                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalSnippetResponse](../../models/operations/getportalsnippetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalSnippet

Updates the configuration of a single custom snippet for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-snippet" method="patch" path="/v3/portals/{portalId}/snippets/{snippetId}" -->
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

    res, err := s.Snippets.UpdatePortalSnippet(ctx, operations.UpdatePortalSnippetRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        SnippetID: "ebbac5b0-ac89-45c3-9d2e-c4542c657e79",
        UpdatePortalSnippetRequest: components.UpdatePortalSnippetRequest{
            Name: sdkkonnectgo.Pointer("my-snippet"),
            Title: sdkkonnectgo.Pointer("My Snippet"),
            Content: sdkkonnectgo.Pointer("# Welcome to My Snippet"),
            Visibility: components.VisibilityStatusPublic.ToPointer(),
            Status: components.PublishedStatusPublished.ToPointer(),
            Description: sdkkonnectgo.Pointer("A custom page about developer portals"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalSnippetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdatePortalSnippetRequest](../../models/operations/updateportalsnippetrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdatePortalSnippetResponse](../../models/operations/updateportalsnippetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalSnippet

Deletes a single custom snippet for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-snippet" method="delete" path="/v3/portals/{portalId}/snippets/{snippetId}" -->
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

    res, err := s.Snippets.DeletePortalSnippet(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "ebbac5b0-ac89-45c3-9d2e-c4542c657e79")
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
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `snippetID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the snippet.                                       | ebbac5b0-ac89-45c3-9d2e-c4542c657e79                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalSnippetResponse](../../models/operations/deleteportalsnippetresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |