# CreateIntegrationInstanceAuthCredential


## Supported Types

### CreateOAuthCredential

```go
createIntegrationInstanceAuthCredential := components.CreateCreateIntegrationInstanceAuthCredentialCreateOAuthCredential(components.CreateOAuthCredential{/* values here */})
```

### CreateGitHubAppInstallationCredential

```go
createIntegrationInstanceAuthCredential := components.CreateCreateIntegrationInstanceAuthCredentialCreateGitHubAppInstallationCredential(components.CreateGitHubAppInstallationCredential{/* values here */})
```

### MultiKeyAuth1

```go
createIntegrationInstanceAuthCredential := components.CreateCreateIntegrationInstanceAuthCredentialMultiKeyAuth1(components.MultiKeyAuth1{/* values here */})
```

### CreateAWSRoleDelegationAuthCredential

```go
createIntegrationInstanceAuthCredential := components.CreateCreateIntegrationInstanceAuthCredentialCreateAWSRoleDelegationAuthCredential(components.CreateAWSRoleDelegationAuthCredential{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createIntegrationInstanceAuthCredential.Type {
	case components.CreateIntegrationInstanceAuthCredentialTypeCreateOAuthCredential:
		// createIntegrationInstanceAuthCredential.CreateOAuthCredential is populated
	case components.CreateIntegrationInstanceAuthCredentialTypeCreateGitHubAppInstallationCredential:
		// createIntegrationInstanceAuthCredential.CreateGitHubAppInstallationCredential is populated
	case components.CreateIntegrationInstanceAuthCredentialTypeMultiKeyAuth1:
		// createIntegrationInstanceAuthCredential.MultiKeyAuth1 is populated
	case components.CreateIntegrationInstanceAuthCredentialTypeCreateAWSRoleDelegationAuthCredential:
		// createIntegrationInstanceAuthCredential.CreateAWSRoleDelegationAuthCredential is populated
}
```
