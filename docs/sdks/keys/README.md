# Keys

## Overview

A key object holds a representation of asymmetric keys in various formats. When Kong Gateway or a Kong plugin requires a specific public or private key to perform certain operations, it can use this entity.


### Available Operations

* [ListKeyWithKeySet](#listkeywithkeyset) - List all Keys associated with a KeySet in a workspace
* [CreateKeyWithKeySet](#createkeywithkeyset) - Create a new Key associated with a KeySet in a workspace
* [DeleteKeyWithKeySet](#deletekeywithkeyset) - Delete a a Key associated with a KeySet in a workspace
* [GetKeyWithKeySet](#getkeywithkeyset) - Get a Key associated with a KeySet in a workspace
* [UpsertKeyWithKeySet](#upsertkeywithkeyset) - Upsert a Key associated with a KeySet in a workspace
* [ListKey](#listkey) - List all Keys in a workspace
* [CreateKey](#createkey) - Create a new Key in a workspace
* [DeleteKey](#deletekey) - Delete a Key in a workspace
* [GetKey](#getkey) - Get a Key in a workspace
* [UpsertKey](#upsertkey) - Upsert a Key in a workspace

## ListKeyWithKeySet

List all Keys associated with a KeySet in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-key-with-key-set" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}/keys" -->
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

    res, err := s.Keys.ListKeyWithKeySet(ctx, operations.ListKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListKeyWithKeySetRequest](../../models/operations/listkeywithkeysetrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListKeyWithKeySetResponse](../../models/operations/listkeywithkeysetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateKeyWithKeySet

Create a new Key associated with a KeySet in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-key-with-key-set" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}/keys" -->
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

    res, err := s.Keys.CreateKeyWithKeySet(ctx, operations.CreateKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        Workspace: "team-payments",
        KeyWithoutParents: components.KeyWithoutParents{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.KeyWithoutParentsPem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.KeyWithoutParentsSet{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateKeyWithKeySetRequest](../../models/operations/createkeywithkeysetrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateKeyWithKeySetResponse](../../models/operations/createkeywithkeysetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteKeyWithKeySet

Delete a a Key associated with a KeySet using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-key-with-key-set" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}/keys/{KeyId}" -->
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

    res, err := s.Keys.DeleteKeyWithKeySet(ctx, operations.DeleteKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteKeyWithKeySetRequest](../../models/operations/deletekeywithkeysetrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.DeleteKeyWithKeySetResponse](../../models/operations/deletekeywithkeysetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetKeyWithKeySet

Get a Key associated with a KeySet using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-key-with-key-set" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}/keys/{KeyId}" -->
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

    res, err := s.Keys.GetKeyWithKeySet(ctx, operations.GetKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetKeyWithKeySetRequest](../../models/operations/getkeywithkeysetrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetKeyWithKeySetResponse](../../models/operations/getkeywithkeysetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertKeyWithKeySet

Create or Update a Key associated with a KeySet using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-key-with-key-set" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/key-sets/{KeySetId}/keys/{KeyId}" -->
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

    res, err := s.Keys.UpsertKeyWithKeySet(ctx, operations.UpsertKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        Workspace: "team-payments",
        KeyWithoutParents: components.KeyWithoutParents{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.KeyWithoutParentsPem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.KeyWithoutParentsSet{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpsertKeyWithKeySetRequest](../../models/operations/upsertkeywithkeysetrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpsertKeyWithKeySetResponse](../../models/operations/upsertkeywithkeysetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListKey

List all Keys in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-key" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys" -->
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

    res, err := s.Keys.ListKey(ctx, operations.ListKeyRequest{
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.ListKeyRequest](../../models/operations/listkeyrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.ListKeyResponse](../../models/operations/listkeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateKey

Create a new Key in a workspace

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-key" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys" example="DuplicateApiKey" -->
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

    res, err := s.Keys.CreateKey(ctx, operations.CreateKeyRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Key: components.Key{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-key" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys" example="InvalidAuthCred" -->
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

    res, err := s.Keys.CreateKey(ctx, operations.CreateKeyRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Key: components.Key{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-key" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys" example="NoAPIKey" -->
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

    res, err := s.Keys.CreateKey(ctx, operations.CreateKeyRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Key: components.Key{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.CreateKeyRequest](../../models/operations/createkeyrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateKeyResponse](../../models/operations/createkeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteKey

Delete a Key in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-key" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys/{KeyId}" -->
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

    res, err := s.Keys.DeleteKey(ctx, operations.DeleteKeyRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DeleteKeyRequest](../../models/operations/deletekeyrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.DeleteKeyResponse](../../models/operations/deletekeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetKey

Get a Key using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-key" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys/{KeyId}" -->
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

    res, err := s.Keys.GetKey(ctx, operations.GetKeyRequest{
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [operations.GetKeyRequest](../../models/operations/getkeyrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.GetKeyResponse](../../models/operations/getkeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertKey

Create or Update Key using ID or name in a workspace.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-key" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys/{KeyId}" example="DuplicateApiKey" -->
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

    res, err := s.Keys.UpsertKey(ctx, operations.UpsertKeyRequest{
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Key: components.Key{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-key" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys/{KeyId}" example="InvalidAuthCred" -->
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

    res, err := s.Keys.UpsertKey(ctx, operations.UpsertKeyRequest{
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Key: components.Key{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-key" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/keys/{KeyId}" example="NoAPIKey" -->
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

    res, err := s.Keys.UpsertKey(ctx, operations.UpsertKeyRequest{
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Key: components.Key{
            ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.Pointer("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.UpsertKeyRequest](../../models/operations/upsertkeyrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.UpsertKeyResponse](../../models/operations/upsertkeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |