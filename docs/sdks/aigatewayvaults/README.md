# AIGatewayVaults

## Overview

API related to the management of AI Gateway vaults for storing secrets.

### Available Operations

* [ListAiGatewayVaults](#listaigatewayvaults) - List AI Gateway Vaults
* [CreateAiGatewayVault](#createaigatewayvault) - Create an AI Gateway Vault
* [GetAiGatewayVault](#getaigatewayvault) - Get an AI Gateway Vault
* [UpdateAiGatewayVault](#updateaigatewayvault) - Update an AI Gateway Vault
* [DeleteAiGatewayVault](#deleteaigatewayvault) - Delete an AI Gateway Vault

## ListAiGatewayVaults

Returns a list of vaults associated with the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ai-gateway-vaults" method="get" path="/v1/ai-gateways/{gatewayId}/vaults" -->
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

    res, err := s.AIGatewayVaults.ListAiGatewayVaults(ctx, operations.ListAiGatewayVaultsRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAIGatewayVaultsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAiGatewayVaultsRequest](../../models/operations/listaigatewayvaultsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAiGatewayVaultsResponse](../../models/operations/listaigatewayvaultsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## CreateAiGatewayVault

Registers a new vault for the AI Gateway.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-gateway-vault" method="post" path="/v1/ai-gateways/{gatewayId}/vaults" -->
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

    res, err := s.AIGatewayVaults.CreateAiGatewayVault(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", components.CreateCreateAIGatewayVaultRequestGcp(
        components.GoogleSecretManagerVault{
            Name: "my-awesome-vault",
            Type: components.GoogleSecretManagerVaultTypeGcp,
            Config: components.GoogleSecretManagerVaultConfig{
                ProjectID: "<id>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayVault != nil {
        switch res.AIGatewayVault.Type {
            case components.AIGatewayVaultTypeKonnect:
                // res.AIGatewayVault.AIGatewayVaultKonnectConfigStoreVault is populated
            case components.AIGatewayVaultTypeEnv:
                // res.AIGatewayVault.AIGatewayVaultEnvironmentVariableVault is populated
            case components.AIGatewayVaultTypeAws:
                // res.AIGatewayVault.AIGatewayVaultAwsSecretsManagerVault is populated
            case components.AIGatewayVaultTypeGcp:
                // res.AIGatewayVault.AIGatewayVaultGoogleSecretManagerVault is populated
            case components.AIGatewayVaultTypeAzure:
                // res.AIGatewayVault.AIGatewayVaultAzureKeyVault is populated
            case components.AIGatewayVaultTypeConjur:
                // res.AIGatewayVault.AIGatewayVaultConjurVault is populated
            case components.AIGatewayVaultTypeHcv:
                // res.AIGatewayVault.AIGatewayVaultHashiCorpVault is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `gatewayID`                                                                                      | `string`                                                                                         | :heavy_check_mark:                                                                               | The unique ID of the AI Gateway.                                                                 | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                             |
| `createAIGatewayVaultRequest`                                                                    | [components.CreateAIGatewayVaultRequest](../../models/components/createaigatewayvaultrequest.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CreateAiGatewayVaultResponse](../../models/operations/createaigatewayvaultresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## GetAiGatewayVault

Returns the details of a specific AI Gateway vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-gateway-vault" method="get" path="/v1/ai-gateways/{gatewayId}/vaults/{vaultIdOrName}" -->
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

    res, err := s.AIGatewayVaults.GetAiGatewayVault(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayVault != nil {
        switch res.AIGatewayVault.Type {
            case components.AIGatewayVaultTypeKonnect:
                // res.AIGatewayVault.AIGatewayVaultKonnectConfigStoreVault is populated
            case components.AIGatewayVaultTypeEnv:
                // res.AIGatewayVault.AIGatewayVaultEnvironmentVariableVault is populated
            case components.AIGatewayVaultTypeAws:
                // res.AIGatewayVault.AIGatewayVaultAwsSecretsManagerVault is populated
            case components.AIGatewayVaultTypeGcp:
                // res.AIGatewayVault.AIGatewayVaultGoogleSecretManagerVault is populated
            case components.AIGatewayVaultTypeAzure:
                // res.AIGatewayVault.AIGatewayVaultAzureKeyVault is populated
            case components.AIGatewayVaultTypeConjur:
                // res.AIGatewayVault.AIGatewayVaultConjurVault is populated
            case components.AIGatewayVaultTypeHcv:
                // res.AIGatewayVault.AIGatewayVaultHashiCorpVault is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `vaultIDOrName`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway Vault.           | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiGatewayVaultResponse](../../models/operations/getaigatewayvaultresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## UpdateAiGatewayVault

Updates the configuration of an existing AI Gateway vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-gateway-vault" method="put" path="/v1/ai-gateways/{gatewayId}/vaults/{vaultIdOrName}" -->
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

    res, err := s.AIGatewayVaults.UpdateAiGatewayVault(ctx, operations.UpdateAiGatewayVaultRequest{
        GatewayID: "bf138ba2-c9b1-4229-b268-04d9d8a6410b",
        VaultIDOrName: "my-entity-name",
        UpdateAIGatewayVaultRequest: components.CreateUpdateAIGatewayVaultRequestEnv(
            components.EnvironmentVariableVault{
                Name: "my-awesome-vault",
                Type: components.EnvironmentVariableVaultTypeEnv,
                Config: components.EnvironmentVariableVaultConfig{
                    Prefix: sdkkonnectgo.Pointer("MY_SECRET_"),
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIGatewayVault != nil {
        switch res.AIGatewayVault.Type {
            case components.AIGatewayVaultTypeKonnect:
                // res.AIGatewayVault.AIGatewayVaultKonnectConfigStoreVault is populated
            case components.AIGatewayVaultTypeEnv:
                // res.AIGatewayVault.AIGatewayVaultEnvironmentVariableVault is populated
            case components.AIGatewayVaultTypeAws:
                // res.AIGatewayVault.AIGatewayVaultAwsSecretsManagerVault is populated
            case components.AIGatewayVaultTypeGcp:
                // res.AIGatewayVault.AIGatewayVaultGoogleSecretManagerVault is populated
            case components.AIGatewayVaultTypeAzure:
                // res.AIGatewayVault.AIGatewayVaultAzureKeyVault is populated
            case components.AIGatewayVaultTypeConjur:
                // res.AIGatewayVault.AIGatewayVaultConjurVault is populated
            case components.AIGatewayVaultTypeHcv:
                // res.AIGatewayVault.AIGatewayVaultHashiCorpVault is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAiGatewayVaultRequest](../../models/operations/updateaigatewayvaultrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateAiGatewayVaultResponse](../../models/operations/updateaigatewayvaultresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## DeleteAiGatewayVault

Removes a specific AI Gateway vault.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-gateway-vault" method="delete" path="/v1/ai-gateways/{gatewayId}/vaults/{vaultIdOrName}" -->
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

    res, err := s.AIGatewayVaults.DeleteAiGatewayVault(ctx, "bf138ba2-c9b1-4229-b268-04d9d8a6410b", "my-entity-name")
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
| `gatewayID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique ID of the AI Gateway.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `vaultIDOrName`                                          | `string`                                                 | :heavy_check_mark:                                       | The unique ID or name of the AI Gateway Vault.           | my-entity-name                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiGatewayVaultResponse](../../models/operations/deleteaigatewayvaultresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |