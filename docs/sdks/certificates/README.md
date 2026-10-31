# Certificates

## Overview

A certificate object represents a public certificate, and can be optionally paired with the corresponding private key. These objects are used by Kong Gateway to handle SSL/TLS termination for encrypted requests, or for use as a trusted CA store when validating peer certificate of client/service.
<br><br>
Certificates are optionally associated with SNI objects to tie a cert/key pair to one or more hostnames.
<br><br>
If intermediate certificates are required in addition to the main certificate, they should be concatenated together into one string.


### Available Operations

* [ListCertificate](#listcertificate) - List all Certificates in a workspace
* [CreateCertificate](#createcertificate) - Create a new Certificate in a workspace
* [DeleteCertificate](#deletecertificate) - Delete a Certificate in a workspace
* [GetCertificate](#getcertificate) - Get a Certificate in a workspace
* [UpsertCertificate](#upsertcertificate) - Upsert a Certificate in a workspace

## ListCertificate

List all Certificates in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-certificate" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates" -->
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

    res, err := s.Certificates.ListCertificate(ctx, operations.ListCertificateRequest{
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

Create a new Certificate in a workspace

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-certificate" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates" example="DuplicateApiKey" -->
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

    res, err := s.Certificates.CreateCertificate(ctx, operations.CreateCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CertificateRequest: components.CreateCertificateRequestCertificateWithCertAndKey(
            components.CertificateWithCertAndKey{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
                Key: "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Certificate != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-certificate" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates" example="InvalidAuthCred" -->
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

    res, err := s.Certificates.CreateCertificate(ctx, operations.CreateCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CertificateRequest: components.CreateCertificateRequestCertificateWithCertAndKey(
            components.CertificateWithCertAndKey{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
                Key: "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Certificate != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-certificate" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates" example="NoAPIKey" -->
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

    res, err := s.Certificates.CreateCertificate(ctx, operations.CreateCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CertificateRequest: components.CreateCertificateRequestCertificateWithCertAndKey(
            components.CertificateWithCertAndKey{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
                Key: "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----",
            },
        ),
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
| `request`                                                                                  | [operations.CreateCertificateRequest](../../models/operations/createcertificaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateCertificateResponse](../../models/operations/createcertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteCertificate

Delete a Certificate in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-certificate" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}" -->
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

    res, err := s.Certificates.DeleteCertificate(ctx, operations.DeleteCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteCertificateRequest](../../models/operations/deletecertificaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.DeleteCertificateResponse](../../models/operations/deletecertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetCertificate

Get a Certificate using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-certificate" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}" -->
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

    res, err := s.Certificates.GetCertificate(ctx, operations.GetCertificateRequest{
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetCertificateRequest](../../models/operations/getcertificaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetCertificateResponse](../../models/operations/getcertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertCertificate

Create or Update Certificate using ID in a workspace.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-certificate" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}" example="DuplicateApiKey" -->
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

    res, err := s.Certificates.UpsertCertificate(ctx, operations.UpsertCertificateRequest{
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CertificateRequest: components.CreateCertificateRequestCertificateWithCertAndKey(
            components.CertificateWithCertAndKey{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
                Key: "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Certificate != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-certificate" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}" example="InvalidAuthCred" -->
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

    res, err := s.Certificates.UpsertCertificate(ctx, operations.UpsertCertificateRequest{
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CertificateRequest: components.CreateCertificateRequestCertificateWithCertAndKey(
            components.CertificateWithCertAndKey{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
                Key: "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Certificate != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-certificate" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}" example="NoAPIKey" -->
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

    res, err := s.Certificates.UpsertCertificate(ctx, operations.UpsertCertificateRequest{
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CertificateRequest: components.CreateCertificateRequestCertificateWithCertAndKey(
            components.CertificateWithCertAndKey{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
                Key: "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----",
            },
        ),
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