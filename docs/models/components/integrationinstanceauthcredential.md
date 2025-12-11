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

