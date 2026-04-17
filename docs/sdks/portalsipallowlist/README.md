# PortalsIPAllowList

## Overview

APIs related to Konnect Portal IP Allow List.

### Available Operations

* [CreatePortalIPAllowList](#createportalipallowlist) - Create an IP allow list for a portal
* [ListPortalIPAllowList](#listportalipallowlist) - List the IP allow list for portal
* [PutPortalIPAllowList](#putportalipallowlist) - Replace an IP allow list for a portal
* [UpdatePortalIPAllowList](#updateportalipallowlist) - Update an IP allow list for a portal
* [DeletePortalIPAllowList](#deleteportalipallowlist) - Delete an IP allow list from a portal

## CreatePortalIPAllowList

Create an IP allow list for a portal.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-portal-ip-allow-list" method="post" path="/v3/portals/{portalId}/ip-allow-list" -->
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

    res, err := s.PortalsIPAllowList.CreatePortalIPAllowList(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.CreatePortalSourceIPRestriction{
        AllowedIps: []string{
            "192.168.1.1",
            "192.168.1.0/22",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IPEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               | Example                                                                                                   |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |                                                                                                           |
| `portalID`                                                                                                | `string`                                                                                                  | :heavy_check_mark:                                                                                        | ID of the portal.                                                                                         | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                      |
| `createPortalSourceIPRestriction`                                                                         | [*components.CreatePortalSourceIPRestriction](../../models/components/createportalsourceiprestriction.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |                                                                                                           |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |                                                                                                           |

### Response

**[*operations.CreatePortalIPAllowListResponse](../../models/operations/createportalipallowlistresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPortalIPAllowList

Lists the IP allow list configuration for a portal.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-portal-ip-allow-list" method="get" path="/v3/portals/{portalId}/ip-allow-list" -->
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

    res, err := s.PortalsIPAllowList.ListPortalIPAllowList(ctx, operations.ListPortalIPAllowListRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageBefore: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalSourceIPRestrictionPaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListPortalIPAllowListRequest](../../models/operations/listportalipallowlistrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListPortalIPAllowListResponse](../../models/operations/listportalipallowlistresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PutPortalIPAllowList

Replace an IP allow list for a portal.


### Example Usage

<!-- UsageSnippet language="go" operationID="put-portal-ip-allow-list" method="put" path="/v3/portals/{portalId}/ip-allow-list/{id}" -->
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

    res, err := s.PortalsIPAllowList.PutPortalIPAllowList(ctx, operations.PutPortalIPAllowListRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "257419c5-063c-450f-9743-115008a0b2de",
        CreatePortalSourceIPRestriction: &components.CreatePortalSourceIPRestriction{
            AllowedIps: []string{
                "192.168.1.1",
                "192.168.1.0/22",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IPEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PutPortalIPAllowListRequest](../../models/operations/putportalipallowlistrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PutPortalIPAllowListResponse](../../models/operations/putportalipallowlistresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalIPAllowList

Update an IP allow list for a portal.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-ip-allow-list" method="patch" path="/v3/portals/{portalId}/ip-allow-list/{id}" -->
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

    res, err := s.PortalsIPAllowList.UpdatePortalIPAllowList(ctx, operations.UpdatePortalIPAllowListRequest{
        PortalID: "f32d905a-ed33-46a3-a093-d8f536af9a8a",
        ID: "ad9aaa6e-8b26-4d71-9e8d-c715e0dbe4cd",
        UpdatePortalSourceIPRestriction: &components.UpdatePortalSourceIPRestriction{
            AllowedIps: []string{
                "192.168.1.1",
                "192.168.1.0/22",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IPEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdatePortalIPAllowListRequest](../../models/operations/updateportalipallowlistrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdatePortalIPAllowListResponse](../../models/operations/updateportalipallowlistresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalIPAllowList

Delete the IP allow list for a portal.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-ip-allow-list" method="delete" path="/v3/portals/{portalId}/ip-allow-list/{id}" -->
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

    res, err := s.PortalsIPAllowList.DeletePortalIPAllowList(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", "3feb9c9c-57e9-412a-a917-d031b9750059")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | ID of the allow list.                                    |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalIPAllowListResponse](../../models/operations/deleteportalipallowlistresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |