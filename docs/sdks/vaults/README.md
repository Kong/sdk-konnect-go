# Vaults

## Overview

Vault objects are used to configure different vault connectors for [managing secrets](https://developer.konghq.com/gateway/secrets-management/).
Configuring a vault lets you reference secrets from other entities.
This allows for a proper separation of secrets and configuration and prevents secret sprawl.
<br><br>
For example, you could store a certificate and a key in a vault, then reference them from a certificate entity. This way, the certificate and key are not stored in the entity directly and are more secure.
<br><br>
Secrets rotation can be managed using [TTLs](https://developer.konghq.com/gateway/entities/vault/).


### Available Operations

* [ListVault](#listvault) - List all Vaults in a workspace
* [CreateVault](#createvault) - Create a new Vault in a workspace
* [DeleteVault](#deletevault) - Delete a Vault in a workspace
* [GetVault](#getvault) - Get a Vault in a workspace
* [UpsertVault](#upsertvault) - Upsert a Vault in a workspace

## ListVault

List all Vaults in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-vault" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults" -->
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

    res, err := s.Vaults.ListVault(ctx, operations.ListVaultRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        Workspace: "team-payments",
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

Create a new Vault in a workspace

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-vault" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults" example="DuplicateApiKey" -->
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

    res, err := s.Vaults.CreateVault(ctx, operations.CreateVaultRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Vault: components.Vault{
            Config: map[string]any{
                "prefix": "ENV_PREFIX",
            },
            Description: sdkkonnectgo.Pointer("environment variable based vault"),
            ID: sdkkonnectgo.Pointer("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
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
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-vault" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults" example="InvalidAuthCred" -->
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

    res, err := s.Vaults.CreateVault(ctx, operations.CreateVaultRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Vault: components.Vault{
            Config: map[string]any{
                "prefix": "ENV_PREFIX",
            },
            Description: sdkkonnectgo.Pointer("environment variable based vault"),
            ID: sdkkonnectgo.Pointer("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
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
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-vault" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults" example="NoAPIKey" -->
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

    res, err := s.Vaults.CreateVault(ctx, operations.CreateVaultRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Vault: components.Vault{
            Config: map[string]any{
                "prefix": "ENV_PREFIX",
            },
            Description: sdkkonnectgo.Pointer("environment variable based vault"),
            ID: sdkkonnectgo.Pointer("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
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
| `request`                                                                      | [operations.CreateVaultRequest](../../models/operations/createvaultrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateVaultResponse](../../models/operations/createvaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteVault

Delete a Vault in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-vault" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults/{VaultId}" -->
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

    res, err := s.Vaults.DeleteVault(ctx, operations.DeleteVaultRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        VaultID: "9d4d6d19-77c6-428e-a965-9bc9647633e9",
        Workspace: "team-payments",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteVaultRequest](../../models/operations/deletevaultrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.DeleteVaultResponse](../../models/operations/deletevaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetVault

Get a Vault using ID or prefix in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-vault" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults/{VaultId}" -->
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

    res, err := s.Vaults.GetVault(ctx, operations.GetVaultRequest{
        VaultID: "9d4d6d19-77c6-428e-a965-9bc9647633e9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetVaultRequest](../../models/operations/getvaultrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.GetVaultResponse](../../models/operations/getvaultresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertVault

Create or Update Vault using ID or prefix in a workspace.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-vault" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults/{VaultId}" example="DuplicateApiKey" -->
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

    res, err := s.Vaults.UpsertVault(ctx, operations.UpsertVaultRequest{
        VaultID: "9d4d6d19-77c6-428e-a965-9bc9647633e9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Vault: components.Vault{
            Config: map[string]any{
                "prefix": "ENV_PREFIX",
            },
            Description: sdkkonnectgo.Pointer("environment variable based vault"),
            ID: sdkkonnectgo.Pointer("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
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
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-vault" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults/{VaultId}" example="InvalidAuthCred" -->
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

    res, err := s.Vaults.UpsertVault(ctx, operations.UpsertVaultRequest{
        VaultID: "9d4d6d19-77c6-428e-a965-9bc9647633e9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Vault: components.Vault{
            Config: map[string]any{
                "prefix": "ENV_PREFIX",
            },
            Description: sdkkonnectgo.Pointer("environment variable based vault"),
            ID: sdkkonnectgo.Pointer("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
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
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-vault" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/vaults/{VaultId}" example="NoAPIKey" -->
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

    res, err := s.Vaults.UpsertVault(ctx, operations.UpsertVaultRequest{
        VaultID: "9d4d6d19-77c6-428e-a965-9bc9647633e9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Vault: components.Vault{
            Config: map[string]any{
                "prefix": "ENV_PREFIX",
            },
            Description: sdkkonnectgo.Pointer("environment variable based vault"),
            ID: sdkkonnectgo.Pointer("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
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