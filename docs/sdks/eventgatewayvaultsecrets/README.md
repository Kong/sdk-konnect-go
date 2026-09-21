# EventGatewayVaultSecrets

## Overview

### Available Operations

* [ListEventGatewayVaultSecrets](#listeventgatewayvaultsecrets) - List Vault Secrets
* [CreateEventGatewayVaultSecret](#createeventgatewayvaultsecret) - Create Vault Secret
* [GetEventGatewayVaultSecret](#geteventgatewayvaultsecret) - Get Vault Secret
* [UpdateEventGatewayVaultSecret](#updateeventgatewayvaultsecret) - Update Vault Secret
* [DeleteEventGatewayVaultSecret](#deleteeventgatewayvaultsecret) - Delete Vault Secret

## ListEventGatewayVaultSecrets

Returns a list of secrets associated with the specified Event Gateway Vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-event-gateway-vault-secrets" method="get" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}/secrets" -->
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

    res, err := s.EventGatewayVaultSecrets.ListEventGatewayVaultSecrets(ctx, operations.ListEventGatewayVaultSecretsRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VaultID: "a8934008-5576-4651-b15f-7b33a0daf9b3",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEventGatewayVaultSecretsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListEventGatewayVaultSecretsRequest](../../models/operations/listeventgatewayvaultsecretsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListEventGatewayVaultSecretsResponse](../../models/operations/listeventgatewayvaultsecretsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateEventGatewayVaultSecret

Creates a new secret associated with the specified Event Gateway Vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-event-gateway-vault-secret" method="post" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}/secrets" -->
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

    res, err := s.EventGatewayVaultSecrets.CreateEventGatewayVaultSecret(ctx, operations.CreateEventGatewayVaultSecretRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VaultID: "c876ab53-2444-4193-8992-ee0c953d9f36",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayVaultSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateEventGatewayVaultSecretRequest](../../models/operations/createeventgatewayvaultsecretrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateEventGatewayVaultSecretResponse](../../models/operations/createeventgatewayvaultsecretresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetEventGatewayVaultSecret

Returns information about a secret contained within the specified Event Gateway Vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-event-gateway-vault-secret" method="get" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}/secrets/{secretId}" -->
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

    res, err := s.EventGatewayVaultSecrets.GetEventGatewayVaultSecret(ctx, operations.GetEventGatewayVaultSecretRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VaultID: "b852d062-dcb3-4a34-a11a-470e61474be1",
        SecretID: "befd7f3f-a205-4bdf-9bd1-cab579290f4a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayVaultSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetEventGatewayVaultSecretRequest](../../models/operations/geteventgatewayvaultsecretrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetEventGatewayVaultSecretResponse](../../models/operations/geteventgatewayvaultsecretresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateEventGatewayVaultSecret

Updates a secret contained within the specified Event Gateway Vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-event-gateway-vault-secret" method="put" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}/secrets/{secretId}" -->
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

    res, err := s.EventGatewayVaultSecrets.UpdateEventGatewayVaultSecret(ctx, operations.UpdateEventGatewayVaultSecretRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VaultID: "3642a0d8-3cdf-405c-b433-1e1bb6fcf682",
        SecretID: "1d1e0fac-b3fa-4fa8-a538-4c603684b279",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGatewayVaultSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateEventGatewayVaultSecretRequest](../../models/operations/updateeventgatewayvaultsecretrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.UpdateEventGatewayVaultSecretResponse](../../models/operations/updateeventgatewayvaultsecretresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteEventGatewayVaultSecret

Deletes a secret contained within the specified Event Gateway Vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-event-gateway-vault-secret" method="delete" path="/v1/event-gateways/{gatewayId}/vaults/{vaultId}/secrets/{secretId}" -->
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

    res, err := s.EventGatewayVaultSecrets.DeleteEventGatewayVaultSecret(ctx, operations.DeleteEventGatewayVaultSecretRequest{
        GatewayID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VaultID: "32de71b8-7a56-413e-9aab-da5d3f25c9c1",
        SecretID: "b5d9e935-6cc3-4083-8a61-48f928fd176e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteEventGatewayVaultSecretRequest](../../models/operations/deleteeventgatewayvaultsecretrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.DeleteEventGatewayVaultSecretResponse](../../models/operations/deleteeventgatewayvaultsecretresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |