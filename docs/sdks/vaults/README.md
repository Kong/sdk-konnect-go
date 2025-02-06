# Vaults
(*Vaults*)

## Overview

Vault objects are used to configure different vault connectors for [managing secrets](https://docs.konghq.com/gateway/latest/kong-enterprise/secrets-management/).
Configuring a vault lets you reference secrets from other entities.
This allows for a proper separation of secrets and configuration and prevents secret sprawl.
<br><br>
For example, you could store a certificate and a key in a vault, then reference them from a certificate entity. This way, the certificate and key are not stored in the entity directly and are more secure.
<br><br>
Secrets rotation can be managed using [TTLs](https://docs.konghq.com/gateway/latest/kong-enterprise/secrets-management/advanced-usage/).


### Available Operations

* [ListVault](#listvault) - List all Vaults
* [CreateVault](#createvault) - Create a new Vault
* [DeleteVault](#deletevault) - Delete a Vault
* [GetVault](#getvault) - Fetch a Vault
* [UpsertVault](#upsertvault) - Upsert a Vault

## ListVault

List all Vaults

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

    res, err := s.Vaults.ListVault(ctx, operations.ListVaultRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.String("tag1,tag2"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListVaultRequest](../../models/operations/listvaultrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.ListVaultResponse](../../models/operations/listvaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateVault

Create a new Vault

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

    res, err := s.Vaults.CreateVault(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.VaultInput{
        Config: map[string]any{
            "prefix": "ENV_PREFIX",
        },
        Description: sdkkonnectgo.String("environment variable based vault"),
        ID: sdkkonnectgo.String("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
        Name: "env",
        Prefix: "env",
        Tags: []string{
            "foo",
            "bar",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Vault != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         | Example                                                                                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |                                                                                                                                                                                                     |
| `controlPlaneID`                                                                                                                                                                                    | *string*                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                  | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                  | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                |
| `vault`                                                                                                                                                                                             | [components.VaultInput](../../models/components/vaultinput.md)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                  | Description of the new Vault for creation                                                                                                                                                           | {<br/>"config": {<br/>"prefix": "ENV_PREFIX"<br/>},<br/>"description": "environment variable based vault",<br/>"id": "2747d1e5-8246-4f65-a939-b392f1ee17f8",<br/>"name": "env",<br/>"prefix": "env",<br/>"tags": [<br/>"foo",<br/>"bar"<br/>]<br/>} |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |                                                                                                                                                                                                     |

### Response

**[*operations.CreateVaultResponse](../../models/operations/createvaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteVault

Delete a Vault

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

    res, err := s.Vaults.DeleteVault(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "9d4d6d19-77c6-428e-a965-9bc9647633e9")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `vaultID`                                                                          | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Vault to lookup                                                          | 9d4d6d19-77c6-428e-a965-9bc9647633e9                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteVaultResponse](../../models/operations/deletevaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetVault

Get a Vault using ID or prefix.

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

    res, err := s.Vaults.GetVault(ctx, "9d4d6d19-77c6-428e-a965-9bc9647633e9", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Vault != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `vaultID`                                                                          | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Vault to lookup                                                          | 9d4d6d19-77c6-428e-a965-9bc9647633e9                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetVaultResponse](../../models/operations/getvaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertVault

Create or Update Vault using ID or prefix.

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

    res, err := s.Vaults.UpsertVault(ctx, operations.UpsertVaultRequest{
        VaultID: "9d4d6d19-77c6-428e-a965-9bc9647633e9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Vault: components.VaultInput{
            Config: map[string]any{
                "prefix": "ENV_PREFIX",
            },
            Description: sdkkonnectgo.String("environment variable based vault"),
            ID: sdkkonnectgo.String("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
            Name: "env",
            Prefix: "env",
            Tags: []string{
                "foo",
                "bar",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Vault != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpsertVaultRequest](../../models/operations/upsertvaultrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpsertVaultResponse](../../models/operations/upsertvaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |