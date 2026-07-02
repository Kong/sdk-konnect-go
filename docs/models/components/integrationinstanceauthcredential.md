# IntegrationInstanceAuthCredential

Object containing metadata for an integration instance auth credential.


## Supported Types

### Oauth1

```go
integrationInstanceAuthCredential := components.CreateIntegrationInstanceAuthCredentialOauth1(components.Oauth1{/* values here */})
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
	case components.IntegrationInstanceAuthCredentialTypeOauth1:
		// integrationInstanceAuthCredential.Oauth1 is populated
	case components.IntegrationInstanceAuthCredentialTypeGithubAppInstallation:
		// integrationInstanceAuthCredential.GithubAppInstallation is populated
	case components.IntegrationInstanceAuthCredentialTypeMultiKeyAuthCredential:
		// integrationInstanceAuthCredential.MultiKeyAuthCredential is populated
	case components.IntegrationInstanceAuthCredentialTypeAWSRoleDelegationAuthCredential:
		// integrationInstanceAuthCredential.AWSRoleDelegationAuthCredential is populated
}
```
