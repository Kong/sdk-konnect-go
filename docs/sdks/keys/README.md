# Keys
(*Keys*)

## Overview

A key object holds a representation of asymmetric keys in various formats. When Kong Gateway or a Kong plugin requires a specific public or private key to perform certain operations, it can use this entity.


### Available Operations

* [ListKeyWithKeySet](#listkeywithkeyset) - List all Keys associated with a KeySet
* [CreateKeyWithKeySet](#createkeywithkeyset) - Create a new Key associated with a KeySet
* [DeleteKeyWithKeySet](#deletekeywithkeyset) - Delete a a Key associated with a KeySet
* [GetKeyWithKeySet](#getkeywithkeyset) - Fetch a Key associated with a KeySet
* [UpsertKeyWithKeySet](#upsertkeywithkeyset) - Upsert a Key associated with a KeySet
* [ListKey](#listkey) - List all Keys
* [CreateKey](#createkey) - Create a new Key
* [DeleteKey](#deletekey) - Delete a Key
* [GetKey](#getkey) - Fetch a Key
* [UpsertKey](#upsertkey) - Upsert a Key

## ListKeyWithKeySet

List all Keys associated with a KeySet

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

    res, err := s.Keys.ListKeyWithKeySet(ctx, operations.ListKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
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

Create a new Key associated with a KeySet

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

    res, err := s.Keys.CreateKeyWithKeySet(ctx, operations.CreateKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        KeyWithoutParents: components.KeyWithoutParents{
            ID: sdkkonnectgo.String("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.String("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.String("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.String("-----BEGIN"),
                PublicKey: sdkkonnectgo.String("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.String("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
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

Delete a a Key associated with a KeySet using ID or name.

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

    res, err := s.Keys.DeleteKeyWithKeySet(ctx, operations.DeleteKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
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

Get a Key associated with a KeySet using ID or name.

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

    res, err := s.Keys.GetKeyWithKeySet(ctx, operations.GetKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
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

Create or Update a Key associated with a KeySet using ID or name.

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

    res, err := s.Keys.UpsertKeyWithKeySet(ctx, operations.UpsertKeyWithKeySetRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        KeySetID: "6cc34248-50b4-4a81-9201-3bdf7a83f712",
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        KeyWithoutParents: components.KeyWithoutParents{
            ID: sdkkonnectgo.String("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.String("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.String("a-key"),
            Pem: &components.Pem{
                PrivateKey: sdkkonnectgo.String("-----BEGIN"),
                PublicKey: sdkkonnectgo.String("-----BEGIN"),
            },
            Set: &components.Set{
                ID: sdkkonnectgo.String("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
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

List all Keys

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

    res, err := s.Keys.ListKey(ctx, operations.ListKeyRequest{
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

Create a new Key

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

    res, err := s.Keys.CreateKey(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.KeyInput{
        ID: sdkkonnectgo.String("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
        Jwk: sdkkonnectgo.String("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
        Kid: "42",
        Name: sdkkonnectgo.String("a-key"),
        Pem: &components.KeyPem{
            PrivateKey: sdkkonnectgo.String("-----BEGIN"),
            PublicKey: sdkkonnectgo.String("-----BEGIN"),
        },
        Set: &components.KeySet1{
            ID: sdkkonnectgo.String("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
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

| Parameter                                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                 |
| `controlPlaneID`                                                                                                                                                                                                                                                | *string*                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                              | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                                                                              | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                                                                            |
| `key`                                                                                                                                                                                                                                                           | [components.KeyInput](../../models/components/keyinput.md)                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                              | Description of the new Key for creation                                                                                                                                                                                                                         | {<br/>"id": "d958f66b-8e99-44d2-b0b4-edd5bbf24658",<br/>"jwk": "{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}",<br/>"kid": "42",<br/>"name": "a-key",<br/>"pem": {<br/>"private_key": "-----BEGIN",<br/>"public_key": "-----BEGIN"<br/>},<br/>"set": {<br/>"id": "b86b331c-dcd0-4b3e-97ce-47c5a9543031"<br/>}<br/>} |
| `opts`                                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                 |

### Response

**[*operations.CreateKeyResponse](../../models/operations/createkeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteKey

Delete a Key

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

    res, err := s.Keys.DeleteKey(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "bba22c06-a632-42be-a018-1b9ff357b5b9")
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
| `keyID`                                                                            | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Key to lookup                                                            | bba22c06-a632-42be-a018-1b9ff357b5b9                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteKeyResponse](../../models/operations/deletekeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetKey

Get a Key using ID or name.

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

    res, err := s.Keys.GetKey(ctx, "bba22c06-a632-42be-a018-1b9ff357b5b9", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `keyID`                                                                            | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Key to lookup                                                            | bba22c06-a632-42be-a018-1b9ff357b5b9                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetKeyResponse](../../models/operations/getkeyresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertKey

Create or Update Key using ID or name.

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

    res, err := s.Keys.UpsertKey(ctx, operations.UpsertKeyRequest{
        KeyID: "bba22c06-a632-42be-a018-1b9ff357b5b9",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Key: components.KeyInput{
            ID: sdkkonnectgo.String("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
            Jwk: sdkkonnectgo.String("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
            Kid: "42",
            Name: sdkkonnectgo.String("a-key"),
            Pem: &components.KeyPem{
                PrivateKey: sdkkonnectgo.String("-----BEGIN"),
                PublicKey: sdkkonnectgo.String("-----BEGIN"),
            },
            Set: &components.KeySet1{
                ID: sdkkonnectgo.String("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
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