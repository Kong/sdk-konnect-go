# EventGatewayVaults

## Overview

### Available Operations

* [ListEventGatewayVaults](#listeventgatewayvaults) - List Vaults
* [CreateEventGatewayVault](#createeventgatewayvault) - Create Vault
* [GetEventGatewayVault](#geteventgatewayvault) - Get a Vault
* [UpdateEventGatewayVault](#updateeventgatewayvault) - Update Vault
* [DeleteEventGatewayVault](#deleteeventgatewayvault) - Delete Vault

## ListEventGatewayVaults

Returns a list of vaults associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-vaults" method="get" path="/v1/event-gateways/{gatewayId}/vaults" -->
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

    res, err := s.EventGatewayVaults.ListEventGatewayVaults(ctx, operations.ListEventGatewayVaultsRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEventGatewayVaultsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListEventGatewayVaultsRequest](../../models/operations/listeventgatewayvaultsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListEventGatewayVaultsResponse](../../models/operations/listeventgatewayvaultsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayVault

Creates a new vault associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-vault" method="post" path="/v1/event-gateways/{gatewayId}/vaults" -->
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

    res, err := s.EventGatewayVaults.CreateEventGatewayVault(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", sdkkonnectgo.Pointer(components.CreateEventGatewayModifyVaultKonnect(
        components.EventGatewayKonnectVault{
            Name: "<value>",
        },
    )))
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayVault != nil {
        switch res.EventGatewayVault.Type {
            case components.EventGatewayVaultTypeEnv:
                // res.EventGatewayVault.EventGatewayVaultEventGatewayEnvVault is populated
            case components.EventGatewayVaultTypeKonnect:
                // res.EventGatewayVault.EventGatewayVaultEventGatewayKonnectVault is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `gatewayID`                                                                               | `string`                                                                                  | :heavy_check_mark:                                                                        | The UUID of your Gateway.                                                                 | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                      |
| `eventGatewayModifyVault`                                                                 | [*components.EventGatewayModifyVault](../../models/components/eventgatewaymodifyvault.md) | :heavy_minus_sign:                                                                        | The request schema for creating a vault.                                                  |                                                                                           |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.CreateEventGatewayVaultResponse](../../models/operations/createeventgatewayvaultresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVault

Returns information about a specific vault associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-vault" method="get" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}" -->
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

    res, err := s.EventGatewayVaults.GetEventGatewayVault(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "85a372b5-93ec-4ff9-9525-81c91606ab6a")
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayVault != nil {
        switch res.EventGatewayVault.Type {
            case components.EventGatewayVaultTypeEnv:
                // res.EventGatewayVault.EventGatewayVaultEventGatewayEnvVault is populated
            case components.EventGatewayVaultTypeKonnect:
                // res.EventGatewayVault.EventGatewayVaultEventGatewayKonnectVault is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `vaultID`                                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the Vault.                                     |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetEventGatewayVaultResponse](../../models/operations/geteventgatewayvaultresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVault

Updates an existing vault associated with the specified Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-vault" method="put" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}" -->
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

    res, err := s.EventGatewayVaults.UpdateEventGatewayVault(ctx, operations.UpdateEventGatewayVaultRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VaultID: "eb40b7d3-8efe-42d8-b0e5-e6cbba0c19a7",
        EventGatewayModifyVault: sdkkonnectgo.Pointer(components.CreateEventGatewayModifyVaultKonnect(
            components.EventGatewayKonnectVault{
                Name: "<value>",
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayVault != nil {
        switch res.EventGatewayVault.Type {
            case components.EventGatewayVaultTypeEnv:
                // res.EventGatewayVault.EventGatewayVaultEventGatewayEnvVault is populated
            case components.EventGatewayVaultTypeKonnect:
                // res.EventGatewayVault.EventGatewayVaultEventGatewayKonnectVault is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateEventGatewayVaultRequest](../../models/operations/updateeventgatewayvaultrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateEventGatewayVaultResponse](../../models/operations/updateeventgatewayvaultresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayVault

Deletes a specific vault associated with the Event Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-vault" method="delete" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}" -->
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

    res, err := s.EventGatewayVaults.DeleteEventGatewayVault(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "b0d99e41-8352-4982-97ac-0a4f2a3a79c3")
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
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The UUID of your Gateway.                                | 9524ec7d-36d9-465d-a8c5-83a3c9390458                     |
| `vaultID`                                                | `string`                                                 | :heavy_check_mark:                                       | The ID of the Vault.                                     |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteEventGatewayVaultResponse](../../models/operations/deleteeventgatewayvaultresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |