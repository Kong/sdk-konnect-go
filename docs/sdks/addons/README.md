# AddOns

## Overview

Optional services that extend the capabilities of a Cloud Gateway data plane group. The currently supported add-on type is managed cache (`managed-cache.v0`), which provisions a Redis-compatible in-memory cache co-located with your data planes. Each add-on is scoped to a control plane or control plane group, and is automatically deployed across all data plane groups that belong to that owner.


### Available Operations

* [CreateAddOn](#createaddon) - Create Add-On
* [GetAddOn](#getaddon) - Get Add-On
* [DeleteAddOn](#deleteaddon) - Delete Add-On
* [UpdateAddOn](#updateaddon) - Update Add-On

## CreateAddOn

Creates a new add-on for a control plane or control plane group. The add-on type is
determined by the `config.kind` field — currently only `managed-cache.v0` is supported,
which provisions a Redis-compatible cache co-located with your data planes. After it's created,
the add-on transitions through `initializing → ready` as it deploys across data plane groups.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-add-on" method="post" path="/v2/cloud-gateways/add-ons" -->
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

    res, err := s.AddOns.CreateAddOn(ctx, components.CreateAddOnRequest{
        Name: "my-add-on",
        Owner: components.CreateAddOnOwnerControlPlaneGroupAddOnOwner(
            components.ControlPlaneGroupAddOnOwner{
                ControlPlaneGroupID: "123e4567-e89b-12d3-a456-426614174000",
                ControlPlaneGroupGeo: components.ControlPlaneGeoSg,
            },
        ),
        Config: components.CreateCreateAddOnConfigManagedCache(
            components.ManagedCache{
                CapacityConfig: components.CreateManagedCacheCapacityConfigTiered(
                    components.Tiered{
                        Tier: components.TierSmall,
                    },
                ),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddOnResponse != nil {
        switch res.AddOnResponse.Owner.Type {
            case components.AddOnOwnerTypeControlPlaneAddOnOwner:
                // res.AddOnResponse.Owner.ControlPlaneAddOnOwner is populated
            case components.AddOnOwnerTypeControlPlaneGroupAddOnOwner:
                // res.AddOnResponse.Owner.ControlPlaneGroupAddOnOwner is populated
        }

    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.CreateAddOnRequest](../../models/components/createaddonrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateAddOnResponse](../../models/operations/createaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAddOn

Retrieves a single add-on by ID, including its current lifecycle state and per data plane group deployment status.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-add-on" method="get" path="/v2/cloud-gateways/add-ons/{addOnId}" -->
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

    res, err := s.AddOns.GetAddOn(ctx, "550e8400-e29b-41d4-a716-446655440000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AddOnResponse != nil {
        switch res.AddOnResponse.Owner.Type {
            case components.AddOnOwnerTypeControlPlaneAddOnOwner:
                // res.AddOnResponse.Owner.ControlPlaneAddOnOwner is populated
            case components.AddOnOwnerTypeControlPlaneGroupAddOnOwner:
                // res.AddOnResponse.Owner.ControlPlaneGroupAddOnOwner is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `addOnID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the add-on to operate on.                          | 550e8400-e29b-41d4-a716-446655440000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAddOnResponse](../../models/operations/getaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAddOn

Deletes an add-on by ID. The request is rejected if any Kong plugins are still referencing
the managed cache add-on — remove those plugin references before deleting.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-add-on" method="delete" path="/v2/cloud-gateways/add-ons/{addOnId}" -->
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

    res, err := s.AddOns.DeleteAddOn(ctx, "550e8400-e29b-41d4-a716-446655440000")
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
| `addOnID`                                                | `string`                                                 | :heavy_check_mark:                                       | ID of the add-on to operate on.                          | 550e8400-e29b-41d4-a716-446655440000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAddOnResponse](../../models/operations/deleteaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAddOn

Updates the configuration of an existing add-on, such as changing the managed cache
capacity tier. Tier upgrades are supported; downgrades are not.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-add-on" method="patch" path="/v2/cloud-gateways/add-ons/{addOnId}" -->
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

    res, err := s.AddOns.UpdateAddOn(ctx, "550e8400-e29b-41d4-a716-446655440000", components.UpdateAddOnRequest{
        Config: components.CreateUpdateAddOnConfigManagedCache(
            components.ManagedCache{
                CapacityConfig: components.CreateManagedCacheCapacityConfigTiered(
                    components.Tiered{
                        Tier: components.TierSmall,
                    },
                ),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddOnResponse != nil {
        switch res.AddOnResponse.Owner.Type {
            case components.AddOnOwnerTypeControlPlaneAddOnOwner:
                // res.AddOnResponse.Owner.ControlPlaneAddOnOwner is populated
            case components.AddOnOwnerTypeControlPlaneGroupAddOnOwner:
                // res.AddOnResponse.Owner.ControlPlaneGroupAddOnOwner is populated
        }

    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `addOnID`                                                                      | `string`                                                                       | :heavy_check_mark:                                                             | ID of the add-on to operate on.                                                | 550e8400-e29b-41d4-a716-446655440000                                           |
| `updateAddOnRequest`                                                           | [components.UpdateAddOnRequest](../../models/components/updateaddonrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.UpdateAddOnResponse](../../models/operations/updateaddonresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |