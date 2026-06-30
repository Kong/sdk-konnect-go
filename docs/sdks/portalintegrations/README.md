# PortalIntegrations

## Overview

APIs to configure Konnect Developer Portal integrations.

### Available Operations

* [GetPortalIntegrations](#getportalintegrations) - Get Portal Integration Configurations
* [UpsertPortalIntegrations](#upsertportalintegrations) - Replace Integration Configurations
* [UpdatePortalIntegrations](#updateportalintegrations) - Update Integration Configurations

## GetPortalIntegrations

Returns the portal integration configurations.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-portal-integrations" method="get" path="/v3/portals/{portalId}/integrations" -->
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

    res, err := s.PortalIntegrations.GetPortalIntegrations(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalIntegrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalIntegrationsResponse](../../models/operations/getportalintegrationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertPortalIntegrations

Replace the portal integration configurations.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-portal-integrations" method="put" path="/v3/portals/{portalId}/integrations" -->
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

    res, err := s.PortalIntegrations.UpsertPortalIntegrations(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalIntegrations{
        GoogleTagManager: &components.GoogleTagManagerIntegration{
            Enabled: false,
            Type: components.GoogleTagManagerIntegrationTypeTracking,
            ConfigData: components.ConfigData{
                ID: "GTM-XXXXXX",
                L: sdkkonnectgo.Pointer("myNewName"),
                Preview: sdkkonnectgo.Pointer("preview"),
                DataLayer: sdkkonnectgo.Pointer("myNewName"),
                EnvName: sdkkonnectgo.Pointer("production"),
                AuthReferrerPolicy: sdkkonnectgo.Pointer("no-referrer"),
            },
        },
        GoogleAnalytics4: &components.GoogleAnalytics4Integration{
            Enabled: true,
            Type: components.GoogleAnalytics4IntegrationTypeAnalytics,
            ConfigData: components.GoogleAnalytics4IntegrationConfigData{
                ID: "G-XXXXXXX",
                L: sdkkonnectgo.Pointer("myNewName"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalIntegrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |                                                                                 |
| `portalID`                                                                      | `string`                                                                        | :heavy_check_mark:                                                              | ID of the portal.                                                               | f32d905a-ed33-46a3-a093-d8f536af9a8a                                            |
| `portalIntegrations`                                                            | [*components.PortalIntegrations](../../models/components/portalintegrations.md) | :heavy_minus_sign:                                                              | N/A                                                                             |                                                                                 |
| `opts`                                                                          | [][operations.Option](../../models/operations/option.md)                        | :heavy_minus_sign:                                                              | The options for this request.                                                   |                                                                                 |

### Response

**[*operations.UpsertPortalIntegrationsResponse](../../models/operations/upsertportalintegrationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalIntegrations

Update the portal integration configurations, merging properties.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-portal-integrations" method="patch" path="/v3/portals/{portalId}/integrations" -->
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

    res, err := s.PortalIntegrations.UpdatePortalIntegrations(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.UpdatePortalIntegrations{
        GoogleTagManager: &components.UpdateGoogleTagManagerIntegration{
            ConfigData: &components.GoogleTagManagerIntegrationConfigProperties{
                ID: sdkkonnectgo.Pointer("GTM-XXXXXX"),
                L: sdkkonnectgo.Pointer("myNewName"),
                Preview: sdkkonnectgo.Pointer("preview"),
                DataLayer: sdkkonnectgo.Pointer("myNewName"),
                EnvName: sdkkonnectgo.Pointer("production"),
                AuthReferrerPolicy: sdkkonnectgo.Pointer("no-referrer"),
            },
        },
        GoogleAnalytics4: &components.UpdateGoogleAnalytics4Integration{
            ConfigData: &components.GoogleAnalytics4IntegrationConfigProperties{
                ID: sdkkonnectgo.Pointer("G-XXXXXXX"),
                L: sdkkonnectgo.Pointer("myNewName"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalIntegrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `portalID`                                                                                  | `string`                                                                                    | :heavy_check_mark:                                                                          | ID of the portal.                                                                           | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                        |
| `updatePortalIntegrations`                                                                  | [*components.UpdatePortalIntegrations](../../models/components/updateportalintegrations.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.UpdatePortalIntegrationsResponse](../../models/operations/updateportalintegrationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |