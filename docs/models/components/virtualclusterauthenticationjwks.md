# VirtualClusterAuthenticationJWKS

JSON Web Key Set configuration for verifying token signatures.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Endpoint`                                                                        | *string*                                                                          | :heavy_check_mark:                                                                | URL for JWKS endpoint.                                                            |
| `Timeout`                                                                         | **string*                                                                         | :heavy_minus_sign:                                                                | Total time from establishing connection to receive a response from JWKS endpoint. |
| `CacheExpiration`                                                                 | **string*                                                                         | :heavy_minus_sign:                                                                | Duration after which the gateway will fetch and cache JWKS.                       |