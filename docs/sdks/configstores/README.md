# ConfigStores

## Overview

Config Stores

### Available Operations

* [ListConfigStores](#listconfigstores) - List all config stores for a control plane
* [CreateConfigStore](#createconfigstore) - Create Config Store
* [GetConfigStore](#getconfigstore) - Get a Config Store
* [UpdateConfigStore](#updateconfigstore) - Update an individual Config Store
* [DeleteConfigStore](#deleteconfigstore) - Delete Config Store
* [ListConfigStoresInWorkspace](#listconfigstoresinworkspace) - List all Config Stores for a workspace
* [CreateConfigStoreInWorkspace](#createconfigstoreinworkspace) - Create Config Store in a workspace
* [GetConfigStoreInWorkspace](#getconfigstoreinworkspace) - Get a Config Store in a workspace
* [UpdateConfigStoreInWorkspace](#updateconfigstoreinworkspace) - Update a Config Store in a workspace
* [DeleteConfigStoreInWorkspace](#deleteconfigstoreinworkspace) - Delete Config Store in a workspace

## ListConfigStores

List all config stores for a control plane

### Example Usage

<!-- UsageSnippet language="go" operationID="list-config-stores" method="get" path="/v2/control-planes/{controlPlaneId}/config-stores" example="Config Stores" -->
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

    res, err := s.ConfigStores.ListConfigStores(ctx, operations.ListConfigStoresRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListConfigStoresResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListConfigStoresRequest](../../models/operations/listconfigstoresrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListConfigStoresResponse](../../models/operations/listconfigstoresresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateConfigStore

Create a Config Store

### Example Usage: ConfigStoreNameBadRequestExample

<!-- UsageSnippet language="go" operationID="create-config-store" method="post" path="/v2/control-planes/{controlPlaneId}/config-stores" example="ConfigStoreNameBadRequestExample" -->
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

    res, err := s.ConfigStores.CreateConfigStore(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreateConfigStore{
        Name: sdkkonnectgo.Pointer("Config Store"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: Create Config Store Response

<!-- UsageSnippet language="go" operationID="create-config-store" method="post" path="/v2/control-planes/{controlPlaneId}/config-stores" example="Create Config Store Response" -->
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

    res, err := s.ConfigStores.CreateConfigStore(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreateConfigStore{
        Name: sdkkonnectgo.Pointer("Config Store"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: CreateConfigStoreRequestExample

<!-- UsageSnippet language="go" operationID="create-config-store" method="post" path="/v2/control-planes/{controlPlaneId}/config-stores" example="CreateConfigStoreRequestExample" -->
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

    res, err := s.ConfigStores.CreateConfigStore(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreateConfigStore{
        Name: sdkkonnectgo.Pointer("Config Store"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="create-config-store" method="post" path="/v2/control-planes/{controlPlaneId}/config-stores" example="UnauthorizedExample" -->
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

    res, err := s.ConfigStores.CreateConfigStore(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreateConfigStore{
        Name: sdkkonnectgo.Pointer("Config Store"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="create-config-store" method="post" path="/v2/control-planes/{controlPlaneId}/config-stores" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.ConfigStores.CreateConfigStore(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreateConfigStore{
        Name: sdkkonnectgo.Pointer("Config Store"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `createConfigStore`                                                                | [components.CreateConfigStore](../../models/components/createconfigstore.md)       | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreateConfigStoreResponse](../../models/operations/createconfigstoreresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## GetConfigStore

Returns a Config Store

### Example Usage

<!-- UsageSnippet language="go" operationID="get-config-store" method="get" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" example="Create Config Store Response" -->
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

    res, err := s.ConfigStores.GetConfigStore(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `configStoreID`                                                                    | `string`                                                                           | :heavy_check_mark:                                                                 | Config Store identifier                                                            | d32d905a-ed33-46a3-a093-d8f536af9a8a                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetConfigStoreResponse](../../models/operations/getconfigstoreresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateConfigStore

Updates a Config Store

### Example Usage: ConfigStoreNameBadRequestExample

<!-- UsageSnippet language="go" operationID="update-config-store" method="put" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" example="ConfigStoreNameBadRequestExample" -->
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

    res, err := s.ConfigStores.UpdateConfigStore(ctx, operations.UpdateConfigStoreRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateConfigStore: components.UpdateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: Create Config Store Response

<!-- UsageSnippet language="go" operationID="update-config-store" method="put" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" example="Create Config Store Response" -->
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

    res, err := s.ConfigStores.UpdateConfigStore(ctx, operations.UpdateConfigStoreRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateConfigStore: components.UpdateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: NotFoundExample

<!-- UsageSnippet language="go" operationID="update-config-store" method="put" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" example="NotFoundExample" -->
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

    res, err := s.ConfigStores.UpdateConfigStore(ctx, operations.UpdateConfigStoreRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateConfigStore: components.UpdateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: UnauthorizedExample

<!-- UsageSnippet language="go" operationID="update-config-store" method="put" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" example="UnauthorizedExample" -->
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

    res, err := s.ConfigStores.UpdateConfigStore(ctx, operations.UpdateConfigStoreRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateConfigStore: components.UpdateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: UnsupportedMediaTypeExample

<!-- UsageSnippet language="go" operationID="update-config-store" method="put" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" example="UnsupportedMediaTypeExample" -->
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

    res, err := s.ConfigStores.UpdateConfigStore(ctx, operations.UpdateConfigStoreRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateConfigStore: components.UpdateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```
### Example Usage: UpdateConfigStoreRequestExample

<!-- UsageSnippet language="go" operationID="update-config-store" method="put" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" example="UpdateConfigStoreRequestExample" -->
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

    res, err := s.ConfigStores.UpdateConfigStore(ctx, operations.UpdateConfigStoreRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        UpdateConfigStore: components.UpdateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store with Updated Name"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateConfigStoreRequest](../../models/operations/updateconfigstorerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateConfigStoreResponse](../../models/operations/updateconfigstoreresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## DeleteConfigStore

Removes a config store

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-config-store" method="delete" path="/v2/control-planes/{controlPlaneId}/config-stores/{configStoreId}" -->
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

    res, err := s.ConfigStores.DeleteConfigStore(ctx, operations.DeleteConfigStoreRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteConfigStoreRequest](../../models/operations/deleteconfigstorerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.DeleteConfigStoreResponse](../../models/operations/deleteconfigstoreresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListConfigStoresInWorkspace

Returns a collection of all Config Stores associated with a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-config-stores-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/{workspace}/config-stores" -->
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

    res, err := s.ConfigStores.ListConfigStoresInWorkspace(ctx, operations.ListConfigStoresInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListConfigStoresResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListConfigStoresInWorkspaceRequest](../../models/operations/listconfigstoresinworkspacerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListConfigStoresInWorkspaceResponse](../../models/operations/listconfigstoresinworkspaceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateConfigStoreInWorkspace

Create a Config Store in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-config-store-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/{workspace}/config-stores" -->
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

    res, err := s.ConfigStores.CreateConfigStoreInWorkspace(ctx, operations.CreateConfigStoreInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CreateConfigStore: components.CreateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateConfigStoreInWorkspaceRequest](../../models/operations/createconfigstoreinworkspacerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateConfigStoreInWorkspaceResponse](../../models/operations/createconfigstoreinworkspaceresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## GetConfigStoreInWorkspace

Returns a Config Store in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-config-store-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/{workspace}/config-stores/{configStoreId}" -->
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

    res, err := s.ConfigStores.GetConfigStoreInWorkspace(ctx, operations.GetConfigStoreInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetConfigStoreInWorkspaceRequest](../../models/operations/getconfigstoreinworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetConfigStoreInWorkspaceResponse](../../models/operations/getconfigstoreinworkspaceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateConfigStoreInWorkspace

Updates a Config Store in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-config-store-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/{workspace}/config-stores/{configStoreId}" -->
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

    res, err := s.ConfigStores.UpdateConfigStoreInWorkspace(ctx, operations.UpdateConfigStoreInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        Workspace: "team-payments",
        UpdateConfigStore: components.UpdateConfigStore{
            Name: sdkkonnectgo.Pointer("Config Store"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigStore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateConfigStoreInWorkspaceRequest](../../models/operations/updateconfigstoreinworkspacerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateConfigStoreInWorkspaceResponse](../../models/operations/updateconfigstoreinworkspaceresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## DeleteConfigStoreInWorkspace

Removes a Config Store in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-config-store-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/{workspace}/config-stores/{configStoreId}" -->
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

    res, err := s.ConfigStores.DeleteConfigStoreInWorkspace(ctx, operations.DeleteConfigStoreInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        ConfigStoreID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DeleteConfigStoreInWorkspaceRequest](../../models/operations/deleteconfigstoreinworkspacerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.DeleteConfigStoreInWorkspaceResponse](../../models/operations/deleteconfigstoreinworkspaceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |