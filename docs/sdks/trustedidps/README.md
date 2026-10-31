# TrustedIdPs

## Overview

Trusted IdPs configure the external identity providers a token-vault enabled directory trusts when
validating tokens. A directory can trust several IdPs, each with a distinct issuer.


### Available Operations

* [ListTrustedIdps](#listtrustedidps) - List all trusted IdPs
* [CreateTrustedIdp](#createtrustedidp) - Create a new trusted IdP
* [GetTrustedIdp](#gettrustedidp) - Get a trusted IdP
* [UpdateTrustedIdp](#updatetrustedidp) - Update a trusted IdP
* [PatchTrustedIdp](#patchtrustedidp) - Partially update a trusted IdP
* [DeleteTrustedIdp](#deletetrustedidp) - Delete a trusted IdP
* [ListTrustedIdpPrincipals](#listtrustedidpprincipals) - List principals of a trusted IdP

## ListTrustedIdps

List all trusted IdPs for a given directory.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTrustedIdps" method="get" path="/v2/directories/{directoryId}/trusted-idps" -->
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

    res, err := s.TrustedIDPs.ListTrustedIdps(ctx, "7a28bb19-81bd-4a8a-9bf7-3a3c6c0bd583", &components.CursorPageParameters{
        Size: sdkkonnectgo.Pointer[int64](10),
        After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTrustedIdps != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `directoryID`                                                                       | `string`                                                                            | :heavy_check_mark:                                                                  | ID of the directory                                                                 |
| `page`                                                                              | [*components.CursorPageParameters](../../models/components/cursorpageparameters.md) | :heavy_minus_sign:                                                                  | Determines which page of the collection to retrieve.                                |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*operations.ListTrustedIdpsResponse](../../models/operations/listtrustedidpsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateTrustedIdp

Create a trusted IdP for a given directory. The `issuer_url` must be unique within the directory.


### Example Usage

<!-- UsageSnippet language="go" operationID="createTrustedIdp" method="post" path="/v2/directories/{directoryId}/trusted-idps" -->
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

    res, err := s.TrustedIDPs.CreateTrustedIdp(ctx, "f42f2194-e6f7-4697-8ff4-6a5d8c7d2c44", components.CreateTrustedIdpPayload{
        IssuerURL: "https://acme.okta.com/oauth2/default",
        JwksURI: sdkkonnectgo.Pointer("https://acme.okta.com/oauth2/default/v1/keys"),
        Jit: &components.TrustedIdpJit{
            Mapping: &components.TrustedIdpJitMapping{
                Metadata: map[string]string{
                    "department": "department",
                },
                Email: sdkkonnectgo.Pointer("email"),
                Groups: sdkkonnectgo.Pointer("groups"),
                DisplayName: sdkkonnectgo.Pointer("name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrustedIdp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `directoryID`                                                                            | `string`                                                                                 | :heavy_check_mark:                                                                       | ID of the directory                                                                      |
| `createTrustedIdpPayload`                                                                | [components.CreateTrustedIdpPayload](../../models/components/createtrustedidppayload.md) | :heavy_check_mark:                                                                       | Trusted IdP to be created                                                                |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateTrustedIdpResponse](../../models/operations/createtrustedidpresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetTrustedIdp

Get a trusted IdP for a directory.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTrustedIdp" method="get" path="/v2/directories/{directoryId}/trusted-idps/{trustedIdpId}" -->
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

    res, err := s.TrustedIDPs.GetTrustedIdp(ctx, "40c7d944-42f2-47b5-bf90-f17e82e85b39", "74c5bad7-b186-4687-9ad2-f297b4d5df29")
    if err != nil {
        log.Fatal(err)
    }
    if res.TrustedIdp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `trustedIdpID`                                           | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the trusted idp for the vault       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTrustedIdpResponse](../../models/operations/gettrustedidpresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateTrustedIdp

Replace a trusted IdP for a given directory. The request body is the full representation; omitted optional fields are reset to their defaults.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateTrustedIdp" method="put" path="/v2/directories/{directoryId}/trusted-idps/{trustedIdpId}" -->
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

    res, err := s.TrustedIDPs.UpdateTrustedIdp(ctx, operations.UpdateTrustedIdpRequest{
        DirectoryID: "c17c9189-718a-4343-8a8b-08f3fa35146e",
        TrustedIdpID: "1b5d60be-286b-444b-bb4d-cd6f56baf5b1",
        ReplaceTrustedIdpBody: components.ReplaceTrustedIdpBody{
            IssuerURL: "https://acme.okta.com/oauth2/default",
            JwksURI: sdkkonnectgo.Pointer("https://acme.okta.com/oauth2/default/v1/keys"),
            Jit: &components.TrustedIdpJit{
                Mapping: &components.TrustedIdpJitMapping{
                    Metadata: map[string]string{
                        "department": "department",
                    },
                    Email: sdkkonnectgo.Pointer("email"),
                    Groups: sdkkonnectgo.Pointer("groups"),
                    DisplayName: sdkkonnectgo.Pointer("name"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrustedIdp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateTrustedIdpRequest](../../models/operations/updatetrustedidprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateTrustedIdpResponse](../../models/operations/updatetrustedidpresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchTrustedIdp

Partially updates a trusted IdP. Only the fields present in the request body are changed; omitted fields keep their current values. This is the only way to modify `vault_access_enabled`.


### Example Usage

<!-- UsageSnippet language="go" operationID="patchTrustedIdp" method="patch" path="/v2/directories/{directoryId}/trusted-idps/{trustedIdpId}" -->
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

    res, err := s.TrustedIDPs.PatchTrustedIdp(ctx, operations.PatchTrustedIdpRequest{
        DirectoryID: "c7b45f73-9402-4a9e-8366-ca30b6d201cd",
        TrustedIdpID: "cd3f8a56-28ba-49e1-b0c2-304576a4c09f",
        UpdateTrustedIdpBody: components.UpdateTrustedIdpBody{
            IssuerURL: sdkkonnectgo.Pointer("https://acme.okta.com/oauth2/default"),
            JwksURI: sdkkonnectgo.Pointer("https://acme.okta.com/oauth2/default/v1/keys"),
            Jit: &components.TrustedIdpJit{
                Mapping: &components.TrustedIdpJitMapping{
                    Metadata: map[string]string{
                        "department": "department",
                    },
                    Email: sdkkonnectgo.Pointer("email"),
                    Groups: sdkkonnectgo.Pointer("groups"),
                    DisplayName: sdkkonnectgo.Pointer("name"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrustedIdp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchTrustedIdpRequest](../../models/operations/patchtrustedidprequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchTrustedIdpResponse](../../models/operations/patchtrustedidpresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteTrustedIdp

Delete a trusted IdP for a given directory.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTrustedIdp" method="delete" path="/v2/directories/{directoryId}/trusted-idps/{trustedIdpId}" -->
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

    res, err := s.TrustedIDPs.DeleteTrustedIdp(ctx, "b261b26b-3847-44af-9b8b-c31f267d7088", "9ead0b52-75a7-4863-913f-224362c2878c")
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
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `trustedIdpID`                                           | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the trusted idp for the vault       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteTrustedIdpResponse](../../models/operations/deletetrustedidpresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListTrustedIdpPrincipals

Retrieves a paginated list of principals whose identity is linked to this trusted IdP.


### Example Usage

<!-- UsageSnippet language="go" operationID="listTrustedIdpPrincipals" method="get" path="/v2/directories/{directoryId}/trusted-idps/{trustedIdpId}/principals" -->
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

    res, err := s.TrustedIDPs.ListTrustedIdpPrincipals(ctx, operations.ListTrustedIdpPrincipalsRequest{
        DirectoryID: "8d928b9e-9939-480d-9c46-10a6a17a8dab",
        TrustedIdpID: "7d8708e1-2b46-4d4c-9d69-e0f793701069",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongPrincipals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListTrustedIdpPrincipalsRequest](../../models/operations/listtrustedidpprincipalsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListTrustedIdpPrincipalsResponse](../../models/operations/listtrustedidpprincipalsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |