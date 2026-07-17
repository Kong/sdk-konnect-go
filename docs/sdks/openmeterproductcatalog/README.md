# OpenMeterProductCatalog

## Overview

Product catalog manages plans, add-ons, and their associations for subscription-based billing.

### Available Operations

* [ListProductCatalogAddons](#listproductcatalogaddons) - List add-ons
* [CreateProductCatalogAddon](#createproductcatalogaddon) - Create add-on
* [UpdateProductCatalogAddon](#updateproductcatalogaddon) - Update add-on
* [GetProductCatalogAddon](#getproductcatalogaddon) - Get add-on
* [DeleteProductCatalogAddon](#deleteproductcatalogaddon) - Soft delete add-on
* [ArchiveProductCatalogAddon](#archiveproductcatalogaddon) - Archive add-on version
* [PublishProductCatalogAddon](#publishproductcatalogaddon) - Publish add-on version
* [ListPlans](#listplans) - List plans
* [CreatePlan](#createplan) - Create plan
* [UpdatePlan](#updateplan) - Update plan
* [GetPlan](#getplan) - Get plan
* [DeletePlan](#deleteplan) - Delete plan
* [ListPlanAddons](#listplanaddons) - List add-ons for plan
* [CreatePlanAddon](#createplanaddon) - Add add-on to plan
* [GetPlanAddon](#getplanaddon) - Get add-on association for plan
* [UpdatePlanAddon](#updateplanaddon) - Update add-on association for plan
* [DeletePlanAddon](#deleteplanaddon) - Remove add-on from plan
* [ArchivePlan](#archiveplan) - Archive plan version
* [PublishPlan](#publishplan) - Publish plan version

## ListProductCatalogAddons

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List all add-ons.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-product-catalog-addons" method="get" path="/v3/openmeter/addons" -->
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

    res, err := s.OpenMeterProductCatalog.ListProductCatalogAddons(ctx, operations.ListProductCatalogAddonsRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListAddonsParamsFilter{
            ID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddonPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListProductCatalogAddonsRequest](../../models/operations/listproductcatalogaddonsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListProductCatalogAddonsResponse](../../models/operations/listproductcatalogaddonsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateProductCatalogAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a new add-on.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-product-catalog-addon" method="post" path="/v3/openmeter/addons" -->
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

    res, err := s.OpenMeterProductCatalog.CreateProductCatalogAddon(ctx, components.CreateAddonRequest1{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Key: "resource_key",
        InstanceType: components.CreateAddonRequestTheInstanceTypeOfTheAddOnsCanBeSingleOrMultipleMultiple,
        Currency: "USD",
        RateCards: []components.BillingRateCard{
            components.BillingRateCard{
                Name: "<value>",
                Labels: map[string]string{
                    "env": "test",
                },
                Key: "resource_key",
                Feature: &components.FeatureReference{
                    ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                },
                BillingCadence: sdkkonnectgo.Pointer("P1Y"),
                Price: components.CreatePriceFlat(
                    components.BillingPriceFlat{
                        Type: components.BillingPriceFlatTypeFlat,
                        Amount: "902.15",
                    },
                ),
                PaymentTerm: components.PaymentTermInArrears.ToPointer(),
                TaxConfig: &components.TaxConfig{
                    Code: components.TaxCodeReference{
                        ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                    },
                },
                Entitlement: sdkkonnectgo.Pointer(components.CreateEntitlementTemplateMetered(
                    components.BillingRateCardMeteredEntitlement{
                        Type: components.BillingRateCardMeteredEntitlementTypeMetered,
                        IsSoftLimit: sdkkonnectgo.Pointer(false),
                    },
                )),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Addon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.CreateAddonRequest1](../../models/components/createaddonrequest1.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateProductCatalogAddonResponse](../../models/operations/createproductcatalogaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateProductCatalogAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update an add-on by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-product-catalog-addon" method="put" path="/v3/openmeter/addons/{addonId}" -->
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

    res, err := s.OpenMeterProductCatalog.UpdateProductCatalogAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpsertAddonRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        InstanceType: components.UpsertAddonRequestTheInstanceTypeOfTheAddOnsCanBeSingleOrMultipleMultiple,
        RateCards: []components.BillingRateCard{
            components.BillingRateCard{
                Name: "<value>",
                Labels: map[string]string{
                    "env": "test",
                },
                Key: "resource_key",
                Feature: &components.FeatureReference{
                    ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                },
                BillingCadence: sdkkonnectgo.Pointer("P1Y"),
                Price: components.CreatePriceVolume(
                    components.BillingPriceVolume{
                        Type: components.BillingPriceVolumeTypeVolume,
                        Tiers: []components.BillingPriceTier{
                            components.BillingPriceTier{},
                        },
                    },
                ),
                PaymentTerm: components.PaymentTermInArrears.ToPointer(),
                TaxConfig: &components.TaxConfig{
                    Code: components.TaxCodeReference{
                        ID: "01G65Z755AFWAKHE12NY0CQ9FH",
                    },
                },
                Entitlement: sdkkonnectgo.Pointer(components.CreateEntitlementTemplateStatic(
                    components.BillingRateCardStaticEntitlement{
                        Type: components.BillingRateCardStaticEntitlementTypeStatic,
                        Config: "<value>",
                    },
                )),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Addon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `addonID`                                                                      | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | 01G65Z755AFWAKHE12NY0CQ9FH                                                     |
| `upsertAddonRequest`                                                           | [components.UpsertAddonRequest](../../models/components/upsertaddonrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.UpdateProductCatalogAddonResponse](../../models/operations/updateproductcatalogaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetProductCatalogAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get add-on by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-product-catalog-addon" method="get" path="/v3/openmeter/addons/{addonId}" -->
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

    res, err := s.OpenMeterProductCatalog.GetProductCatalogAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.Addon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `addonID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetProductCatalogAddonResponse](../../models/operations/getproductcatalogaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteProductCatalogAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Soft delete add-on by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-product-catalog-addon" method="delete" path="/v3/openmeter/addons/{addonId}" -->
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

    res, err := s.OpenMeterProductCatalog.DeleteProductCatalogAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `addonID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteProductCatalogAddonResponse](../../models/operations/deleteproductcatalogaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ArchiveProductCatalogAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Archive an add-on version.

### Example Usage

<!-- UsageSnippet language="go" operationID="archive-product-catalog-addon" method="post" path="/v3/openmeter/addons/{addonId}/archive" -->
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

    res, err := s.OpenMeterProductCatalog.ArchiveProductCatalogAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.Addon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `addonID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ArchiveProductCatalogAddonResponse](../../models/operations/archiveproductcatalogaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PublishProductCatalogAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Publish an add-on version.

### Example Usage

<!-- UsageSnippet language="go" operationID="publish-product-catalog-addon" method="post" path="/v3/openmeter/addons/{addonId}/publish" -->
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

    res, err := s.OpenMeterProductCatalog.PublishProductCatalogAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.Addon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `addonID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PublishProductCatalogAddonResponse](../../models/operations/publishproductcatalogaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPlans

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List all plans.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plans" method="get" path="/v3/openmeter/plans" -->
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

    res, err := s.OpenMeterProductCatalog.ListPlans(ctx, operations.ListPlansRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListPlansRequest](../../models/operations/listplansrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.ListPlansResponse](../../models/operations/listplansresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePlan

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a new plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plan" method="post" path="/v3/openmeter/plans" -->
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

    res, err := s.OpenMeterProductCatalog.CreatePlan(ctx, components.CreatePlanRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Key: "resource_key",
        Currency: "USD",
        BillingCadence: "P1Y",
        Phases: []components.BillingPlanPhase{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingPlan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.CreatePlanRequest](../../models/components/createplanrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreatePlanResponse](../../models/operations/createplanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePlan

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update a plan by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-plan" method="put" path="/v3/openmeter/plans/{planId}" -->
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

    res, err := s.OpenMeterProductCatalog.UpdatePlan(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpsertPlanRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Phases: []components.BillingPlanPhase{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingPlan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `planID`                                                                     | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          | 01G65Z755AFWAKHE12NY0CQ9FH                                                   |
| `upsertPlanRequest`                                                          | [components.UpsertPlanRequest](../../models/components/upsertplanrequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.UpdatePlanResponse](../../models/operations/updateplanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPlan

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get a plan by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plan" method="get" path="/v3/openmeter/plans/{planId}" -->
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

    res, err := s.OpenMeterProductCatalog.GetPlan(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingPlan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `planID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPlanResponse](../../models/operations/getplanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.GoneError         | 410                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePlan

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Delete a plan by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plan" method="delete" path="/v3/openmeter/plans/{planId}" -->
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

    res, err := s.OpenMeterProductCatalog.DeletePlan(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `planID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePlanResponse](../../models/operations/deleteplanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPlanAddons

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List add-ons associated with a plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-plan-addons" method="get" path="/v3/openmeter/plans/{planId}/addons" -->
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

    res, err := s.OpenMeterProductCatalog.ListPlanAddons(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanAddonPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |                                                                                   |
| `planID`                                                                          | `string`                                                                          | :heavy_check_mark:                                                                | N/A                                                                               | 01G65Z755AFWAKHE12NY0CQ9FH                                                        |
| `page`                                                                            | [*components.PagePaginationQuery](../../models/components/pagepaginationquery.md) | :heavy_minus_sign:                                                                | Determines which page of the collection to retrieve.                              |                                                                                   |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |                                                                                   |

### Response

**[*operations.ListPlanAddonsResponse](../../models/operations/listplanaddonsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePlanAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Add an add-on to a plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-plan-addon" method="post" path="/v3/openmeter/plans/{planId}/addons" -->
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

    res, err := s.OpenMeterProductCatalog.CreatePlanAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreatePlanAddonRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Addon: components.CreatePlanAddonRequestAddOn{
            ID: "01G65Z755AFWAKHE12NY0CQ9FH",
        },
        FromPlanPhase: "resource_key",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanAddon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `planID`                                                                               | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | 01G65Z755AFWAKHE12NY0CQ9FH                                                             |
| `createPlanAddonRequest`                                                               | [components.CreatePlanAddonRequest](../../models/components/createplanaddonrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.CreatePlanAddonResponse](../../models/operations/createplanaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPlanAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get an add-on association for a plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-plan-addon" method="get" path="/v3/openmeter/plans/{planId}/addons/{planAddonId}" -->
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

    res, err := s.OpenMeterProductCatalog.GetPlanAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanAddon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `planID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `planAddonID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPlanAddonResponse](../../models/operations/getplanaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePlanAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update an add-on association for a plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-plan-addon" method="put" path="/v3/openmeter/plans/{planId}/addons/{planAddonId}" -->
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

    res, err := s.OpenMeterProductCatalog.UpdatePlanAddon(ctx, operations.UpdatePlanAddonRequest{
        PlanID: "01G65Z755AFWAKHE12NY0CQ9FH",
        PlanAddonID: "01G65Z755AFWAKHE12NY0CQ9FH",
        UpsertPlanAddonRequest: components.UpsertPlanAddonRequest{
            Name: "<value>",
            Labels: map[string]string{
                "env": "test",
            },
            FromPlanPhase: "resource_key",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanAddon != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdatePlanAddonRequest](../../models/operations/updateplanaddonrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdatePlanAddonResponse](../../models/operations/updateplanaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePlanAddon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Remove an add-on from a plan.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-plan-addon" method="delete" path="/v3/openmeter/plans/{planId}/addons/{planAddonId}" -->
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

    res, err := s.OpenMeterProductCatalog.DeletePlanAddon(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `planID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `planAddonID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePlanAddonResponse](../../models/operations/deleteplanaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ArchivePlan

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Archive a plan version.

### Example Usage

<!-- UsageSnippet language="go" operationID="archive-plan" method="post" path="/v3/openmeter/plans/{planId}/archive" -->
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

    res, err := s.OpenMeterProductCatalog.ArchivePlan(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingPlan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `planID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ArchivePlanResponse](../../models/operations/archiveplanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PublishPlan

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Publish a plan version.

### Example Usage

<!-- UsageSnippet language="go" operationID="publish-plan" method="post" path="/v3/openmeter/plans/{planId}/publish" -->
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

    res, err := s.OpenMeterProductCatalog.PublishPlan(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingPlan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `planID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PublishPlanResponse](../../models/operations/publishplanresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |