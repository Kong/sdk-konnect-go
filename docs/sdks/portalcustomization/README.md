# PortalCustomization
(*PortalCustomization*)

## Overview

APIs related to customization of Konnect Developer Portals.

### Available Operations

* [GetPortalCustomization](#getportalcustomization) - Get Customization
* [ReplacePortalCustomization](#replaceportalcustomization) - Replace Customization
* [UpdatePortalCustomization](#updateportalcustomization) - Update Customization

## GetPortalCustomization

Returns the portal customization options.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-customization" method="get" path="/v3/portals/{portalId}/customization" -->
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

    res, err := s.PortalCustomization.GetPortalCustomization(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomization != nil {
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

**[*operations.GetPortalCustomizationResponse](../../models/operations/getportalcustomizationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePortalCustomization

Replace the portal customization options.

### Example Usage

<!-- UsageSnippet language="go" operationID="replace-portal-customization" method="put" path="/v3/portals/{portalId}/customization" -->
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

    res, err := s.PortalCustomization.ReplacePortalCustomization(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalCustomization{
        Theme: &components.Theme{
            Colors: &components.Colors{
                Primary: sdkkonnectgo.Pointer("#000000"),
            },
        },
        Menu: &components.Menu{
            Main: []components.PortalMenuItem{
                components.PortalMenuItem{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: false,
                },
            },
            FooterSections: []components.PortalFooterMenuSection{
                components.PortalFooterMenuSection{
                    Title: "<value>",
                    Items: []components.PortalMenuItem{},
                },
            },
            FooterBottom: []components.PortalMenuItem{
                components.PortalMenuItem{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: false,
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |                                                                                   |
| `portalID`                                                                        | *string*                                                                          | :heavy_check_mark:                                                                | ID of the portal.                                                                 | f32d905a-ed33-46a3-a093-d8f536af9a8a                                              |
| `portalCustomization`                                                             | [*components.PortalCustomization](../../models/components/portalcustomization.md) | :heavy_minus_sign:                                                                | N/A                                                                               |                                                                                   |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |                                                                                   |

### Response

**[*operations.ReplacePortalCustomizationResponse](../../models/operations/replaceportalcustomizationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalCustomization

Update the portal customization options, merging properties.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-customization" method="patch" path="/v3/portals/{portalId}/customization" -->
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

    res, err := s.PortalCustomization.UpdatePortalCustomization(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalCustomization{
        Theme: &components.Theme{
            Colors: &components.Colors{
                Primary: sdkkonnectgo.Pointer("#000000"),
            },
        },
        Menu: &components.Menu{
            Main: []components.PortalMenuItem{
                components.PortalMenuItem{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: false,
                },
            },
            FooterSections: []components.PortalFooterMenuSection{
                components.PortalFooterMenuSection{
                    Title: "<value>",
                    Items: []components.PortalMenuItem{
                        components.PortalMenuItem{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: true,
                        },
                    },
                },
            },
            FooterBottom: []components.PortalMenuItem{
                components.PortalMenuItem{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: true,
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |                                                                                   |
| `portalID`                                                                        | *string*                                                                          | :heavy_check_mark:                                                                | ID of the portal.                                                                 | f32d905a-ed33-46a3-a093-d8f536af9a8a                                              |
| `portalCustomization`                                                             | [*components.PortalCustomization](../../models/components/portalcustomization.md) | :heavy_minus_sign:                                                                | N/A                                                                               |                                                                                   |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |                                                                                   |

### Response

**[*operations.UpdatePortalCustomizationResponse](../../models/operations/updateportalcustomizationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |