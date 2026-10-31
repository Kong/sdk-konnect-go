# CatalogIntegrationAuthorization

Defines the authorization strategy for an integration.


## Supported Types

### CatalogIntegrationAuthorization1

```go
catalogIntegrationAuthorization := components.CreateCatalogIntegrationAuthorizationCatalogIntegrationAuthorization1(components.CatalogIntegrationAuthorization1{/* values here */})
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

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogIntegrationAuthorization.Type {
	case components.CatalogIntegrationAuthorizationTypeCatalogIntegrationAuthorization1:
		// catalogIntegrationAuthorization.CatalogIntegrationAuthorization1 is populated
	case components.CatalogIntegrationAuthorizationTypeOAuth1:
		// catalogIntegrationAuthorization.OAuth1 is populated
	case components.CatalogIntegrationAuthorizationTypeMultiKeyAuth1:
		// catalogIntegrationAuthorization.MultiKeyAuth1 is populated
	case components.CatalogIntegrationAuthorizationTypeGitHubAppInstallationAuth:
		// catalogIntegrationAuthorization.GitHubAppInstallationAuth is populated
}
```
