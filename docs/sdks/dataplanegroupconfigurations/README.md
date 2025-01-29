# DataPlaneGroupConfigurations
(*DataPlaneGroupConfigurations*)

## Overview

### Available Operations

* [ListConfigurations](#listconfigurations) - List Configurations
* [CreateConfiguration](#createconfiguration) - Create Configuration
* [GetConfiguration](#getconfiguration) - Get Configuration

## ListConfigurations

Returns a paginated collection of configurations across control-planes for an organization (restricted by
permitted control-plane reads).


### Example Usage

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

    res, err := s.DataPlaneGroupConfigurations.ListConfigurations(ctx, operations.ListConfigurationsRequest{
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

    res, err := s.DataPlaneGroupConfigurations.CreateConfiguration(ctx, components.CreateConfigurationRequest{
        ControlPlaneID: "0949471e-b759-45ba-87ab-ee63fb781388",
        ControlPlaneGeo: components.ControlPlaneGeoEu,
        Version: "3.2",
        DataplaneGroups: []components.CreateConfigurationDataPlaneGroup{
            components.CreateConfigurationDataPlaneGroup{
                Provider: components.ProviderNameAws,
                Region: "us-east-2",
                CloudGatewayNetworkID: "36ae63d3-efd1-4bec-b246-62aa5d3f5695",
                Autoscale: components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleStatic(
                    components.ConfigurationDataPlaneGroupAutoscaleStatic{
                        Kind: components.KindStatic,
                        InstanceType: components.InstanceTypeNameMedium,
                        RequestedInstances: 3,
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

Retrieves a configuration by ID (restricted by permitted control-plane read).

### Example Usage

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

    res, err := s.DataPlaneGroupConfigurations.GetConfiguration(ctx, "edaf40f9-9fb0-4ffe-bb74-4e763a6bd471")
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