# OpenMeterApps

## Overview

Apps enable you to extend and customize billing and usage workflows by integrating with external systems and services. Apps can automate and enhance your billing ecosystem by supporting capabilities such as synchronizing usage data with third-party platforms, calculating taxes, generating and delivering invoices, handling payment collection, and other billing-related tasks.

### Available Operations

* [ListAppCatalog](#listappcatalog) - List app catalog
* [InstallApp](#installapp) - Install app from the catalog
* [GetAppCatalogItem](#getappcatalogitem) - Get app catalog item by type
* [ListApps](#listapps) - List apps
* [GetApp](#getapp) - Get app
* [UninstallApp](#uninstallapp) - Uninstall app
* [UpdateApp](#updateapp) - Update app

## ListAppCatalog

List available apps.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-app-catalog" method="get" path="/v3/openmeter/app-catalog" -->
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

    res, err := s.OpenMeterApps.ListAppCatalog(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppCatalogItemPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `page`                                                                            | [*components.PagePaginationQuery](../../models/components/pagepaginationquery.md) | :heavy_minus_sign:                                                                | Determines which page of the collection to retrieve.                              |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.ListAppCatalogResponse](../../models/operations/listappcatalogresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## InstallApp

Install an app from the catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="install-app" method="post" path="/v3/openmeter/app-catalog/install" -->
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

    res, err := s.OpenMeterApps.InstallApp(ctx, components.CreateBillingInstallAppRequestStripe(
        components.BillingInstallAppStripeWithAPIKey{
            Type: components.BillingInstallAppStripeWithAPIKeyTypeStripe,
            Name: "<value>",
            CreateBillingProfile: false,
            APIKey: "hL8x61xtp5RsIFG",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingInstallAppResponse != nil {
        switch res.BillingInstallAppResponse.Type {
            case components.BillingInstallAppResponseTypeStripe:
                // res.BillingInstallAppResponse.BillingInstalledAppStripe is populated
            case components.BillingInstallAppResponseTypeSandbox:
                // res.BillingInstallAppResponse.BillingInstalledAppSandbox is populated
            case components.BillingInstallAppResponseTypeExternalInvoicing:
                // res.BillingInstallAppResponse.BillingInstalledAppExternalInvoicing is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.BillingInstallAppRequest](../../models/components/billinginstallapprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.InstallAppResponse](../../models/operations/installappresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAppCatalogItem

Get an app catalog item by type.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-app-catalog-item" method="get" path="/v3/openmeter/app-catalog/{appType}" -->
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

    res, err := s.OpenMeterApps.GetAppCatalogItem(ctx, components.BillingAppTypeStripe)
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingAppCatalogItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `appType`                                                              | [components.BillingAppType](../../models/components/billingapptype.md) | :heavy_check_mark:                                                     | The type of the app.                                                   |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.GetAppCatalogItemResponse](../../models/operations/getappcatalogitemresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListApps

List installed apps.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-apps" method="get" path="/v3/openmeter/apps" -->
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

    res, err := s.OpenMeterApps.ListApps(ctx, operations.ListAppsRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListAppsParamsFilter{
            ID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListAppsRequest](../../models/operations/listappsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.ListAppsResponse](../../models/operations/listappsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetApp

Get an installed app.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-app" method="get" path="/v3/openmeter/apps/{appId}" -->
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

    res, err := s.OpenMeterApps.GetApp(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingApp != nil {
        switch res.BillingApp.Type {
            case components.BillingAppUnionTypeStripe:
                // res.BillingApp.BillingAppStripe is populated
            case components.BillingAppUnionTypeSandbox:
                // res.BillingApp.BillingAppSandbox is populated
            case components.BillingAppUnionTypeExternalInvoicing:
                // res.BillingApp.BillingAppExternalInvoicing is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `appID`                                                  | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAppResponse](../../models/operations/getappresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UninstallApp

Uninstall an app by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="uninstall-app" method="delete" path="/v3/openmeter/apps/{appId}" -->
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

    res, err := s.OpenMeterApps.UninstallApp(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `appID`                                                  | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.UninstallAppResponse](../../models/operations/uninstallappresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateApp

Update an installed app.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-app" method="put" path="/v3/openmeter/apps/{appId}" -->
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

    res, err := s.OpenMeterApps.UpdateApp(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateBillingUpdateAppRequestExternalInvoicing(
        components.UpdateAppExternalInvoicingRequest{
            Name: "<value>",
            Type: components.UpdateAppExternalInvoicingRequestTypeExternalInvoicing,
            EnableDraftSyncHook: true,
            EnableIssuingSyncHook: false,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingApp != nil {
        switch res.BillingApp.Type {
            case components.BillingAppUnionTypeStripe:
                // res.BillingApp.BillingAppStripe is populated
            case components.BillingAppUnionTypeSandbox:
                // res.BillingApp.BillingAppSandbox is populated
            case components.BillingAppUnionTypeExternalInvoicing:
                // res.BillingApp.BillingAppExternalInvoicing is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `appID`                                                                                  | `string`                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                                                               |
| `billingUpdateAppRequest`                                                                | [components.BillingUpdateAppRequest](../../models/components/billingupdateapprequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.UpdateAppResponse](../../models/operations/updateappresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |