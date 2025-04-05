# Certificates
(*Certificates*)

## Overview

A certificate object represents a public certificate, and can be optionally paired with the corresponding private key. These objects are used by Kong Gateway to handle SSL/TLS termination for encrypted requests, or for use as a trusted CA store when validating peer certificate of client/service.
<br><br>
Certificates are optionally associated with SNI objects to tie a cert/key pair to one or more hostnames.
<br><br>
If intermediate certificates are required in addition to the main certificate, they should be concatenated together into one string.


### Available Operations

* [ListCertificate](#listcertificate) - List all Certificates
* [CreateCertificate](#createcertificate) - Create a new Certificate
* [DeleteCertificate](#deletecertificate) - Delete a Certificate
* [GetCertificate](#getcertificate) - Fetch a Certificate
* [UpsertCertificate](#upsertcertificate) - Upsert a Certificate

## ListCertificate

List all Certificates

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

    res, err := s.Certificates.ListCertificate(ctx, operations.ListCertificateRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListCertificateRequest](../../models/operations/listcertificaterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListCertificateResponse](../../models/operations/listcertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateCertificate

Create a new Certificate

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

    res, err := s.Certificates.CreateCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Certificate{
        Cert: "-----BEGIN CERTIFICATE-----\n" +
        "certificate-content\n" +
        "-----END CERTIFICATE-----",
        ID: sdkkonnectgo.String("b2f34145-0343-41a4-9602-4c69dec2f269"),
        Key: "-----BEGIN PRIVATE KEY-----\n" +
        "private-key-content\n" +
        "-----END PRIVATE KEY-----",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Certificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                     | Type                                                                                                                                                                                                                          | Required                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                   | Example                                                                                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                            | The context to use for the request.                                                                                                                                                                                           |                                                                                                                                                                                                                               |
| `controlPlaneID`                                                                                                                                                                                                              | *string*                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                            | The UUID of your control plane. This variable is available in the Konnect manager.                                                                                                                                            | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                                                                                                          |
| `certificate`                                                                                                                                                                                                                 | [components.Certificate](../../models/components/certificate.md)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                            | Description of the new Certificate for creation                                                                                                                                                                               | {<br/>"cert": "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",<br/>"id": "b2f34145-0343-41a4-9602-4c69dec2f269",<br/>"key": "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----"<br/>} |
| `opts`                                                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                            | The options for this request.                                                                                                                                                                                                 |                                                                                                                                                                                                                               |

### Response

**[*operations.CreateCertificateResponse](../../models/operations/createcertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteCertificate

Delete a Certificate

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

    res, err := s.Certificates.DeleteCertificate(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "ddf3cdaa-3329-4961-822a-ce6dbd38eff7")
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
| `certificateID`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Certificate to lookup                                                    | ddf3cdaa-3329-4961-822a-ce6dbd38eff7                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteCertificateResponse](../../models/operations/deletecertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetCertificate

Get a Certificate using ID.

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

    res, err := s.Certificates.GetCertificate(ctx, "ddf3cdaa-3329-4961-822a-ce6dbd38eff7", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Certificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `certificateID`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | ID of the Certificate to lookup                                                    | ddf3cdaa-3329-4961-822a-ce6dbd38eff7                                               |
| `controlPlaneID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetCertificateResponse](../../models/operations/getcertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertCertificate

Create or Update Certificate using ID.

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

    res, err := s.Certificates.UpsertCertificate(ctx, operations.UpsertCertificateRequest{
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Certificate: components.Certificate{
            Cert: "-----BEGIN CERTIFICATE-----\n" +
            "certificate-content\n" +
            "-----END CERTIFICATE-----",
            ID: sdkkonnectgo.String("b2f34145-0343-41a4-9602-4c69dec2f269"),
            Key: "-----BEGIN PRIVATE KEY-----\n" +
            "private-key-content\n" +
            "-----END PRIVATE KEY-----",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Certificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpsertCertificateRequest](../../models/operations/upsertcertificaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpsertCertificateResponse](../../models/operations/upsertcertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |