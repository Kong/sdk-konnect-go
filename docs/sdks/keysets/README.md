# KeySets

## Overview

A JSON Web key set. Key sets are the preferred way to expose keys to plugins because they tell the plugin where to look for keys or have a scoping mechanism to restrict plugins to specific keys.


### Available Operations

* [ListKeySet](#listkeyset) - List all KeySets in a workspace
* [CreateKeySet](#createkeyset) - Create a new KeySet in a workspace
* [DeleteKeySet](#deletekeyset) - Delete a KeySet in a workspace
* [GetKeySet](#getkeyset) - Get a KeySet in a workspace
* [UpsertKeySet](#upsertkeyset) - Upsert a KeySet in a workspace

## ListKeySet

List all KeySets in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-key-set" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets" -->
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

    res, err := s.KeySets.ListKeySet(ctx, operations.ListKeySetRequest{
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListKeySetRequest](../../models/operations/listkeysetrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListKeySetResponse](../../models/operations/listkeysetresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateKeySet

Create a new KeySet in a workspace

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-key-set" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets" example="DuplicateApiKey" -->
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

    res, err := s.KeySets.CreateKeySet(ctx, operations.CreateKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        KeySet: &components.KeySet{
            ID: sdkkonnectgo.Pointer("b58c7d9d-e54f-444c-b24d-cdfc4159f61e"),
            Name: sdkkonnectgo.Pointer("example-key-set"),
            Tags: []string{
                "idp-keys",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KeySet != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-key-set" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets" example="InvalidAuthCred" -->
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

    res, err := s.KeySets.CreateKeySet(ctx, operations.CreateKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        KeySet: &components.KeySet{
            ID: sdkkonnectgo.Pointer("b58c7d9d-e54f-444c-b24d-cdfc4159f61e"),
            Name: sdkkonnectgo.Pointer("example-key-set"),
            Tags: []string{
                "idp-keys",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KeySet != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-key-set" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets" example="NoAPIKey" -->
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

    res, err := s.KeySets.CreateKeySet(ctx, operations.CreateKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        KeySet: &components.KeySet{
            ID: sdkkonnectgo.Pointer("b58c7d9d-e54f-444c-b24d-cdfc4159f61e"),
            Name: sdkkonnectgo.Pointer("example-key-set"),
            Tags: []string{
                "idp-keys",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KeySet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateKeySetRequest](../../models/operations/createkeysetrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateKeySetResponse](../../models/operations/createkeysetresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteKeySet

Delete a KeySet in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-key-set" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}" -->
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

    res, err := s.KeySets.DeleteKeySet(ctx, operations.DeleteKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteKeySetRequest](../../models/operations/deletekeysetrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.DeleteKeySetResponse](../../models/operations/deletekeysetresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetKeySet

Get a KeySet using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-key-set" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}" -->
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

    res, err := s.KeySets.GetKeySet(ctx, operations.GetKeySetRequest{
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KeySet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetKeySetRequest](../../models/operations/getkeysetrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetKeySetResponse](../../models/operations/getkeysetresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertKeySet

Create or Update KeySet using ID or name in a workspace.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-key-set" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}" example="DuplicateApiKey" -->
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

    res, err := s.KeySets.UpsertKeySet(ctx, operations.UpsertKeySetRequest{
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        KeySet: &components.KeySet{
            ID: sdkkonnectgo.Pointer("b58c7d9d-e54f-444c-b24d-cdfc4159f61e"),
            Name: sdkkonnectgo.Pointer("example-key-set"),
            Tags: []string{
                "idp-keys",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KeySet != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-key-set" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}" example="InvalidAuthCred" -->
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

    res, err := s.KeySets.UpsertKeySet(ctx, operations.UpsertKeySetRequest{
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        KeySet: &components.KeySet{
            ID: sdkkonnectgo.Pointer("b58c7d9d-e54f-444c-b24d-cdfc4159f61e"),
            Name: sdkkonnectgo.Pointer("example-key-set"),
            Tags: []string{
                "idp-keys",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KeySet != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-key-set" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}" example="NoAPIKey" -->
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

    res, err := s.KeySets.UpsertKeySet(ctx, operations.UpsertKeySetRequest{
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        KeySet: &components.KeySet{
            ID: sdkkonnectgo.Pointer("b58c7d9d-e54f-444c-b24d-cdfc4159f61e"),
            Name: sdkkonnectgo.Pointer("example-key-set"),
            Tags: []string{
                "idp-keys",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KeySet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpsertKeySetRequest](../../models/operations/upsertkeysetrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.UpsertKeySetResponse](../../models/operations/upsertkeysetresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |