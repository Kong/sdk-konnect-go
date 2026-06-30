# BackendClusterAuthenticationScheme


## Supported Types

### BackendClusterAuthenticationAnonymous

```go
backendClusterAuthenticationScheme := components.CreateBackendClusterAuthenticationSchemeAnonymous(components.BackendClusterAuthenticationAnonymous{/* values here */})
```

### BackendClusterAuthenticationSaslPlain

```go
backendClusterAuthenticationScheme := components.CreateBackendClusterAuthenticationSchemeSaslPlain(components.BackendClusterAuthenticationSaslPlain{/* values here */})
```

### BackendClusterAuthenticationSaslScram

```go
backendClusterAuthenticationScheme := components.CreateBackendClusterAuthenticationSchemeSaslScram(components.BackendClusterAuthenticationSaslScram{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch backendClusterAuthenticationScheme.Type {
	case components.BackendClusterAuthenticationSchemeTypeAnonymous:
		// backendClusterAuthenticationScheme.BackendClusterAuthenticationAnonymous is populated
	case components.BackendClusterAuthenticationSchemeTypeSaslPlain:
		// backendClusterAuthenticationScheme.BackendClusterAuthenticationSaslPlain is populated
	case components.BackendClusterAuthenticationSchemeTypeSaslScram:
		// backendClusterAuthenticationScheme.BackendClusterAuthenticationSaslScram is populated
}
```
