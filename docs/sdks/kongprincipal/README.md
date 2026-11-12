# KongPrincipal

## Overview

Principals represent identities within a directory that can be authenticated and authorized.
Each principal can have associated metadata and API keys for authentication purposes.


### Available Operations

* [ImportOrgChart](#importorgchart) - Import user principals from a canonical CSV
* [CreatePrincipal](#createprincipal) - Create a Principal in a Directory
* [ListPrincipals](#listprincipals) - List all the Principals of a Directory
* [UpdatePrincipal](#updateprincipal) - Update a Principal
* [ReplacePrincipal](#replaceprincipal) - Replace a Principal
* [GetPrincipal](#getprincipal) - Get a Principal
* [DeletePrincipal](#deleteprincipal) - Delete a Principal
* [ListPrincipalGroups](#listprincipalgroups) - List the groups of a principal

## ImportOrgChart

Imports user principals in bulk from a canonical CSV file. Each row
creates one user principal — tagged hris_import, stamped with the
row's external_id — plus its email custom identity. The optional
cost_center column is stored in the principal's metadata as
cost_center; rows matched to existing principals adopt the file's
value, and a blank cell clears the stored one, while a file without
the column leaves existing cost centers untouched. The file is
validated in full before anything is written; the response is a
report of counts and per-row errors, never the created rows.


### Example Usage

<!-- UsageSnippet language="go" operationID="importOrgChart" method="post" path="/v2/directories/{directoryId}/org-chart/import" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.KongPrincipal.ImportOrgChart(ctx, "b565a055-9789-4949-a616-988a1ec14254", components.ImportPrincipalsRequest{
        File: components.File{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongImportReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `directoryID`                                                                            | `string`                                                                                 | :heavy_check_mark:                                                                       | ID of the directory                                                                      |
| `importPrincipalsRequest`                                                                | [components.ImportPrincipalsRequest](../../models/components/importprincipalsrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ImportOrgChartResponse](../../models/operations/importorgchartresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePrincipal

Creates a new principal identity within the specified directory.
Principals can be configured with metadata and labels for organizational purposes.


### Example Usage

<!-- UsageSnippet language="go" operationID="createPrincipal" method="post" path="/v2/directories/{directoryId}/principals" -->
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

    res, err := s.KongPrincipal.CreatePrincipal(ctx, "6efc4226-e90b-4528-8b65-14a10e1e71ae", components.CreatePrincipalBody{
        Labels: map[string]string{
            "env": "test",
        },
        ManagedBy: map[string]string{
            "owner": "terraform",
        },
        Active: sdkkonnectgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `directoryID`                                                                    | `string`                                                                         | :heavy_check_mark:                                                               | ID of the directory                                                              |
| `createPrincipalBody`                                                            | [components.CreatePrincipalBody](../../models/components/createprincipalbody.md) | :heavy_check_mark:                                                               | Principal to be created                                                          |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreatePrincipalResponse](../../models/operations/createprincipalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPrincipals

Retrieves a paginated list of all principals within the specified directory.
Results can be sorted by display_name, created_at, or updated_at.


### Example Usage

<!-- UsageSnippet language="go" operationID="listPrincipals" method="get" path="/v2/directories/{directoryId}/principals" -->
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

    res, err := s.KongPrincipal.ListPrincipals(ctx, operations.ListPrincipalsRequest{
        DirectoryID: "1f95709d-56eb-496c-9d91-77d5bd915d57",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListPrincipalsRequest](../../models/operations/listprincipalsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListPrincipalsResponse](../../models/operations/listprincipalsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePrincipal

Partially updates an existing principal by modifying only the provided fields.
Unspecified fields will retain their current values.


### Example Usage

<!-- UsageSnippet language="go" operationID="updatePrincipal" method="patch" path="/v2/directories/{directoryId}/principals/{principalId}" -->
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

    res, err := s.KongPrincipal.UpdatePrincipal(ctx, operations.UpdatePrincipalRequest{
        DirectoryID: "f70a780b-3a26-40e2-9e06-ccc359478e7c",
        PrincipalID: "5936947d-8d16-41c1-9de5-d82fcd9a9095",
        UpdatePrincipalBody: components.UpdatePrincipalBody{
            Labels: map[string]string{
                "env": "test",
            },
            ManagedBy: map[string]*string{
                "owner": sdkkonnectgo.Pointer("terraform"),
            },
            Active: sdkkonnectgo.Pointer(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdatePrincipalRequest](../../models/operations/updateprincipalrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdatePrincipalResponse](../../models/operations/updateprincipalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePrincipal

Completely replaces an existing principal with the provided data.
All fields must be specified; any omitted fields will be reset to their defaults.


### Example Usage

<!-- UsageSnippet language="go" operationID="replacePrincipal" method="put" path="/v2/directories/{directoryId}/principals/{principalId}" -->
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

    res, err := s.KongPrincipal.ReplacePrincipal(ctx, operations.ReplacePrincipalRequest{
        DirectoryID: "57853b11-93dc-4d44-bfe1-42bbf6d45a2c",
        PrincipalID: "93010d8b-955f-40c9-84aa-58994f474da7",
        ReplacePrincipalBody: components.ReplacePrincipalBody{
            Labels: map[string]string{
                "env": "test",
            },
            ManagedBy: map[string]string{
                "owner": "terraform",
            },
            Active: sdkkonnectgo.Pointer(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ReplacePrincipalRequest](../../models/operations/replaceprincipalrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ReplacePrincipalResponse](../../models/operations/replaceprincipalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPrincipal

Retrieves detailed information about a specific principal by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPrincipal" method="get" path="/v2/directories/{directoryId}/principals/{principalId}" -->
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

    res, err := s.KongPrincipal.GetPrincipal(ctx, "d9f4f5d9-3748-4d42-abd7-009b59289bda", "4b099dd9-8e28-4467-8220-56decc9f8657")
    if err != nil {
        log.Fatal(err)
    }
    if res.KongPrincipal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `principalID`                                            | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the principal (UUID format)         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPrincipalResponse](../../models/operations/getprincipalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePrincipal

Permanently deletes a principal from the directory.


### Example Usage

<!-- UsageSnippet language="go" operationID="deletePrincipal" method="delete" path="/v2/directories/{directoryId}/principals/{principalId}" -->
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

    res, err := s.KongPrincipal.DeletePrincipal(ctx, "3eeff6d6-e96f-4603-8b24-493017ab2376", "97e8cb58-fa27-4945-aa16-45282c603271")
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
| `principalID`                                            | `string`                                                 | :heavy_check_mark:                                       | Unique identifier of the principal (UUID format)         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePrincipalResponse](../../models/operations/deleteprincipalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPrincipalGroups

Retrieves a paginated list of every group the principal belongs to, of any type, including memberships inherited through nested teams. Results are ordered by when the membership was created, oldest first.


### Example Usage

<!-- UsageSnippet language="go" operationID="listPrincipalGroups" method="get" path="/v2/directories/{directoryId}/principals/{principalId}/groups" -->
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

    res, err := s.KongPrincipal.ListPrincipalGroups(ctx, operations.ListPrincipalGroupsRequest{
        DirectoryID: "1457529a-db52-4b5b-857a-c536ce8607ce",
        PrincipalID: "b896b80a-6ab2-4265-b9ae-c2a40b548fce",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKongPrincipalGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListPrincipalGroupsRequest](../../models/operations/listprincipalgroupsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListPrincipalGroupsResponse](../../models/operations/listprincipalgroupsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |