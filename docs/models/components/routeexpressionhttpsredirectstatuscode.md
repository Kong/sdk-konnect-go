# RouteExpressionHTTPSRedirectStatusCode

The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RouteExpressionHTTPSRedirectStatusCodeThreeHundredAndOne

// Open enum: custom values can be created with a direct type cast
custom := components.RouteExpressionHTTPSRedirectStatusCode(999)
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `RouteExpressionHTTPSRedirectStatusCodeThreeHundredAndOne`      | 301                                                             |
| `RouteExpressionHTTPSRedirectStatusCodeThreeHundredAndTwo`      | 302                                                             |
| `RouteExpressionHTTPSRedirectStatusCodeThreeHundredAndSeven`    | 307                                                             |
| `RouteExpressionHTTPSRedirectStatusCodeThreeHundredAndEight`    | 308                                                             |
| `RouteExpressionHTTPSRedirectStatusCodeFourHundredAndTwentySix` | 426                                                             |