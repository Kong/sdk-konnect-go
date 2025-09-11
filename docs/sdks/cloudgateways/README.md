# CloudGateways
(*CloudGateways*)

## Overview

### Available Operations

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

## GetAvailabilityJSON

Get Cloud Gateways Availability JSON document for describing cloud provider and region availability, pricing,
gateway version availability, and instance type information.


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

Returns a paginated collection of configurations across control-planes for an organization (restricted by
permitted control-plane reads).


### Example Usage

<!-- UsageSnippet language="go" operationID="list-configurations" method="get" path="/v2/cloud-gateways/configurations" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.ListConfigurations(ctx, operations.ListConfigurationsRequest{
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
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

Creates a new configuration for a control-plane (restricted by permitted control-plane permissions for
configurations). This request will replace any existing configuration for the requested control_plane_id and
control_plane_geo by performing a diff. From this diff, new resources detected in the requested configuration
will be added, resources not found in the request configuration but in the previous will be deleted, and
resources found in both will be updated to the requested configuration. Networks referenced in this request that
are in an offline state will automatically initialize (i.e. move to an initializing state).


### Example Usage

<!-- UsageSnippet language="go" operationID="create-configuration" method="put" path="/v2/cloud-gateways/configurations" -->
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

    res, err := s.CloudGateways.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoIn,
        Version: "3.2",
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
                Autoscale: components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                        MaxRps: sdkkonnectgo.Int64(1000),
                    },
                ),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "INFO",
                    },
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "INFO",
                    },
                },
            },
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
                Autoscale: components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(
                    components.ConfigurationDataPlaneGroupAutoscaleAutopilot{
                        Kind: components.ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot,
                        BaseRps: 1,
                        MaxRps: sdkkonnectgo.Int64(1000),
                    },
                ),
                Environment: []components.ConfigurationDataPlaneGroupEnvironmentField{
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "INFO",
                    },
                    components.ConfigurationDataPlaneGroupEnvironmentField{
                        Name: "KONG_LOG_LEVEL",
                        Value: "INFO",
                    },
                },
            },
        },
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

Retrieves a configuration by ID (restricted by permitted control-plane read).

### Example Usage

<!-- UsageSnippet language="go" operationID="get-configuration" method="get" path="/v2/cloud-gateways/configurations/{configurationId}" -->
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
| `configurationID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ID of the configuration to operate on.               | edaf40f9-9fb0-4ffe-bb74-4e763a6bd471                     |
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

Returns a paginated collection of custom domains across control-planes for an organization (restricted by
permitted control-plane reads).


### Example Usage

<!-- UsageSnippet language="go" operationID="list-custom-domains" method="get" path="/v2/cloud-gateways/custom-domains" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.ListCustomDomains(ctx, operations.ListCustomDomainsRequest{
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
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

Creates a new custom domain for a control-plane (restricted by permitted control-plane associate-custom-domain
action).


### Example Usage

<!-- UsageSnippet language="go" operationID="create-custom-domains" method="post" path="/v2/cloud-gateways/custom-domains" -->
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

Retrieves a custom domain by ID (restricted by permitted control-plane reads).

### Example Usage

<!-- UsageSnippet language="go" operationID="get-custom-domain" method="get" path="/v2/cloud-gateways/custom-domains/{customDomainId}" -->
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
| `customDomainID`                                         | *string*                                                 | :heavy_check_mark:                                       | ID of the custom domain to operate on.                   | 39ed3790-085d-4605-9627-f96d86aaf425                     |
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

Deletes a custom domain by ID (restricted by permitted control-plane reads).

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-custom-domain" method="delete" path="/v2/cloud-gateways/custom-domains/{customDomainId}" -->
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
| `customDomainID`                                         | *string*                                                 | :heavy_check_mark:                                       | ID of the custom domain to operate on.                   | 39ed3790-085d-4605-9627-f96d86aaf425                     |
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

Retrieves the CNAME and SSL status of a custom domain.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-custom-domain-online-status" method="get" path="/v2/cloud-gateways/custom-domains/{customDomainId}/online-status" -->
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
| `customDomainID`                                         | *string*                                                 | :heavy_check_mark:                                       | ID of the custom domain to operate on.                   | 39ed3790-085d-4605-9627-f96d86aaf425                     |
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

Returns a paginated collection of default resource configurations for cloud-gateways, along with
organizationally-defined overrides for those resource configurations.
Resource configurations are settings that are applied to all cloud gateway resources in an organization.
For example, the "data-plane-group-idle-timeout-minutes" resource configuration sets the idle timeout for all data plane groups in an organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-default-resource-configurations" method="get" path="/v2/cloud-gateways/default-resource-configurations" -->
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

    res, err := s.CloudGateways.ListDefaultResourceConfigurations(ctx, sdkkonnectgo.Int64(10), sdkkonnectgo.Int64(1))
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
| `pageSize`                                                                                              | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
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

Returns a paginated collection of default resource quotas for cloud-gateways, along with
organizationally-defined overrides for those resource quotas.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-default-resource-quotas" method="get" path="/v2/cloud-gateways/default-resource-quotas" -->
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

    res, err := s.CloudGateways.ListDefaultResourceQuotas(ctx, sdkkonnectgo.Int64(10), sdkkonnectgo.Int64(1))
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
| `pageSize`                                                                                              | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
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

Returns a paginated list of networks.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-networks" method="get" path="/v2/cloud-gateways/networks" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.ListNetworks(ctx, operations.ListNetworksRequest{
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
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

Creates a new network for a given provider account.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-network" method="post" path="/v2/cloud-gateways/networks" -->
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

Retrieves a network by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-network" method="get" path="/v2/cloud-gateways/networks/{networkId}" -->
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
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
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

Updates a network by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-network" method="patch" path="/v2/cloud-gateways/networks/{networkId}" -->
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

    res, err := s.CloudGateways.UpdateNetwork(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.PatchNetworkRequest{
        Name: sdkkonnectgo.String("us-east-2 network"),
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
| `networkID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | The network to operate on.                                                       | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                                             |
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

Deletes a network by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-network" method="delete" path="/v2/cloud-gateways/networks/{networkId}" -->
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
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
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

Returns a paginated collection of configurations that reference a network.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-network-configurations" method="get" path="/v2/cloud-gateways/networks/{networkId}/configuration-references" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.ListNetworkConfigurations(ctx, operations.ListNetworkConfigurationsRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
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

Returns a paginated collection of Private DNS for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-private-dns" method="get" path="/v2/cloud-gateways/networks/{networkId}/private-dns" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.ListPrivateDNS(ctx, operations.ListPrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
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

Creates a new Private DNS for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-private-dns" method="post" path="/v2/cloud-gateways/networks/{networkId}/private-dns" -->
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

    res, err := s.CloudGateways.CreatePrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", components.CreatePrivateDNSRequest{
        Name: sdkkonnectgo.String("us-east-2 private dns"),
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
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `networkID`                                                                              | *string*                                                                                 | :heavy_check_mark:                                                                       | The network to operate on.                                                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                                                     |
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

Retrieves a Private DNS by ID for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-private-dns" method="get" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" -->
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

    res, err := s.CloudGateways.GetPrivateDNS(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "1850820b-c69f-4a2a-b9be-bbcdbc5cd618")
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateDNSResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `privateDNSID`                                           | *string*                                                 | :heavy_check_mark:                                       | The ID of the Private DNS to operate on.                 | 1850820b-c69f-4a2a-b9be-bbcdbc5cd618                     |
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

Updates a Private DNS by ID for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-private-dns" method="patch" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.UpdatePrivateDNS(ctx, operations.UpdatePrivateDNSRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PrivateDNSID: "1850820b-c69f-4a2a-b9be-bbcdbc5cd618",
        PatchPrivateDNSRequest: components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(
            components.PatchAwsPrivateDNSResolver{
                Name: sdkkonnectgo.String("us-east-2 private dns"),
                PrivateDNSAttachmentConfig: &components.AwsPrivateDNSResolverAttachmentConfig{
                    Kind: components.AWSPrivateDNSResolverTypeAwsOutboundResolver,
                    DNSConfig: map[string]components.PrivateDNSResolverConfig{
                        "global.api.konghq.com": components.PrivateDNSResolverConfig{
                            RemoteDNSServerIPAddresses: []string{
                                "10.0.0.2",
                            },
                        },
                        "us.api.konghq.dev": components.PrivateDNSResolverConfig{
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
        // handle response
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

Deletes a Private DNS by ID for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-private-dns" method="delete" path="/v2/cloud-gateways/networks/{networkId}/private-dns/{privateDnsId}" -->
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
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `privateDNSID`                                           | *string*                                                 | :heavy_check_mark:                                       | The ID of the Private DNS to operate on.                 | 1850820b-c69f-4a2a-b9be-bbcdbc5cd618                     |
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

Returns a paginated collection of transit gateways for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-transit-gateways" method="get" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.ListTransitGateways(ctx, operations.ListTransitGatewaysRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
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

Creates a new transit gateway for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-transit-gateway" method="post" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways" -->
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
                Kind: components.AWSVPCPeeringAttachmentConfigAwsVpcPeeringAttachment,
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
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `networkID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The network to operate on.                                                                       | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                                                             |
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

Retrieves a transit gateway by ID for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-transit-gateway" method="get" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways/{transitGatewayId}" -->
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

    res, err := s.CloudGateways.GetTransitGateway(ctx, "36ae63d3-efd1-4bec-b246-62aa5d3f5695", "0850820b-d153-4a2a-b9be-7d2204779139")
    if err != nil {
        log.Fatal(err)
    }
    if res.TransitGatewayResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `transitGatewayID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the transit gateway to operate on.             | 0850820b-d153-4a2a-b9be-7d2204779139                     |
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

Updates a transit gateway by ID for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-transit-gateway" method="patch" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways/{transitGatewayId}" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.UpdateTransitGateway(ctx, operations.UpdateTransitGatewayRequest{
        NetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
        TransitGatewayID: "0850820b-d153-4a2a-b9be-7d2204779139",
        PatchTransitGatewayRequest: components.CreatePatchTransitGatewayRequestPatchAwsResourceEndpointGateway(
            components.PatchAwsResourceEndpointGateway{
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
        // handle response
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

Deletes a transit gateway by ID for a given network.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-transit-gateway" method="delete" path="/v2/cloud-gateways/networks/{networkId}/transit-gateways/{transitGatewayId}" -->
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
| `networkID`                                              | *string*                                                 | :heavy_check_mark:                                       | The network to operate on.                               | 36ae63d3-efd1-4bec-b246-62aa5d3f5695                     |
| `transitGatewayID`                                       | *string*                                                 | :heavy_check_mark:                                       | The ID of the transit gateway to operate on.             | 0850820b-d153-4a2a-b9be-7d2204779139                     |
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

Returns a a paginated collection of provider accounts for an organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-provider-accounts" method="get" path="/v2/cloud-gateways/provider-accounts" -->
```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CloudGateways.ListProviderAccounts(ctx, operations.ListProviderAccountsRequest{
        PageSize: sdkkonnectgo.Int64(10),
        PageNumber: sdkkonnectgo.Int64(1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListProviderAccountsResponse != nil {
        // handle response
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

Retrieves a provider account by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-provider-account" method="get" path="/v2/cloud-gateways/provider-accounts/{providerAccountId}" -->
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
| `providerAccountID`                                      | *string*                                                 | :heavy_check_mark:                                       | The ID of the provider account to operate on.            | 929b2449-c69f-44c4-b6ad-9ecec6f811ae                     |
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

Returns a paginated collection of resource configurations for an organization.
Resource configurations are settings that are applied to all cloud gateway resources in an organization.
For example, the "data-plane-group-idle-timeout-minutes" resource configuration sets the idle timeout for all data plane groups in an organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-resource-configurations" method="get" path="/v2/cloud-gateways/resource-configurations" -->
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

    res, err := s.CloudGateways.ListResourceConfigurations(ctx, sdkkonnectgo.Int64(10), sdkkonnectgo.Int64(1))
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
| `pageSize`                                                                                              | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
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

Retrieves a resource configuration by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-resource-configuration" method="get" path="/v2/cloud-gateways/resource-configurations/{resourceConfigurationId}" -->
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
| `resourceConfigurationID`                                | *string*                                                 | :heavy_check_mark:                                       | The ID of the resource configuration to operate on.      | 9678f205-49a1-47bb-82d9-d01cafa42a0d                     |
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

Returns a paginated collection of resource quotas for an organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-resource-quotas" method="get" path="/v2/cloud-gateways/resource-quotas" -->
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

    res, err := s.CloudGateways.ListResourceQuotas(ctx, sdkkonnectgo.Int64(10), sdkkonnectgo.Int64(1))
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
| `pageSize`                                                                                              | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | Determines which page of the entities to retrieve.                                                      | 1                                                                                                       |
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

Retrieves a resource quota by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-resource-quota" method="get" path="/v2/cloud-gateways/resource-quotas/{resourceQuotaId}" -->
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
| `resourceQuotaID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ID of the resource quota to operate on.              | 9678f205-49a1-47bb-82d9-d01cafa42a0d                     |
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