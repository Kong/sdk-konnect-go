# CatalogIntegrationAuthorization

Defines the authorization strategy for an integration.


## Supported Types

### One

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationOne(components.One{/* values here */})
```

### OAuth

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationOAuth(components.OAuth{/* values here */})
```

### MultiKeyAuth

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationMultiKeyAuth(components.MultiKeyAuth{/* values here */})
```

### GitHubAppInstallationAuth

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationGitHubAppInstallationAuth(components.GitHubAppInstallationAuth{/* values here */})
```

