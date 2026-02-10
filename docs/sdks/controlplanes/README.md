# ControlPlanes

## Overview

### Available Operations

* [ListControlPlanes](#listcontrolplanes) - List Control Planes
* [CreateControlPlane](#createcontrolplane) - Create Control Plane
* [GetControlPlane](#getcontrolplane) - Get a Control Plane
* [UpdateControlPlane](#updatecontrolplane) - Update Control Plane
* [DeleteControlPlane](#deletecontrolplane) - Delete Control Plane

## ListControlPlanes

Returns an array of control plane objects containing information about the Konnect Control Planes.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-control-planes" method="get" path="/v2/control-planes" -->
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

    res, err := s.ControlPlanes.ListControlPlanes(ctx, operations.ListControlPlanesRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Filter: &components.ControlPlaneFilterParameters{
            CloudGateway: sdkkonnectgo.Pointer(true),
        },
        FilterLabels: sdkkonnectgo.Pointer("key:value,existCheck"),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListControlPlanesResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListControlPlanesRequest](../../models/operations/listcontrolplanesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListControlPlanesResponse](../../models/operations/listcontrolplanesresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.BaseError          | 500                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CreateControlPlane

Create a control plane in the Konnect Organization.

### Example Usage: Cannot Be Blank

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Cannot Be Blank" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Conflict

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Conflict" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Create Control Plane Response

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Create Control Plane Response" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Example Request Body

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Example Request Body" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeK8SIngressController.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Internal Server Error

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Internal Server Error" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Invalid ID Format

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Invalid ID Format" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Permission Denied

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Permission Denied" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Request Format Invalid

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Request Format Invalid" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Service Unavailable

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Service Unavailable" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Unauthorized" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Unknown Property

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Unknown Property" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Usage Limits Reached

<!-- UsageSnippet language="go" operationID="create-control-plane" method="post" path="/v2/control-planes" example="Usage Limits Reached" -->
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

    res, err := s.ControlPlanes.CreateControlPlane(ctx, components.CreateControlPlaneRequest{
        Name: "Test Control Plane",
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ClusterType: components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane.ToPointer(),
        AuthType: components.AuthTypePinnedClientCerts.ToPointer(),
        CloudGateway: sdkkonnectgo.Pointer(false),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateControlPlaneRequest](../../models/components/createcontrolplanerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateControlPlaneResponse](../../models/operations/createcontrolplaneresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.ConflictError       | 409                           | application/problem+json      |
| sdkerrors.InternalServerError | 500                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetControlPlane

Returns information about an individual control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-control-plane" method="get" path="/v2/control-planes/{id}" example="Single control plane response" -->
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

    res, err := s.ControlPlanes.GetControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The control plane ID                                     | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetControlPlaneResponse](../../models/operations/getcontrolplaneresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.BadRequestError    | 400                          | application/problem+json     |
| sdkerrors.UnauthorizedError  | 401                          | application/problem+json     |
| sdkerrors.ForbiddenError     | 403                          | application/problem+json     |
| sdkerrors.NotFoundError      | 404                          | application/problem+json     |
| sdkerrors.BaseError          | 500                          | application/problem+json     |
| sdkerrors.ServiceUnavailable | 503                          | application/problem+json     |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## UpdateControlPlane

Update an individual control plane.

### Example Usage: Cannot Be Blank

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Cannot Be Blank" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Example Request Body

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Example Request Body" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "development",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Internal Server Error

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Internal Server Error" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Invalid ID Format

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Invalid ID Format" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Not Found" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Permission Denied

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Permission Denied" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Request Format Invalid

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Request Format Invalid" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Service Unavailable

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Service Unavailable" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Unauthorized" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Unknown Property

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Unknown Property" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Update Control Plane Response

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Update Control Plane Response" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```
### Example Usage: Usage Limits Reached

<!-- UsageSnippet language="go" operationID="update-control-plane" method="patch" path="/v2/control-planes/{id}" example="Usage Limits Reached" -->
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

    res, err := s.ControlPlanes.UpdateControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.UpdateControlPlaneRequest{
        Name: sdkkonnectgo.Pointer("Test Control Plane"),
        Description: sdkkonnectgo.Pointer("A test control plane for exploration."),
        AuthType: components.UpdateControlPlaneRequestAuthTypePinnedClientCerts.ToPointer(),
        ProxyUrls: []components.ProxyURL{
            components.ProxyURL{
                Host: "example.com",
                Port: 443,
                Protocol: "https",
            },
        },
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlane != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | The control plane ID                                                                         | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                         |
| `updateControlPlaneRequest`                                                                  | [components.UpdateControlPlaneRequest](../../models/components/updatecontrolplanerequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.UpdateControlPlaneResponse](../../models/operations/updatecontrolplaneresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.InternalServerError | 500                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## DeleteControlPlane

Delete an individual control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-control-plane" method="delete" path="/v2/control-planes/{id}" -->
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

    res, err := s.ControlPlanes.DeleteControlPlane(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The control plane ID                                     | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteControlPlaneResponse](../../models/operations/deletecontrolplaneresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequestError     | 400                           | application/problem+json      |
| sdkerrors.UnauthorizedError   | 401                           | application/problem+json      |
| sdkerrors.ForbiddenError      | 403                           | application/problem+json      |
| sdkerrors.NotFoundError       | 404                           | application/problem+json      |
| sdkerrors.InternalServerError | 500                           | application/problem+json      |
| sdkerrors.ServiceUnavailable  | 503                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |