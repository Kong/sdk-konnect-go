# CACertificates
(*CACertificates*)

## Overview

A CA certificate object represents a trusted certificate authority. 
These objects are used by Kong Gateway to verify the validity of a client or server certificate.

### Available Operations

* [ListCaCertificate](#listcacertificate) - List all CA Certificates
* [CreateCaCertificate](#createcacertificate) - Create a new CA Certificate
* [DeleteCaCertificate](#deletecacertificate) - Delete a CA Certificate
* [GetCaCertificate](#getcacertificate) - Fetch a CA Certificate
* [UpsertCaCertificate](#upsertcacertificate) - Upsert a CA Certificate

## ListCaCertificate

List all CA Certificates

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.CACertificates.ListCaCertificate(ctx, operations.ListCaCertificateRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListCaCertificateRequest](../../models/operations/listcacertificaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListCaCertificateResponse](../../models/operations/listcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateCaCertificate

Create a new CA Certificate

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.CACertificates.CreateCaCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CACertificateInput{
        Cert: "-----BEGIN CERTIFICATE-----\n" +
        "certificate-content\n" +
        "-----END CERTIFICATE-----",
        ID: sdkkonnectgo.String("b2f34145-0343-41a4-9602-4c69dec2f260"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CACertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                               | Type                                                                                                                                    | Required                                                                                                                                | Description                                                                                                                             | Example                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                   | :heavy_check_mark:                                                                                                                      | The context to use for the request.                                                                                                     |                                                                                                                                         |
| `controlPlaneID`                                                                                                                        | *string*                                                                                                                                | :heavy_check_mark:                                                                                                                      | The UUID of your control plane. This variable is available in the Konnect manager.                                                      | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                    |
| `caCertificate`                                                                                                                         | [components.CACertificateInput](../../models/components/cacertificateinput.md)                                                          | :heavy_check_mark:                                                                                                                      | Description of the new CA Certificate for creation                                                                                      | {<br/>"cert": "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",<br/>"id": "b2f34145-0343-41a4-9602-4c69dec2f260"<br/>} |
| `opts`                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                | :heavy_minus_sign:                                                                                                                      | The options for this request.                                                                                                           |                                                                                                                                         |

### Response

**[*operations.CreateCaCertificateResponse](../../models/operations/createcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteCaCertificate

Delete a CA Certificate

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.CACertificates.DeleteCaCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "3c31f18a-f27a-4f9b-8cd4-bf841554612f")
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
| `caCertificateID`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the CA Certificate to lookup                                                 | 3c31f18a-f27a-4f9b-8cd4-bf841554612f                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteCaCertificateResponse](../../models/operations/deletecacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetCaCertificate

Get a CA Certificate using ID.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.CACertificates.GetCaCertificate(ctx, "3c31f18a-f27a-4f9b-8cd4-bf841554612f", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.CACertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `caCertificateID`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the CA Certificate to lookup                                                 | 3c31f18a-f27a-4f9b-8cd4-bf841554612f                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetCaCertificateResponse](../../models/operations/getcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertCaCertificate

Create or Update CA Certificate using ID.

### Example Usage

```go
package main

import(
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.CACertificates.UpsertCaCertificate(ctx, operations.UpsertCaCertificateRequest{
        CACertificateID: "3c31f18a-f27a-4f9b-8cd4-bf841554612f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CACertificate: components.CACertificateInput{
            Cert: "-----BEGIN CERTIFICATE-----\n" +
            "certificate-content\n" +
            "-----END CERTIFICATE-----",
            ID: sdkkonnectgo.String("b2f34145-0343-41a4-9602-4c69dec2f260"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CACertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpsertCaCertificateRequest](../../models/operations/upsertcacertificaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpsertCaCertificateResponse](../../models/operations/upsertcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |