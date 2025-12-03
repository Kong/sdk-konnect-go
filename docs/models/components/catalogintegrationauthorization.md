# CatalogIntegrationAuthorization

Defines the authorization strategy for an integration.


## Supported Types

### One

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationOne(components.One{/* values here */})
```

### OAuth1

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationOAuth1(components.OAuth1{/* values here */})
```

### MultiKeyAuth1

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationMultiKeyAuth1(components.MultiKeyAuth1{/* values here */})
```

### GitHubAppInstallationAuth

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationGitHubAppInstallationAuth(components.GitHubAppInstallationAuth{/* values here */})
```

