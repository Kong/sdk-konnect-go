# BackendClusterAuthenticationSensitiveDataAwareScheme


## Supported Types

### BackendClusterAuthenticationAnonymous

```go
backendClusterAuthenticationSensitiveDataAwareScheme := components.CreateBackendClusterAuthenticationSensitiveDataAwareSchemeAnonymous(components.BackendClusterAuthenticationAnonymous{/* values here */})
```

### BackendClusterAuthenticationSaslPlainSensitiveDataAware

```go
backendClusterAuthenticationSensitiveDataAwareScheme := components.CreateBackendClusterAuthenticationSensitiveDataAwareSchemeSaslPlain(components.BackendClusterAuthenticationSaslPlainSensitiveDataAware{/* values here */})
```

### BackendClusterAuthenticationSaslScramSensitiveDataAware

```go
backendClusterAuthenticationSensitiveDataAwareScheme := components.CreateBackendClusterAuthenticationSensitiveDataAwareSchemeSaslScram(components.BackendClusterAuthenticationSaslScramSensitiveDataAware{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch backendClusterAuthenticationSensitiveDataAwareScheme.Type {
	case components.BackendClusterAuthenticationSensitiveDataAwareSchemeTypeAnonymous:
		// backendClusterAuthenticationSensitiveDataAwareScheme.BackendClusterAuthenticationAnonymous is populated
	case components.BackendClusterAuthenticationSensitiveDataAwareSchemeTypeSaslPlain:
		// backendClusterAuthenticationSensitiveDataAwareScheme.BackendClusterAuthenticationSaslPlainSensitiveDataAware is populated
	case components.BackendClusterAuthenticationSensitiveDataAwareSchemeTypeSaslScram:
		// backendClusterAuthenticationSensitiveDataAwareScheme.BackendClusterAuthenticationSaslScramSensitiveDataAware is populated
}
```
