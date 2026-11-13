# PortalNavigation

## Overview

APIs related to Konnect Developer Portal sidebar navigation configuration.

### Available Operations

* [ListPortalNavigations](#listportalnavigations) - List Navigation configurations
* [CreatePortalNavigation](#createportalnavigation) - Create Navigation configuration
* [GetPortalNavigation](#getportalnavigation) - Get a Navigation configuration
* [ReplacePortalNavigation](#replaceportalnavigation) - Replace Navigation configuration
* [DeletePortalNavigation](#deleteportalnavigation) - Delete Navigation configuration

## ListPortalNavigations

Returns a paginated list of the portal's sidebar navigation configurations. Each navigation is matched to pages by its `show_on_path_prefixes` and displayed automatically when using a named page layout that supports the sidebar UX.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-portal-navigations" method="get" path="/v3/portals/{portalId}/customization/navigations" -->
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

    res, err := s.PortalNavigation.ListPortalNavigations(ctx, operations.ListPortalNavigationsRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalNavigationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListPortalNavigationsRequest](../../models/operations/listportalnavigationsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListPortalNavigationsResponse](../../models/operations/listportalnavigationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortalNavigation

Creates a sidebar navigation configuration for this portal. A navigation claims every page under its `show_on_path_prefixes`, and the longest matching entry wins when more than one navigation is configured.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-portal-navigation" method="post" path="/v3/portals/{portalId}/customization/navigations" -->
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

    res, err := s.PortalNavigation.CreatePortalNavigation(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreatePortalNavigationRequest{
        ShowOnPathPrefixes: []string{
            "/guides",
        },
        Name: "Main",
        Groups: []components.NavigationLinkGroup{
            components.NavigationLinkGroup{
                Header: sdkkonnectgo.Pointer("Getting started"),
                Links: []components.NavigationLink{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalNavigationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `portalID`                                                                                           | `string`                                                                                             | :heavy_check_mark:                                                                                   | ID of the portal.                                                                                    | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                 |
| `createPortalNavigationRequest`                                                                      | [components.CreatePortalNavigationRequest](../../models/components/createportalnavigationrequest.md) | :heavy_check_mark:                                                                                   | Create a sidebar navigation configuration for a portal.                                              |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CreatePortalNavigationResponse](../../models/operations/createportalnavigationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalNavigation

Returns the configuration of a single sidebar navigation for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-navigation" method="get" path="/v3/portals/{portalId}/customization/navigations/{navigationId}" -->
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

    res, err := s.PortalNavigation.GetPortalNavigation(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "ebbac5b0-ac89-45c3-9d2e-c4542c657e79")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalNavigationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `navigationID`                                           | `string`                                                 | :heavy_check_mark:                                       | ID of the navigation configuration.                      | ebbac5b0-ac89-45c3-9d2e-c4542c657e79                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalNavigationResponse](../../models/operations/getportalnavigationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePortalNavigation

Replaces the configuration of a single sidebar navigation for this portal. This is a full replacement; partial updates are not supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="replace-portal-navigation" method="put" path="/v3/portals/{portalId}/customization/navigations/{navigationId}" -->
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

    res, err := s.PortalNavigation.ReplacePortalNavigation(ctx, operations.ReplacePortalNavigationRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        NavigationID: "ebbac5b0-ac89-45c3-9d2e-c4542c657e79",
        ReplacePortalNavigationRequest: components.ReplacePortalNavigationRequest{
            ShowOnPathPrefixes: []string{
                "/guides",
            },
            Name: "Main",
            Groups: []components.NavigationLinkGroup{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalNavigationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ReplacePortalNavigationRequest](../../models/operations/replaceportalnavigationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ReplacePortalNavigationResponse](../../models/operations/replaceportalnavigationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalNavigation

Deletes a single sidebar navigation configuration for this portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-navigation" method="delete" path="/v3/portals/{portalId}/customization/navigations/{navigationId}" -->
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

    res, err := s.PortalNavigation.DeletePortalNavigation(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "ebbac5b0-ac89-45c3-9d2e-c4542c657e79")
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
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `navigationID`                                           | `string`                                                 | :heavy_check_mark:                                       | ID of the navigation configuration.                      | ebbac5b0-ac89-45c3-9d2e-c4542c657e79                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalNavigationResponse](../../models/operations/deleteportalnavigationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |