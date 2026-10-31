# AuthServerClients

## Overview

Clients represent the identity of machines, such as microservices, mobile apps, or scripts entity. The management API will give you the ability to create, configure and manage multiple Clients per Auth Server.

### Available Operations

* [ListAuthServerClients](#listauthserverclients) - List all auth server clients
* [CreateAuthServerClient](#createauthserverclient) - Create a new auth server client
* [GetAuthServerClient](#getauthserverclient) - Get an auth server client
* [UpdateAuthServerClient](#updateauthserverclient) - Update an auth server client
* [ReplaceAuthServerClient](#replaceauthserverclient) - Replace an auth server client
* [DeleteAuthServerClient](#deleteauthserverclient) - Delete an auth server client
* [ListClientSecrets](#listclientsecrets) - List all client secrets
* [CreateClientSecret](#createclientsecret) - Create a new client secret
* [GetClientSecret](#getclientsecret) - Get a client secret
* [UpdateClientSecret](#updateclientsecret) - Update a client secret
* [DeleteClientSecret](#deleteclientsecret) - Delete a client secret
* [TestClaimForClient](#testclaimforclient) - Test a claim template for a client
* [ListAuthServerClientTokens](#listauthserverclienttokens) - List all auth server client tokens
* [DeleteAllAuthServerClientTokens](#deleteallauthserverclienttokens) - Delete all auth server client tokens
* [GetAuthServerClientToken](#getauthserverclienttoken) - Get an auth server client token
* [DeleteAuthServerClientToken](#deleteauthserverclienttoken) - Delete an auth server client token

## ListAuthServerClients

List all the clients for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAuthServerClients" method="get" path="/v1/auth-servers/{authServerId}/clients" -->
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

    res, err := s.AuthServerClients.ListAuthServerClients(ctx, operations.ListAuthServerClientsRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListAuthServerClientsRequest](../../models/operations/listauthserverclientsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListAuthServerClientsResponse](../../models/operations/listauthserverclientsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## CreateAuthServerClient

Create a new client for a given auth server. The client ID and secret will be generated automatically. The `name` attribute must be unique within the auth server. The `grant_types` and `response_types` attributes define the OAuth 2.0 and OpenID Connect capabilities of the client.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAuthServerClient" method="post" path="/v1/auth-servers/{authServerId}/clients" -->
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

    res, err := s.AuthServerClients.CreateAuthServerClient(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateClient{
        Name: "<value>",
        GrantTypes: []components.AuthServerGrantType{
            components.AuthServerGrantTypeUrnIetfParamsOauthGrantTypeJwtBearer,
        },
        ResponseTypes: []components.ResponseType{},
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatedClient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `authServerID`                                                     | `string`                                                           | :heavy_check_mark:                                                 | The auth server ID                                                 | d32d905a-ed33-46a3-a093-d8f536af9a8a                               |
| `createClient`                                                     | [components.CreateClient](../../models/components/createclient.md) | :heavy_check_mark:                                                 | Client to be created                                               |                                                                    |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[*operations.CreateAuthServerClientResponse](../../models/operations/createauthserverclientresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetAuthServerClient

Get a client for given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthServerClient" method="get" path="/v1/auth-servers/{authServerId}/clients/{clientId}" -->
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

    res, err := s.AuthServerClients.GetAuthServerClient(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "kYa9iQFU5xPDSIUH9z1z")
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerClient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | The OAuth 2.0 client ID                                  | kYa9iQFU5xPDSIUH9z1z                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAuthServerClientResponse](../../models/operations/getauthserverclientresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## UpdateAuthServerClient

Update a client for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthServerClient" method="patch" path="/v1/auth-servers/{authServerId}/clients/{clientId}" -->
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

    res, err := s.AuthServerClients.UpdateAuthServerClient(ctx, operations.UpdateAuthServerClientRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        UpdateClient: components.UpdateClient{
            Labels: map[string]*string{
                "env": sdkkonnectgo.Pointer("test"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerClient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateAuthServerClientRequest](../../models/operations/updateauthserverclientrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateAuthServerClientResponse](../../models/operations/updateauthserverclientresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ReplaceAuthServerClient

Replace a client for a given auth server. The client ID and secret can be provided to update an existing client, or to create a client with the matching ID and secret.

### Example Usage

<!-- UsageSnippet language="go" operationID="replaceAuthServerClient" method="put" path="/v1/auth-servers/{authServerId}/clients/{clientId}" -->
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

    res, err := s.AuthServerClients.ReplaceAuthServerClient(ctx, operations.ReplaceAuthServerClientRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        ReplaceClient: components.ReplaceClient{
            Name: "<value>",
            ClientSecret: "<value>",
            GrantTypes: []components.AuthServerGrantType{
                components.AuthServerGrantTypeClientCredentials,
            },
            ResponseTypes: []components.ResponseType{
                components.ResponseTypeToken,
            },
            Labels: map[string]string{
                "env": "test",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerClient != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ReplaceAuthServerClientRequest](../../models/operations/replaceauthserverclientrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ReplaceAuthServerClientResponse](../../models/operations/replaceauthserverclientresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeleteAuthServerClient

Delete a client for a given auth server. All tokens associated with the client will also be deleted. This action is irreversible.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAuthServerClient" method="delete" path="/v1/auth-servers/{authServerId}/clients/{clientId}" -->
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

    res, err := s.AuthServerClients.DeleteAuthServerClient(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "kYa9iQFU5xPDSIUH9z1z")
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
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | The OAuth 2.0 client ID                                  | kYa9iQFU5xPDSIUH9z1z                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAuthServerClientResponse](../../models/operations/deleteauthserverclientresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## ListClientSecrets

List all the client secrets for a given client.

### Example Usage

<!-- UsageSnippet language="go" operationID="listClientSecrets" method="get" path="/v1/auth-servers/{authServerId}/clients/{clientId}/secrets" -->
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

    res, err := s.AuthServerClients.ListClientSecrets(ctx, operations.ListClientSecretsRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientSecretList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListClientSecretsRequest](../../models/operations/listclientsecretsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListClientSecretsResponse](../../models/operations/listclientsecretsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateClientSecret

Create a new client secret for a given client. Multiple secrets can be active simultaneously for a single client.

### Example Usage

<!-- UsageSnippet language="go" operationID="createClientSecret" method="post" path="/v1/auth-servers/{authServerId}/clients/{clientId}/secrets" -->
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

    res, err := s.AuthServerClients.CreateClientSecret(ctx, operations.CreateClientSecretRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        CreateSecretPayload: components.CreateSecretPayload{
            Secret: sdkkonnectgo.Pointer("YAzsyUlNZ5gNGeKS9H3VAdxVPzhPo4ae"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientSecretItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateClientSecretRequest](../../models/operations/createclientsecretrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateClientSecretResponse](../../models/operations/createclientsecretresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.ForbiddenError  | 403                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetClientSecret

Get a specific client secret for a given client.

### Example Usage

<!-- UsageSnippet language="go" operationID="getClientSecret" method="get" path="/v1/auth-servers/{authServerId}/clients/{clientId}/secrets/{clientSecretId}" -->
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

    res, err := s.AuthServerClients.GetClientSecret(ctx, operations.GetClientSecretRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        ClientSecretID: "07d05309-45cc-4b37-92fb-1524846deec3",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientSecretItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetClientSecretRequest](../../models/operations/getclientsecretrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetClientSecretResponse](../../models/operations/getclientsecretresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## UpdateClientSecret

Update a client secret for a given client, such as enabling or disabling the secret.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateClientSecret" method="patch" path="/v1/auth-servers/{authServerId}/clients/{clientId}/secrets/{clientSecretId}" -->
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

    res, err := s.AuthServerClients.UpdateClientSecret(ctx, operations.UpdateClientSecretRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        ClientSecretID: "07d05309-45cc-4b37-92fb-1524846deec3",
        UpdateSecretPayload: components.UpdateSecretPayload{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClientSecretItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateClientSecretRequest](../../models/operations/updateclientsecretrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateClientSecretResponse](../../models/operations/updateclientsecretresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeleteClientSecret

Delete a client secret for a given client. This action is irreversible.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteClientSecret" method="delete" path="/v1/auth-servers/{authServerId}/clients/{clientId}/secrets/{clientSecretId}" -->
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

    res, err := s.AuthServerClients.DeleteClientSecret(ctx, operations.DeleteClientSecretRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        ClientSecretID: "07d05309-45cc-4b37-92fb-1524846deec3",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteClientSecretRequest](../../models/operations/deleteclientsecretrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeleteClientSecretResponse](../../models/operations/deleteclientsecretresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## TestClaimForClient

Test a claim template for a client. Utility function to test claim templates against a client and verify the resulting claim values.

### Example Usage

<!-- UsageSnippet language="go" operationID="testClaimForClient" method="post" path="/v1/auth-servers/{authServerId}/clients/{clientId}/test-claim" -->
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

    res, err := s.AuthServerClients.TestClaimForClient(ctx, operations.TestClaimForClientRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        TestClaimPayload: components.TestClaimPayload{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TestClaimResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TestClaimForClientRequest](../../models/operations/testclaimforclientrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.TestClaimForClientResponse](../../models/operations/testclaimforclientresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListAuthServerClientTokens

List all the tokens for a given client.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAuthServerClientTokens" method="get" path="/v1/auth-servers/{authServerId}/clients/{clientId}/tokens" -->
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

    res, err := s.AuthServerClients.ListAuthServerClientTokens(ctx, operations.ListAuthServerClientTokensRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TokenList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListAuthServerClientTokensRequest](../../models/operations/listauthserverclienttokensrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListAuthServerClientTokensResponse](../../models/operations/listauthserverclienttokensresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## DeleteAllAuthServerClientTokens

Revoke all tokens for a given client. Revoked tokens can no longer be used to access resources when introspection is enabled on the client.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAllAuthServerClientTokens" method="delete" path="/v1/auth-servers/{authServerId}/clients/{clientId}/tokens" -->
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

    res, err := s.AuthServerClients.DeleteAllAuthServerClientTokens(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "kYa9iQFU5xPDSIUH9z1z")
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
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `clientID`                                               | `string`                                                 | :heavy_check_mark:                                       | The OAuth 2.0 client ID                                  | kYa9iQFU5xPDSIUH9z1z                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAllAuthServerClientTokensResponse](../../models/operations/deleteallauthserverclienttokensresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetAuthServerClientToken

Get a token for given client and auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthServerClientToken" method="get" path="/v1/auth-servers/{authServerId}/clients/{clientId}/tokens/{tokenId}" -->
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

    res, err := s.AuthServerClients.GetAuthServerClientToken(ctx, operations.GetAuthServerClientTokenRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        TokenID: "2cec70a4-d197-495c-8f3e-c0a1199d1da4",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Token != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetAuthServerClientTokenRequest](../../models/operations/getauthserverclienttokenrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetAuthServerClientTokenResponse](../../models/operations/getauthserverclienttokenresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## DeleteAuthServerClientToken

Revoke a token for a given client and auth server. Revoked tokens can no longer be used to access resources when introspection is enabled on the client.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAuthServerClientToken" method="delete" path="/v1/auth-servers/{authServerId}/clients/{clientId}/tokens/{tokenId}" -->
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

    res, err := s.AuthServerClients.DeleteAuthServerClientToken(ctx, operations.DeleteAuthServerClientTokenRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        ClientID: "kYa9iQFU5xPDSIUH9z1z",
        TokenID: "2cec70a4-d197-495c-8f3e-c0a1199d1da4",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteAuthServerClientTokenRequest](../../models/operations/deleteauthserverclienttokenrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.DeleteAuthServerClientTokenResponse](../../models/operations/deleteauthserverclienttokenresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |