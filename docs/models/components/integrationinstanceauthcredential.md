# IntegrationInstanceAuthCredential

Object containing metadata for an integration instance auth credential.


## Supported Types

### Oauth

```go
integrationInstanceAuthCredential := components.CreateIntegrationInstanceAuthCredentialOauth(components.Oauth{/* values here */})
```

### GithubAppInstallation

```go
integrationInstanceAuthCredential := components.CreateIntegrationInstanceAuthCredentialGithubAppInstallation(components.GithubAppInstallation{/* values here */})
```

### MultiKeyAuthCredential

```go
integrationInstanceAuthCredential := components.CreateIntegrationInstanceAuthCredentialMultiKeyAuthCredential(components.MultiKeyAuthCredential{/* values here */})
```

### AWSRoleDelegationAuthCredential

```go
integrationInstanceAuthCredential := components.CreateIntegrationInstanceAuthCredentialAWSRoleDelegationAuthCredential(components.AWSRoleDelegationAuthCredential{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch integrationInstanceAuthCredential.Type {
	case components.IntegrationInstanceAuthCredentialTypeOauth:
		// integrationInstanceAuthCredential.Oauth is populated
	case components.IntegrationInstanceAuthCredentialTypeGithubAppInstallation:
		// integrationInstanceAuthCredential.GithubAppInstallation is populated
	case components.IntegrationInstanceAuthCredentialTypeMultiKeyAuthCredential:
		// integrationInstanceAuthCredential.MultiKeyAuthCredential is populated
	case components.IntegrationInstanceAuthCredentialTypeAWSRoleDelegationAuthCredential:
		// integrationInstanceAuthCredential.AWSRoleDelegationAuthCredential is populated
}
```
