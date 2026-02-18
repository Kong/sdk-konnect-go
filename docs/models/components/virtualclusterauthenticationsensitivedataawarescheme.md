# VirtualClusterAuthenticationSensitiveDataAwareScheme


## Supported Types

### VirtualClusterAuthenticationAnonymous

```go
virtualClusterAuthenticationSensitiveDataAwareScheme := components.CreateVirtualClusterAuthenticationSensitiveDataAwareSchemeAnonymous(components.VirtualClusterAuthenticationAnonymous{/* values here */})
```

### VirtualClusterAuthenticationSaslPlainSensitiveDataAware

```go
virtualClusterAuthenticationSensitiveDataAwareScheme := components.CreateVirtualClusterAuthenticationSensitiveDataAwareSchemeSaslPlain(components.VirtualClusterAuthenticationSaslPlainSensitiveDataAware{/* values here */})
```

### VirtualClusterAuthenticationSaslScram

```go
virtualClusterAuthenticationSensitiveDataAwareScheme := components.CreateVirtualClusterAuthenticationSensitiveDataAwareSchemeSaslScram(components.VirtualClusterAuthenticationSaslScram{/* values here */})
```

### VirtualClusterAuthenticationOauthBearer

```go
virtualClusterAuthenticationSensitiveDataAwareScheme := components.CreateVirtualClusterAuthenticationSensitiveDataAwareSchemeOauthBearer(components.VirtualClusterAuthenticationOauthBearer{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualClusterAuthenticationSensitiveDataAwareScheme.Type {
	case components.VirtualClusterAuthenticationSensitiveDataAwareSchemeTypeAnonymous:
		// virtualClusterAuthenticationSensitiveDataAwareScheme.VirtualClusterAuthenticationAnonymous is populated
	case components.VirtualClusterAuthenticationSensitiveDataAwareSchemeTypeSaslPlain:
		// virtualClusterAuthenticationSensitiveDataAwareScheme.VirtualClusterAuthenticationSaslPlainSensitiveDataAware is populated
	case components.VirtualClusterAuthenticationSensitiveDataAwareSchemeTypeSaslScram:
		// virtualClusterAuthenticationSensitiveDataAwareScheme.VirtualClusterAuthenticationSaslScram is populated
	case components.VirtualClusterAuthenticationSensitiveDataAwareSchemeTypeOauthBearer:
		// virtualClusterAuthenticationSensitiveDataAwareScheme.VirtualClusterAuthenticationOauthBearer is populated
}
```
