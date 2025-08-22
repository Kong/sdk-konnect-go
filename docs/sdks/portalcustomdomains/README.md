# PortalCustomDomains
(*PortalCustomDomains*)

## Overview

APIs related to configuration of Konnect Developer Portals custom domains.

### Available Operations

* [GetPortalCustomDomain](#getportalcustomdomain) - Get Custom Domain
* [CreatePortalCustomDomain](#createportalcustomdomain) - Create Custom Domain
* [UpdatePortalCustomDomain](#updateportalcustomdomain) - Enable or Disable Domain
* [DeletePortalCustomDomain](#deleteportalcustomdomain) - Remove Domain

## GetPortalCustomDomain

Get the custom domain associated to the portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-custom-domain" method="get" path="/v3/portals/{portalId}/custom-domain" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalCustomDomains.GetPortalCustomDomain(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomDomain != nil {
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

**[*operations.GetPortalCustomDomainResponse](../../models/operations/getportalcustomdomainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortalCustomDomain

Creates the custom domain associated with the portal. Only one custom domain can be associated with a portal at a time.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-portal-custom-domain" method="post" path="/v3/portals/{portalId}/custom-domain" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalCustomDomains.CreatePortalCustomDomain(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreatePortalCustomDomainRequest{
        Hostname: "lavish-custody.org",
        Enabled: true,
        Ssl: components.CreateCreatePortalCustomDomainSSLHTTP(
            components.HTTP{},
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `portalID`                                                                                               | *string*                                                                                                 | :heavy_check_mark:                                                                                       | ID of the portal.                                                                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                     |
| `createPortalCustomDomainRequest`                                                                        | [components.CreatePortalCustomDomainRequest](../../models/components/createportalcustomdomainrequest.md) | :heavy_check_mark:                                                                                       | Create a portal custom domain.                                                                           |                                                                                                          |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[*operations.CreatePortalCustomDomainResponse](../../models/operations/createportalcustomdomainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalCustomDomain

Updates the portal domain associated with the portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-custom-domain" method="patch" path="/v3/portals/{portalId}/custom-domain" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalCustomDomains.UpdatePortalCustomDomain(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdatePortalCustomDomainRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `portalID`                                                                                               | *string*                                                                                                 | :heavy_check_mark:                                                                                       | ID of the portal.                                                                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                                     |
| `updatePortalCustomDomainRequest`                                                                        | [components.UpdatePortalCustomDomainRequest](../../models/components/updateportalcustomdomainrequest.md) | :heavy_check_mark:                                                                                       | Create a portal custom domain.                                                                           |                                                                                                          |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[*operations.UpdatePortalCustomDomainResponse](../../models/operations/updateportalcustomdomainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortalCustomDomain

Deletes the custom domain associated with the portal.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-portal-custom-domain" method="delete" path="/v3/portals/{portalId}/custom-domain" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PortalCustomDomains.DeletePortalCustomDomain(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePortalCustomDomainResponse](../../models/operations/deleteportalcustomdomainresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |