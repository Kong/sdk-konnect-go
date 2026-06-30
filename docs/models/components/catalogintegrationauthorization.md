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

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogIntegrationAuthorization.Type {
	case components.CatalogIntegrationAuthorizationTypeOne:
		// catalogIntegrationAuthorization.One is populated
	case components.CatalogIntegrationAuthorizationTypeOAuth:
		// catalogIntegrationAuthorization.OAuth is populated
	case components.CatalogIntegrationAuthorizationTypeMultiKeyAuth:
		// catalogIntegrationAuthorization.MultiKeyAuth is populated
	case components.CatalogIntegrationAuthorizationTypeGitHubAppInstallationAuth:
		// catalogIntegrationAuthorization.GitHubAppInstallationAuth is populated
}
```
