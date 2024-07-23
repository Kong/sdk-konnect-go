# sdk-konnect-go

This is a prototype and should not be used. See [CONTRIBUTING.md](https://github.com/Kong/sdk-konnect-go/blob/main/CONTRIBUTING.md) for information on how this SDK is generated.

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
	"log"
	"os"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String(os.Getenv("PERSONAL_ACCESS_TOKEN")),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request)
	if err != nil {

		var e *sdkerrors.UnauthorizedError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ForbiddenError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://global.api.konghq.com` | None |
| 1 | `https://us.api.konghq.com` | None |
| 2 | `https://eu.api.konghq.com` | None |
| 3 | `https://au.api.konghq.com` | None |

#### Example

```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
	"os"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithServerIndex(3),
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String(os.Getenv("PERSONAL_ACCESS_TOKEN")),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListServerlessCloudGatewaysResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
	"os"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithServerURL("https://global.api.konghq.com"),
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String(os.Getenv("PERSONAL_ACCESS_TOKEN")),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListServerlessCloudGatewaysResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
	"os"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String(os.Getenv("PERSONAL_ACCESS_TOKEN")),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request, operations.WithServerURL("https://global.api.konghq.com/"))
	if err != nil {
		log.Fatal(err)
	}
	if res.ListServerlessCloudGatewaysResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name                       | Type                       | Scheme                     |
| -------------------------- | -------------------------- | -------------------------- |
| `PersonalAccessToken`      | http                       | HTTP Bearer                |
| `SystemAccountAccessToken` | http                       | HTTP Bearer                |
| `KonnectAccessToken`       | http                       | HTTP Bearer                |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
	"os"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String(os.Getenv("PERSONAL_ACCESS_TOKEN")),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListServerlessCloudGatewaysResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"github.com/Kong/sdk-konnect-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String(os.Getenv("PERSONAL_ACCESS_TOKEN")),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ListServerlessCloudGatewaysResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"github.com/Kong/sdk-konnect-go/retry"
	"log"
	"os"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.String(os.Getenv("PERSONAL_ACCESS_TOKEN")),
		}),
	)
	request := operations.ListServerlessCloudGatewayRequest{
		PageSize:   sdkkonnectgo.Int64(10),
		PageNumber: sdkkonnectgo.Int64(1),
		Labels:     sdkkonnectgo.String("filter[labels][eq]=env:prod"),
	}
	ctx := context.Background()
	res, err := s.ServerlessCloudGateways.ListServerlessCloudGateway(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListServerlessCloudGatewaysResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

d