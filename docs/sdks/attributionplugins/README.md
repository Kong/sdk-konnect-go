# AttributionPlugins

## Overview

The attribution plugin turns request headers and identity claims into cost tags on an AI Gateway. These operations install, update, and remove it per AI Gateway, and list the installations.

### Available Operations

* [GetCostAttributionPlugin](#getcostattributionplugin) - Get an AI Gateway's attribution plugin
* [UpsertCostAttributionPlugin](#upsertcostattributionplugin) - Install or replace an attribution plugin
* [UpdateCostAttributionPlugin](#updatecostattributionplugin) - Update an AI Gateway's attribution plugin
* [DeleteCostAttributionPlugin](#deletecostattributionplugin) - Remove an AI Gateway's attribution plugin
* [ListCostAttributionPlugins](#listcostattributionplugins) - List attribution plugin installations

## GetCostAttributionPlugin

Returns the attribution plugin installed on an AI Gateway, with its version, the newest version, and its configuration. The installation is what this product last installed: a change made to the gateway's policies elsewhere doesn't show here. If the gateway has no installation, or the identifier isn't a valid ID, the request fails with `404`.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-cost-attribution-plugin" method="get" path="/v1/ai-cost-management/ai-gateways/{gatewayId}/attribution-plugin" -->
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

    res, err := s.AttributionPlugins.GetCostAttributionPlugin(ctx, "527bc328-6950-4662-a473-78dd7a5beb84")
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostAttributionPlugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the AI Gateway.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCostAttributionPluginResponse](../../models/operations/getcostattributionpluginresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpsertCostAttributionPlugin

Installs the latest attribution plugin version on an AI Gateway with this configuration, or replaces the installed plugin and its configuration, and returns the installation.

- The plugin becomes a streaming custom policy named `ai-cost-attribution` and a global policy that turns it on. The request writes both through the AI Gateway API as the caller, so the caller also needs the AI Gateways Admin role on the gateway.
- Every request installs the latest version, so a `PUT` of the stored configuration updates an outdated installation.
- The plugin tags requests only on data planes that run AI Gateway 2.3 or later. Each data plane must run with `KONG_CUSTOM_PLUGIN_STREAMING_ENABLED=on`. A data plane with streaming off refuses the gateway's whole configuration while the plugin is installed.
- If the gateway doesn't have the plugin yet and a data plane the AI Gateway API lists reports streaming off, or doesn't report the setting because it runs an older release, the request fails with `409`, names those data planes, and changes nothing. The API lists a data plane for five minutes after its last ping, so a data plane replaced in that time can still be named. A data plane that connects later isn't checked, and neither is a gateway that already has the plugin.
- If the request fails partway, repeating it finishes the installation.
- If the body is invalid, the request fails with `400`. A header mapping that names a credential header, or repeats a header or tag key, counts as invalid.
- If the gateway already has a custom policy or policy of this name that this product didn't install, the request fails with `409` and changes nothing.
- If another change to the gateway's plugin is running, the request fails with `409` and changes nothing. Repeat it once that change has finished.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-cost-attribution-plugin" method="put" path="/v1/ai-cost-management/ai-gateways/{gatewayId}/attribution-plugin" -->
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

    res, err := s.AttributionPlugins.UpsertCostAttributionPlugin(ctx, "4b2efecb-9f38-4d74-9073-5ab0f83a19c6", components.AICostAttributionPluginConfiguration{
        Integrations: []components.AICostAttributionPluginConfigurationIntegrations{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostAttributionPlugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `gatewayID`                                                                                                        | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | The ID of the AI Gateway.                                                                                          |
| `aiCostAttributionPluginConfiguration`                                                                             | [components.AICostAttributionPluginConfiguration](../../models/components/aicostattributionpluginconfiguration.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpsertCostAttributionPluginResponse](../../models/operations/upsertcostattributionpluginresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateCostAttributionPlugin

Changes part of an installation's configuration with a JSON Merge Patch, installs the latest plugin version with the result, and returns the installation.

- An omitted field keeps its stored value, and an array, such as `headers`, replaces the stored one whole.
- `operator_tags` merges per key, and a key set to `null` is removed.
- `null` returns `headers`, `prefixes`, `operator_tags`, or `group_patterns` to its empty default. `integrations` can't be `null`.
- Like `PUT`, the request writes the plugin through the AI Gateway API as the caller, and it updates an outdated installation to the latest version.
- If the gateway has no installation, or the identifier isn't a valid ID, the request fails with `404`. The request never installs the plugin.
- The body is `application/merge-patch+json` or `application/json`. Any other media type fails with `415`.
- If the patch, or the configuration it produces, is invalid, the request fails with `400`.
- If another change to the gateway's plugin is running, the request fails with `409` and changes nothing. Repeat it once that change has finished.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-cost-attribution-plugin" method="patch" path="/v1/ai-cost-management/ai-gateways/{gatewayId}/attribution-plugin" -->
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

    res, err := s.AttributionPlugins.UpdateCostAttributionPlugin(ctx, "062b78aa-4607-4c6d-8eab-e8c23d829b86", components.AICostAttributionPluginPatch{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AICostAttributionPlugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `gatewayID`                                                                                        | `string`                                                                                           | :heavy_check_mark:                                                                                 | The ID of the AI Gateway.                                                                          |
| `aiCostAttributionPluginPatch`                                                                     | [components.AICostAttributionPluginPatch](../../models/components/aicostattributionpluginpatch.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateCostAttributionPluginResponse](../../models/operations/updatecostattributionpluginresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteCostAttributionPlugin

Removes the attribution plugin from an AI Gateway: the global policy that turns it on, the custom policy that holds it, and the installation. The gateway's requests stop getting the plugin's tags. The request writes through the AI Gateway API as the caller, so the caller also needs the AI Gateways Admin role on the gateway, and a policy that is already gone counts as removed.

If the gateway has no installation, or the identifier isn't a valid ID, the request fails with `404`. This includes removing the plugin again. If another change to the gateway's plugin is running, the request fails with `409` and changes nothing. Repeat it once that change has finished.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-cost-attribution-plugin" method="delete" path="/v1/ai-cost-management/ai-gateways/{gatewayId}/attribution-plugin" -->
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

    res, err := s.AttributionPlugins.DeleteCostAttributionPlugin(ctx, "86a243ff-0e4d-40f7-aa8d-58fdd7d5d387")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID of the AI Gateway.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCostAttributionPluginResponse](../../models/operations/deletecostattributionpluginresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListCostAttributionPlugins

Returns a page of the attribution plugin installations on the organization's AI Gateways, sorted by `gateway_id` unless the request sets `sort`. An AI Gateway without the plugin has no entry.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-cost-attribution-plugins" method="get" path="/v1/ai-cost-management/attribution-plugins" -->
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

    res, err := s.AttributionPlugins.ListCostAttributionPlugins(ctx, operations.ListCostAttributionPluginsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCostAttributionPluginsRequest](../../models/operations/listcostattributionpluginsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListCostAttributionPluginsResponse](../../models/operations/listcostattributionpluginsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |