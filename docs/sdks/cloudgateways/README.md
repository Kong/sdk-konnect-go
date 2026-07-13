# CloudGateways

## Overview

### Available Operations

* [CreateAddOn](#createaddon) - Create Add-On
* [ListAddOns](#listaddons) - List Add-Ons
* [GetAddOn](#getaddon) - Get Add-On
* [DeleteAddOn](#deleteaddon) - Delete Add-On
* [UpdateAddOn](#updateaddon) - Update Add-On
* [GetAvailabilityJSON](#getavailabilityjson) - Get Resource Availability JSON
* [ListConfigurations](#listconfigurations) - List Configurations
* [CreateConfiguration](#createconfiguration) - Create Configuration
* [GetConfiguration](#getconfiguration) - Get Configuration
* [ListCustomDomains](#listcustomdomains) - List Custom Domains
* [CreateCustomDomains](#createcustomdomains) - Create Custom Domain
* [GetCustomDomain](#getcustomdomain) - Get Custom Domain
* [DeleteCustomDomain](#deletecustomdomain) - Delete Custom Domain
* [GetCustomDomainOnlineStatus](#getcustomdomainonlinestatus) - Get Custom Domain Online Status
* [ListDefaultResourceConfigurations](#listdefaultresourceconfigurations) - List Default Resource Configurations
* [ListDefaultResourceQuotas](#listdefaultresourcequotas) - List Default Resource Quotas
* [ListNetworks](#listnetworks) - List Networks
* [CreateNetwork](#createnetwork) - Create Network
* [GetNetwork](#getnetwork) - Get Network
* [UpdateNetwork](#updatenetwork) - Update Network
* [DeleteNetwork](#deletenetwork) - Delete Network
* [ListNetworkConfigurations](#listnetworkconfigurations) - List Network Configuration References
* [ListPrivateDNS](#listprivatedns) - List Private DNS
* [CreatePrivateDNS](#createprivatedns) - Create Private DNS
* [GetPrivateDNS](#getprivatedns) - Get Private DNS
* [UpdatePrivateDNS](#updateprivatedns) - Update Private DNS
* [DeletePrivateDNS](#deleteprivatedns) - Delete Private DNS
* [ListTransitGateways](#listtransitgateways) - List Transit Gateways
* [CreateTransitGateway](#createtransitgateway) - Create Transit Gateway
* [GetTransitGateway](#gettransitgateway) - Get Transit Gateway
* [UpdateTransitGateway](#updatetransitgateway) - Update Transit Gateway
* [DeleteTransitGateway](#deletetransitgateway) - Delete Transit Gateway
* [ListProviderAccounts](#listprovideraccounts) - List Provider Accounts
* [GetProviderAccount](#getprovideraccount) - Get Provider Account
* [ListResourceConfigurations](#listresourceconfigurations) - List Resource Configurations
* [GetResourceConfiguration](#getresourceconfiguration) - Get Resource Configuration
* [ListResourceQuotas](#listresourcequotas) - List Resource Quotas
* [GetResourceQuota](#getresourcequota) - Get Resource Quota

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

    res, err := s.CloudGateways.CreateAddOn(ctx, components.CreateAddOnRequest{
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

## ListAddOns

Returns a paginated list of add-ons for the organization. Use filter parameters to narrow
results by owner (control plane or control plane group), lifecycle state, or config kind.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-add-ons" method="get" path="/v2/cloud-gateways/add-ons" -->
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

    res, err := s.CloudGateways.ListAddOns(ctx, operations.ListAddOnsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAddOnsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListAddOnsRequest](../../models/operations/listaddonsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListAddOnsResponse](../../models/operations/listaddonsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
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

    res, err := s.CloudGateways.GetAddOn(ctx, "550e8400-e29b-41d4-a716-446655440000")
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

    res, err := s.CloudGateways.DeleteAddOn(ctx, "550e8400-e29b-41d4-a716-446655440000")
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

    res, err := s.CloudGateways.UpdateAddOn(ctx, "550e8400-e29b-41d4-a716-446655440000", components.UpdateAddOnRequest{
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

## GetAvailabilityJSON

Returns a public JSON document describing the supported cloud providers, regions, gateway versions, and instance types available for Kong Cloud Gateways. Authentication isn't required — query this endpoint to discover valid inputs for network and configuration requests.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-availability-json" method="get" path="/v2/cloud-gateways/availability.json" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New()

    res, err := s.CloudGateways.GetAvailabilityJSON(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AvailabilityDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAvailabilityJSONResponse](../../models/operations/getavailabilityjsonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListConfigurations

Returns a paginated collection of Cloud Gateway configurations visible to the caller.
Results are scoped to control planes the caller has permission to read.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-configurations" method="get" path="/v2/cloud-gateways/configurations" -->
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

    res, err := s.CloudGateways.ListConfigurations(ctx, operations.ListConfigurationsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListConfigurationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListConfigurationsRequest](../../models/operations/listconfigurationsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListConfigurationsResponse](../../models/operations/listconfigurationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateConfiguration

Creates or replaces the Cloud Gateway configuration for a control plane and geo. The request fully describes
the desired state — Kong diffs it against the current configuration, then adds, removes, or updates data plane
groups to match. Any network referenced in the request that is currently `offline` automatically transitions
to `initializing`.

Use `kind: dedicated.v0` (default) for dedicated Cloud Gateways — `version`, `cloud_gateway_network_id`, and
`autoscale` are required. Use `kind: serverless.v1` for serverless Cloud Gateways — those three fields must
be omitted.


### Example Usage: Configuration Api Access Enum Validation

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Configuration Api Access Enum Validation" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Conflict

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Configuration Conflict" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Autopilot Base RPS Minimum

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Configuration Data-Plane Group Autopilot Base RPS Minimum" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Static Request Instances Minimum

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Configuration Data-Plane Group Static Request Instances Minimum" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Custom Domain Domain Length Too Short

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Custom Domain Domain Length Too Short" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Network Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Network Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="NotFoundExample" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Permission Denied

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Permission Denied" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Quota Exceeded

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Quota Exceeded" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Transit Gateway Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Transit Gateway Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="Unauthorized" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" example="UnauthorizedExample" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: sdkkonnectgo.Pointer("3.2"),
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: sdkkonnectgo.Pointer("36ae63d3-efd1-4bec-b246-62aa5d3f5695"),
                Autoscale: sdkkonnectgo.Pointer(components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                    },
                )),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "info",
                    },
                },
            },
        },
        APIAccess: components.APIAccessPrivatePlusPublic.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.CreateConfigurationRequest](../../models/components/createconfigurationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateConfigurationResponse](../../models/operations/createconfigurationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetConfiguration

Retrieves a single Cloud Gateway configuration by ID, including the current state of each deployed
data plane group. Access is restricted to control planes the caller has permission to read.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-configuration" method="get" path="/v2/cloud-gateways/configurations/{configurationId}" -->
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

    res, err := s.CloudGateways.GetConfiguration(ctx, "edaf40f9-9fb0-4ffe-bb74-4e763a6bd471")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationManifest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `configurationID`                                        | `string`                                                 | :heavy_check_mark:                                       | The ID of the configuration to operate on.               | edaf40f9-9fb0-4ffe-bb74-4e763a6bd471                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetConfigurationResponse](../../models/operations/getconfigurationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCustomDomains

Returns a paginated list of custom domains across all control planes in the organization,
scoped to control planes you have read access to.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-custom-domains" method="get" path="/v2/cloud-gateways/custom-domains" -->
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

    res, err := s.CloudGateways.ListCustomDomains(ctx, operations.ListCustomDomainsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCustomDomainsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListCustomDomainsRequest](../../models/operations/listcustomdomainsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListCustomDomainsResponse](../../models/operations/listcustomdomainsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateCustomDomains

Registers a custom domain for a control plane. After creation, Konnect provisions a TLS
certificate and configures SNI routing, transitioning the domain through
`initializing → ready`. To complete setup, configure two CNAME records at your DNS
registrar: one pointing your domain to the Konnect gateway hostname, and one pointing
`_acme-challenge.<your-domain>` to the ACME challenge hostname provided by Konnect.
Use the online-status endpoint to verify both records are correctly configured.


### Example Usage: Configuration Api Access Enum Validation

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="Configuration Api Access Enum Validation" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Autopilot Base RPS Minimum

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="Configuration Data-Plane Group Autopilot Base RPS Minimum" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Static Request Instances Minimum

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="Configuration Data-Plane Group Static Request Instances Minimum" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: Custom Domain Domain Length Too Short

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="Custom Domain Domain Length Too Short" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: Network Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="Network Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="NotFoundExample" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: Transit Gateway Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="Transit Gateway Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="Unauthorized" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" example="UnauthorizedExample" -->
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

    res, err := s.CloudGateways.CreateCustomDomains(ctx, components.CreateCustomDomainRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoUs,
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateCustomDomainRequest](../../models/components/createcustomdomainrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateCustomDomainsResponse](../../models/operations/createcustomdomainsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomDomain

Retrieves a single custom domain by ID, including its current lifecycle state and any error metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-custom-domain" method="get" path="/v2/cloud-gateways/custom-domains/{customDomainId}" -->
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

    res, err := s.CloudGateways.GetCustomDomain(ctx, "39ed3790-085d-4605-9627-f96d86aaf425")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customDomainID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of the custom domain to operate on.                   | 39ed3790-085d-4605-9627-f96d86aaf425                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCustomDomainResponse](../../models/operations/getcustomdomainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCustomDomain

Deletes a custom domain by ID, removing the associated TLS certificate and SNI configuration from the control plane's data planes.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-custom-domain" method="delete" path="/v2/cloud-gateways/custom-domains/{customDomainId}" -->
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

    res, err := s.CloudGateways.DeleteCustomDomain(ctx, "39ed3790-085d-4605-9627-f96d86aaf425")
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
| `customDomainID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of the custom domain to operate on.                   | 39ed3790-085d-4605-9627-f96d86aaf425                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCustomDomainResponse](../../models/operations/deletecustomdomainresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetCustomDomainOnlineStatus

Checks whether the primary domain CNAME and ACME challenge CNAME records are correctly configured at your DNS registrar. Returns `verified` or `unverified` for each.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-custom-domain-online-status" method="get" path="/v2/cloud-gateways/custom-domains/{customDomainId}/online-status" -->
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

    res, err := s.CloudGateways.GetCustomDomainOnlineStatus(ctx, "39ed3790-085d-4605-9627-f96d86aaf425")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomDomainOnlineStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customDomainID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of the custom domain to operate on.                   | 39ed3790-085d-4605-9627-f96d86aaf425                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCustomDomainOnlineStatusResponse](../../models/operations/getcustomdomainonlinestatusresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListDefaultResourceConfigurations

Returns the platform-default resource configurations for Cloud Gateways, along with any
organization-level overrides. Resource configurations are behavioral settings applied to
Cloud Gateway resources — for example, the idle timeout for data plane groups.
Use this to view the effective settings for your organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-default-resource-configurations" method="get" path="/v2/cloud-gateways/default-resource-configurations" -->
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

    res, err := s.CloudGateways.ListDefaultResourceConfigurations(ctx, sdkkonnectgo.Pointer[int64](10), sdkkonnectgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDefaultResourceConfigurationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `pageSize`                                                                                              | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.ListDefaultResourceConfigurationsResponse](../../models/operations/listdefaultresourceconfigurationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListDefaultResourceQuotas

Returns the platform-default resource quotas for Cloud Gateways, along with any
organization-level overrides. Use this to view the effective limits for your organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-default-resource-quotas" method="get" path="/v2/cloud-gateways/default-resource-quotas" -->
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

    res, err := s.CloudGateways.ListDefaultResourceQuotas(ctx, sdkkonnectgo.Pointer[int64](10), sdkkonnectgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDefaultResourceQuotasResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `pageSize`                                                                                              | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.ListDefaultResourceQuotasResponse](../../models/operations/listdefaultresourcequotasresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListNetworks

Returns a paginated list of Cloud Gateway networks visible to the caller. Filter by
`state` to narrow results — for example, poll for `state=ready` to find networks
available for data plane group configurations.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-networks" method="get" path="/v2/cloud-gateways/networks" -->
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

    res, err := s.CloudGateways.ListNetworks(ctx, operations.ListNetworksRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListNetworksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListNetworksRequest](../../models/operations/listnetworksrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListNetworksResponse](../../models/operations/listnetworksresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateNetwork

Creates a new Cloud Gateway network in the specified provider account and region. Network creation is
asynchronous — the network starts in `initializing` state and transitions to `ready` once provisioned.


### Example Usage: Configuration Api Access Enum Validation

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Configuration Api Access Enum Validation" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Autopilot Base RPS Minimum

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Configuration Data-Plane Group Autopilot Base RPS Minimum" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Static Request Instances Minimum

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Configuration Data-Plane Group Static Request Instances Minimum" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Custom Domain Domain Length Too Short

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Custom Domain Domain Length Too Short" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Network Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Network Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Permission Denied

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Permission Denied" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Quota Exceeded

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Quota Exceeded" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Transit Gateway Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Transit Gateway Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="Unauthorized" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" example="UnauthorizedExample" -->
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

    res, err := s.CloudGateways.CreateNetwork(ctx, components.CreateNetworkRequest{
        Name: "us-east-2 network",
        CloudGatewayProviderAccountID: "929b2449-c69f-44c4-b6ad-9ecec6f811ae",
        Region: "us-east-2",
        AvailabilityZones: []string{
            "use2-az1",
            "use2-az2",
            "use2-az3",
        },
        CidrBlock: "10.0.0.0/8",
        State: components.NetworkCreateStateInitializing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.CreateNetworkRequest](../../models/components/createnetworkrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateNetworkResponse](../../models/operations/createnetworkresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetNetwork

Retrieves a Cloud Gateway network by ID, including its current state and provider metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-network" method="get" path="/v2/cloud-gateways/networks/{networkId}" -->
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

    res, err := s.CloudGateways.GetNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695")
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `networkID`                                              | `string`                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetNetworkResponse](../../models/operations/getnetworkresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateNetwork

Updates a Cloud Gateway network by ID. You can also rename the network.

### Example Usage: Configuration Api Access Enum Validation

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Configuration Api Access Enum Validation" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Autopilot Base RPS Minimum

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Configuration Data-Plane Group Autopilot Base RPS Minimum" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Configuration Data-Plane Group Static Request Instances Minimum

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Configuration Data-Plane Group Static Request Instances Minimum" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Custom Domain Domain Length Too Short

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Custom Domain Domain Length Too Short" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Network Name Length Exceeded

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Network Name Length Exceeded" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="NotFoundExample" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Permission Denied

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Permission Denied" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Quota Exceeded

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Quota Exceeded" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Transit Gateway Name Length Exceeded

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Transit Gateway Name Length Exceeded" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="Unauthorized" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" example="UnauthorizedExample" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 network"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `networkID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | The network to operate on.                                                       | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                                             |
| `patchNetworkRequest`                                                            | [components.PatchNetworkRequest](../../models/components/patchnetworkrequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.UpdateNetworkResponse](../../models/operations/updatenetworkresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteNetwork

Deletes a Cloud Gateway network by ID. The network cannot be referenced by any active configuration before it can be deleted.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-network" method="delete" path="/v2/cloud-gateways/networks/{networkId}" -->
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

    res, err := s.CloudGateways.DeleteNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695")
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
| `networkID`                                              | `string`                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteNetworkResponse](../../models/operations/deletenetworkresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListNetworkConfigurations

Returns a paginated list of data plane group configurations that reference a given network.
Use this to determine which control planes are using a network before renaming or deleting it.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-network-configurations" method="get" path="/v2/cloud-gateways/networks/{networkId}/configuration-references" -->
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

    res, err := s.CloudGateways.ListNetworkConfigurations(ctx, operations.ListNetworkConfigurationsRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListNetworkConfigurationReferencesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListNetworkConfigurationsRequest](../../models/operations/listnetworkconfigurationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListNetworkConfigurationsResponse](../../models/operations/listnetworkconfigurationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPrivateDNS

Returns a paginated collection of private DNS attachments for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-private-dns" method="get" path="/v2/cloud-gateways/networks/{networkId}/private-dns" -->
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

    res, err := s.CloudGateways.ListPrivateDNS(ctx, operations.ListPrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPrivateDNSResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListPrivateDNSRequest](../../models/operations/listprivatednsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListPrivateDNSResponse](../../models/operations/listprivatednsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePrivateDNS

Creates a new private DNS attachment for a given network. The attachment type is determined by
`private_dns_attachment_config.kind`. Supported types: `aws-private-hosted-zone-attachment`,
`aws-outbound-resolver`, `gcp-private-hosted-zone-attachment`, `azure-private-hosted-zone-attachment`,
and `azure-outbound-resolver`.


### Example Usage: Configuration Api Access Enum Validation

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="Configuration Api Access Enum Validation" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Configuration Data-Plane Group Autopilot Base RPS Minimum

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="Configuration Data-Plane Group Autopilot Base RPS Minimum" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Configuration Data-Plane Group Static Request Instances Minimum

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="Configuration Data-Plane Group Static Request Instances Minimum" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Custom Domain Domain Length Too Short

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="Custom Domain Domain Length Too Short" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Network Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="Network Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="NotFoundExample" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Transit Gateway Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="Transit Gateway Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="Unauthorized" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" example="UnauthorizedExample" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
        PrivateDNSAttachmentConfig: sdkkonnectgo.Pointer(components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(
            components.AwsPrivateHostedZoneAttachmentConfig{
                Kind: components.AWSPrivateHostedZoneTypeAwsPrivateHostedZoneAttachment,
                HostedZoneID: "<id>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `networkID`                                                                              | `string`                                                                                 | :heavy_check_mark:                                                                       | The network to operate on.                                                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                                                     |
| `createPrivateDNSRequest`                                                                | [components.CreatePrivateDNSRequest](../../models/components/createprivatednsrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreatePrivateDNSResponse](../../models/operations/createprivatednsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPrivateDNS

Retrieves a private DNS attachment by ID, including its current state and attachment configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-private-dns" method="get" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" -->
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

    res, err := s.CloudGateways.GetPrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "1850820b-c69f-4a2a-b9be-bbcdbc5cd618")
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `networkID`                                              | `string`                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `privateDNSID`                                           | `string`                                                 | :heavy_check_mark:                                       | The ID of the Private DNS to operate on.                 | 1850820b-c69f-4a2a-b9be-bbcdbc5cd618                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPrivateDNSResponse](../../models/operations/getprivatednsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePrivateDNS

Updates a private DNS attachment by ID. Supports updating the name or DNS resolver configuration
for AWS Outbound Resolver and Azure Outbound Resolver attachment types.


### Example Usage: Configuration Api Access Enum Validation

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="Configuration Api Access Enum Validation" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Configuration Data-Plane Group Autopilot Base RPS Minimum

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="Configuration Data-Plane Group Autopilot Base RPS Minimum" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Configuration Data-Plane Group Static Request Instances Minimum

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="Configuration Data-Plane Group Static Request Instances Minimum" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Custom Domain Domain Length Too Short

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="Custom Domain Domain Length Too Short" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Network Name Length Exceeded

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="Network Name Length Exceeded" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="NotFoundExample" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Transit Gateway Name Length Exceeded

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="Transit Gateway Name Length Exceeded" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="Unauthorized" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" example="UnauthorizedExample" -->
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

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.Pointer("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfigObject{
                        "global.api.konghq.com": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfigObject{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.8",
                            },
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        switch res.PrivateDNSResponse.Type {
            case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AwsPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AwsPrivateDNSResolverResponse is populated
            case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
                // res.PrivateDNSResponse.GcpPrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
                // res.PrivateDNSResponse.AzurePrivateHostedZoneResponse is populated
            case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
                // res.PrivateDNSResponse.AzurePrivateDNSResolverResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdatePrivateDNSRequest](../../models/operations/updateprivatednsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdatePrivateDNSResponse](../../models/operations/updateprivatednsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePrivateDNS

Deletes a Private DNS attachment by ID. The attachment must be in a stable state (`ready`
or `error`) before deletion — requests against attachments in a transitional state
(`initializing`, `terminating`) will be rejected.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-private-dns" method="delete" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" -->
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

    res, err := s.CloudGateways.DeletePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "1850820b-c69f-4a2a-b9be-bbcdbc5cd618")
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
| `networkID`                                              | `string`                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `privateDNSID`                                           | `string`                                                 | :heavy_check_mark:                                       | The ID of the Private DNS to operate on.                 | 1850820b-c69f-4a2a-b9be-bbcdbc5cd618                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePrivateDNSResponse](../../models/operations/deleteprivatednsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListTransitGateways

Returns a paginated collection of transit gateways attached to a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-transit-gateways" method="get" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" -->
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

    res, err := s.CloudGateways.ListTransitGateways(ctx, operations.ListTransitGatewaysRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTransitGatewaysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListTransitGatewaysRequest](../../models/operations/listtransitgatewaysrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListTransitGatewaysResponse](../../models/operations/listtransitgatewaysresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateTransitGateway

Creates a new transit gateway attachment for a given network. The attachment type is determined by the
`transit_gateway_attachment_config.kind` field. Supported types: `aws-transit-gateway-attachment`,
`aws-vpc-peering-attachment`, `aws-resource-endpoint-attachment`, `azure-vnet-peering-attachment`,
`azure-vhub-peering-attachment`, and `gcp-vpc-peering-attachment`. Creation is asynchronous —
the transit gateway starts in `initializing` state and transitions to `ready` once provisioned.


### Example Usage: Configuration Api Access Enum Validation

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="Configuration Api Access Enum Validation" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSTransitGateway(
        components.AWSTransitGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsTransitGatewayAttachmentConfig{
                Kind: components.AWSTransitGatewayAttachmentConfigAWSTransitGatewayAttachmentTypeAwsTransitGatewayAttachment,
                TransitGatewayID: "<id>",
                RAMShareArn: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: Configuration Data-Plane Group Autopilot Base RPS Minimum

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="Configuration Data-Plane Group Autopilot Base RPS Minimum" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSVpcPeeringGateway(
        components.AWSVpcPeeringGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsVpcPeeringGatewayAttachmentConfig{
                Kind: components.AWSVPCPeeringGatewayAttachmentConfigAWSVPCPeeringAttachmentConfigAwsVpcPeeringAttachment,
                PeerAccountID: "<id>",
                PeerVpcID: "<id>",
                PeerVpcRegion: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: Configuration Data-Plane Group Static Request Instances Minimum

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="Configuration Data-Plane Group Static Request Instances Minimum" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSVpcPeeringGateway(
        components.AWSVpcPeeringGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsVpcPeeringGatewayAttachmentConfig{
                Kind: components.AWSVPCPeeringGatewayAttachmentConfigAWSVPCPeeringAttachmentConfigAwsVpcPeeringAttachment,
                PeerAccountID: "<id>",
                PeerVpcID: "<id>",
                PeerVpcRegion: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: Custom Domain Domain Length Too Short

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="Custom Domain Domain Length Too Short" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSTransitGateway(
        components.AWSTransitGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsTransitGatewayAttachmentConfig{
                Kind: components.AWSTransitGatewayAttachmentConfigAWSTransitGatewayAttachmentTypeAwsTransitGatewayAttachment,
                TransitGatewayID: "<id>",
                RAMShareArn: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: Network Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="Network Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSVpcPeeringGateway(
        components.AWSVpcPeeringGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsVpcPeeringGatewayAttachmentConfig{
                Kind: components.AWSVPCPeeringGatewayAttachmentConfigAWSVPCPeeringAttachmentConfigAwsVpcPeeringAttachment,
                PeerAccountID: "<id>",
                PeerVpcID: "<id>",
                PeerVpcRegion: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="NotFoundExample" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSVpcPeeringGateway(
        components.AWSVpcPeeringGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsVpcPeeringGatewayAttachmentConfig{
                Kind: components.AWSVPCPeeringGatewayAttachmentConfigAWSVPCPeeringAttachmentConfigAwsVpcPeeringAttachment,
                PeerAccountID: "<id>",
                PeerVpcID: "<id>",
                PeerVpcRegion: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: Transit Gateway Name Length Exceeded

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="Transit Gateway Name Length Exceeded" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSVpcPeeringGateway(
        components.AWSVpcPeeringGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsVpcPeeringGatewayAttachmentConfig{
                Kind: components.AWSVPCPeeringGatewayAttachmentConfigAWSVPCPeeringAttachmentConfigAwsVpcPeeringAttachment,
                PeerAccountID: "<id>",
                PeerVpcID: "<id>",
                PeerVpcRegion: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="Unauthorized" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAWSTransitGateway(
        components.AWSTransitGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            CidrBlocks: []string{
                "10.0.0.0/8",
                "100.64.0.0/10",
                "172.16.0.0/12",
            },
            TransitGatewayAttachmentConfig: components.AwsTransitGatewayAttachmentConfig{
                Kind: components.AWSTransitGatewayAttachmentConfigAWSTransitGatewayAttachmentTypeAwsTransitGatewayAttachment,
                TransitGatewayID: "<id>",
                RAMShareArn: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" example="UnauthorizedExample" -->
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

    res, err := s.CloudGateways.CreateTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreateCreateTransitGatewayRequestAzureTransitGateway(
        components.AzureTransitGateway{
            Name: "us-east-2 transit gateway",
            DNSConfig: []components.TransitGatewayDNSConfig{
                components.TransitGatewayDNSConfig{
                    RemoteDNSServerIPAddresses: []string{
                        "10.0.0.2",
                    },
                    DomainProxyList: []string{
                        "foobar.com",
                    },
                },
            },
            TransitGatewayAttachmentConfig: components.AzureVNETPeeringAttachmentConfig{
                Kind: components.AzureVNETPeeringAttachmentTypeAzureVnetPeeringAttachment,
                TenantID: "<id>",
                SubscriptionID: "<id>",
                ResourceGroupName: "<value>",
                VnetName: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `networkID`                                                                                      | `string`                                                                                         | :heavy_check_mark:                                                                               | The network to operate on.                                                                       | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                                                             |
| `createTransitGatewayRequest`                                                                    | [components.CreateTransitGatewayRequest](../../models/components/createtransitgatewayrequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateTransitGatewayResponse](../../models/operations/createtransitgatewayresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetTransitGateway

Retrieves a transit gateway by ID, including its current state and attachment configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-transit-gateway" method="get" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways/{transitGatewayId}" -->
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

    res, err := s.CloudGateways.GetTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "0850820b-d153-4a2a-b9be-7d2204779139")
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        switch res.TransitGatewayResponse.Type {
            case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.TransitGatewayResponse.AwsTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
                // res.TransitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
                // res.TransitGatewayResponse.AzureTransitGatewayResponse is populated
            case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
                // res.TransitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
                // res.TransitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
            case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.TransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `networkID`                                              | `string`                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `transitGatewayID`                                       | `string`                                                 | :heavy_check_mark:                                       | The ID of the transit gateway to operate on.             | 0850820b-d153-4a2a-b9be-7d2204779139                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetTransitGatewayResponse](../../models/operations/gettransitgatewayresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateTransitGateway

Updates a transit gateway by ID. Supports updating CIDR blocks on an AWS Transit Gateway, or updating
the resource endpoint configuration on an AWS Resource Endpoint gateway.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-transit-gateway" method="patch" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways/{transitGatewayId}" -->
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

    res, err := s.CloudGateways.UpdateTransitGateway(ctx, operations.UpdateTransitGatewayRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        TransitGatewayID: "0850820b-d153-4a2a-b9be-7d2204779139",
        PatchTransitGatewayRequest: components.CreatePatchTransitGatewayRequestPatchAWSResourceEndpointGatewayAWSResourceEndpointGateway(
            components.PatchAWSResourceEndpointGatewayAWSResourceEndpointGateway{
                TransitGatewayAttachmentConfig: components.TransitGatewayAttachmentConfig{
                    Kind: components.PatchAWSResourceEndpointGatewayAWSResourceEndpointAttachmentTypeAwsResourceEndpointAttachment,
                    ResourceConfig: []components.AwsResourceEndpointConfig{
                        components.AwsResourceEndpointConfig{
                            ResourceConfigID: "<id>",
                            DomainName: "delectable-molasses.com",
                        },
                    },
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PatchTransitGatewayResponse != nil {
        switch res.PatchTransitGatewayResponse.Type {
            case components.PatchTransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
                // res.PatchTransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
            case components.PatchTransitGatewayResponseTypeAwsTransitGatewayResponse:
                // res.PatchTransitGatewayResponse.AwsTransitGatewayResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateTransitGatewayRequest](../../models/operations/updatetransitgatewayrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateTransitGatewayResponse](../../models/operations/updatetransitgatewayresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteTransitGateway

Deletes a transit gateway by ID. The transit gateway must be in a non-transitional state before deletion.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-transit-gateway" method="delete" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways/{transitGatewayId}" -->
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

    res, err := s.CloudGateways.DeleteTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "0850820b-d153-4a2a-b9be-7d2204779139")
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
| `networkID`                                              | `string`                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `transitGatewayID`                                       | `string`                                                 | :heavy_check_mark:                                       | The ID of the transit gateway to operate on.             | 0850820b-d153-4a2a-b9be-7d2204779139                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteTransitGatewayResponse](../../models/operations/deletetransitgatewayresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListProviderAccounts

Returns a paginated list of provider accounts linked to the organization. Filter by cloud provider to see accounts for a specific CSP.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-provider-accounts" method="get" path="/v2/cloud-gateways/provider-accounts" -->
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

    res, err := s.CloudGateways.ListProviderAccounts(ctx, operations.ListProviderAccountsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListProviderAccountsResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListProviderAccountsRequest](../../models/operations/listprovideraccountsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListProviderAccountsResponse](../../models/operations/listprovideraccountsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetProviderAccount

Retrieves a single provider account by ID, including the associated cloud provider and cloud account ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-provider-account" method="get" path="/v2/cloud-gateways/provider-accounts/{providerAccountId}" -->
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

    res, err := s.CloudGateways.GetProviderAccount(ctx, "929b2449-c69f-44c4-b6ad-9ecec6f811ae")
    if err != nil {
        log.Fatal(err)
    }
    if res.ProviderAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `providerAccountID`                                      | `string`                                                 | :heavy_check_mark:                                       | The ID of the provider account to operate on.            | 929b2449-c69f-44c4-b6ad-9ecec6f811ae                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetProviderAccountResponse](../../models/operations/getprovideraccountresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListResourceConfigurations

Returns organization-specific resource configuration overrides that take precedence over
platform defaults. Each configuration controls a behavioral setting for a specific Cloud
Gateway resource type — for example, `data-plane-group-idle-timeout-minutes` sets how long
a data plane group can remain idle before it scales to zero instances.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-resource-configurations" method="get" path="/v2/cloud-gateways/resource-configurations" -->
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

    res, err := s.CloudGateways.ListResourceConfigurations(ctx, sdkkonnectgo.Pointer[int64](10), sdkkonnectgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceConfigurationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `pageSize`                                                                                              | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.ListResourceConfigurationsResponse](../../models/operations/listresourceconfigurationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetResourceConfiguration

Retrieves a single organization-level resource configuration override by ID. Returns the qualifier, override value, and description for the configuration setting.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-resource-configuration" method="get" path="/v2/cloud-gateways/resource-configurations/{resourceConfigurationId}" -->
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

    res, err := s.CloudGateways.GetResourceConfiguration(ctx, "9678f205-49a1-47bb-82d9-d01cafa42a0d")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResourceConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `resourceConfigurationID`                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the resource configuration to operate on.      | 9678f205-49a1-47bb-82d9-d01cafa42a0d                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetResourceConfigurationResponse](../../models/operations/getresourceconfigurationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListResourceQuotas

Returns organization-specific resource quota overrides that take precedence over platform
defaults. Each quota enforces a count limit on a specific Cloud Gateway resource type —
including active networks, data planes, serverless data planes, data plane groups, and
provider accounts per cloud provider.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-resource-quotas" method="get" path="/v2/cloud-gateways/resource-quotas" -->
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

    res, err := s.CloudGateways.ListResourceQuotas(ctx, sdkkonnectgo.Pointer[int64](10), sdkkonnectgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceQuotasResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `pageSize`                                                                                              | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | `*int64`                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.ListResourceQuotasResponse](../../models/operations/listresourcequotasresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetResourceQuota

Retrieves a single organization-level resource quota override by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-resource-quota" method="get" path="/v2/cloud-gateways/resource-quotas/{resourceQuotaId}" -->
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

    res, err := s.CloudGateways.GetResourceQuota(ctx, "9678f205-49a1-47bb-82d9-d01cafa42a0d")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResourceQuota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `resourceQuotaID`                                        | `string`                                                 | :heavy_check_mark:                                       | The ID of the resource quota to operate on.              | 9678f205-49a1-47bb-82d9-d01cafa42a0d                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetResourceQuotaResponse](../../models/operations/getresourcequotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |