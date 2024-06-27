# CreateRouteWithoutParentsHTTPSRedirectStatusCode

The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol.


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `CreateRouteWithoutParentsHTTPSRedirectStatusCodeFourHundredAndTwentySix` | 426                                                                       |
| `CreateRouteWithoutParentsHTTPSRedirectStatusCodeThreeHundredAndOne`      | 301                                                                       |
| `CreateRouteWithoutParentsHTTPSRedirectStatusCodeThreeHundredAndTwo`      | 302                                                                       |
| `CreateRouteWithoutParentsHTTPSRedirectStatusCodeThreeHundredAndSeven`    | 307                                                                       |
| `CreateRouteWithoutParentsHTTPSRedirectStatusCodeThreeHundredAndEight`    | 308                                                                       |