# SaslAwsIam


## Supported Types

### BackendClusterAuthenticationSaslAwsIamDefaultProviderChain

```go
saslAwsIam := components.CreateSaslAwsIamDefaultProviderChain(components.BackendClusterAuthenticationSaslAwsIamDefaultProviderChain{/* values here */})
```

### BackendClusterAuthenticationSaslAwsIamAssumeRole

```go
saslAwsIam := components.CreateSaslAwsIamAssumeRole(components.BackendClusterAuthenticationSaslAwsIamAssumeRole{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch saslAwsIam.Type {
	case components.SaslAwsIamTypeDefaultProviderChain:
		// saslAwsIam.BackendClusterAuthenticationSaslAwsIamDefaultProviderChain is populated
	case components.SaslAwsIamTypeAssumeRole:
		// saslAwsIam.BackendClusterAuthenticationSaslAwsIamAssumeRole is populated
}
```
