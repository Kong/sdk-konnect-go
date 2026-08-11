# AIGatewayMCPServerProtectedResourceMetadata

**Pre-release Feature**
This feature is currently in beta and is subject to change.

OAuth 2.0 Protected Resource Metadata (RFC 9728) advertised for this MCP
server, allowing clients to discover the authorization servers that
protect it.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `DiscoveryEndpoint`                                                               | `*string`                                                                         | :heavy_minus_sign:                                                                | The authorization server metadata discovery URL.                                  |
| `Endpoint`                                                                        | `*string`                                                                         | :heavy_minus_sign:                                                                | The URL path where the OAuth 2.0 Protected Resource Metadata is served.           |
| `AuthorizationServers`                                                            | []`string`                                                                        | :heavy_minus_sign:                                                                | List of authorization server issuer URLs that can issue tokens for this resource. |
| `Resource`                                                                        | `*string`                                                                         | :heavy_minus_sign:                                                                | The protected resource's identifier (resource URI).                               |
| `ScopesSupported`                                                                 | []`string`                                                                        | :heavy_minus_sign:                                                                | List of OAuth scopes supported by the protected resource.                         |