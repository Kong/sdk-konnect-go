# Pages
(*Pages*)

## Overview

APIs related to Konnect Developer Portal Custom Pages.

### Available Operations

* [CreateDefaultContent](#createdefaultcontent) - Creates Default Pages
* [ListPortalPages](#listportalpages) - List Pages
* [CreatePortalPage](#createportalpage) - Create Page
* [GetPortalPage](#getportalpage) - Get a Page
* [UpdatePortalPage](#updateportalpage) - Update Page
* [DeletePortalPage](#deleteportalpage) - Delete Page
* [MovePortalPages](#moveportalpages) - Move Page

## CreateDefaultContent

Creates the default pages for a portal if they do not already exists.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-default-content" method="post" path="/v3/portals/{portalId}/default-content" -->
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

    res, err := s.Pages.CreateDefaultContent(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.DefaultContent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.CreateDefaultContentResponse](../../models/operations/createdefaultcontentresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPortalPages

Returns the tree view of custom pages that have been created for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-portal-pages" method="get" path="/v3/portals/{portalId}/pages" -->
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

    res, err := s.Pages.ListPortalPages(ctx, operations.ListPortalPagesRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        Filter: &components.PortalPagesFilterParameters{
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalPagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListPortalPagesRequest](../../models/operations/listportalpagesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListPortalPagesResponse](../../models/operations/listportalpagesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortalPage

Creates a new custom page for this portal. Custom pages can be used to display static content, documentation, or other information to developers. Title and Description properties may be provided in the frontmatter section of `content`. If you set values in both the `POST` request _and_ in the frontmatter, the values in frontmatter will take precedence.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-portal-page" method="post" path="/v3/portals/{portalId}/pages" -->
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

    res, err := s.Pages.CreatePortalPage(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreatePortalPageRequest{
        Slug: "/my-page",
        Title: sdkkonnectgo.Pointer("My Page"),
        Content: "# Welcome to My Page",
        Visibility: components.PageVisibilityStatusPublic.ToPointer(),
        Status: components.PublishedStatusPublished.ToPointer(),
        Description: sdkkonnectgo.Pointer("A custom page about developer portals"),
        ParentPageID: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalPageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `portalID`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the portal.                                                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                     |
| `createPortalPageRequest`                                                                | [components.CreatePortalPageRequest](../../models/components/createportalpagerequest.md) | :heavy_check_mark:                                                                       | Create a page in a portal.                                                               |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreatePortalPageResponse](../../models/operations/createportalpageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalPage

Returns the configuration of a single custom page for this portal. Custom pages can be used to display static content, documentation, or other information to developers.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-page" method="get" path="/v3/portals/{portalId}/pages/{pageId}" -->
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

    res, err := s.Pages.GetPortalPage(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "ebbac5b0-ac89-45c3-9d2e-c4542c657e79")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalPageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `pageID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the page.                                          | ebbac5b0-ac89-45c3-9d2e-c4542c657e79                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalPageResponse](../../models/operations/getportalpageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalPage

Updates the configuration of a single custom page for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-page" method="patch" path="/v3/portals/{portalId}/pages/{pageId}" -->
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

    res, err := s.Pages.UpdatePortalPage(ctx, operations.UpdatePortalPageRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageID: "ebbac5b0-ac89-45c3-9d2e-c4542c657e79",
        UpdatePortalPageRequest: components.UpdatePortalPageRequest{
            Slug: sdkkonnectgo.Pointer("/my-page"),
            Title: sdkkonnectgo.Pointer("My Page"),
            Content: sdkkonnectgo.Pointer("# Welcome to My Page"),
            Visibility: components.VisibilityStatusPublic.ToPointer(),
            Status: components.PublishedStatusPublished.ToPointer(),
            Description: sdkkonnectgo.Pointer("A custom page about developer portals"),
            ParentPageID: nil,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalPageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdatePortalPageRequest](../../models/operations/updateportalpagerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdatePortalPageResponse](../../models/operations/updateportalpageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalPage

Deletes a single custom page for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-page" method="delete" path="/v3/portals/{portalId}/pages/{pageId}" -->
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

    res, err := s.Pages.DeletePortalPage(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "ebbac5b0-ac89-45c3-9d2e-c4542c657e79")
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
| `pageID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the page.                                          | ebbac5b0-ac89-45c3-9d2e-c4542c657e79                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalPageResponse](../../models/operations/deleteportalpageresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## MovePortalPages

This api allows the user to move a page within the page tree using the parameters parent_page_id and index. If parent_page_id is not provided, the page will be placed at the top level of the page tree. index represents a zero-indexed page order relative to its siblings under the same parent. For example, if we want to put the page at top level in first position we would send parent_page_id: null and index: 0. This api also supports using a negative index to count backwards from the end of the page list, which means you can put the page in last position by using index: -1.

### Example Usage

<!-- UsageSnippet language="go" operationID="move-portal-pages" method="post" path="/v3/portals/{portalId}/pages/{pageId}/move" -->
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

    res, err := s.Pages.MovePortalPages(ctx, operations.MovePortalPagesRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageID: "ebbac5b0-ac89-45c3-9d2e-c4542c657e79",
        MovePageRequestPayload: &components.MovePageRequestPayload{
            ParentPageID: sdkkonnectgo.Pointer("dd4e1b98-3629-4dd3-acc0-759a726ffee2"),
            Index: sdkkonnectgo.Pointer[int64](1),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.MovePortalPagesRequest](../../models/operations/moveportalpagesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.MovePortalPagesResponse](../../models/operations/moveportalpagesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |