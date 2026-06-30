# SNIs

## Overview

An SNI object represents a many-to-one mapping of hostnames to a certificate.
<br><br>
A certificate object can have many hostnames associated with it. When Kong Gateway receives an SSL request, it uses the SNI field in the Client Hello to look up the certificate object based on the SNI associated with the certificate.

### Available Operations

* [ListSniWithCertificateInWorkspace](#listsniwithcertificateinworkspace) - List all SNIs associated with a Certificate in a workspace
* [CreateSniWithCertificateInWorkspace](#createsniwithcertificateinworkspace) - Create a new SNI associated with a Certificate in a workspace
* [DeleteSniWithCertificateInWorkspace](#deletesniwithcertificateinworkspace) - Delete a an SNI associated with a Certificate in a workspace
* [GetSniWithCertificateInWorkspace](#getsniwithcertificateinworkspace) - Get an SNI associated with a Certificate in a workspace
* [UpsertSniWithCertificateInWorkspace](#upsertsniwithcertificateinworkspace) - Upsert an SNI associated with a Certificate in a workspace
* [ListSniInWorkspace](#listsniinworkspace) - List all SNIs in a workspace
* [CreateSniInWorkspace](#createsniinworkspace) - Create a new SNI in a workspace
* [DeleteSniInWorkspace](#deletesniinworkspace) - Delete an SNI in a workspace
* [GetSniInWorkspace](#getsniinworkspace) - Get an SNI in a workspace
* [UpsertSniInWorkspace](#upsertsniinworkspace) - Upsert a SNI in a workspace
* [ListSniWithCertificate](#listsniwithcertificate) - List all SNIs associated with a Certificate
* [CreateSniWithCertificate](#createsniwithcertificate) - Create a new SNI associated with a Certificate
* [DeleteSniWithCertificate](#deletesniwithcertificate) - Delete a an SNI associated with a Certificate
* [GetSniWithCertificate](#getsniwithcertificate) - Get an SNI associated with a Certificate
* [UpsertSniWithCertificate](#upsertsniwithcertificate) - Upsert an SNI associated with a Certificate
* [ListSni](#listsni) - List all SNIs
* [CreateSni](#createsni) - Create a new SNI
* [DeleteSni](#deletesni) - Delete an SNI
* [GetSni](#getsni) - Get an SNI
* [UpsertSni](#upsertsni) - Upsert a SNI

## ListSniWithCertificateInWorkspace

List all SNIs associated with a Certificate in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-sni-with-certificate-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}/snis" -->
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

    res, err := s.SNIs.ListSniWithCertificateInWorkspace(ctx, operations.ListSniWithCertificateInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ListSniWithCertificateInWorkspaceRequest](../../models/operations/listsniwithcertificateinworkspacerequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ListSniWithCertificateInWorkspaceResponse](../../models/operations/listsniwithcertificateinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSniWithCertificateInWorkspace

Create a new SNI associated with a Certificate in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-sni-with-certificate-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}/snis" -->
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

    res, err := s.SNIs.CreateSniWithCertificateInWorkspace(ctx, operations.CreateSniWithCertificateInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        Workspace: "team-payments",
        SNIWithoutParents: components.SNIWithoutParents{
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.CreateSniWithCertificateInWorkspaceRequest](../../models/operations/createsniwithcertificateinworkspacerequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.CreateSniWithCertificateInWorkspaceResponse](../../models/operations/createsniwithcertificateinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSniWithCertificateInWorkspace

Delete a an SNI associated with a Certificate using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-sni-with-certificate-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}/snis/{SNIId}" -->
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

    res, err := s.SNIs.DeleteSniWithCertificateInWorkspace(ctx, operations.DeleteSniWithCertificateInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteSniWithCertificateInWorkspaceRequest](../../models/operations/deletesniwithcertificateinworkspacerequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.DeleteSniWithCertificateInWorkspaceResponse](../../models/operations/deletesniwithcertificateinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSniWithCertificateInWorkspace

Get an SNI associated with a Certificate using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-sni-with-certificate-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}/snis/{SNIId}" -->
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

    res, err := s.SNIs.GetSniWithCertificateInWorkspace(ctx, operations.GetSniWithCertificateInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetSniWithCertificateInWorkspaceRequest](../../models/operations/getsniwithcertificateinworkspacerequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.GetSniWithCertificateInWorkspaceResponse](../../models/operations/getsniwithcertificateinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertSniWithCertificateInWorkspace

Create or Update an SNI associated with a Certificate using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-sni-with-certificate-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/certificates/{CertificateId}/snis/{SNIId}" -->
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

    res, err := s.SNIs.UpsertSniWithCertificateInWorkspace(ctx, operations.UpsertSniWithCertificateInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        Workspace: "team-payments",
        SNIWithoutParents: components.SNIWithoutParents{
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.UpsertSniWithCertificateInWorkspaceRequest](../../models/operations/upsertsniwithcertificateinworkspacerequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.UpsertSniWithCertificateInWorkspaceResponse](../../models/operations/upsertsniwithcertificateinworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSniInWorkspace

List all SNIs in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-sni-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/snis" -->
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

    res, err := s.SNIs.ListSniInWorkspace(ctx, operations.ListSniInWorkspaceRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListSniInWorkspaceRequest](../../models/operations/listsniinworkspacerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListSniInWorkspaceResponse](../../models/operations/listsniinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateSniInWorkspace

Create a new SNI in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-sni-in-workspace" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/snis" -->
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

    res, err := s.SNIs.CreateSniInWorkspace(ctx, operations.CreateSniInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Sni: components.Sni{
            Certificate: components.SNICertificate{
                ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
            },
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateSniInWorkspaceRequest](../../models/operations/createsniinworkspacerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateSniInWorkspaceResponse](../../models/operations/createsniinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteSniInWorkspace

Delete an SNI in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-sni-in-workspace" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/snis/{SNIId}" -->
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

    res, err := s.SNIs.DeleteSniInWorkspace(ctx, operations.DeleteSniInWorkspaceRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteSniInWorkspaceRequest](../../models/operations/deletesniinworkspacerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.DeleteSniInWorkspaceResponse](../../models/operations/deletesniinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetSniInWorkspace

Get an SNI using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-sni-in-workspace" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/snis/{SNIId}" -->
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

    res, err := s.SNIs.GetSniInWorkspace(ctx, operations.GetSniInWorkspaceRequest{
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSniInWorkspaceRequest](../../models/operations/getsniinworkspacerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetSniInWorkspaceResponse](../../models/operations/getsniinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertSniInWorkspace

Create or Update SNI using ID or name in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-sni-in-workspace" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/snis/{SNIId}" -->
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

    res, err := s.SNIs.UpsertSniInWorkspace(ctx, operations.UpsertSniInWorkspaceRequest{
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        Sni: components.Sni{
            Certificate: components.SNICertificate{
                ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
            },
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpsertSniInWorkspaceRequest](../../models/operations/upsertsniinworkspacerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpsertSniInWorkspaceResponse](../../models/operations/upsertsniinworkspaceresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## ListSniWithCertificate

List all SNIs associated with a Certificate

### Example Usage

<!-- UsageSnippet language="go" operationID="list-sni-with-certificate" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/certificates/{CertificateId}/snis" -->
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

    res, err := s.SNIs.ListSniWithCertificate(ctx, operations.ListSniWithCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListSniWithCertificateRequest](../../models/operations/listsniwithcertificaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListSniWithCertificateResponse](../../models/operations/listsniwithcertificateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSniWithCertificate

Create a new SNI associated with a Certificate

### Example Usage

<!-- UsageSnippet language="go" operationID="create-sni-with-certificate" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/certificates/{CertificateId}/snis" -->
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

    res, err := s.SNIs.CreateSniWithCertificate(ctx, operations.CreateSniWithCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        SNIWithoutParents: components.SNIWithoutParents{
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateSniWithCertificateRequest](../../models/operations/createsniwithcertificaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateSniWithCertificateResponse](../../models/operations/createsniwithcertificateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSniWithCertificate

Delete a an SNI associated with a Certificate using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-sni-with-certificate" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/certificates/{CertificateId}/snis/{SNIId}" -->
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

    res, err := s.SNIs.DeleteSniWithCertificate(ctx, operations.DeleteSniWithCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeleteSniWithCertificateRequest](../../models/operations/deletesniwithcertificaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.DeleteSniWithCertificateResponse](../../models/operations/deletesniwithcertificateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSniWithCertificate

Get an SNI associated with a Certificate using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-sni-with-certificate" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/certificates/{CertificateId}/snis/{SNIId}" -->
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

    res, err := s.SNIs.GetSniWithCertificate(ctx, operations.GetSniWithCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetSniWithCertificateRequest](../../models/operations/getsniwithcertificaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetSniWithCertificateResponse](../../models/operations/getsniwithcertificateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpsertSniWithCertificate

Create or Update an SNI associated with a Certificate using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-sni-with-certificate" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/certificates/{CertificateId}/snis/{SNIId}" -->
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

    res, err := s.SNIs.UpsertSniWithCertificate(ctx, operations.UpsertSniWithCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CertificateID: "ddf3cdaa-3329-4961-822a-ce6dbd38eff7",
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        SNIWithoutParents: components.SNIWithoutParents{
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpsertSniWithCertificateRequest](../../models/operations/upsertsniwithcertificaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpsertSniWithCertificateResponse](../../models/operations/upsertsniwithcertificateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSni

List all SNIs

### Example Usage

<!-- UsageSnippet language="go" operationID="list-sni" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/snis" -->
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

    res, err := s.SNIs.ListSni(ctx, operations.ListSniRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
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
| `request`                                                              | [operations.ListSniRequest](../../models/operations/listsnirequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.ListSniResponse](../../models/operations/listsniresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateSni

Create a new SNI

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="create-sni" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/snis" example="DuplicateApiKey" -->
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

    res, err := s.SNIs.CreateSni(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Sni{
        Certificate: components.SNICertificate{
            ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
        },
        ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
        Name: "some.example.org",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="create-sni" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/snis" example="InvalidAuthCred" -->
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

    res, err := s.SNIs.CreateSni(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Sni{
        Certificate: components.SNICertificate{
            ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
        },
        ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
        Name: "some.example.org",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="create-sni" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/snis" example="NoAPIKey" -->
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

    res, err := s.SNIs.CreateSni(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.Sni{
        Certificate: components.SNICertificate{
            ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
        },
        ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
        Name: "some.example.org",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `controlPlaneID`                                                                                                                              | `string`                                                                                                                                      | :heavy_check_mark:                                                                                                                            | The UUID of your control plane. This variable is available in the Konnect manager.                                                            | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                                                          |
| `sni`                                                                                                                                         | [components.Sni](../../models/components/sni.md)                                                                                              | :heavy_check_mark:                                                                                                                            | Description of the new SNI for creation                                                                                                       | {<br/>"certificate": {<br/>"id": "bd380f99-659d-415e-b0e7-72ea05df3218"<br/>},<br/>"id": "36c4566c-14cc-4132-9011-4139fcbbe50a",<br/>"name": "some.example.org"<br/>} |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.CreateSniResponse](../../models/operations/createsniresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteSni

Delete an SNI

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-sni" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/snis/{SNIId}" -->
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

    res, err := s.SNIs.DeleteSni(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f")
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
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `sniID`                                                                            | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the SNI to lookup                                                            | 64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteSniResponse](../../models/operations/deletesniresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetSni

Get an SNI using ID or name.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-sni" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/snis/{SNIId}" -->
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

    res, err := s.SNIs.GetSni(ctx, "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `sniID`                                                                            | `string`                                                                           | :heavy_check_mark:                                                                 | ID of the SNI to lookup                                                            | 64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f                                               |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetSniResponse](../../models/operations/getsniresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertSni

Create or Update SNI using ID or name.

### Example Usage: DuplicateApiKey

<!-- UsageSnippet language="go" operationID="upsert-sni" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/snis/{SNIId}" example="DuplicateApiKey" -->
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

    res, err := s.SNIs.UpsertSni(ctx, operations.UpsertSniRequest{
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Sni: components.Sni{
            Certificate: components.SNICertificate{
                ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
            },
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```
### Example Usage: InvalidAuthCred

<!-- UsageSnippet language="go" operationID="upsert-sni" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/snis/{SNIId}" example="InvalidAuthCred" -->
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

    res, err := s.SNIs.UpsertSni(ctx, operations.UpsertSniRequest{
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Sni: components.Sni{
            Certificate: components.SNICertificate{
                ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
            },
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```
### Example Usage: NoAPIKey

<!-- UsageSnippet language="go" operationID="upsert-sni" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/snis/{SNIId}" example="NoAPIKey" -->
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

    res, err := s.SNIs.UpsertSni(ctx, operations.UpsertSniRequest{
        SNIID: "64c17a1a-b7d7-4a65-a5a4-42e4a7016e7f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Sni: components.Sni{
            Certificate: components.SNICertificate{
                ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
            },
            ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
            Name: "some.example.org",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sni != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.UpsertSniRequest](../../models/operations/upsertsnirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.UpsertSniResponse](../../models/operations/upsertsniresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |